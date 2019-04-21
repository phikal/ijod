// Code generated by go-bindata.
// sources:
// index.html
// DO NOT EDIT!

package main

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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x58\xdd\x8e\xdb\xb8\x15\xbe\xd6\x3c\xc5\x19\x2d\x76\x61\xef\x58\xf2\x4c\xd0\x14\x0b\x8d\xed\x02\x0d\xb2\x48\x16\xc9\x6e\xd1\x04\xed\x45\x51\x14\x34\x79\x2c\x33\xa6\x78\x54\x92\x1a\x8f\x13\xa4\xe8\x6b\xf4\xf5\xfa\x24\x05\x49\xc9\x96\x6d\xd9\x8b\xed\xee\xc5\x50\xe4\xf9\xf9\xce\x77\x7e\x48\x67\xb6\x76\x95\x5a\xdc\x00\xcc\xd6\xc8\xc4\xe2\x26\x99\x39\xe9\x14\x2e\xde\x7e\x22\x31\x9b\xc6\xf5\x4d\x32\xab\xd0\x31\xe0\x6b\x66\x2c\xba\x79\xda\xb8\x55\xf6\x43\x0a\x53\xaf\x07\x10\x0f\x35\xab\x70\x9e\x3e\x49\xdc\xd6\x64\x5c\x0a\x9c\xb4\x43\xed\xe6\xe9\x56\x0a\xb7\x9e\x0b\x7c\x92\x1c\xb3\xf0\x31\xa8\x69\x70\x85\xc6\xa0\xe9\x69\x6a\xd2\x18\x64\x93\x99\x75\xbb\x00\x04\xa0\x28\xb2\x2d\x2e\x37\xd2\x65\x96\x1b\x52\x6a\xc9\x0c\x7c\x81\x60\xb7\x80\x1f\xea\xe7\x47\xf8\x7a\x41\x2c\x73\x86\xf1\x0d\x7c\x81\x25\xe3\x9b\xd2\x50\xa3\x45\x01\xa6\x5c\xb2\xd1\xfd\x24\xfc\x9f\x3f\x8c\xaf\x69\xaf\x9b\x6a\x79\x4d\xfb\xc5\xcb\xa0\xee\xf5\x3d\xa9\x13\x58\x92\xd8\xc1\x97\x9b\x24\x01\xa8\x98\x29\xa5\x2e\xe0\xfe\x31\x7c\xd6\x4c\x08\xa9\xcb\xfd\x77\x0b\xff\xe1\xfe\xfe\x69\x1b\x77\xd6\x28\xcb\xb5\x8b\x5b\xeb\xb8\x25\xa4\xad\x15\xdb\x15\xb0\x52\xf8\x1c\xb7\xfc\x2a\x5b\x29\xda\x16\xc0\x49\x35\x95\x6e\xb7\x49\xbb\x6c\xc5\x2a\xa9\x76\x05\xbc\x22\x6d\x49\x31\x3b\x81\x37\x8c\x6f\x26\xf0\x9e\x34\xe3\x34\x81\x8a\x34\xd9\x9a\x71\xec\x2d\x7b\xea\x56\x7e\xc6\x02\x1e\x5e\xd4\xad\xab\x4a\xea\xac\x85\xf9\xb2\xdd\x85\x36\xda\x8a\x49\x3d\x01\x66\xa5\xc0\x36\xdc\x9a\xac\x74\x92\x74\x01\x6c\x69\x49\x35\x0e\x4f\xc4\x5b\xb9\x3e\x97\x4b\xc5\xf8\xa6\x17\x56\x69\x7c\x58\x0f\x67\x74\x7c\x7b\x44\x19\x67\x8a\x8f\xfc\x2e\x64\xf0\xe2\xe5\xef\xeb\xe7\xf1\x19\xdc\xb0\x7d\xea\x7f\x01\x4f\x52\x20\xb5\x38\x0e\xfc\x7f\x3b\xe8\xaf\x53\xed\xc7\xd8\xc7\xbe\x5d\x4b\x87\xb6\xa2\x0d\x1e\x81\xeb\x3c\x0f\x06\xb0\x24\x23\xd0\x64\x0a\x57\x7e\xbb\x7e\x06\x4b\x4a\x0a\x28\x0d\xdb\xfd\xd6\x7c\x33\x25\x4b\x9d\x59\x54\xab\x28\x9c\xa1\x16\x67\xa8\x17\xe0\x3b\x1c\xcd\xaf\x54\xe4\x21\xe3\x2d\xc0\x25\x39\x47\xd5\x00\xc4\x4b\xe6\x8f\x98\x71\x8d\xf9\x67\x43\xd2\x62\xdb\x59\xc7\xc2\x05\x5b\xb9\xa0\xd2\xf6\x7c\x01\xe9\x07\x54\xc8\x1d\x30\xf8\x8b\xcf\xcf\xbf\xd2\xc7\x41\x2f\xc5\x9a\x9e\xce\x7c\x71\x32\x4c\x3d\x02\x6f\x8c\x25\x53\x40\x4d\x52\x3b\x34\xc3\x8e\x83\xfe\x80\xfb\x1f\xc9\x70\x84\x0f\x3b\xcd\x6f\x2f\xb8\xce\x15\x31\x4f\xd5\x40\x19\x1c\x82\x0d\x47\x1d\x10\x81\x2b\xd6\x28\xd7\x52\x76\xd1\x60\x87\x26\xea\xee\x11\xbd\x8b\xa7\x79\x9e\xa7\x03\xa4\x37\xea\x1b\x25\xad\x6b\xb5\xfc\x32\x0b\xb3\xb2\x00\x3f\x3d\x0f\x55\xb3\xef\xa4\xd3\xd1\x73\x52\x07\x9e\x16\x5f\x60\xd9\xae\x00\xd6\x38\x3a\x17\xba\x84\x60\x01\x4a\x76\xed\xdf\xf9\xf8\xa1\xab\xa5\x2d\x19\x91\x2d\x0d\xb2\x4d\x01\xe1\x4f\xc6\x94\x3a\x27\xa4\x6f\xab\x4b\xf1\x19\xcb\x25\x93\xda\x2e\xc9\xd0\x31\xcb\x5d\xba\xaf\x1a\xcd\x6d\x28\x2f\x14\xad\xdd\x30\xea\xb6\x6d\x6f\x2e\x49\x0d\xb5\xcd\x37\x8a\x86\x92\xdd\x9b\x57\x9c\x94\x47\x10\xa6\xc0\x20\xab\x03\xdd\x35\x34\xbf\x7e\xb7\x1f\x5f\xfb\x71\xf1\x62\x4f\xe1\x3e\x33\x5b\xc3\xea\x8e\x45\xcf\xeb\xb5\xc4\xb5\x3d\xec\xa8\x1e\x9e\x31\x17\x26\xc7\xe9\x1d\x72\xe1\xae\x18\x64\xaa\x57\x07\x03\xe5\x78\x51\xe5\x28\xdd\x2d\xa1\x57\x07\xcd\xa4\x97\xe0\xc3\x7a\x45\x74\x68\xa2\xee\x12\x6f\x2c\x9a\x2c\x66\xbe\x05\x02\xd3\xef\xe1\x03\x5b\x31\x23\xe1\xfb\x69\x94\xad\xe8\xf3\x05\xc1\x1f\xa5\xc1\x15\x3d\x1f\x24\xed\x05\xc1\xb7\xaf\x1f\xee\xef\xa6\xaf\x45\x89\x9d\xec\x05\xcf\x8e\x69\xc1\x8c\x08\x52\xa1\x5a\x67\xd3\xf6\x85\x03\x30\x9b\xc6\x97\x18\xc0\xcc\x3f\x20\xc2\xeb\x8b\x49\x1d\x1e\x3f\xb3\xdb\x2c\xf3\x97\xab\x43\x33\x4f\xd7\xce\xd5\xb6\x98\x4e\x3f\x97\x98\x37\x36\x77\x34\x95\x55\x39\x65\x5b\xb6\xc1\xfc\x53\x5d\xa6\x90\x65\x51\x29\xde\x74\x52\xf8\xe7\x99\x40\x8a\x2f\x2c\x43\xca\x2e\x6e\x92\xe4\x97\x35\x68\xba\x85\x8f\x6b\x69\x61\x69\x68\x6b\xd1\x80\x20\xb4\xfa\xbf\xff\xfe\x8f\x03\xdb\xd4\xfe\x39\x07\x6f\x3e\xbe\x7f\x17\x6f\xcc\x3c\x98\x9c\x86\xb5\xc7\x36\x6d\xc1\xcd\x42\x06\xa2\xc3\xf6\x26\xf0\x1e\xed\x4e\xf3\x74\x11\x63\x42\x13\x8f\x1b\x15\x8e\x7c\x75\xf8\xa3\x46\x1d\x62\x6b\xc7\x1e\x0a\xe9\xd8\x52\xe1\x3c\x75\xa6\xc1\x43\x24\x9d\x26\x95\xad\x62\x92\x78\x08\xad\x6b\x8f\x2b\x52\x06\x30\xb3\xdc\xc8\xda\x2d\x6e\x92\xb4\xb1\x08\xd6\x19\xc9\x9d\x1f\xa0\x9c\xb4\x75\xed\xe5\x3f\x07\x41\xbc\xa9\x50\xbb\xbc\x44\xf7\x5a\xa1\x5f\xfe\x71\xf7\x56\x8c\x5a\xa6\xc6\x7b\x85\x30\x41\xae\xc8\x87\x60\x0e\xe2\x2d\x03\x57\x14\x02\x31\x3d\xfb\x54\x5e\x35\x4f\xa5\x17\xee\xa4\x2d\xf1\x0d\x7a\x3c\x1a\xb7\xf0\x57\x5c\x7e\x08\xdf\xa3\x9b\x24\x49\xb7\xbe\x24\x52\xb8\x03\x45\x9c\xf9\x37\x58\xbe\x26\xeb\xfc\x1b\x1b\xee\x20\x2d\x8e\x4e\x42\x6a\xef\x20\x9d\x46\x7b\xe9\x4d\x12\x7c\x3c\x31\x03\x8d\x14\x30\x87\xec\xc1\x7f\xaf\x1a\xcd\xbd\x3c\x58\xd4\x62\xe4\x4d\x4d\x40\x30\xc7\xc6\xa1\xcd\x78\x78\x5f\x62\xae\xa8\xec\x9f\xf9\x29\x12\xcd\xe6\x41\xed\xa7\x0f\xbf\xfc\x9c\xfb\x34\xe8\x52\xae\x76\x23\xaf\x99\xa4\x5e\x3e\x2d\xc2\x2f\x80\x49\xd8\xf0\xaa\x69\x11\x2c\xdc\x24\xc9\xd7\xb1\x37\xf3\xb5\x0f\x41\x51\x59\xd9\x72\x54\xd9\x32\x7a\x57\xe8\x73\xd3\xa7\x8e\x1b\x64\x0e\x5b\xf6\x7c\x62\x02\xcb\x89\x92\xb9\xd4\x1a\x4d\xa8\xe4\x39\x54\xb6\x0c\xbb\x54\xe6\xac\xae\x51\x8b\x57\x6b\xa9\xc4\x48\xc9\x71\xb7\x1d\x5f\xfe\x1f\xa9\x86\x39\x1c\xbe\xdf\x84\xc1\x0c\x19\x3c\xbc\x38\x85\x16\xdb\x7c\x54\x33\xb7\x8e\xd8\x62\xbf\x58\xc3\x61\x0e\xe9\xd4\xc7\x14\x12\xe3\x05\x1e\xf7\xc7\xfe\xf2\x1f\x8d\x0f\xdf\xbc\x31\x06\xb5\xfb\x28\x2b\x84\x79\xbc\x3f\xda\x67\x02\x57\xcc\xda\x77\xd2\xba\x9c\x89\x50\x12\xe1\x5d\x10\xa3\x5b\x91\x81\x91\xcf\x9b\x0c\x4a\x20\x61\x16\x4a\x36\xe7\x3e\x2c\x83\x3a\x57\xa8\x4b\xb7\x7e\x04\x79\x77\x17\xd1\x85\x34\x4b\x87\x95\x8f\xaf\x2f\xfa\x37\xf9\x77\x6f\x32\x91\x2b\x18\xf9\xf3\xc8\xdb\x47\x7c\x76\x30\x9f\x07\xf4\x63\x7f\x9c\x84\xb3\x13\x50\xdd\xf5\x1a\x51\x25\xa8\x2c\x02\x9c\x08\x1a\xac\xe8\x09\x4f\x65\xbf\x46\x36\xdb\x92\x21\x5d\xa1\xb5\xac\xf4\x1c\xe0\x93\x83\xf9\x22\x32\xca\x8c\xcf\x1c\xcc\x21\x94\x53\xed\x7f\x8f\x8e\xf0\xc9\xe5\xfb\x92\xeb\x57\x63\x5a\x92\x7f\x40\x4d\xbc\x4a\x38\x74\x26\xfe\x1e\x4b\xec\x56\x3a\xbe\x06\x5f\x46\xb9\x2f\xbe\x96\x11\xce\x2c\x42\x6c\xe5\x22\x84\xb8\x5d\x4b\x85\xa3\xc0\xce\x4a\x1a\xeb\x42\x91\xc4\xe8\x13\x75\x88\xa5\xab\x9d\x63\xb1\xc0\x40\xe2\x5d\x78\x74\x79\xc5\xea\xd0\x1f\x5d\x2c\xc9\x6f\x28\xde\xe0\xee\xa8\x50\x4f\x34\x7c\x76\x7e\x26\x81\xc1\xc3\xb8\xa7\x43\x9a\x2b\xc9\x37\x30\x87\x7f\x78\xc7\xa1\x15\x5b\xe6\xd3\x49\x68\xbc\x83\xb0\x4f\xe2\x79\x2f\xf8\x36\x8c\x7f\xc3\x93\xe3\xf1\x40\x54\xcd\x1a\x8b\x2d\x53\xbe\x58\x7c\xa8\x2b\x43\x15\xdc\xce\xc3\xfc\xf8\xee\xbb\x68\xd9\xaf\x6f\xfd\x30\xd9\xef\xdc\xc6\x5a\x0f\x06\xc4\x78\x4f\x47\xec\xed\xf4\x4f\x61\x1b\x0e\x23\xd8\xff\xd7\xd3\x18\x75\xb8\x86\x51\x29\xb6\xfb\xbf\x40\x5d\xc5\xf4\x8a\xb4\x93\xba\x41\x01\xde\xbe\x7f\x07\x9e\x22\x53\x6c\x77\x0c\xec\xac\x6d\xbb\xc2\x3f\xea\xdc\xf3\x00\x9c\xac\x4e\x59\xf5\x05\xb4\xc7\xf9\x9e\xb9\x75\xce\x96\xf6\x70\x92\xc1\xd9\xec\x18\xc3\x02\xee\xf3\x87\x97\xa7\x81\xec\x75\x66\xe7\x3a\xf0\x87\x28\x9a\x00\x40\xfa\x53\x53\xd5\x28\xc2\x8b\x77\xcb\x8c\xb0\x29\x14\xe7\xa7\x2b\x32\xf1\xf0\x98\x8b\xe3\x19\xd6\xb9\xbc\x96\xb4\x55\xb8\x0e\xa3\x87\xe9\x14\xa4\xe6\x54\xd5\x0a\x1d\x0e\x4a\x37\x52\xa4\x45\x7b\x45\xed\xad\xc7\x17\xf1\x41\xa8\x2d\xf2\xa2\x9b\xcb\x9d\xe0\xb8\x95\xec\xd9\xab\x6c\x99\x16\x70\xc2\xd0\x80\x5c\x4d\x36\xd8\xf3\x3d\xe4\xd7\x93\x01\xde\xfb\x5a\xed\x0f\xbf\x18\x56\x57\x49\x7f\x46\x2e\xf1\x09\x05\x34\x7a\xa3\x69\xab\xa1\x9d\x74\x05\xcc\x38\x09\x5c\xa4\x70\x77\xd3\xf1\xec\x67\xdd\x1d\xa4\xb3\x69\x3c\x89\x2c\x7b\x06\xbf\x02\x67\x61\x82\xed\x47\x57\x6f\xee\xe1\xe0\x40\xa5\x1a\xf5\xf1\x18\xf8\x44\x52\x07\x9b\x7b\x19\x34\x86\x4c\xbc\xef\xe2\x05\xb9\x95\x5a\xd0\x36\x4c\x11\xb2\xd8\xa9\x77\xa1\xcc\xb0\x5a\x70\xd2\x1a\xe3\x05\xc8\x96\x64\x1c\x8a\xdb\xd9\x14\xab\x08\x36\xd2\x43\xda\x37\xc7\xb1\xef\xd0\xa5\x83\xfc\x1d\xb4\x2c\xe2\x06\x7d\x8e\xf7\x66\x7c\x6f\x9e\xd8\x09\x33\xe8\x57\x0c\x71\x36\x80\xc0\x20\x13\xbb\x80\xb2\x6d\xd4\xc1\x51\x19\x0b\x73\xd8\xbe\x7f\x6f\x76\xcf\xcc\xd9\x34\xfe\x43\xea\xff\x02\x00\x00\xff\xff\xd5\x06\x40\xff\x50\x15\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 5456, mode: os.FileMode(420), modTime: time.Unix(1555875276, 0)}
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
	"index.html": indexHtml,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
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

