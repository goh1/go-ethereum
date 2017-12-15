// Code generated by go-bindata.
// sources:
// 4byte_tracer.js
// call_tracer.js
// emvdis_tracer.js
// genesis_tracer.js
// noop_tracer.js
// opcount_tracer.js
// DO NOT EDIT!

package tracers

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __4byte_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\xdd\x6e\xdb\x38\x13\xbd\x96\x9e\x62\x3e\xdf\xc4\x4a\x64\xc9\xf2\x4f\xfc\x13\xa4\xf8\xb2\x6d\xd0\x16\x68\xd1\x05\x12\xec\xcd\x62\x2f\x68\x72\x64\x71\x2d\x93\x02\x39\x72\xec\xa4\x7e\xf7\x05\x29\xb9\xb6\x93\x60\x81\xd5\x85\x2d\x91\x33\x87\x47\x67\xe6\x8c\xd2\xcb\xcb\x30\x84\xd1\x62\x47\x68\xc1\x22\x33\xbc\x40\x0b\xb9\x36\xcd\x5a\x4f\x0a\x54\x24\x73\x89\xc6\xc6\xc0\x94\x00\xae\xcb\x12\x39\x59\xa0\x02\xd7\x3e\xb0\xd2\x96\x7a\x95\xd1\x1c\xad\x95\x6a\x99\x84\xf0\x95\xce\xc2\x60\x8d\x54\x68\x61\xe1\x04\x0c\x58\xa9\xd5\x12\x9e\x24\x15\x3e\xc4\xca\x67\x04\x9d\x37\xf7\x75\x55\x95\x12\x05\x08\x46\x2c\x06\xab\x81\x0a\x46\x21\x30\x30\xb8\x41\x63\x51\x80\x95\x4b\xc5\xa8\x36\x08\x9c\x29\x58\x20\xac\x19\xf1\x02\x05\xb0\x25\x93\xca\xd2\x1b\x4c\x07\x95\x84\xe1\xfd\x96\xad\xab\x12\xe7\x61\x08\x00\xf0\x01\x04\x2e\xea\x65\x42\x86\x71\x7c\x34\x4c\x59\xc6\x49\x6a\xd5\x85\x4e\x7f\x3b\xc8\x46\x38\x9e\x4d\x70\x38\x16\xac\x3f\x1d\x5e\xcf\x06\xf9\x78\x38\xbd\xce\x46\x19\x5e\xcf\xf2\xd1\x04\x67\x93\xe1\x62\xc0\xc7\xd7\x38\x61\xd3\xfe\x64\xb8\xc8\x90\xf5\xa7\xb9\x98\x8c\x27\x19\xce\x04\x76\x62\x78\xf1\xc0\x66\x0e\x9d\x46\xe0\xce\x3e\xf2\xe7\xbe\xf8\x5f\x80\xfe\x76\x30\x11\x7c\x30\x9b\x60\x2f\x1b\x4c\xe7\x90\xc5\xbf\x36\x86\x53\xce\x47\xd3\x61\xd6\xeb\xcf\x61\x70\x5c\x1e\x0f\x46\xf9\x70\x3a\x9d\xf5\x66\xd7\x67\xe1\x4c\xe4\xe3\x59\x3e\x9b\xf5\x06\xd3\x73\x1c\x3e\x98\x66\x22\x9b\xa1\xc3\xc9\xfc\xea\x3e\x0c\xff\xcf\x6a\x2a\xb4\x81\xef\x5f\xe0\xe1\x09\x95\xc0\x10\x2e\x2f\xd3\xf0\x25\x0c\xa4\xb0\x30\x87\x97\x7d\x1c\x86\x41\x9a\x02\x67\x65\xf9\xb8\xab\x10\x0c\x52\x6d\x94\x85\x8b\x9c\x95\x16\x2f\x7c\xdd\x95\x56\x3d\x17\x60\x63\xd0\xc6\x89\xec\x53\x2a\xc4\x55\x4f\x2a\x81\x5b\x1f\xe4\xb4\xcf\xa5\xb1\x04\x15\x33\x6c\x0d\x2c\x27\x34\x70\xb1\x61\x65\x8d\x17\x31\xc8\x04\x13\x58\xe3\xda\xd5\x8c\x19\xf2\x04\x7f\x1d\x3a\x87\xbc\x56\x4d\x41\x74\x65\xc9\x44\x07\xdd\x00\xec\x93\x24\x5e\xbc\x59\x6e\xd2\x2d\x42\xe7\xe3\xdd\xb7\x6f\x9d\x79\x18\x04\xc1\xf1\xf9\xe3\x8f\x4f\xf7\x9d\xf9\x59\x70\x90\xa6\xcb\x98\xc5\x9b\xb8\x21\xd1\xfe\xd9\x67\x7f\xa3\x6b\x3a\xfc\xdb\xe7\xf3\xb4\x46\x0f\x18\x42\x9a\x5a\x62\x7c\x05\x15\x19\x20\xdd\xa4\xbf\x43\xe7\xd3\xfd\xb7\xfb\xcf\x77\x8f\xf7\x2d\xad\xb7\x01\x0f\x8f\x77\x8f\x5f\x3f\xbe\xb3\xed\xae\x86\xe4\x7f\x62\xe8\xae\x96\xe4\xe0\xe6\x5f\x49\xee\xc3\x57\xf1\xbe\xc4\x37\x4d\xaf\xb4\x7d\x60\x49\x1b\x04\xcb\x36\xe8\x0b\xba\x94\x1b\x54\xe0\x8a\xdc\xda\xd9\x4f\x06\x67\x31\x67\xb9\x30\xf0\xe1\x27\xd5\x93\x22\xf6\x66\x8c\x5e\xc2\x20\xd8\x30\x03\x2b\xdc\xc1\x2d\x74\x3a\x57\x52\x5c\x75\x7a\x9d\x2b\xb7\x79\x13\x06\x01\x15\xd2\x26\x52\xd8\x3f\x57\xb8\xfb\x0b\x6e\xe1\xec\xf9\x2a\x83\x9f\x3f\x21\xbb\x09\x83\x23\x2d\xac\x40\x5a\x90\x6a\xa3\x57\x28\x7c\xc3\xb9\x09\xb1\x03\x5d\x71\x2d\xd0\xcf\x0d\xcf\xf8\x8f\xef\x80\x5b\xe4\x35\xa1\x4d\x1c\x3f\xac\x4e\xe8\x95\x7a\x19\x83\x58\x44\x70\x6c\x24\x47\x92\xd3\x81\xc1\xa1\x23\x5d\x64\xa2\xab\x84\xf4\x03\x19\xa9\x96\xdd\x28\xba\xf9\x95\x22\x73\xe8\xfe\x8f\x53\xd4\xca\x78\xdc\x48\x53\x78\x58\xc9\x0a\x98\xda\x41\x65\xb0\xc7\xf5\xba\x92\x25\x7a\xd6\x9c\x39\x06\x36\x06\x2a\xb4\x45\x60\x06\xe1\xef\xda\x12\xe4\x4c\xf1\xc3\x5b\xd8\xb3\x23\xa4\xfd\xdd\x60\x0b\x21\xba\xa4\xef\x84\x30\x68\xad\xa7\xe6\xab\x9c\x38\x07\x76\xb3\x28\xf9\xcd\x8d\x9c\x6e\x14\x45\xa7\x2f\x76\xac\xf3\x3b\x0d\x90\xa6\xf0\x99\x51\x81\x06\xa4\x22\x34\x8a\x95\xde\x8c\x20\x90\x98\x2c\xed\x99\x3a\x52\x3d\x3c\xc3\x2d\xbc\x3a\x96\x13\x5c\x41\x16\x25\x5f\x15\x5d\x8f\xba\xaf\xd4\xf1\x29\x1f\x6e\x61\xf4\x9a\x50\x83\xf7\x23\xcf\xdf\x03\x7c\x07\xcc\x5d\xbe\x30\xbe\xd1\xba\xa4\xbf\xe0\xd6\x0b\xe0\xac\x60\x76\x89\x2d\x25\xc7\xae\x47\x8c\x5b\xe0\x2b\x18\x45\x51\xec\x59\xf7\x46\x27\x50\xfb\x63\x37\x19\xb4\x75\x49\xa7\xfd\xf4\x54\xa0\x02\x27\x80\x6b\xa1\xb6\x18\x50\x38\x17\x2c\xd0\x19\x80\xd0\x30\x42\x01\x7a\xd3\x5a\xa0\x1d\x93\x1e\xae\x99\x7c\x4e\xc3\x16\xb8\xfd\x12\xb9\x6f\x82\xff\x4c\x06\xcd\xfa\x49\x23\x72\xda\x3a\x69\x82\x20\x4d\x1f\x0e\x5e\xd3\xb5\x1b\x98\xae\x0c\xce\x5f\xc0\x4a\xab\xc3\x20\x90\xb9\x0b\x4e\xa4\xaa\x6a\x4a\x4a\x54\x4b\x2a\xe0\x03\x8c\xbc\xc1\x82\x37\xd2\x1c\x43\x1b\x65\xfa\xb1\x13\xe3\x35\x40\x6f\x04\x51\x18\x04\xfb\x30\x38\x8c\xb7\x83\xfd\x1a\xcb\xed\xc3\x7f\x02\x00\x00\xff\xff\x1b\x64\x99\x95\x2d\x08\x00\x00")

func _4byte_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		__4byte_tracerJs,
		"4byte_tracer.js",
	)
}

func _4byte_tracerJs() (*asset, error) {
	bytes, err := _4byte_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4byte_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _call_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\xdf\x73\xdb\x36\xf2\x7f\x96\xfe\x8a\x4d\x1e\x6a\x69\xa2\x48\x4a\xd2\x6f\x1e\xe4\xaf\x7a\xe3\x71\xd4\x5e\x66\x7c\xe7\x4e\xe2\xb6\x0f\x19\x3f\x40\xe4\x52\x42\x0c\x02\x2c\x00\x4a\xd6\xa5\xfe\xdf\x6f\x76\x01\x52\xa4\x44\x3b\xba\xdc\xdc\x4d\xef\x45\x23\x02\xbb\xcb\xc5\xee\x67\x7f\x81\x93\x09\x24\x42\xa9\x1b\x2b\x12\xb4\x20\x1d\x08\xc8\x4a\xa5\x60\xa9\xcc\x56\x83\xb7\x42\x3b\x91\x78\x69\xf8\x3f\x91\xf8\xb5\xf0\x80\xf7\xf4\xe4\x1d\x08\x9d\x82\xc5\xc2\x58\xfa\xaf\x54\x7f\x32\x01\xbf\x46\x90\xda\xa3\xd5\x42\xb1\x6c\x07\xb9\x48\x11\x96\x3b\x10\x4d\x81\x23\x10\xca\xe8\x15\x6c\xa5\x5f\x83\xd0\x3b\x28\x1d\x66\xa5\x02\xa9\x33\x63\x73\x41\x24\xe3\xfe\x97\x7e\x2f\x6a\xe8\xbc\x48\xee\x48\x41\x92\x9f\x94\xd6\xa2\xf6\x60\x31\x29\xad\x93\x1b\x64\x12\x08\x34\x26\x63\x9a\xc5\xaf\x7f\x03\xbc\xc7\xa4\x0c\x92\x7a\xb5\x90\x19\x7c\xfa\xf2\x70\x3b\xea\xb3\xe8\x14\x5d\x82\x3a\xc5\x94\xcf\x77\xe7\x60\xbb\x46\xbf\x46\x0b\x5b\x3c\xdb\x20\x7c\x2e\x9d\x6f\xd0\x64\xd6\xe4\x20\x34\x98\xd2\x93\x29\x1a\xd6\x91\xda\x1b\x16\x28\xe8\xbf\x46\xcb\x1a\x8d\xfb\xbd\x9a\x79\x06\x99\x50\x0e\xe3\x7b\x9d\xc7\x82\x4e\x23\xf5\xc6\xdc\x91\x64\x63\x01\x37\x68\x77\x60\x8a\xc4\xa4\x18\xec\x4c\xe7\xa8\x8f\x81\x6e\xdc\xef\x11\xdf\x0c\xb2\x52\xf3\x6b\x07\xca\xac\x46\x90\x2e\x87\xf0\xa5\xdf\x23\xb1\x97\xa2\xf0\xa5\x45\xb6\x27\x5a\x6b\xac\x03\x99\xe7\x98\x4a\xe1\x51\xed\xfa\xbd\xde\x46\xd8\xb0\x01\x73\x50\x66\x35\x5e\xa1\x5f\xd0\xe3\x60\x78\xde\xef\xf5\x64\x06\x83\xb0\xfb\x6c\x3e\x87\x52\xa7\x98\x49\x8d\x69\x10\xdf\xf3\x6b\xe9\xc6\x99\x28\x95\xaf\xdf\x4b\x4c\x3d\x8b\xbe\xb4\x9a\xfe\x3e\x04\x2d\x7e\x43\x30\x5a\xed\x20\x11\xa4\xca\xd2\x94\x1e\xdc\xce\x79\xcc\xe3\xe1\xdc\x08\x32\xe1\xc8\x84\x32\x83\x2d\x42\x61\xf1\x65\xb2\x46\xf2\x9d\x4e\x30\x6a\xe9\x76\x8e\x9d\x3a\x07\x7a\xdb\xd8\x14\x63\x6f\xfe\x5e\xe6\x4b\xb4\x83\x21\x7c\x07\xd3\xfb\x6c\x3a\x84\xf9\x9c\xff\x54\xba\x47\x9e\xa8\x2f\x49\x31\x45\x3c\x28\xf3\x7f\xf4\x56\xea\x55\x38\x6b\xd4\xf5\x7d\x06\x02\x34\x6e\x21\x31\x9a\x41\x4d\x5e\x59\xa2\xd4\x2b\x48\x2c\x0a\x8f\xe9\x08\x44\x9a\x82\x37\x01\x79\x35\xce\xda\xaf\x84\xef\xbe\xe3\x77\xcd\xe1\xec\xf2\xc3\xe2\xe2\x66\x71\xd6\x50\x42\xea\xeb\x2c\x8b\x7a\x30\xef\xb8\x40\xbc\x1b\xbc\x1a\x8e\x37\x42\x95\x78\x9d\x05\x8d\x22\xed\x42\xa7\x30\x8f\x3c\x2f\x0e\x79\x5e\xb7\x78\x88\x69\x32\x81\x0b\xe7\x30\x5f\x2a\x3c\x8e\xbd\x18\x9c\x1c\xa7\xce\x1b\x8b\x0c\xb4\xc4\xe4\x85\x42\x02\x50\xf5\xd6\x68\x69\xd6\xb8\xe7\x77\x05\xce\x00\x00\x4c\x31\xe2\x05\x82\x3d\x2f\x78\xf3\x57\xbc\x67\x77\x54\xd6\x22\x00\x5d\xa4\xa9\x45\xe7\x06\xc3\x61\x20\x97\xba\x28\xfd\xac\x45\x9e\x63\x6e\xec\x6e\xec\x94\x4c\x70\xc0\x47\x1b\x85\x93\x56\x3c\x2b\xe1\xde\x6b\xe2\x89\xa0\xfc\x49\xb8\xc1\x7e\xeb\xd2\x38\x3f\xab\xb6\xe8\xa1\xda\x63\x5b\x10\xdb\xd9\xf4\xfe\xec\xd8\x5a\xd3\xe1\xde\xe9\xaf\xde\x0e\x89\xe5\xe1\xbc\x86\x72\x9d\x11\xc6\x45\xe9\xd6\x03\x46\xce\x7e\x77\x1f\xf5\x73\xf0\xb6\xc4\x4e\xa4\x33\x7a\x8e\x91\xe3\x50\x65\x94\x36\xbc\x2d\x13\x46\xd0\x4a\x70\x52\xe1\xa0\x16\x94\x64\x5d\xb9\x64\x9b\x7b\x63\x1e\x05\xd2\xc7\xc5\xd5\x8f\xef\x16\x1f\x6f\x3e\xfc\x72\x79\xd3\x84\x93\xc2\xcc\x93\x52\xed\x33\x28\xd4\x2b\xbf\x66\xfd\x49\x5c\x7b\xf7\x13\xf1\xbc\x7c\x75\x1b\x56\x60\xde\x11\xdd\xbd\xa7\x39\xe0\xd3\x2d\xcb\x7e\x38\x36\x5f\x9b\x34\x18\xf3\x4b\x00\x91\x29\x1e\x9a\x39\xa2\x23\xec\x72\xf4\x6b\x93\x72\x1e\x4c\x44\x48\xa5\x95\x15\x53\xa3\xf1\xe4\xe0\x1b\x54\xd1\x77\x71\x75\x75\x06\x7f\xfc\x01\x8d\xe7\xcb\xeb\x77\x8b\xe6\xda\xbb\xc5\xd5\xe2\xa7\x8b\x9b\xc5\x21\xed\xc7\x9b\x8b\x9b\xf7\x97\xbc\x3a\x8c\x56\x99\x4c\xe0\xe3\x9d\x2c\x38\xa1\x72\x9a\x32\x79\x21\x15\x36\xf4\x75\x23\xf0\x6b\xe3\x10\x28\xd9\x71\xbd\xc8\x84\x4e\xaa\x3c\xee\x2a\xa7\x79\x43\x2e\x33\x55\xac\x1c\xa7\x82\x26\x50\x87\xb5\x1b\xa5\xfb\xd9\x62\x7c\x69\x3a\xf0\xa6\xd2\x6b\x6f\xd0\xe0\x11\xce\x75\x9c\x64\x06\xa7\x1f\x12\xfe\x02\x53\x98\xc1\xab\x98\x49\x9e\x48\x55\xaf\xe1\x05\x89\xff\x86\x84\xf5\xa6\x83\xf3\xcf\x99\xb6\xbc\x61\xe2\x8a\xdc\x9b\xff\x7e\x3a\x33\xa5\xbf\xce\xb2\x19\x1c\x1a\xf1\xfb\x23\x23\xd6\xf4\x57\xa8\x8f\xe9\xff\xef\x88\x7e\x9f\xfa\x08\x55\xa6\x80\x67\x47\x10\x09\x89\xe7\xd9\x41\x1c\x44\xe3\x72\x37\xc3\xd2\x60\xfe\x48\xb2\x7d\xdd\xc6\xf0\x63\xd9\xe2\xdf\x4a\xb6\x9d\x5d\x19\xf5\x5e\xed\xbe\x6b\x04\x16\xbd\x95\xb8\x41\x90\xfe\xcc\xb1\x48\xea\x4f\xcd\x56\xe8\x04\xc7\xf0\x1b\x06\x89\x1a\x91\x93\x4b\xec\x67\xa9\x1d\xe1\x16\x8f\x7a\x52\xa9\xf7\x39\x47\x70\xdb\x69\x11\x72\xb1\x83\x25\x52\xff\x75\xb7\x83\x95\x70\x90\xee\xb4\xc8\x65\xe2\x82\x3c\xee\x65\x2d\xae\x84\x65\xb1\x16\x7f\x2f\xd1\x79\x4c\x19\xc8\x22\xf1\xa5\x50\x6a\x07\x2b\xb9\x41\xcd\xdc\x83\xd7\x6f\xa6\x53\x70\x5e\x16\xa8\xd3\x11\xbc\x7d\x33\x79\xfb\x3d\xd8\x52\xe1\x70\xdc\x6f\xa4\xf1\xfa\xa8\xd1\x1b\xb4\x11\xd1\xf3\x0e\x0b\xbf\x1e\x0c\xe1\x87\x47\xea\xc1\x23\xc9\xbd\x93\x16\x5e\xc2\xab\xdb\x31\xe9\x35\x6f\xe1\x36\x78\x12\x50\x39\x8c\xd2\x26\x13\xb8\xb9\x7e\x77\x3d\xb8\x13\x56\x28\xb1\xc4\xe1\x0c\x6e\x2a\x5b\x6d\x45\x6c\xf8\xc9\x29\x50\x28\x21\x35\x88\x24\x31\xa5\xf6\x64\xf8\xaa\x77\x57\x3b\xca\xef\x67\xbe\x92\xb7\x16\x1b\x24\x3a\x74\xae\x4a\xf7\xec\x35\x52\x47\xe4\xc4\x0d\x52\x3b\x99\x62\xc3\x2b\x94\x1d\x0c\xa7\xe6\x48\xb1\x95\x4a\x55\x02\x73\xe3\xe8\x25\x4b\x84\xad\xa5\x39\xc3\x49\x9d\x10\x1c\x20\x45\xb2\xb6\x03\xa3\x41\x80\x32\x9e\x06\x06\x8e\x71\x10\x76\xe5\xc6\x21\xdf\xd3\x6b\x29\xe7\x68\xb3\x1d\xb7\x81\xdc\x84\x2a\x77\xf4\x07\xed\x80\x06\xbc\x97\xce\x73\x03\x49\x5a\x4a\x07\x01\xc9\x52\xaf\x46\x50\x98\x82\xf3\xf4\x89\xbd\xe4\x87\xc5\xaf\x8b\x0f\x75\xf1\x3f\xdd\x89\x55\x8b\xff\xbc\x9e\x80\xc0\xd2\x78\xe1\x31\x7d\xde\xd1\xb3\x77\x00\x6a\xfe\x08\xa0\x48\xfe\xbe\x36\xfe\xdc\x38\x8e\x12\xce\xef\x1d\xb3\xc2\x30\xbe\x34\x15\x70\xa5\xf2\xee\x20\x77\x1f\x26\x07\x53\x54\x15\x82\x94\xe2\xb4\x43\x89\xbd\xa3\xb3\x8e\x06\xf7\x4d\xe0\x09\x08\x34\x8d\x04\xc0\xfb\x55\x87\x26\x42\xce\x67\x0d\x4d\xe9\xc9\xe9\x54\xa5\xf7\x29\x6e\x25\xdc\x2f\x8e\x7d\x1b\x93\xdc\x52\xae\xde\x6b\x3f\xa8\x36\xdf\x6b\x78\x09\xd5\x03\xa5\x6e\x78\xd9\x8a\x95\x8e\x1c\xd8\x4b\x51\xa1\x47\xd8\x8b\x38\x87\x83\x25\x12\x14\x0e\xcd\xa6\xb1\xe8\x8f\x4b\xf0\x34\x4a\x23\xb3\x3c\xb3\xe8\xc7\xf8\x7b\x29\x94\x1b\x4c\xeb\x96\x20\x9c\xc0\x1b\x2e\x62\xf3\xba\x8c\x55\x75\x8e\x78\x5a\x4d\x46\x14\x18\xd8\xa2\x35\x2a\xb6\x74\x19\x6a\x53\x8a\x4f\x4a\x88\x22\x62\x72\xa8\x3d\x16\xe1\xd7\xd5\x65\xf6\x9a\x04\xf0\xbc\x2e\xfb\x99\x90\xaa\xb4\xf8\xfc\x1c\x3a\x92\x8b\x2b\x6d\x26\x12\xf6\xa5\x43\xe0\x11\xd4\x81\x33\x39\xae\xcd\x36\x28\xd0\x95\xa2\x8e\xc1\x51\xe3\xe0\xa0\x48\x10\x19\x45\x7c\xe9\xc4\x0a\x1b\xe0\xa8\x0d\x5e\x39\xaa\x73\x2e\xfe\x66\xe8\xbc\xa8\x1f\xbf\x82\xa2\xf0\x96\xaf\x42\xe3\x29\x6c\x74\x7a\xf9\xa8\x97\xa9\x88\xb8\xa3\x69\x3c\x54\xaa\x86\x86\xa3\x46\xce\xbf\xe2\xf7\xff\x8c\xe3\x83\xe7\xe3\xef\xa9\x81\x76\x48\x1b\xce\xd8\x26\x0e\x27\xdd\x37\x31\x5f\x47\x41\xbd\xfb\x18\x00\x1e\xeb\x8f\x08\xaa\xfa\x33\x26\x7e\x0f\x57\x6e\x69\xe8\xa9\xb0\xb8\x91\xa6\xa4\x6a\x85\xff\x4b\xf3\x5f\xdd\xdf\x3d\xf4\x7b\x0f\xf1\xce\x8b\xdd\xd7\xbc\xf4\xda\xae\x31\x34\x59\xa1\x35\x6a\xd4\x0a\xc3\x85\x34\x5e\x85\x05\xb7\x8f\xfb\x3d\xfe\xf3\xc4\xed\x57\x0c\x78\x6f\x0a\x2a\xfe\xb1\x16\x29\x8b\x22\xdd\xd5\xe5\x6f\x14\xda\x0e\x58\x0b\x9d\xc6\xd1\x43\xa4\xa9\x24\x79\x0c\x46\x52\x51\xac\x84\xd4\xfd\x4e\x3b\x7e\xb5\xe6\x76\x41\xe3\xa8\x93\x6d\x96\xcd\x38\x32\xd2\x7c\xc7\x1a\xf7\x4f\x28\x8f\x07\xc1\x74\x78\x91\x17\xef\x02\x8d\x76\x65\xce\x7d\x2f\x88\x8d\x90\x4a\xd0\xac\xc5\xfd\x94\x4e\x21\x51\x28\x34\xf7\x4e\xe4\x3d\xb3\x41\xeb\xfa\x27\xa0\xfc\x5b\x40\x7e\x90\x1d\xab\xc7\x68\x8e\xd3\x83\xf6\xd4\x90\x0d\xc7\xff\x51\x09\xef\x23\xbe\x1a\xe6\x0d\xa1\x25\xbd\x83\x42\x50\x1f\xda\x3f\x2d\xa6\xb8\x43\x22\x9a\x1f\x60\xda\xe8\xc2\xff\x2c\x51\x76\x0c\xb1\xab\xba\x1b\x8b\x87\xf7\xc6\x8c\x40\xa1\xe0\x99\x08\xe2\x74\x53\x75\x9f\x4f\x8d\x68\x55\xf8\x86\xfe\xed\x28\x7e\xf9\x16\x6b\x8d\xd5\x7d\x47\x68\xe4\x97\x88\x1a\xa4\x47\x2b\x68\xfa\x21\x74\xc5\x4f\x05\xa4\xa5\x63\x71\xec\x17\x49\x41\x17\x05\xc7\x7b\x7b\x2a\xd0\x52\xaf\xc6\xfd\x5e\x58\x6f\xc4\x7b\xe2\xef\x83\x19\x43\x25\x64\xae\x38\xfd\xd7\xc3\x7f\xe2\xef\xb9\x61\xe4\x01\xf9\xe0\x06\x80\xf6\x68\x29\x4c\xcf\x07\xf3\x3e\x33\xc6\x99\xff\xf0\x5a\x91\xf6\x78\xad\x05\x6e\x26\x5d\x09\x17\xc4\x1c\x84\x83\xbf\x3f\x8e\x86\x8a\x81\x02\x61\xd6\xcd\x40\x5b\x1d\x4c\x07\x77\x10\x44\xcc\x4b\x61\x37\x54\xf5\x59\x73\x37\x2c\xc5\x83\xca\xbc\x61\x1b\x99\xb3\x6d\x1e\xce\xbb\x13\xdc\xb4\xc2\x62\x77\x22\x23\x9b\xd7\x60\x7d\x84\xb5\x39\x55\x1c\x93\x3c\x95\x26\x59\x7a\x95\xd5\x1e\x61\x65\xe9\x8d\xbe\xc3\xdf\x9f\x2e\xb2\x26\x6e\xaa\xd8\xa2\xe9\x12\x12\x73\x4c\xa4\x0b\x96\xad\x04\x04\x44\x07\x5d\x19\xcd\xf2\x1f\x18\x25\x36\x63\xa7\xda\x02\x8b\xe1\xab\x02\x77\xa3\x14\x3a\x66\xc9\x95\xbf\x74\x34\x30\xee\x63\x22\x45\x27\x2d\xa6\x90\x49\x54\x29\x98\x14\x2d\x8f\xa3\x9f\x9d\xd1\xe1\xfb\x11\x5a\x49\x12\xc3\x77\x32\xb8\x59\x53\x58\x92\x50\x2d\x13\xf4\x3b\xc8\x50\xf0\x87\x20\x6f\xa0\x10\xce\x41\x8e\x82\x06\xd0\xac\x54\x6a\x07\xc6\xa6\x48\xc2\xeb\x89\x8c\xc2\xd1\x40\xe9\xd0\x3a\xd8\xae\x4d\x2c\x91\xdc\xa2\x15\xd4\x71\x4a\x3f\x8a\x97\x2e\xd2\x15\x4a\xec\x40\x7a\x2a\xc7\xf1\x50\xcd\x08\xad\xbf\xbe\xf0\x27\x1c\x43\x15\xf7\x38\x44\xab\xa1\xae\x1d\xa3\xbc\x4c\x4f\xed\xe8\x8c\x43\x4d\x3b\x2e\xf7\xd7\x51\xed\x20\xac\x4a\x46\x3b\xd2\x9a\x05\xa8\x1d\x4e\xbc\xc3\x4f\xed\x40\x6a\x34\xcb\xbc\xc1\xe0\xa8\x19\xf8\xe9\x20\xb4\x58\xcb\x18\x5b\xe1\x5b\x63\x4d\xce\x4f\xa3\x08\x18\xf2\xe2\x80\x8c\x73\x87\x3b\xca\xc2\xc1\x46\x8d\x92\x12\x16\x3e\xdd\xe1\xee\xb6\xbb\x82\x44\x38\x36\xe8\xea\x92\x51\x41\x3a\xec\x3d\x11\xc8\xb5\x16\x72\x3e\x3d\x07\xf9\xff\x4d\x86\xaa\xea\x81\x7c\xf1\xa2\x7a\x67\x73\xff\x93\xbc\xad\xa2\xb3\x46\xfc\xc1\xfe\xb0\xa5\x51\x8c\x91\x40\x43\x41\xd1\x7f\xe8\xff\x33\x00\x00\xff\xff\xfe\xdb\x25\x4e\x78\x1e\x00\x00")

func call_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		_call_tracerJs,
		"call_tracer.js",
	)
}

func call_tracerJs() (*asset, error) {
	bytes, err := call_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "call_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _emvdis_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4b\x6f\xdb\x38\x10\xbe\xfb\x57\x0c\x7c\xb2\x12\x57\x15\xf5\x8a\x2d\xaf\x17\x5b\xa4\xc6\x22\x45\x5f\x68\xdc\x93\xe1\x83\xe2\xd0\xb1\x10\x5b\x14\x28\x2a\x6d\x36\xf0\x7f\x5f\x0c\x67\xf4\x70\xec\x4d\xd0\xc5\xee\xa1\x9c\x58\x9c\xf9\xe6\xf1\x7d\x24\xfb\xf6\xec\xac\x37\xdf\x64\x25\x18\x9d\xae\xa4\x06\x2d\x4d\xa5\xf3\x12\xca\x6a\xbd\xce\x56\x32\x37\x90\xe5\x6b\xa5\x77\xa9\xc9\x54\x0e\x6b\xad\x76\x90\x92\x2f\x18\x05\xbd\x42\x6a\xdc\x05\xf9\xb0\xbb\xcd\xca\x37\xa5\x79\xdc\x4a\xb8\xcd\xca\xb4\x2c\xe5\xee\x66\xfb\xd8\xeb\xfd\x91\x56\x66\xa3\x34\x7c\xce\x56\xf7\xf0\x41\x6d\xf2\x52\xe5\x43\xf8\x94\x6a\x93\xe5\x70\xfd\x43\xe6\xb7\xb2\x77\x76\xf6\xb6\xd7\x7b\xea\x01\xf4\x4b\x93\xae\xee\xfb\x09\x2c\x9e\xfa\xaa\x28\xf1\x8f\xe5\x7e\x39\xc4\x9d\xbc\xa8\xca\x8d\xc4\x4f\x4f\x5e\x02\xde\x10\x44\x02\x62\x08\xbe\x5d\x03\xbb\x86\x76\x8d\xec\x1a\xdb\xf5\xc2\xae\x23\xbb\x8e\xed\x2a\x3c\x32\x14\x2d\xc8\x4d\x90\x9f\x20\x47\x41\x9e\x3e\x79\xfa\x9c\x87\x12\xf9\x94\xc9\xa7\x54\x3e\xe5\xf2\x09\x25\x20\x97\x90\x50\x42\x42\x89\x08\x25\x22\x94\x88\x5c\x22\x42\x89\xb8\xe0\xc8\xf6\x13\x11\x4a\x74\x41\xbf\x08\x25\x22\x94\x98\x5a\x8e\x29\x20\xe6\x16\x29\x20\xa6\xe2\x63\x0a\x88\x29\x60\x44\x01\x23\x4a\x3b\xf2\xe9\x57\x40\x86\x50\x46\x94\x76\x14\x93\xa1\xb4\x23\x42\x19\x11\xca\x98\x8a\x1f\x0b\xbb\x37\xa6\x7c\x63\xca\x37\xe6\xa9\xd6\x63\xe5\xb9\x7a\x3c\x58\xcf\x67\x1b\xb0\x0d\xd9\x46\x6c\x79\xf2\x1e\x8f\xde\xe3\xd9\x7b\x8c\xd7\xf0\xc4\x78\x82\xf1\x04\xe3\x09\xc6\x13\x8c\x57\x33\x59\x53\x59\x73\xc9\x64\x0a\x66\x53\x30\x9d\x82\xf9\x14\x4c\xa8\x60\x46\x05\x53\x2a\x98\x53\xe1\x33\x9e\x3f\x4a\xc0\x47\x3b\x4e\x20\x18\x82\x08\xbc\x04\x42\xb4\x22\x81\x08\xad\x9f\x40\x8c\x36\x48\xe0\x02\x6d\x98\xc0\x08\x6d\x94\xc0\x18\x2d\xe2\xa1\x6a\x03\x04\x44\xc4\x00\x2b\x44\xc8\x00\x4b\x44\xcc\x10\x6b\x44\xd0\x10\x8b\x44\xd4\x10\xab\x44\xd8\x10\xcb\x44\xdc\x30\xa4\x3a\xc2\x88\xea\x08\x63\xaa\x23\xbc\xa0\x3a\x50\x7d\x36\x60\x4c\x75\xa0\xfe\xb0\x0e\x14\x20\xd6\x61\x15\x88\x75\x58\x0d\x62\x1d\x56\x85\x08\x89\x3a\xb4\x75\x58\x25\x22\x28\x6a\xd1\xd6\x61\xd5\x88\xb0\x56\x8f\x88\xcb\x8a\x14\xb1\x60\xeb\xb3\x0d\xd8\x86\xd6\xfa\x21\x9f\xa2\x90\x8f\x51\xc8\xe7\x28\x0c\x78\x9f\xfd\xec\x21\xd8\xdb\x93\xae\x65\x59\x6d\x4d\x3f\x81\x75\x95\xaf\xf0\xda\x19\x38\xf0\xc4\x17\x13\x98\x4d\x56\xba\xf6\x96\x58\x78\x4b\x57\x15\xe5\x04\x28\xaa\x34\xb2\xe8\xc6\x6c\xd5\xdd\x10\x6e\x6f\x1c\xc0\x7b\x05\xe0\x21\xd5\xb0\xd6\xe9\x4e\xc2\xb4\x8b\xd1\xfe\xe9\x6e\x65\x7e\x67\x36\xf0\x06\xc4\x72\x62\x43\xb2\x35\x82\xb8\x52\xeb\x1a\x04\x08\x62\xd1\x97\x5a\x2b\xdd\x5f\xc2\x14\xd8\xc3\x35\xea\xda\xe8\x2c\xbf\x1b\x38\x14\xbc\x07\xb9\x2d\x65\x8d\x71\x2b\x0b\xb3\x81\x69\x37\x35\xe7\x6b\xa1\x55\x81\x37\x2d\x4c\x9b\x0f\x00\x7d\x85\x2d\x21\x80\x2a\x5c\xa3\x3e\x57\xbb\x1b\xa9\x07\xce\xb0\x75\xb0\xc0\x7d\x20\x27\xfb\xa3\xb3\xd9\x0c\x72\xb1\xac\xbf\xee\x27\xfc\x47\xb6\x1e\xd8\x5e\x70\x84\x75\xeb\xbf\x83\xe7\x74\xb2\xe3\xc8\x0a\x2d\x1f\x54\x01\x53\x68\x9c\x17\x47\x61\xed\xc4\xec\x84\x94\x1e\x60\x64\x06\x53\xf0\x26\x90\xc1\x6f\xd4\x34\xdf\xdf\x0b\x42\x74\x55\xb1\x9c\x40\x76\x7e\xee\x34\x81\xc0\xc9\x5c\x2a\xdb\x45\x7f\x3b\x3c\x1a\x57\x21\xe5\xfd\x20\x73\xdc\xb9\xfc\x69\x06\x22\x76\x9c\x3a\xe5\x9e\x6d\xf9\x23\x33\x2b\x8a\xb0\xd3\xaa\xf9\x68\x5b\x5a\xa5\xa5\x84\xfe\xe5\xbb\x8f\x1f\xfb\xc9\xd1\xa7\xcb\x2f\xef\x67\xcd\x67\x6a\x3e\xcb\x4b\x93\x6a\xc3\x2c\x77\xca\x08\x1c\xf7\x2a\x37\x71\x38\x70\x26\xcf\x03\xb2\xbf\xe4\xb1\x7f\x78\xc2\x9f\xe8\x5e\xf4\xef\xd2\xb2\x11\x52\x27\xc4\x7b\x21\xc4\xa8\x53\x11\xa2\x9d\xcd\x71\xc8\x43\xba\xad\xe4\xa9\x28\xdf\x71\x0f\x85\xdb\x8d\xca\xf2\xa2\x32\x4d\xd4\x4e\xee\x94\x7e\x74\xcb\x6d\xb6\x92\x03\x9e\xcd\xb0\x19\xd2\x39\x77\x7f\x02\xa6\x3d\x2d\x79\xb5\xdd\x1e\xef\xd3\xd1\x7e\xc1\x01\xff\x2b\x80\xbb\x8b\x8e\xce\x3a\x27\xc9\x2a\x85\x7c\x3b\xd9\x6f\xb4\x4c\xef\x27\x07\x3c\xbf\x9f\x7d\x9c\xfd\xf9\x6e\x3e\x3b\x21\x81\xeb\xf9\xbb\xf9\xd5\xe5\xc1\xc6\x2b\x22\xf0\x7f\x51\x04\xa7\x44\xd3\x36\x68\xfb\x83\x23\xf1\xfe\xb3\x5e\xfe\x85\x60\x7e\x49\x31\x2d\xf7\xff\x15\xf9\xaf\xb3\xff\x3f\xd3\xff\x6d\x36\xff\xfe\xed\xf3\x33\x7e\x55\x75\x82\xdb\x53\xc3\x64\xe7\xd3\xe4\x8a\x13\x01\x74\x51\xf2\xbb\x75\xe2\x08\xa9\xca\x0c\x6d\xfa\xf3\x1a\xf7\xe5\xf2\xaf\xe7\x5f\xbe\x3e\x57\xed\xf7\xab\xcb\xab\x83\x7b\xeb\xb5\xa4\xde\x10\xbc\x97\xd3\x7c\xf8\xfe\xe9\xeb\xfb\xd9\xf5\xbc\x03\x5a\x13\x51\xac\x9a\xcb\xa0\x58\x3d\xbf\x7f\xf9\xa9\x53\x85\x9b\x95\x5f\x91\x12\xa7\xfb\x98\xd4\x10\x5b\x99\x37\x18\x07\x8f\x1a\xbc\x01\xef\x67\x24\x9f\xa3\xb6\xaf\xcd\x31\xcd\xfc\xc6\xd6\x49\x5a\x49\x1c\xbc\xf1\x6d\xe3\xcd\x1b\x59\xc7\xf7\xf0\xdf\xbe\xf7\x77\x00\x00\x00\xff\xff\xbc\x1e\x80\x65\x04\x0d\x00\x00")

func emvdis_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		_emvdis_tracerJs,
		"emvdis_tracer.js",
	)
}

func emvdis_tracerJs() (*asset, error) {
	bytes, err := emvdis_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "emvdis_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _genesis_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xfb\x8f\xda\xc6\x13\xff\xdd\x7f\xc5\x7c\x4f\xfa\xca\x70\xf2\x81\x2f\xc9\x5d\x22\x4e\x57\x95\x7b\xb4\x17\x95\x1c\x55\xa1\xea\xe3\x84\xaa\xb5\x3d\xd8\x2b\xcc\x2e\xda\x1d\xf3\x68\xc4\xff\x5e\xed\x2e\x10\xdb\x98\x34\x91\x1a\xff\x02\x78\x3e\x33\xf3\x99\xd7\xee\xd0\x3d\x3f\xf7\xc6\x19\xd7\x40\x8a\xc5\xa8\x2e\x74\xac\xf8\x82\x40\x16\xb4\x28\x48\x83\x2e\xa6\x53\x1e\x73\x14\x04\x5c\x4c\xa5\x9a\x33\xe2\x52\x00\x49\x88\x15\x32\x42\x60\x90\xcb\x98\xe5\x1e\xae\x31\x2e\xac\x4c\x4e\x81\x32\x34\xf6\x84\x66\xb1\x79\xd5\xf1\xbc\xef\x59\x41\x99\x54\xf0\x81\x29\xe2\x02\x9e\x64\xae\x09\x46\x2b\x14\x09\x7a\xe7\xe7\x5d\xef\xa3\x07\x00\xd0\xed\x8e\x33\x84\x14\x05\x6a\xc3\x28\x63\x04\x2b\xf4\x15\x42\x54\xf0\x3c\xe1\x22\xb5\xa8\x85\x42\x4d\xc6\x77\x0f\x44\x91\xe7\x81\x67\xdf\xe6\x52\xce\x8a\x45\x3f\x8e\x65\x21\xa8\x07\xd3\x42\x58\xdf\x2d\x16\xc7\x41\x12\xb5\x9d\x03\xf3\x2c\x99\x02\x96\x24\x0a\x6e\x81\xe4\x13\xae\x0d\xa2\x7d\x73\x10\xf3\x69\x8b\x32\xae\x3b\x7b\x2f\x2f\x06\x3b\x69\x83\x42\x2a\x94\xb8\xf1\x0e\xc0\x06\x14\xdc\x02\x7c\x72\x64\x1e\x3f\x62\x39\x13\x31\xfa\x3d\x48\xa2\x4e\x8a\x74\xe7\x7e\x5b\xa7\x41\x15\x2a\xa4\x01\xc2\x1e\xf9\x2c\x4f\xe0\x62\x99\x58\x98\x63\xef\xc0\xf7\x32\x71\xd8\x3a\x58\x93\x54\x2c\xb5\xf8\x8f\xdb\x83\x68\xeb\xe2\xdd\x56\x72\x37\x72\xd0\x72\xee\x92\x44\x05\xc0\x45\x82\xeb\x00\x2a\x49\xec\x76\x9f\x25\xc8\x25\xaa\x95\xe2\x84\xa5\xe4\x41\xeb\x7f\x0d\x89\x79\x39\xf0\x98\xbc\x58\x73\x93\x76\x39\x4f\x5f\xa4\x01\xb7\xbb\xc4\x8c\x0c\xca\x71\x9b\xe1\xa6\x54\xba\x6d\x25\x28\x85\xba\xc8\xcb\x9d\x10\xd3\xda\x74\x02\x94\xa3\x80\x3e\x59\xef\xb0\x90\x5c\x50\x00\x2b\x04\x81\x98\x98\x06\x4f\x30\x29\x62\xb2\xbd\xec\x2f\x59\x5e\xa0\x0f\x53\x25\xe7\xe6\x45\xd9\x80\x2c\x08\x55\xb9\xdb\x03\x60\x22\x81\xb9\x5c\x22\x70\x82\x88\xc5\x33\x63\xcd\x98\x91\x8a\xa7\x5c\x54\x1a\xd1\x5a\x34\x8d\xd8\x4f\x12\x85\x5a\x1b\x92\x1d\xf3\xb2\x14\x96\xc1\x91\x3c\x42\x91\xac\x61\x96\x70\x0b\x46\x60\xc9\xde\x54\xfb\xb4\x32\x1d\x2d\xe3\xc0\x16\xb4\xd4\xce\x26\x17\x49\xe2\xf8\x5a\x0b\xc6\xa7\x36\x03\xaa\xa0\x15\x19\xe2\xa4\xdb\x47\xdc\xff\x8a\x58\x6e\x98\x55\x0a\xe8\x1a\xd3\x46\x31\x79\x39\x4c\xc0\xa4\x1e\xd1\x4e\xb7\x51\x99\xe4\x09\x55\xa7\xd6\x19\x15\x51\xcb\x7d\x0d\x60\x59\xca\xc3\x9e\x53\xa7\x9f\x24\xad\xfd\x0f\x07\x29\x47\x3a\xc2\x5d\x65\xec\xcc\x41\xab\xe0\x82\xae\xdf\xb4\x4f\xcc\x76\x35\x1e\x37\xa6\x93\x8b\x8b\x8a\xc9\x67\x49\x18\xd8\x2a\xbb\x9c\x5d\x30\x97\x6b\xd0\xc4\xf3\x1c\xe6\x6c\x03\x42\x12\x44\x08\x97\x61\xf8\x7f\x88\xa5\x52\x18\x53\x50\x26\xa5\xb9\xe1\xb2\x42\xc8\xd8\x12\x85\x4f\xb0\x33\x81\x09\x4c\xa5\xb2\xb6\x53\xa6\x3b\x25\x95\x27\x39\x47\x4d\xc8\x92\xc3\xbb\x58\x8a\x29\x4f\xe1\xb6\x76\x0c\x9d\x21\x5f\x5c\x5e\x85\x77\xb9\x8c\x67\x67\x3d\x78\x15\x86\x61\xd0\x00\x78\xf7\x6f\x80\xab\xcf\x00\xb2\x3d\x99\x3d\xa6\x0e\x48\x98\xfc\x41\xaa\xd9\x29\x71\xb4\xf9\x9b\x09\xe2\xc5\xdc\x01\xa0\xee\x64\x5b\x69\x9f\x28\x17\xc5\xdc\x35\xfc\x8b\x1f\x19\x0d\x51\xcc\x23\x54\xfe\xa4\x72\x18\x39\xd8\x77\xf0\xea\xcd\xf5\xeb\x30\x0c\xdb\xd5\xb4\x74\xbb\x30\x66\x22\x45\xc5\x05\xf6\x20\xe1\x9a\x45\x39\x56\x10\x2e\x9f\x2f\x95\xfc\x99\x93\x28\xac\x9f\x3a\x47\x0e\xaf\xdf\x5e\x35\x3a\x1c\x2d\x0a\xc5\x65\xa1\x4f\xbb\xb9\x6a\x74\x73\x04\x7b\xf7\x15\x6c\xde\xbc\x7e\x1b\x36\xb1\xb9\xdb\xe7\xbc\xd1\x4d\xad\x22\x27\x3d\xed\x2f\xec\xdb\xaa\x79\x3b\x28\xbd\x5d\x75\xc3\x75\x58\x7b\xce\xaa\xf5\x4f\xb8\xd9\x34\x8a\x9c\x36\x3d\x8b\x7e\xd5\x80\x99\xf3\x75\xc6\x74\xd6\x3b\x61\xf1\x6b\x9f\x9a\xf5\x58\x72\x11\x31\x6d\x29\x7f\x85\xf5\x9a\x15\xe2\x66\x0c\xd8\x7c\xd1\xdb\x59\xa9\xc9\x5d\x9b\x96\xd3\x52\x03\x2c\x98\x42\x41\x4f\x36\xce\xff\x3e\x48\x5c\x93\x62\x0f\x8c\x58\xef\x5b\xa4\x30\x65\x7a\xc0\xe7\x9c\x6c\x78\xee\xd8\xb4\x03\x9a\x32\xed\x4f\x3a\x77\x1b\x42\xdd\xaa\xaf\x28\x2c\xcf\x65\xdc\xab\x1e\xb9\xf5\xc2\x98\x7e\xec\xed\x3e\x83\xf2\x16\x73\xf8\xee\xd6\xb3\x7d\x2b\x56\xd7\x1b\x4d\xb8\x28\xed\x01\xb9\x4c\xed\xe5\x07\x1f\xbd\xcf\x2c\x2e\xc7\x93\x6b\xae\x47\x73\x08\xc7\x85\x32\x15\xda\x1f\xcf\xd5\xea\x97\x4d\x98\x63\x78\x7b\x53\x37\xb3\xdb\x00\x61\x65\xee\x85\x85\x24\x14\xc4\x59\x9e\x6f\xcc\xdd\xb0\x52\x52\xa4\x90\xa1\xaa\x65\xe0\x70\x3b\xd8\x5d\xc5\x6a\x72\x11\xe7\x45\x82\x96\x91\xbb\xaf\x35\xd6\xc8\x74\xbb\xc0\x72\x63\x70\xc5\x29\xb3\xc0\x39\x6a\xcd\x52\xac\xa3\x7e\x43\x98\xf2\xb5\x5b\xb8\xb9\x00\xdf\xad\x4e\xad\xb6\x7f\x1c\x5b\x75\x91\xc8\x65\xda\xd9\xa5\x21\xb0\xeb\xc4\xa1\x34\x9f\x52\xab\x57\x9c\xe2\xcc\x42\xe5\xa2\x43\x72\x44\x8a\x8b\xb4\xd5\xae\x25\x38\x66\x1a\xe1\xec\xf1\xf7\xf1\xfd\xf0\xe1\xf1\x7e\xf8\xf3\x1f\x67\xbd\xd3\xf2\xd1\xfb\x3f\x1f\x1b\xe5\x77\xfd\x41\xff\xf9\xbe\x2e\x83\xea\xda\x6f\xa8\x2c\x10\x67\xad\xb0\xdd\x19\xe3\x9a\x5a\x97\xd7\xed\x9b\x23\x85\x86\x70\xed\xce\x59\x89\x73\xff\x44\x0a\xd9\xec\xa6\x81\xd0\x7d\x7f\x30\x68\x64\x6a\x04\x26\x94\x46\xe1\xc3\xe3\xe0\xf1\xc7\xfe\xf8\xf1\xa4\xf6\x68\xdc\x1f\xbf\xbf\x6f\x10\x37\x84\xaa\x89\xc5\x33\x17\xf0\xe5\x37\x0f\xd8\x1f\x8d\xc6\xc3\x5f\x1e\xfd\x06\xd2\xfe\x68\x30\xec\x3f\xf8\xcd\x7c\x67\xb8\x39\xa6\xfb\xa5\xf5\xd9\xfd\x73\x29\xb7\xe3\x41\x31\x30\xa6\x8f\x63\xa8\xf1\xdf\x5f\x66\x5b\x6f\xeb\x75\xcf\xcf\x61\x3c\x7c\x18\x7a\xde\x07\xae\x35\x17\x29\xdc\x0d\x86\xf7\x3f\x3d\xff\xfa\xc1\x33\xcf\x48\xce\x31\x93\x2b\x48\x91\xc0\x77\x2b\xbd\x7f\xd8\xf6\xb8\x30\x73\xa6\xb0\x63\xff\xd5\xfe\x13\x00\x00\xff\xff\x29\xa1\x73\xf0\x5d\x0f\x00\x00")

func genesis_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		_genesis_tracerJs,
		"genesis_tracer.js",
	)
}

func genesis_tracerJs() (*asset, error) {
	bytes, err := genesis_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "genesis_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _noop_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcf\xc1\x6a\x23\x31\x10\x04\xd0\xb3\xf5\x15\x75\xdc\x05\xe3\xb9\xef\x27\x2c\xec\x69\x43\xee\x2d\x4d\x8d\x47\x8e\xa2\x9e\xb4\x7a\x26\x0e\xc6\xff\x1e\x46\x26\xe0\x9b\x68\xa8\x57\xa5\x61\x40\x55\x5d\x5e\x4c\x12\x0d\xb9\xe1\xb2\x36\x87\xcf\x44\x14\x63\xd4\x4a\x44\xcd\x85\xb6\x14\x71\x22\xe9\x48\x18\x3f\xd6\x6c\x1c\x31\x99\xbe\x43\xf0\x57\x36\xf9\x9f\x2c\x2f\x1e\x86\x01\x1a\x2f\x4c\x0e\x57\x44\x62\x6d\x12\x0b\x21\x0d\x02\x37\xa9\x4d\x92\x67\xad\xfb\x3b\xd1\x4e\xe1\x16\x0e\xc3\x80\xe6\x5c\xf6\xee\x5c\x37\x7d\xdb\x5d\x35\x70\xa3\x7d\x41\x97\xde\xe8\xb3\x3c\x46\xbd\xfe\x03\xaf\x4c\xab\xb3\x9d\xc2\x61\xcf\xfd\xc1\xb4\xd6\x8e\xfe\x2a\x7a\x3e\x62\x8c\xbf\x71\xc3\xfd\x18\xba\x6c\x6c\x6b\xf1\x67\xfb\x73\x66\x85\x94\xd2\xb9\x07\xdf\x30\xcb\x46\x44\xb2\x22\x3b\x4d\x9c\x23\x74\xa3\x41\xea\x08\xa3\xaf\x56\x5b\xe7\xf6\xcc\x94\xab\x94\x1f\x58\xa7\x7e\xdb\xbf\x93\xeb\xf9\x14\x0e\x8f\xfb\xd3\xa8\xe4\xd7\x3e\x28\xdc\xc3\x77\x00\x00\x00\xff\xff\x8f\x9c\x5f\x55\x6c\x01\x00\x00")

func noop_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		_noop_tracerJs,
		"noop_tracer.js",
	)
}

func noop_tracerJs() (*asset, error) {
	bytes, err := noop_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "noop_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opcount_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xb1\x6e\xeb\x30\x0c\x45\x67\xeb\x2b\xee\xf8\x1e\x12\xd8\x9d\xb3\x77\xcc\x56\x64\x97\x6d\x3a\x56\xe3\x50\x01\x49\xb9\x09\x82\xfc\x7b\x21\xb9\x2e\x8c\x8e\x22\xc8\x73\xee\x55\xd3\x20\xde\xba\x98\xd8\x3e\xc4\x77\x24\x08\x0a\x0f\xf5\xd7\xdb\x44\xb0\x65\x64\xa3\x37\x7c\x26\x35\x94\x45\x85\x8d\x04\x4e\xd7\x96\x04\x71\x40\x60\x35\x49\x9d\x85\xc8\xea\x9a\x06\x74\xa7\x2e\x19\xf5\x68\x1f\x65\xf3\xfd\x74\x44\x4b\x43\x14\x2a\x4f\x13\xcf\xea\xcb\x3a\x8c\xe4\x1a\xd8\x1b\xf5\xb5\x7b\xba\xaa\x69\x16\x43\x11\x5f\xfe\x7a\x32\x67\xeb\xfa\x15\xd5\xae\x2a\x67\x07\xbc\xed\x5d\xa1\xa8\xd1\x2d\x37\x09\x3c\xc7\x0b\xf5\x18\xa2\x80\x66\x92\x47\x29\xdb\xd3\x52\x29\xe3\x4f\xc7\x15\xa3\xb5\xab\xf2\xdd\x01\x43\xe2\x62\xf8\x37\xc5\xf3\x1e\x7d\xfb\x1f\x4f\xd8\x18\xb4\x2e\x96\xdd\x0e\xaf\x1f\x8d\x90\xa6\xc9\xb6\xa2\xaf\x91\x18\x7e\x9a\x0a\x7b\x71\x29\x46\x3f\x13\x5a\x22\x46\x30\x92\xdc\x16\x71\x26\x81\xe7\x1e\x42\x96\x84\xb5\xe0\xf2\xcd\x10\xd8\x4f\x2b\x38\x0e\xeb\x8f\x75\x81\xcf\xb5\xab\x96\xf9\x26\x61\x67\xf7\x9c\x6e\xa1\x6c\x42\xe2\xe5\x5e\xee\x3b\x00\x00\xff\xff\x6e\xdf\xbf\xab\xdc\x01\x00\x00")

func opcount_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		_opcount_tracerJs,
		"opcount_tracer.js",
	)
}

func opcount_tracerJs() (*asset, error) {
	bytes, err := opcount_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "opcount_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"4byte_tracer.js":   _4byte_tracerJs,
	"call_tracer.js":    call_tracerJs,
	"emvdis_tracer.js":  emvdis_tracerJs,
	"genesis_tracer.js": genesis_tracerJs,
	"noop_tracer.js":    noop_tracerJs,
	"opcount_tracer.js": opcount_tracerJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"4byte_tracer.js":   {_4byte_tracerJs, map[string]*bintree{}},
	"call_tracer.js":    {call_tracerJs, map[string]*bintree{}},
	"emvdis_tracer.js":  {emvdis_tracerJs, map[string]*bintree{}},
	"genesis_tracer.js": {genesis_tracerJs, map[string]*bintree{}},
	"noop_tracer.js":    {noop_tracerJs, map[string]*bintree{}},
	"opcount_tracer.js": {opcount_tracerJs, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
