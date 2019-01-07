// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package trie

import (
	"fmt"
	"io"
	"math/big"
	"sync"
	"time"

	"github.com/allegro/bigcache"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	memcacheCleanHitMeter   = metrics.NewRegisteredMeter("trie/memcache/clean/hit", nil)
	memcacheCleanMissMeter  = metrics.NewRegisteredMeter("trie/memcache/clean/miss", nil)
	memcacheCleanReadMeter  = metrics.NewRegisteredMeter("trie/memcache/clean/read", nil)
	memcacheCleanWriteMeter = metrics.NewRegisteredMeter("trie/memcache/clean/write", nil)

	memcacheFlushTimeTimer  = metrics.NewRegisteredResettingTimer("trie/memcache/flush/time", nil)
	memcacheFlushNodesMeter = metrics.NewRegisteredMeter("trie/memcache/flush/nodes", nil)
	memcacheFlushSizeMeter  = metrics.NewRegisteredMeter("trie/memcache/flush/size", nil)

	memcacheGCTimeTimer  = metrics.NewRegisteredResettingTimer("trie/memcache/gc/time", nil)
	memcacheGCNodesMeter = metrics.NewRegisteredMeter("trie/memcache/gc/nodes", nil)
	memcacheGCSizeMeter  = metrics.NewRegisteredMeter("trie/memcache/gc/size", nil)

	memcachePruneTimeTimer  = metrics.NewRegisteredResettingTimer("trie/memcache/prune/time", nil)
	memcachePruneNodesMeter = metrics.NewRegisteredMeter("trie/memcache/prune/nodes", nil)
	memcachePruneSizeMeter  = metrics.NewRegisteredMeter("trie/memcache/prune/size", nil)

	memcacheCommitTimeTimer  = metrics.NewRegisteredResettingTimer("trie/memcache/commit/time", nil)
	memcacheCommitNodesMeter = metrics.NewRegisteredMeter("trie/memcache/commit/nodes", nil)
	memcacheCommitSizeMeter  = metrics.NewRegisteredMeter("trie/memcache/commit/size", nil)
)

// secureKeyPrefix is the database key prefix used to store trie node preimages.
var secureKeyPrefix = []byte("secure-key-")

// metaRoot is the identifier of the global memcache root that anchors the block
// accounts tries for garbage collection.
const metaRoot = ""

// DatabaseReader wraps the Get and Has method of a backing store for the trie.
type DatabaseReader interface {
	// Get retrieves the value associated with key from the database.
	Get(key []byte) (value []byte, err error)

	// Has retrieves whether a key is present in the database.
	Has(key []byte) (bool, error)
}

// makeNodeKey returns the database key for a trie node.
func makeNodeKey(owner common.Hash, hash common.Hash) string {
	if hash == (common.Hash{}) {
		return metaRoot
	}
	if owner == (common.Hash{}) {
		return string(hash[:])
	}
	return string(append(owner[:], hash[:]...))
}

// splitNodeKey returns the composing hashes of a trie node key.
func splitNodeKey(key string) (common.Hash, common.Hash) {
	switch len(key) {
	case 0:
		return common.Hash{}, common.Hash{}

	case common.HashLength:
		return common.Hash{}, common.BytesToHash([]byte(key))

	case 2 * common.HashLength:
		return common.BytesToHash([]byte(key[:common.HashLength])), common.BytesToHash([]byte(key[common.HashLength:]))

	default:
		panic(fmt.Sprintf("invalid node key: %s", key))
	}
}

// Database is an intermediate write layer between the trie data structures and
// the disk database. The aim is to accumulate trie writes in-memory and only
// periodically flush a couple tries to disk, garbage collecting the remainder.
type Database struct {
	diskdb ethdb.Database // Persistent storage for matured trie nodes

	cleans  *bigcache.BigCache     // GC friendly memory cache of clean node RLPs
	dirties map[string]*cachedNode // Data and references relationships of dirty nodes
	oldest  string                 // Oldest tracked node, flush-list head
	newest  string                 // Newest tracked node, flush-list tail

	preimages map[common.Hash][]byte // Preimages of nodes from the secure trie

	gctime  time.Duration      // Time spent on garbage collection since last commit
	gcnodes uint64             // Nodes garbage collected since last commit
	gcsize  common.StorageSize // Data storage garbage collected since last commit

	prunetime  time.Duration      // Time spend on disk pruning since last commit
	prunenodes uint64             // Nodes pruned from disk since last commit
	prunesize  common.StorageSize // Data storage pruned from disk since last commit

	flushtime  time.Duration      // Time spent on data flushing since last commit
	flushnodes uint64             // Nodes flushed since last commit
	flushsize  common.StorageSize // Data storage flushed since last commit

	dirtiesSize   common.StorageSize // Storage size of the dirty node cache (exc. flushlist)
	preimagesSize common.StorageSize // Storage size of the preimages cache

	lock sync.RWMutex
}

// rawNode is a simple binary blob used to differentiate between collapsed trie
// nodes and already encoded RLP binary blobs (while at the same time store them
// in the same cache fields).
type rawNode []byte

func (n rawNode) canUnload(uint16, uint16) bool { panic("this should never end up in a live trie") }
func (n rawNode) cache() (hashNode, bool)       { panic("this should never end up in a live trie") }
func (n rawNode) fstring(ind string) string     { panic("this should never end up in a live trie") }

// rawFullNode represents only the useful data content of a full node, with the
// caches and flags stripped out to minimize its data storage. This type honors
// the same RLP encoding as the original parent.
type rawFullNode [17]node

func (n rawFullNode) canUnload(uint16, uint16) bool { panic("this should never end up in a live trie") }
func (n rawFullNode) cache() (hashNode, bool)       { panic("this should never end up in a live trie") }

func (n rawFullNode) EncodeRLP(w io.Writer) error {
	var nodes [17]node

	for i, child := range n {
		if child != nil {
			nodes[i] = child
		} else {
			nodes[i] = nilValueNode
		}
	}
	return rlp.Encode(w, nodes)
}

func (n rawFullNode) String() string { return n.fstring("") }

func (n rawFullNode) fstring(ind string) string {
	resp := fmt.Sprintf("[\n%s  ", ind)
	for i, node := range n {
		if node == nil {
			resp += fmt.Sprintf("%s: <nil> ", indices[i])
		} else {
			resp += fmt.Sprintf("%s: %v", indices[i], node.fstring(ind+"  "))
		}
	}
	return resp + fmt.Sprintf("\n%s] ", ind)
}

// rawShortNode represents only the useful data content of a short node, with the
// caches and flags stripped out to minimize its data storage. This type honors
// the same RLP encoding as the original parent.
type rawShortNode struct {
	Key []byte
	Val node
}

func (n *rawShortNode) canUnload(uint16, uint16) bool {
	panic("this should never end up in a live trie")
}
func (n *rawShortNode) cache() (hashNode, bool) { panic("this should never end up in a live trie") }

func (n *rawShortNode) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &shortNode{Key: hexToCompact(n.Key), Val: n.Val})
}

func (n *rawShortNode) String() string { return n.fstring("") }

func (n *rawShortNode) fstring(ind string) string {
	return fmt.Sprintf("{%x: %v} ", n.Key, n.Val.fstring(ind+"  "))
}

// cachedNode is all the information we know about a single cached node in the
// memory database write layer.
type cachedNode struct {
	node node   // Cached collapsed trie node, or raw rlp data
	size uint16 // Byte size of the useful cached data

	parents  uint32            // Number of live nodes referencing this one
	children map[string]uint16 // External children referenced by this node

	flushPrev string // Previous node in the flush-list
	flushNext string // Next node in the flush-list
}

// rlp returns the raw rlp encoded blob of the cached node, either directly from
// the cache, or by regenerating it from the collapsed node.
func (n *cachedNode) rlp() []byte {
	if node, ok := n.node.(rawNode); ok {
		return node
	}
	blob, err := rlp.EncodeToBytes(n.node)
	if err != nil {
		panic(err)
	}
	return blob
}

// obj returns the decoded and expanded trie node, either directly from the cache,
// or by regenerating it from the rlp encoded blob.
func (n *cachedNode) obj(hash common.Hash, cachegen uint16) node {
	if node, ok := n.node.(rawNode); ok {
		return mustDecodeNode(hash[:], node, cachegen)
	}
	return expandNode(hash[:], n.node, cachegen)
}

// iterateRefs walks the embedded children of the cached node, tracking the
// internal path and invoking the provided callback on all hash nodes.
func (n *cachedNode) iterateRefs(path []byte, onHashNode func([]byte, common.Hash) error) error {
	if _, ok := n.node.(rawNode); ok {
		return nil
	}
	return iterateRefs(n.node, path, onHashNode)
}

// iterateRefs traverses the node hierarchy of a cached node and invokes the
// provided callback on all hash nodes.
func iterateRefs(n node, path []byte, onHashNode func([]byte, common.Hash) error) error {
	switch n := n.(type) {
	case *rawShortNode:
		return iterateRefs(n.Val, append(path, n.Key...), onHashNode)

	case *shortNode:
		return iterateRefs(n.Val, append(path, n.Key...), onHashNode)

	case rawFullNode:
		for i := 0; i < 16; i++ {
			if err := iterateRefs(n[i], append(path, byte(i)), onHashNode); err != nil {
				return err
			}
		}
		return nil

	case *fullNode:
		for i := 0; i < 16; i++ {
			if err := iterateRefs(n.Children[i], append(path, byte(i)), onHashNode); err != nil {
				return err
			}
		}
		return nil

	case hashNode:
		return onHashNode(path, common.BytesToHash(n))

	case valueNode, nil:
		return nil

	default:
		panic(fmt.Sprintf("unknown node type: %T", n))
	}
}

// simplifyNode traverses the hierarchy of an expanded memory node and discards
// all the internal caches, returning a node that only contains the raw data.
func simplifyNode(n node) node {
	switch n := n.(type) {
	case *shortNode:
		// Short nodes discard the flags and cascade
		return &rawShortNode{Key: compactToHex(n.Key), Val: simplifyNode(n.Val)}

	case *fullNode:
		// Full nodes discard the flags and cascade
		node := rawFullNode(n.Children)
		for i := 0; i < len(node); i++ {
			if node[i] != nil {
				node[i] = simplifyNode(node[i])
			}
		}
		return node

	case valueNode, hashNode, rawNode:
		return n

	default:
		panic(fmt.Sprintf("unknown node type: %T", n))
	}
}

// expandNode traverses the node hierarchy of a collapsed storage node and converts
// all fields and keys into expanded memory form.
func expandNode(hash hashNode, n node, cachegen uint16) node {
	switch n := n.(type) {
	case *rawShortNode:
		// Short nodes need key and child expansion
		return &shortNode{
			Key: n.Key,
			Val: expandNode(nil, n.Val, cachegen),
			flags: nodeFlag{
				hash: hash,
				gen:  cachegen,
			},
		}

	case rawFullNode:
		// Full nodes need child expansion
		node := &fullNode{
			flags: nodeFlag{
				hash: hash,
				gen:  cachegen,
			},
		}
		for i := 0; i < len(node.Children); i++ {
			if n[i] != nil {
				node.Children[i] = expandNode(nil, n[i], cachegen)
			}
		}
		return node

	case valueNode, hashNode:
		return n

	default:
		panic(fmt.Sprintf("unknown node type: %T", n))
	}
}

// NewDatabase creates a new trie database to store ephemeral trie content before
// its written out to disk or garbage collected. No read cache is created, so all
// data retrievals will hit the underlying disk database.
func NewDatabase(diskdb ethdb.Database) *Database {
	return NewDatabaseWithCache(diskdb, 0)
}

// NewDatabaseWithCache creates a new trie database to store ephemeral trie content
// before its written out to disk or garbage collected. It also acts as a read cache
// for nodes loaded from disk.
func NewDatabaseWithCache(diskdb ethdb.Database, cache int) *Database {
	var cleans *bigcache.BigCache
	if cache > 0 {
		cleans, _ = bigcache.NewBigCache(bigcache.Config{
			Shards:             1024,
			LifeWindow:         time.Hour,
			MaxEntriesInWindow: cache * 1024,
			MaxEntrySize:       512,
			HardMaxCacheSize:   cache,
		})
	}
	return &Database{
		diskdb:    diskdb,
		cleans:    cleans,
		dirties:   map[string]*cachedNode{metaRoot: {}},
		preimages: make(map[common.Hash][]byte),
	}
}

// DiskDB retrieves the persistent storage backing the trie database.
func (db *Database) DiskDB() DatabaseReader {
	return db.diskdb
}

// InsertBlob writes a new reference tracked blob to the memory database if it's
// yet unknown. This method should only be used for non-trie nodes that require
// reference counting, since trie nodes are garbage collected directly through
// their embedded children.
func (db *Database) InsertBlob(owner common.Hash, hash common.Hash, blob []byte) {
	db.lock.Lock()
	defer db.lock.Unlock()

	db.DiskDB().(ethdb.Database).Put([]byte(makeNodeKey(owner, hash)), blob)
	//db.insert(owner, hash, blob, rawNode(blob))
}

// insert inserts a collapsed trie node into the memory database. This method is
// a more generic version of InsertBlob, supporting both raw blob insertions as
// well ex trie node insertions. The blob must always be specified to allow proper
// size tracking.
func (db *Database) insert(owner common.Hash, hash common.Hash, blob []byte, node node) {
	// If the node's already cached, skip
	key := makeNodeKey(owner, hash)
	if _, ok := db.dirties[key]; ok {
		return
	}
	// Create the cached entry for this node
	entry := &cachedNode{
		node:      simplifyNode(node),
		size:      uint16(len(blob)),
		flushPrev: db.newest,
	}
	// Track all the implicit references (explicits must be empty)
	entry.iterateRefs(nil, func(path []byte, child common.Hash) error {
		if c := db.dirties[makeNodeKey(owner, child)]; c != nil {
			c.parents++
		}
		return nil
	})
	db.dirties[key] = entry

	// Update the flush-list endpoints
	if db.oldest == metaRoot {
		db.oldest, db.newest = key, key
	} else {
		db.dirties[db.newest].flushNext, db.newest = key, key
	}
	db.dirtiesSize += common.StorageSize(common.HashLength + entry.size)
}

// insertPreimage writes a new trie node pre-image to the memory database if it's
// yet unknown. The method will make a copy of the slice.
//
// Note, this method assumes that the database's lock is held!
func (db *Database) insertPreimage(hash common.Hash, preimage []byte) {
	if _, ok := db.preimages[hash]; ok {
		return
	}
	db.preimages[hash] = common.CopyBytes(preimage)
	db.preimagesSize += common.StorageSize(common.HashLength + len(preimage))
}

// node retrieves a cached trie node from memory, or returns nil if none can be
// found in the memory cache.
func (db *Database) node(owner common.Hash, hash common.Hash, cachegen uint16) node {
	key := makeNodeKey(owner, hash)

	// Retrieve the node from the clean cache if available
	if db.cleans != nil {
		if enc, err := db.cleans.Get(key); err == nil && enc != nil {
			memcacheCleanHitMeter.Mark(1)
			memcacheCleanReadMeter.Mark(int64(len(enc)))
			return mustDecodeNode(hash[:], enc, cachegen)
		}
	}
	// Retrieve the node from the dirty cache if available
	db.lock.RLock()
	dirty := db.dirties[key]
	db.lock.RUnlock()

	if dirty != nil {
		return dirty.obj(hash, cachegen)
	}
	// Content unavailable in memory, attempt to retrieve from disk
	enc, err := db.diskdb.Get([]byte(key))
	if err != nil || enc == nil {
		return nil
	}
	if db.cleans != nil {
		db.cleans.Set(string(hash[:]), enc)
		memcacheCleanMissMeter.Mark(1)
		memcacheCleanWriteMeter.Mark(int64(len(enc)))
	}
	return mustDecodeNode(hash[:], enc, cachegen)
}

// Node retrieves an encoded cached trie node from memory. If it cannot be found
// cached, the method queries the persistent database for the content.
func (db *Database) Node(owner common.Hash, hash common.Hash) ([]byte, error) {
	key := makeNodeKey(owner, hash)

	// Retrieve the node from the clean cache if available
	if db.cleans != nil {
		if enc, err := db.cleans.Get(key); err == nil && enc != nil {
			memcacheCleanHitMeter.Mark(1)
			memcacheCleanReadMeter.Mark(int64(len(enc)))
			return enc, nil
		}
	}
	// Retrieve the node from the dirty cache if available
	db.lock.RLock()
	dirty := db.dirties[key]
	db.lock.RUnlock()

	if dirty != nil {
		return dirty.rlp(), nil
	}
	// Content unavailable in memory, attempt to retrieve from disk
	enc, err := db.diskdb.Get([]byte(key))
	if err == nil && enc != nil {
		if db.cleans != nil {
			db.cleans.Set(key, enc)
			memcacheCleanMissMeter.Mark(1)
			memcacheCleanWriteMeter.Mark(int64(len(enc)))
		}
	}
	return enc, err
}

// preimage retrieves a cached trie node pre-image from memory. If it cannot be
// found cached, the method queries the persistent database for the content.
func (db *Database) preimage(hash common.Hash) ([]byte, error) {
	// Retrieve the node from cache if available
	db.lock.RLock()
	preimage := db.preimages[hash]
	db.lock.RUnlock()

	if preimage != nil {
		return preimage, nil
	}
	// Content unavailable in memory, attempt to retrieve from disk
	return db.diskdb.Get(db.preimageKey(hash[:]))
}

// preimageKey returns the database key for the preimage of key.
func (db *Database) preimageKey(key []byte) []byte {
	return append(secureKeyPrefix, key...)
}

// Nodes retrieves the hashes of all the nodes cached within the memory database.
// This method is extremely expensive and should only be used to validate internal
// states in test code.
func (db *Database) Nodes() []string {
	db.lock.RLock()
	defer db.lock.RUnlock()

	var keys = make([]string, 0, len(db.dirties))
	for key := range db.dirties {
		if key != metaRoot { // Special case for "root" references/nodes
			keys = append(keys, key)
		}
	}
	return keys
}

// Reference adds a new reference from a parent node to a child node. We're going
// to break genericity here and assume that parent nodes are not owned (account
// trie) whereas child nodes may be owned (storage trie or bytecode).
func (db *Database) Reference(owner common.Hash, child common.Hash, parent common.Hash) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	// If the node does not exist, it's a node pulled from disk, skip
	childKey := makeNodeKey(owner, child)
	node, ok := db.dirties[childKey]
	if !ok {
		return
	}
	// If the reference already exists, only duplicate for roots
	parentKey := makeNodeKey(common.Hash{}, parent)
	if db.dirties[parentKey].children == nil {
		db.dirties[parentKey].children = make(map[string]uint16)
	} else if _, ok = db.dirties[parentKey].children[childKey]; ok && parent != (common.Hash{}) {
		return
	}
	node.parents++
	db.dirties[parentKey].children[childKey]++
}

// Dereference removes an existing reference from a root node.
func (db *Database) Dereference(root common.Hash, prune bool) error {
	// Sanity check to ensure that the meta-root is not removed
	if root == (common.Hash{}) {
		log.Error("Attempted to dereference the trie cache meta root")
		return nil
	}
	db.lock.Lock()
	defer db.lock.Unlock()

	nodes, storage, start := len(db.dirties), db.dirtiesSize, time.Now()
	prunetime, prunenodes, prunesize := db.prunetime, db.prunenodes, db.prunesize
	if err := db.dereference(common.Hash{}, root, common.Hash{}, common.Hash{}, prune, nil, make(map[string]node)); err != nil {
		return err
	}
	db.gcnodes += uint64(nodes - len(db.dirties))
	db.gcsize += storage - db.dirtiesSize
	db.gctime += time.Since(start)

	memcacheGCTimeTimer.Update(time.Since(start))
	memcacheGCSizeMeter.Mark(int64(storage - db.dirtiesSize))
	memcacheGCNodesMeter.Mark(int64(nodes - len(db.dirties)))

	memcachePruneTimeTimer.Update(db.prunetime - prunetime)
	memcachePruneNodesMeter.Mark(int64(db.prunenodes - prunenodes))
	memcachePruneSizeMeter.Mark(int64(db.prunesize - prunesize))

	log.Debug("Dereferenced trie from memory database", "nodes", nodes-len(db.dirties), "size", storage-db.dirtiesSize, "time", common.PrettyDuration(time.Since(start)),
		"gcnodes", db.gcnodes, "gcsize", db.gcsize, "gctime", common.PrettyDuration(db.gctime), "prnodes", db.prunenodes, "prsize", db.prunesize, "prtime", common.PrettyDuration(db.prunetime),
		"livenodes", len(db.dirties), "livesize", db.dirtiesSize)

	return nil
}

// dereference is the private locked version of Dereference.
func (db *Database) dereference(childOwner common.Hash, childHash common.Hash, parentOwner common.Hash, parentHash common.Hash, prune bool, path []byte, cache map[string]node) error {
	// Dereference the parent-child
	parentKey := makeNodeKey(parentOwner, parentHash)
	parent := db.dirties[parentKey]

	childKey := makeNodeKey(childOwner, childHash)
	if parent.children != nil && parent.children[childKey] > 0 {
		parent.children[childKey]--
		if parent.children[childKey] == 0 {
			delete(parent.children, childKey)
		}
	}
	// If the child does not exist, it's a previously committed node.
	child, ok := db.dirties[childKey]
	if !ok {
		if prune {
			batch := db.diskdb.NewBatch()

			start := time.Now()
			db.prune(childOwner, childHash, path, batch, cache)
			db.prunetime += time.Since(start)

			if err := batch.Write(); err != nil {
				return err
			}
		}
		return nil
	}
	// If there are no more references to the child, delete it and cascade
	if child.parents > 0 {
		// This is a special cornercase where a node loaded from disk (i.e. not in the
		// memcache any more) gets reinjected as a new node (short node split into full,
		// then reverted into short), causing a cached node to have no parents. That is
		// no problem in itself, but don't make maxint parents out of it.
		child.parents--
	}
	if child.parents == 0 {
		// Remove the node from the flush-list
		switch childKey {
		case db.oldest:
			db.oldest = child.flushNext
			db.dirties[child.flushNext].flushPrev = metaRoot
		case db.newest:
			db.newest = child.flushPrev
			db.dirties[child.flushPrev].flushNext = metaRoot
		default:
			db.dirties[child.flushPrev].flushNext = child.flushNext
			db.dirties[child.flushNext].flushPrev = child.flushPrev
		}
		// Dereference all children and delete the node
		child.iterateRefs(path, func(path []byte, hash common.Hash) error {
			db.dereference(childOwner, hash, childOwner, childHash, prune, path, cache)
			return nil
		})
		for key := range child.children {
			owner, hash := splitNodeKey(key)
			db.dereference(owner, hash, childOwner, childHash, prune, nil, cache)
		}
		delete(db.dirties, childKey)
		db.dirtiesSize -= common.StorageSize(common.HashLength + int(child.size))
	}
	return nil
}

// prune deletes a trie node from disk if there are no more live references to
// it, cascading until all dangling nodes are removed.
func (db *Database) prune(owner common.Hash, hash common.Hash, path []byte, batch ethdb.Batch, cache map[string]node) {
	// If the node is still live in the memory cache, it's still referenced so we
	// can abort. This case is important when and old trie being pruned references
	// a new node (maybe that node was recreted since), since currently live nodes
	// are stored expanded, not as hashes.
	key := makeNodeKey(owner, hash)
	if db.dirties[key] != nil {
		return
	}
	// Iterate over all the live tries in the cache and check node liveliness
	for key := range db.dirties[metaRoot].children {
		_, root := splitNodeKey(key)

		var paths [][]byte
		if owner != (common.Hash{}) {
			paths = [][]byte{keybytesToHex(owner[:])}
		}
		if db.live(hashNode(root[:]), owner, hash, append(paths, path), cache) {
			return
		}
	}
	// Dead node found, delete it from the database
	dead := []byte(makeNodeKey(owner, hash))
	blob, err := db.diskdb.Get(dead)
	if blob == nil || err != nil {
		log.Error("Missing prune target", "owner", owner, "hash", hash, "path", fmt.Sprintf("%x", path))
		return
	}
	node := mustDecodeNode(hash[:], blob, 0)

	// Prune the node and its children if it's not a bytecode blob
	db.cleans.Delete(key)
	batch.Delete(dead)
	db.prunenodes++
	db.prunesize += common.StorageSize(len(blob))

	iterateRefs(node, path, func(path []byte, hash common.Hash) error {
		db.prune(owner, hash, path, batch, cache)
		return nil
	})
}

// live descends in the trie and returns whether the given hash is part of the
// trie or not.
func (db *Database) live(root node, owner common.Hash, hash common.Hash, paths [][]byte, cache map[string]node) bool {
	// If we reached the end of our path, it should be a hash node
	if len(paths) == 1 && len(paths[0]) == 0 {
		if have, ok := root.(hashNode); ok {
			return common.BytesToHash(have) == hash
		}
		// Not a hash node? It rarely happens that a 32+ byte leaf short node gets
		// turned into a 31- byte one, converting if from a hash node to an embedded
		// one. Allow this case, but reject as a wrong path.

		// TODO(karalabe): get rid of this warning, only curiosity
		log.Warn("Liveness check terminated on non-hash", "type", fmt.Sprintf("%T", root), "node", root.fstring(""))

		return false
	}
	// If we're at a hash node, expand before continuing
	if n, ok := root.(hashNode); ok {
		var (
			key  string
			hash = common.BytesToHash(n)
		)
		if len(paths) > 1 {
			key = makeNodeKey(common.Hash{}, hash)
		} else {
			key = makeNodeKey(owner, hash)
		}
		if enc, err := db.cleans.Get(key); err == nil && enc != nil {
			root = mustDecodeNode(hash[:], enc, 0)
			cache[key] = root
		} else if node := db.dirties[key]; node != nil {
			root = node.node
		} else if node := cache[key]; node != nil {
			root = node
		} else {
			blob, err := db.diskdb.Get([]byte(key))
			if blob == nil || err != nil {
				panic(fmt.Sprintf("missing referenced node %x (searching for %x:%x at %x)", key, owner, hash, paths))
			}
			root = mustDecodeNode(hash[:], blob, 0)
			cache[key] = root
		}
	}
	// If we reached an account node, extract the storage trie root to continue on
	if len(paths) == 2 && len(paths[0]) == 0 {
		if have, ok := root.(valueNode); ok {
			var account struct {
				Nonce    uint64
				Balance  *big.Int
				Root     common.Hash
				CodeHash []byte
			}
			if err := rlp.DecodeBytes(have, &account); err != nil {
				panic(err)
			}
			if account.Root == emptyRoot {
				return false
			}
			return db.live(hashNode(account.Root[:]), owner, hash, paths[1:], cache)
		}
		panic(fmt.Sprintf("liveness check path swap terminated on non value node: %T", root))
	}

	// Descend into the trie following the specified path. This code segment must
	// be able to handle both simplified raw nodes kept in this cache as well as
	// cold nodes loaded directly from disk.
	switch n := root.(type) {
	case *rawShortNode:
		if prefixLen(n.Key, paths[0]) == len(n.Key) {
			return db.live(n.Val, owner, hash, append([][]byte{paths[0][len(n.Key):]}, paths[1:]...), cache)
		}
		return false

	case *shortNode:
		if prefixLen(n.Key, paths[0]) == len(n.Key) {
			return db.live(n.Val, owner, hash, append([][]byte{paths[0][len(n.Key):]}, paths[1:]...), cache)
		}
		return false

	case rawFullNode:
		if child := n[paths[0][0]]; child != nil {
			return db.live(child, owner, hash, append([][]byte{paths[0][1:]}, paths[1:]...), cache)
		}
		return false

	case *fullNode:
		if child := n.Children[paths[0][0]]; child != nil {
			return db.live(child, owner, hash, append([][]byte{paths[0][1:]}, paths[1:]...), cache)
		}
		return false

	default:
		panic(fmt.Sprintf("unknown node type: %T", n))
	}
}

// Cap iteratively flushes old but still referenced trie nodes until the total
// memory usage goes below the given threshold.
func (db *Database) Cap(limit common.StorageSize) error {
	// Create a database batch to flush persistent data out. It is important that
	// outside code doesn't see an inconsistent state (referenced data removed from
	// memory cache during commit but not yet in persistent storage). This is ensured
	// by only uncaching existing data when the database write finalizes.
	db.lock.RLock()

	nodes, storage, start := len(db.dirties), db.dirtiesSize, time.Now()
	batch := db.diskdb.NewBatch()

	// db.dirtiesSize only contains the useful data in the cache, but when reporting
	// the total memory consumption, the maintenance metadata is also needed to be
	// counted. For every useful node, we track 2 extra hashes as the flushlist.
	size := db.dirtiesSize + common.StorageSize((len(db.dirties)-1)*2*common.HashLength)

	// If the preimage cache got large enough, push to disk. If it's still small
	// leave for later to deduplicate writes.
	flushPreimages := db.preimagesSize > 4*1024*1024
	if flushPreimages {
		for hash, preimage := range db.preimages {
			if err := batch.Put(db.preimageKey(hash[:]), preimage); err != nil {
				log.Error("Failed to commit preimage from trie database", "err", err)
				db.lock.RUnlock()
				return err
			}
			if batch.ValueSize() > ethdb.IdealBatchSize {
				if err := batch.Write(); err != nil {
					db.lock.RUnlock()
					return err
				}
				batch.Reset()
			}
		}
	}
	// Keep committing nodes from the flush-list until we're below allowance
	oldest := db.oldest
	for size > limit && oldest != metaRoot {
		// Fetch the oldest referenced node and push into the batch
		node := db.dirties[oldest]
		if err := batch.Put([]byte(oldest), node.rlp()); err != nil {
			db.lock.RUnlock()
			return err
		}
		// If we exceeded the ideal batch size, commit and reset
		if batch.ValueSize() >= ethdb.IdealBatchSize {
			if err := batch.Write(); err != nil {
				log.Error("Failed to write flush list to disk", "err", err)
				db.lock.RUnlock()
				return err
			}
			batch.Reset()
		}
		// Iterate to the next flush item, or abort if the size cap was achieved. Size
		// is the total size, including both the useful cached data (hash -> blob), as
		// well as the flushlist metadata (2*hash). When flushing items from the cache,
		// we need to reduce both.
		size -= common.StorageSize(3*common.HashLength + int(node.size))
		oldest = node.flushNext
	}
	// Flush out any remainder data from the last batch
	if err := batch.Write(); err != nil {
		log.Error("Failed to write flush list to disk", "err", err)
		db.lock.RUnlock()
		return err
	}
	db.lock.RUnlock()

	// Write successful, clear out the flushed data
	db.lock.Lock()
	defer db.lock.Unlock()

	if flushPreimages {
		db.preimages = make(map[common.Hash][]byte)
		db.preimagesSize = 0
	}
	for db.oldest != oldest {
		node := db.dirties[db.oldest]
		delete(db.dirties, db.oldest)
		db.oldest = node.flushNext

		db.dirtiesSize -= common.StorageSize(common.HashLength + int(node.size))
	}
	if db.oldest != metaRoot {
		db.dirties[db.oldest].flushPrev = metaRoot
	}
	db.flushnodes += uint64(nodes - len(db.dirties))
	db.flushsize += storage - db.dirtiesSize
	db.flushtime += time.Since(start)

	memcacheFlushTimeTimer.Update(time.Since(start))
	memcacheFlushSizeMeter.Mark(int64(storage - db.dirtiesSize))
	memcacheFlushNodesMeter.Mark(int64(nodes - len(db.dirties)))

	log.Debug("Persisted nodes from memory database", "nodes", nodes-len(db.dirties), "size", storage-db.dirtiesSize, "time", common.PrettyDuration(time.Since(start)),
		"flnodes", db.flushnodes, "flsize", db.flushsize, "fltime", common.PrettyDuration(db.flushtime), "livenodes", len(db.dirties), "livesize", db.dirtiesSize)

	return nil
}

// Commit iterates over all the children of a particular node, writes them out
// to disk, forcefully tearing down all references in both directions.
//
// As a side effect, all pre-images accumulated up to this point are also written.
func (db *Database) Commit(node common.Hash, report bool) error {
	// Create a database batch to flush persistent data out. It is important that
	// outside code doesn't see an inconsistent state (referenced data removed from
	// memory cache during commit but not yet in persistent storage). This is ensured
	// by only uncaching existing data when the database write finalizes.
	db.lock.RLock()

	start := time.Now()
	batch := db.diskdb.NewBatch()

	// Move all of the accumulated preimages into a write batch
	for hash, preimage := range db.preimages {
		if err := batch.Put(db.preimageKey(hash[:]), preimage); err != nil {
			log.Error("Failed to commit preimage from trie database", "err", err)
			db.lock.RUnlock()
			return err
		}
		if batch.ValueSize() > ethdb.IdealBatchSize {
			if err := batch.Write(); err != nil {
				return err
			}
			batch.Reset()
		}
	}
	// Move the trie itself into the batch, flushing if enough data is accumulated
	nodes, storage := len(db.dirties), db.dirtiesSize
	if err := db.commit(common.Hash{}, node, batch); err != nil {
		log.Error("Failed to commit trie from trie database", "err", err)
		db.lock.RUnlock()
		return err
	}
	// Write batch ready, unlock for readers during persistence
	if err := batch.Write(); err != nil {
		log.Error("Failed to write trie to disk", "err", err)
		db.lock.RUnlock()
		return err
	}
	db.lock.RUnlock()

	// Write successful, clear out the flushed data
	db.lock.Lock()
	defer db.lock.Unlock()

	db.preimages = make(map[common.Hash][]byte)
	db.preimagesSize = 0

	db.uncache(common.Hash{}, node)

	memcacheCommitTimeTimer.Update(time.Since(start))
	memcacheCommitSizeMeter.Mark(int64(storage - db.dirtiesSize))
	memcacheCommitNodesMeter.Mark(int64(nodes - len(db.dirties)))

	logger := log.Info
	if !report {
		logger = log.Debug
	}
	logger("Persisted trie from memory database", "nodes", nodes-len(db.dirties)+int(db.flushnodes), "size", storage-db.dirtiesSize+db.flushsize, "time", common.PrettyDuration(time.Since(start)+db.flushtime),
		"gcnodes", db.gcnodes, "gcsize", db.gcsize, "gctime", common.PrettyDuration(db.gctime), "prnodes", db.prunenodes, "prsize", db.prunesize, "prtime", common.PrettyDuration(db.prunetime),
		"linodes", len(db.dirties), "lisize", db.dirtiesSize)

	// Reset the garbage collection statistics
	db.gcnodes, db.gcsize, db.gctime = 0, 0, 0
	db.prunenodes, db.prunesize, db.prunetime = 0, 0, 0
	db.flushnodes, db.flushsize, db.flushtime = 0, 0, 0

	return nil
}

// commit is the private locked version of Commit.
func (db *Database) commit(owner common.Hash, hash common.Hash, batch ethdb.Batch) error {
	// If the node does not exist, it's a previously committed node
	key := makeNodeKey(owner, hash)

	node, ok := db.dirties[key]
	if !ok {
		return nil
	}
	if err := node.iterateRefs(nil, func(path []byte, child common.Hash) error {
		return db.commit(owner, child, batch)
	}); err != nil {
		return err
	}
	for child := range node.children {
		owner, hash := splitNodeKey(child)
		if err := db.commit(owner, hash, batch); err != nil {
			return err
		}
	}
	if err := batch.Put([]byte(key), node.rlp()); err != nil {
		return err
	}
	// If we've reached an optimal batch size, commit and start over
	if batch.ValueSize() >= ethdb.IdealBatchSize {
		if err := batch.Write(); err != nil {
			return err
		}
		batch.Reset()
	}
	return nil
}

// uncache is the post-processing step of a commit operation where the already
// persisted trie is removed from the cache. The reason behind the two-phase
// commit is to ensure consistent data availability while moving from memory
// to disk.
func (db *Database) uncache(owner common.Hash, hash common.Hash) {
	// If the node does not exist, we're done on this path
	key := makeNodeKey(owner, hash)

	node, ok := db.dirties[key]
	if !ok {
		return
	}
	// Node still exists, remove it from the flush-list
	switch key {
	case db.oldest:
		db.oldest = node.flushNext
		db.dirties[node.flushNext].flushPrev = metaRoot
	case db.newest:
		db.newest = node.flushPrev
		db.dirties[node.flushPrev].flushNext = metaRoot
	default:
		db.dirties[node.flushPrev].flushNext = node.flushNext
		db.dirties[node.flushNext].flushPrev = node.flushPrev
	}
	// Uncache the node's subtries and remove the node itself too
	node.iterateRefs(nil, func(path []byte, child common.Hash) error {
		db.uncache(owner, child)
		return nil
	})
	for child := range node.children {
		db.uncache(splitNodeKey(child))
	}
	delete(db.dirties, key)
	db.dirtiesSize -= common.StorageSize(common.HashLength + int(node.size))
}

// Size returns the current storage size of the memory cache in front of the
// persistent database layer.
func (db *Database) Size() (common.StorageSize, common.StorageSize) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	// db.dirtiesSize only contains the useful data in the cache, but when reporting
	// the total memory consumption, the maintenance metadata is also needed to be
	// counted. For every useful node, we track 2 extra hashes as the flushlist.
	var flushlistSize = common.StorageSize((len(db.dirties) - 1) * 2 * common.HashLength)
	return db.dirtiesSize + flushlistSize, db.preimagesSize
}

// verifyIntegrity is a debug method to iterate over the entire trie stored in
// memory and check whether every node is reachable from the meta root. The goal
// is to find any errors that might cause memory leaks and or trie nodes to go
// missing.
//
// This method is extremely CPU and memory intensive, only use when must.
func (db *Database) verifyIntegrity() {
	// Iterate over all the cached nodes and accumulate them into a set
	reachable := map[string]struct{}{metaRoot: struct{}{}}

	for key := range db.dirties[metaRoot].children {
		_, root := splitNodeKey(key)
		db.accumulate(common.Hash{}, root, reachable)
	}
	// Find any unreachable but cached nodes
	unreachable := []string{}
	for key, node := range db.dirties {
		if _, ok := reachable[key]; !ok {
			unreachable = append(unreachable, fmt.Sprintf("%x: {Node: %v, Parents: %d, Prev: %x, Next: %x}",
				key, node.node, node.parents, node.flushPrev, node.flushNext))
		}
	}
	if len(unreachable) != 0 {
		panic(fmt.Sprintf("trie cache memory leak: %v", unreachable))
	}
}

// accumulate iterates over the trie defined by owner:hash and accumulates all
// the cached children found in memory.
func (db *Database) accumulate(owner common.Hash, hash common.Hash, reachable map[string]struct{}) {
	// Mark the node reachable if present in the memory cache
	key := makeNodeKey(owner, hash)

	node, ok := db.dirties[key]
	if !ok {
		return
	}
	reachable[key] = struct{}{}

	// Iterate over all the children and accumulate them too
	node.iterateRefs(nil, func(path []byte, hash common.Hash) error {
		db.accumulate(owner, hash, reachable)
		return nil
	})
	for key := range node.children {
		owner, hash := splitNodeKey(key)
		db.accumulate(owner, hash, reachable)
	}
}
