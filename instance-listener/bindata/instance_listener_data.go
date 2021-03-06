// Code generated by go-bindata.
// sources:
// instance-listener/Godeps/Godeps.json
// instance-listener/Godeps/Readme
// instance-listener/main.go
// DO NOT EDIT!

package bindata

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

var _instanceListenerGodepsGodepsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcc\xbd\x0a\xc2\x30\x14\x40\xe1\x39\x79\x0a\xb9\xb3\xb9\x41\x84\x0e\xce\x82\xb8\x75\x72\x91\x0e\x31\xb9\xb4\xa1\xe6\x87\x24\xcd\x22\xbe\xbb\x6d\x70\xe8\xfa\x1d\x38\x1f\xce\xe0\xee\x62\x48\xa5\x57\x65\x82\xcb\x01\x46\x5b\xa6\xe5\x85\x3a\x38\x49\x4e\x0b\x65\xaa\xf2\x9a\x8c\x30\x54\xe5\xe2\xed\x2c\xad\xcf\x65\x23\xf1\xb6\xb9\x90\xa7\x04\xc7\x75\x72\x0b\x0f\x4a\xd9\x06\xdf\x1e\xe1\x84\xdd\x9f\x0d\xc5\x5d\xa9\xdd\xb9\x79\xaf\xf4\xac\x46\xca\xab\x3d\x39\x63\x80\x12\x11\x81\xb3\x61\x8b\x57\x8a\x2d\x0c\xfc\xcb\x7f\x01\x00\x00\xff\xff\x2a\x8c\x59\x34\xa1\x00\x00\x00")

func instanceListenerGodepsGodepsJsonBytes() ([]byte, error) {
	return bindataRead(
		_instanceListenerGodepsGodepsJson,
		"instance-listener/Godeps/Godeps.json",
	)
}

func instanceListenerGodepsGodepsJson() (*asset, error) {
	bytes, err := instanceListenerGodepsGodepsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "instance-listener/Godeps/Godeps.json", size: 161, mode: os.FileMode(420), modTime: time.Unix(1463577211, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _instanceListenerGodepsReadme = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x1c\x8d\xb1\x0d\x03\x21\x10\x04\xf3\xaf\xe2\x2a\x80\xdc\x55\x58\xb2\x1b\xe0\xe1\x0c\x48\xc0\xbe\x8e\xfd\x80\xee\x8d\x9d\xad\x46\x9a\x9d\x77\xa9\x53\x52\x35\x8d\x84\x2d\xa1\xa9\xca\x26\x59\x87\x5a\xa0\x26\x09\x37\xd1\x03\x6b\x0c\xad\x2d\x39\x97\x64\x24\xbd\xdc\x71\x3c\x9b\x86\xa9\x92\x20\x03\x14\x4d\x95\x1b\xbe\xb6\x5e\xc8\x6b\x3e\xbc\xcf\x95\xe5\x3e\x5d\x44\xf7\x04\xda\xf4\x7f\x53\x3e\x30\xe9\xb0\x9d\x19\x7b\xfe\xae\x31\xdc\xf1\x0d\x00\x00\xff\xff\x52\x08\x9f\x0e\x88\x00\x00\x00")

func instanceListenerGodepsReadmeBytes() ([]byte, error) {
	return bindataRead(
		_instanceListenerGodepsReadme,
		"instance-listener/Godeps/Readme",
	)
}

func instanceListenerGodepsReadme() (*asset, error) {
	bytes, err := instanceListenerGodepsReadmeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "instance-listener/Godeps/Readme", size: 136, mode: os.FileMode(420), modTime: time.Unix(1463577211, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _instanceListenerMainGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x58\x5b\x73\xdb\x36\x16\x7e\x96\x7e\x05\xca\x99\x66\xa8\x58\xa1\xdc\xd9\x4c\x77\x57\xa9\x1f\xdc\xe6\x52\xcd\xc4\x89\x56\x4e\xa6\x0f\x4e\xd6\x85\x49\xc8\xe6\x86\x22\x18\x02\x94\xe3\xc9\xfa\xbf\xef\x77\x0e\xc0\x9b\x28\xb9\x9d\xbd\x4d\xfd\x60\x52\x38\x17\x9c\x3b\x3e\xb0\x90\xf1\x27\x79\xad\xc4\x46\xa6\xf9\x78\x3c\x9b\xbd\x7b\xfb\xfc\xed\x1c\xbf\x3e\x29\x61\xaa\x52\x09\xab\x85\xcc\x6e\xe5\x9d\x11\x57\x69\x4e\xcc\xc2\xde\xa4\x46\xac\xd3\x8c\x69\x58\x4c\xa4\x95\xb3\x34\x37\x56\xe6\xb1\xba\xcc\x52\x63\x55\xae\xca\x4b\x5a\x16\x3a\x17\xa5\x8a\xf5\xa6\x00\x3b\x94\xcb\x24\x11\xd2\x29\xb7\xb2\xbc\x56\x76\x3c\x4e\x37\x85\x2e\xad\x08\xc7\xa3\x40\xe5\xb1\x4e\xd2\xfc\x7a\xf6\x0f\xa3\xf3\x80\x16\xca\x52\x97\x86\xde\xd6\x1b\x4b\x8f\x54\xcf\x52\x5d\xd9\x34\xa3\x1f\xb9\xb2\xfe\x31\xbb\xb1\xb6\xa0\x77\x63\x4b\xc8\xb3\x84\xb9\xcb\x63\x7a\x66\xfa\x9a\x1e\x36\xdd\x28\x56\x94\x49\xfe\xad\xc1\x34\x19\x8f\x63\x0d\xbb\x05\x4c\xb7\x8a\x3d\x3a\x11\xc1\x8c\xfd\x69\x96\x22\x67\xcc\xd8\xde\x15\xca\x31\xe2\x7f\x59\xc5\x56\x7c\x1d\x8f\xce\x64\xbc\x28\xce\x64\x21\xe0\x53\x71\xe1\x76\xff\xe8\x1e\xa2\xf3\xf7\x2b\xe9\x98\x07\x8b\xc2\x04\xbf\xb2\xd0\x8b\x7c\x4b\x52\x1d\xa1\xa1\xbc\x17\x02\x2b\x49\xdd\x8f\xc7\xeb\x2a\x8f\x39\x4d\xe1\x84\xf6\x46\xfc\x8c\x98\x9f\x08\x6d\xa2\x53\xbc\x8e\x47\x6b\x5d\x8a\x74\x2a\xb0\x4e\xcb\xa5\xcc\x91\x55\x66\x02\xf3\x08\x61\x88\x96\xd0\x6b\xd7\x61\x40\x1c\xdf\x6e\xe7\xe2\x5b\x13\x4c\xbd\xc4\x64\x3c\xba\x1f\x8f\xc8\xf3\x65\x09\xb7\xbf\x90\x06\x0a\x55\x74\xce\xc6\x84\x41\xc1\xcb\xe0\x0f\xaa\x3c\xfd\x74\x49\x2f\x6e\x49\xd0\xb6\x9c\x6b\xa3\x72\x2b\xb6\xa9\x14\x55\x52\x88\x90\xad\x49\xb0\x94\xae\xd3\x58\xda\x14\x95\x50\x54\x65\xa1\x8d\x32\x01\x76\x63\xe5\x4b\x59\x1a\x15\x4e\xf6\x9b\xce\x1c\xe4\x19\xbb\xbb\xe3\x02\x11\x0f\xba\x91\xae\xc5\xe3\x8e\x2b\x27\x27\xb5\xd1\x83\x48\xbc\x58\xad\xde\xae\x50\xed\x15\x8a\xa0\x28\xf5\x16\x06\x8b\x27\xde\x55\xe8\x1a\x95\xca\x56\x65\x7e\x48\xe9\x41\x7d\x5e\x85\x88\x65\x9e\x6b\x2b\xae\x94\xf8\x10\x7c\x08\x06\x1a\xa9\x72\x5e\x6b\x74\x14\x7c\xa6\x72\x8d\x56\xbf\x9c\x55\x56\x7d\xf9\x0a\xa2\xe2\x0a\x39\x44\x35\x72\xab\x7a\xb4\x86\xb2\x95\xa5\x30\xae\x4e\xc1\x16\x35\x05\x7a\xc2\x5d\x17\x0e\xca\x6c\xe2\xb9\x7c\x45\x0e\xd9\xf6\x48\xb8\x42\x99\x0a\xb4\x27\x19\xe0\x3a\x32\x5a\x29\x99\xbc\x44\xc3\x84\x4d\xeb\x4c\x38\x6a\xc4\xf5\xcd\x89\xc8\xd3\x6c\x10\xae\x58\x57\x59\x22\x28\x44\x25\x84\xdb\x36\x9c\xc2\x8a\x3b\x44\x8d\x27\x0d\x0f\x9b\x12\x09\xba\xd2\xda\xce\x45\x70\x04\x85\xd1\x0b\x9a\x0c\xe1\x84\xf2\x2d\x54\x66\x14\xab\xf6\x9b\xc1\x24\x6a\x9d\xe8\x7d\xbe\x41\x7d\xdd\xc8\x2c\x74\xe6\x3e\x32\x93\x67\xbb\xd6\xf4\x8b\x4a\x62\xef\x84\xa6\x5a\x41\x85\xe9\xbb\x9d\xdb\x70\xb0\x2d\x25\x10\x3d\x39\xaa\xc7\xdd\xa2\x68\xe2\x81\xb1\x86\xdc\xc8\x6c\x51\x84\xbf\x19\x01\x5f\x30\xed\xce\x10\x16\x19\x49\x8b\xb4\x40\x69\x6f\x03\x56\xdb\xaf\x9c\x71\x4f\xc5\x39\x26\xa9\xa5\x81\x41\x35\x2e\x92\xd4\xc4\x7a\xab\xca\x3b\x11\x52\x1b\xde\x28\x10\xaf\x94\x44\xf4\x4a\x2d\x93\x58\x1a\x3b\x11\xb7\xa9\xbd\x81\x7a\xd7\x38\xad\x03\x75\xaf\x93\x7b\x69\xbe\xd6\xe4\xcb\xc5\xc7\xab\x3b\xab\xc2\x6e\xe5\x1f\x05\xf3\xe0\xe8\x80\x58\xbb\x7a\x26\x0d\x57\x67\x87\xef\xb9\x5a\xcb\x2a\xb3\x44\xa1\xc0\xfc\xb8\x7a\x7b\xfa\xfc\xa7\xd3\xf3\x77\x97\x8b\xe5\xf6\x29\xf7\xbd\x82\xe1\x46\x31\x43\x37\xae\x7d\xad\x2e\xa6\x3b\x3b\x9d\x3c\x18\xde\x9a\xf9\x89\x8b\xe9\xb3\xee\xc2\x93\x0d\xe4\xdd\x6a\x5b\x8e\x88\x7f\x5c\x65\x94\xfd\x26\x6c\x02\x67\x57\xa9\x4c\x3f\x64\x7b\x6c\xeb\x76\xb8\x41\x8b\x2a\xdb\x14\x06\x4e\xaa\xe8\x79\x2a\xb3\xf7\xcf\x97\x61\x80\xec\x3c\x85\x2e\x58\x8d\xca\x24\x0a\x56\x4f\xb1\x05\xb9\xb0\x58\xce\x71\x6a\xf4\x03\x34\xc5\xfa\x12\x27\xe5\x5c\xfc\xf5\x2f\x7f\xfe\x1e\xbf\xee\x7f\xab\xb8\x70\x66\x46\xe7\x45\x3f\x12\x8d\x3b\x4d\x28\xda\xd2\x4b\x60\x1a\xcf\xee\xd6\x67\x1c\x8f\xb9\x8a\x69\x78\xc3\xd4\xbe\x3d\x93\x1d\x5f\xaf\xb5\xa0\x03\xca\x9d\x4d\xbd\x24\x34\xea\x50\x25\x51\x14\xf1\x18\xa4\x81\xcf\xfd\x77\xe9\xa2\x83\x31\xc6\xb1\x8a\x7e\x29\x53\xd4\x1b\x95\x1f\xb1\xed\x71\x70\x6f\x82\x0f\xba\x75\x0b\x75\xd4\x1b\x84\x54\x1a\xaf\xc8\x47\xb7\xdd\x4e\x5f\x0f\x7c\xe4\xfd\x6a\x27\xb9\xe7\x47\x23\x42\x12\xd1\x79\xa6\x54\x11\x7e\x77\x7c\x7c\x2c\x1e\x0b\x5e\x39\x4b\x33\xd4\x02\xb0\x4e\x9e\xd4\xe3\x81\xaa\x7c\x43\x89\x27\x78\x12\xbd\x51\xb7\xe7\xaa\xdc\xaa\xb3\xea\x0b\x13\xa2\x9f\x65\x9e\x64\xea\x25\x05\x2d\x98\x95\xea\x9a\x2a\xa9\x44\x9c\x39\x8c\x28\x35\x27\xb6\x52\xa6\x00\x48\x51\x1c\x98\x72\x8a\x26\xf9\x2c\x1e\x7b\xca\xe7\x4a\x51\x3f\xfb\xd9\x07\x4a\x74\xa6\xec\x8d\x4e\x28\x5e\xc1\xf2\xed\xf9\x3b\x77\x42\xc1\x03\xe3\x02\xfb\x33\xc6\xac\x2a\x43\x16\xc7\xdc\xb0\x95\x79\xa3\xed\x4b\x5d\x39\x9b\x5b\x4f\xef\x05\xfe\x99\x22\x4b\x2d\x95\x24\x9f\x32\x0e\x59\xa1\xa0\xb0\x18\xd2\x56\x2b\xb5\xd1\x56\x11\x1d\x58\x60\xce\x59\xa5\xbe\x54\x79\xd8\x08\x4e\xc4\x0f\xe2\xbb\xe1\x9c\xed\x4b\x53\xbe\xa6\x9d\xde\x73\xb3\xb7\x64\x3a\x37\x9d\x80\x94\xc6\xbc\x9a\x31\x50\x04\x98\x04\xe6\x94\xbe\x26\xfb\xaa\x76\x9c\x20\x8b\x3c\x2c\x5d\x14\xec\x44\x6d\xd8\xc5\xf1\x47\x10\x37\x32\x3e\x75\x4d\xed\xa6\xcf\xe7\xe8\xfd\xea\x75\xf4\xb7\x0a\xd3\x33\x9c\x44\xaf\x94\x0d\x03\xb0\x5c\xd6\x8d\x3f\xd9\xa9\xeb\x85\xd7\x2d\xea\xdc\xa9\x64\xc0\xd3\x0c\xf1\xd6\x90\x83\x2c\xad\x39\xc4\x32\x9b\xe1\x77\x3d\x74\xd0\x1f\xf6\x46\x35\x4a\x00\xac\xf0\x2e\xb6\xa6\xb8\xc1\xae\xb3\xed\x95\xfe\x02\x89\x7e\x07\xb6\xf0\x22\xa2\x7f\x21\xc7\x26\x51\x6b\x85\x78\x36\x94\xf7\x79\xd6\xd0\x5a\xb0\x70\xd1\x5a\xf2\x11\x3b\xb7\xa6\x13\x1b\x76\x21\xf4\x11\x9a\xa9\xa8\x51\x08\xd7\x3b\x2b\x69\x51\x4b\xb4\x6a\x76\x75\x9b\x76\x49\x9d\x6d\xb1\x3c\x15\xda\x01\x99\x16\x87\x74\x2d\x70\x75\xf5\x0d\x78\xd8\x2d\x08\x3c\x04\x66\xfa\xb1\x85\x9a\xbd\xc1\xed\x31\x41\xa3\x67\xea\x58\x30\xe0\xca\x35\xb9\x00\x90\x6b\x19\xf1\xd6\x41\xa1\x12\x2c\xb2\x3b\x9a\x31\x7c\xac\xaa\x4d\x61\xef\x08\xd8\x07\x7e\x0a\xf4\xb1\x12\x03\x93\x33\x0f\x4b\xa0\x6f\x32\xde\x3b\xe3\xf6\x83\x24\x0f\x67\xd8\x0e\x74\xc4\x01\x60\xd2\x6f\x80\x7e\xdf\xd1\x28\x49\x1a\x5b\xc9\x32\x0f\x9f\xe9\x95\xe7\x32\xce\x8c\x97\xfe\xcc\x00\x3b\x3a\xbb\x4b\xbe\x1f\xce\x2d\xc4\xe3\xb2\xb9\xfb\xc1\xb0\x3f\xc8\xfc\xfa\xf7\xba\xfb\x10\xac\x3d\xcd\x32\x1e\x79\x3f\xea\xe4\xee\x40\xc6\x1e\x30\x71\x91\xc3\xf1\x5c\x66\x3c\xfd\x4b\xce\xd4\xa4\x27\x12\x7a\x94\xd5\x4d\xe4\xd0\x1f\xd7\x46\xb5\x19\xd1\x4f\x99\x76\x77\x27\x86\xfc\x54\x12\x83\x76\xf8\x1d\x90\x98\x4a\x70\x08\x8a\xff\xf7\xce\xf4\xce\xf0\x3d\x6d\x35\x18\xa4\x87\x3a\x79\x7f\x23\xfb\xce\xea\xcc\x9c\xc3\xd3\xa8\x33\x8c\x0e\x4c\x20\xcc\x1b\xf0\x8f\x0f\xce\xbe\x07\xc2\x75\x1a\xc7\xaa\xb0\x2a\xd9\xdf\x3c\xb5\xb7\xe6\xbf\xd7\x35\xaf\x5e\xfc\x47\x4d\xd3\x1e\x0d\xbb\xf3\xbb\x43\xe9\x44\xec\xf0\x70\x6b\x8f\x93\xff\x63\xc7\xdc\x77\xb3\x11\x76\xa6\x56\xb7\x4c\x1c\x64\xa7\x29\x48\xdf\x24\x08\x53\xfc\x09\x38\x8e\x2a\x8e\x4d\x78\xcd\xe4\xd3\x3c\xe1\xdd\xc3\x60\xce\x54\xd4\xdd\xa4\xf9\x08\xd3\xbd\xe1\x89\x90\xc0\xfb\x62\xc9\x51\x80\x9d\xe4\x5c\xba\x96\x31\x0d\xcf\x0e\xee\x67\x97\x78\x79\xff\xa5\xd0\x65\xc1\xb1\x2e\xbf\xde\x7b\x6d\x86\x50\x23\x0d\x6e\xf4\xb3\xda\x92\xc9\x60\xb8\xd5\xe5\x27\xc2\x44\x5e\x5f\x20\x8e\xc4\xce\x95\xd8\x7d\x50\x01\xb4\x66\x43\xda\x4f\x2a\xce\xae\x01\x40\x5f\x53\x2d\x08\x99\xb7\x4a\xa9\x8d\x3e\x10\xbe\x62\x09\x8a\x2c\x8d\xc9\xd6\x23\x5e\x8e\xa8\x3f\xd8\x9d\x7d\xe9\x05\x10\x06\xf2\xae\x94\xcf\x8a\x37\x48\x7a\x38\xe9\xbf\x4e\x91\x82\xe1\x89\x87\xb6\x28\xe8\xd6\x01\x77\xfd\x74\xf6\x6d\x2d\x6b\x8c\x67\x70\x7e\xc5\x37\x62\x4b\xaa\x68\x31\x0a\xe9\xf3\x9c\x07\x3e\xc0\xf8\x4a\x3c\x76\x91\x7c\x03\x8c\xcf\x20\x9e\x20\xc4\x16\x0b\xd1\xc2\xbc\xd6\xba\xb8\x92\x54\xc0\xe2\xd1\x23\xe1\x17\x5f\x65\xfa\x0a\x37\xb3\x3c\xa5\x1b\x42\x87\xf2\x4e\x3f\xc5\xaf\xde\xfd\xa3\xce\x55\x4b\xe7\x6b\x1c\xd3\xee\xeb\x6b\x02\x7f\x1e\x78\x38\xab\xed\xcd\x6b\x9d\x22\xfe\xb8\x90\xa3\x1e\x5d\x8a\x70\xac\x23\xab\xbd\x0b\x9c\x43\x91\x4c\x9d\x70\x25\xce\x66\x62\xd5\x5e\x98\x85\xdb\xca\x30\x4e\x44\xc8\x70\xd1\x16\x7a\x2d\xe8\x72\xcb\xf7\x1f\xac\x2e\x96\x0d\x9c\xc4\x5e\x8c\x00\x88\x1c\xb9\x9a\xee\x5e\xbe\x41\xae\x8b\x9a\x38\xfc\x0f\xbe\xe0\xfa\x77\x0a\x45\xce\xf7\x7b\xe0\xfd\xb4\x70\x25\x9d\x53\x9c\x68\x61\xc3\x9c\xdd\xb2\xa6\xf0\x20\x1e\x38\x52\x49\x88\xd1\x5b\xbd\x43\x5e\x7f\xff\x23\xca\xf1\x33\x3c\x7f\x10\x39\x1e\x47\x47\xac\x01\x22\x17\x29\xc3\xd0\x82\x9e\xff\x14\xe1\xdf\x49\x3d\xde\x27\xdd\x10\x83\xad\x69\x4f\x37\xa3\xdd\xe7\x1b\xdc\xd0\x3b\x5f\xc8\x7c\x77\xd6\x65\xec\xe1\x32\x27\xc5\x75\xc5\xee\x39\x91\x45\xbf\x6f\xd8\x1d\x9c\x71\x6c\x1c\x96\xeb\xf9\xba\xde\x81\x17\x3c\xa9\xfa\x9f\xcd\x1c\xdc\x9a\x8a\xe3\xef\x9f\x3e\xdd\x7b\x38\xef\xe8\xec\xc5\x38\x1c\x4a\xec\xff\xc4\x45\x41\xf2\x5f\xb8\xf8\xcb\x37\xe3\xbc\xee\xa7\xbb\x7b\xc4\xf3\x5f\x01\x00\x00\xff\xff\xb2\xfd\x8c\x92\x1d\x18\x00\x00")

func instanceListenerMainGoBytes() ([]byte, error) {
	return bindataRead(
		_instanceListenerMainGo,
		"instance-listener/main.go",
	)
}

func instanceListenerMainGo() (*asset, error) {
	bytes, err := instanceListenerMainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "instance-listener/main.go", size: 6173, mode: os.FileMode(420), modTime: time.Unix(1463421683, 0)}
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
	"instance-listener/Godeps/Godeps.json": instanceListenerGodepsGodepsJson,
	"instance-listener/Godeps/Readme":      instanceListenerGodepsReadme,
	"instance-listener/main.go":            instanceListenerMainGo,
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
	"instance-listener": &bintree{nil, map[string]*bintree{
		"Godeps": &bintree{nil, map[string]*bintree{
			"Godeps.json": &bintree{instanceListenerGodepsGodepsJson, map[string]*bintree{}},
			"Readme":      &bintree{instanceListenerGodepsReadme, map[string]*bintree{}},
		}},
		"main.go": &bintree{instanceListenerMainGo, map[string]*bintree{}},
	}},
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
