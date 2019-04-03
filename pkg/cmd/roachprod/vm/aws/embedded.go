// Code generated by go-bindata. DO NOT EDIT.
// sources:
// config.json

package aws

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

var _configJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\xdb\x6e\x23\x47\x0e\x7d\x9f\xaf\x30\xfc\x1c\x03\x64\x5d\xc8\x62\x7e\x65\xb1\x08\x58\x2c\x72\x22\xc0\xb1\x07\x96\x94\x45\x10\xcc\xbf\x2f\x64\x6f\x7a\x54\x33\xad\xd1\x65\xa5\xb7\x2e\xf5\x61\x55\x9f\x43\xf2\xb0\xfb\xef\x4f\x0f\x0f\x0f\x0f\x8f\x6f\xfe\x79\xf3\xfa\xb2\x7d\xfc\xf5\xe1\x63\xe1\x7d\x71\xeb\x2f\xdb\xcd\x6e\xf3\xa7\x3f\xfe\xfa\x10\xfa\xbc\xf5\x5f\xbe\xfd\xb7\xfb\xeb\xcb\x61\xf9\xf1\x79\xb3\xdd\x3d\x1e\xad\xff\xa9\xcf\xfb\xc3\x1f\xff\x5a\x96\x0e\xbf\xbf\xa7\xab\xf7\x1b\xf5\x8f\xcd\x6f\x9b\x71\x08\xa1\x7f\x6c\x9e\x4a\xa3\x0c\x96\xfc\x28\xd4\x72\xe7\xc7\xd9\xde\xef\xfc\xf2\xf4\xf2\xfa\xb6\xfb\xdd\x75\xbb\x7b\xc2\xb5\x7b\xb7\x6e\xfb\xb7\xcd\xee\xaf\xdf\x3e\xbf\xbd\xee\xbf\x1c\x30\xdb\xcf\x4f\x00\x40\x5e\x1a\x0c\x66\x45\x40\x28\xab\xc8\x7d\x7f\xf1\xdd\xcc\xc0\x7c\xe0\x69\x73\x7d\x8f\xfd\x8e\x79\x82\x81\xa5\x8c\x9e\x4d\xbc\xb0\x8f\xa8\x2b\xf1\x57\x42\xd8\x71\x88\x14\xa6\xaa\x94\x30\x85\x65\x43\xbd\x2c\xc4\x38\x0e\xe1\x72\x78\xce\x68\x3d\x77\x0f\xc2\x78\xfc\x21\xc2\xd7\x69\xe5\xeb\x2f\xd7\x69\xa4\x9d\x79\x14\x5b\x7b\xba\x53\x1a\xa5\xcb\x35\x72\x00\x4b\xd1\x22\x71\x51\x08\x5f\x23\xe0\x4a\x8d\xd2\xac\x51\x2a\xa5\x40\x4f\xa2\x4c\x9d\xd3\x5a\x0e\xac\x84\x98\x34\x62\x2d\x81\x65\x40\x1b\x9e\x59\x46\xbb\x37\xc1\xad\x78\xee\xc9\xfb\x59\x82\xb7\xaf\xfb\xdd\xef\xd7\x14\x40\x56\x6a\xbd\xc3\x60\xaa\x86\xb9\xae\x96\xd9\x25\xe4\xfe\x6f\xe3\x89\xd8\xd4\x68\x64\x35\x04\xa9\x61\x14\x6b\xa7\xfa\x0e\xde\x8f\xe1\x98\x88\x48\x00\x34\x53\x62\x68\x77\x27\xb5\x73\xe4\xd6\xec\x42\x52\xaf\xed\x2c\x4d\x4a\x2b\x61\x55\x2a\x57\xb4\xe0\xff\x8b\xd8\x95\xce\xd2\x5a\x6d\x42\x8a\x19\xa1\x62\x86\x7c\x8e\xdc\x8f\x10\x13\xc1\x00\x5d\x7b\xc5\xd1\xa3\x78\x43\x84\xcb\x42\x4c\x89\x4f\x64\xac\x2c\xc3\xa3\xb8\x68\xf5\x7b\x6b\x94\xb3\xf6\x9a\xea\x2a\xef\x27\x34\xba\xa2\xb3\x40\xef\xce\x05\x87\x0c\x80\x18\xf9\xe6\xce\x72\xb4\xf9\x5c\x00\xca\x4c\x1a\x26\x46\x6c\xe1\xed\x22\x82\xd3\xa4\x91\x72\x09\xa7\x4a\xa8\x79\x34\xa6\xcb\x64\x9e\x9b\x93\x1a\x50\xa3\xcc\x1e\x23\xa3\x2b\xdf\x5b\x23\xe4\x3c\x7a\xe3\xb5\xa3\x1d\x69\x64\xfa\x64\xfe\xb2\x7b\xd3\xe7\x6b\xaa\x68\x08\x07\xbb\x65\x1f\xd6\x82\x6d\x1d\x79\x4e\xa1\xe3\xad\x67\x7d\x3c\x5a\x8b\xec\x9d\x81\x12\xe3\x5a\xf3\xfb\x21\xc0\xa4\x0e\xa7\xee\x04\x03\x93\x76\x31\x23\xbb\x37\xb5\x15\x6a\xb5\x91\xe3\xe7\xd4\xfa\xfe\x16\x6a\xab\xb0\x18\x36\x1f\x09\x95\xb8\xde\xd6\xa0\x8e\xb7\x9e\x4d\xb5\xd9\x10\x84\x4e\x31\x28\x67\x3d\x49\xed\x71\x80\x89\x5a\x8a\x1c\x89\x90\x35\x4a\x1e\xe2\x27\x5d\xf9\x38\xc0\x94\xf6\x11\xbd\x76\x1c\x09\x3c\xea\x10\xc6\xbb\xa7\x7d\x67\xc1\x46\xab\xed\x66\xd2\xe6\x3f\x7e\x9d\x73\x94\x2c\x07\x37\xae\x6c\x03\x04\x6f\x74\x8e\x65\xdf\x49\x15\x49\x3a\xbc\x6b\x96\x46\x0a\x09\x4f\x36\xa4\x05\x3d\x49\x32\xa0\x72\x43\x6d\xa6\x85\x3b\x57\x3b\x8b\x9e\xad\x22\x77\x2d\x65\x68\x4d\x5a\xa5\xd4\xbb\xeb\x21\x05\x1d\x4a\xac\xb9\xd8\x8a\x1e\x57\xb8\x84\xf5\x4a\x18\x44\x20\xb5\x6a\x92\xb5\xc7\xbe\x5c\x8f\xd9\x20\x4a\xa6\x51\x7b\x6b\x18\x3d\x4b\xb5\xb3\x8c\xce\xde\x10\xde\xa4\xb2\x0b\x91\x42\x55\xff\x59\x8d\x7d\xa0\x27\x3d\x4a\xc3\x24\x46\x38\x70\x38\x23\xa4\x7b\xeb\x61\x87\xa9\x47\xfa\x2a\x5d\x3f\xea\xb1\xea\x1e\x27\x46\xd6\xd4\x4d\x19\xa0\x49\x2e\x9e\xec\xb6\x91\x75\xd9\x77\x36\x04\x0b\xc4\x86\x84\x18\x62\x40\x74\x8e\xd1\xdc\x67\xb4\xb5\x51\x22\x37\x02\x62\xc3\x71\x16\x3d\xe9\x61\x35\xbc\x8a\xf5\x34\x3c\xa8\xb6\xbb\xdb\x74\xef\xd2\x07\x8f\xd5\x76\xf2\x4d\x8f\xad\x3e\x5d\x3b\xe9\x3a\xc7\x08\x49\x96\x3b\x8e\x3e\xf0\x36\x8f\x5e\xf6\xd5\xef\x73\xb4\xd7\xdc\x11\xba\xbb\xb7\x35\x1f\x9c\xd1\x93\x1e\x62\xd6\xc2\x6c\xb4\x52\x39\x5a\x3d\xf9\xe6\xbc\xa0\xe7\x7e\x95\xb2\x84\x69\xee\xae\xa5\x06\xdc\x5d\x0f\xa2\x0a\x64\x78\xa6\x3e\xf6\xdb\xab\xf5\x10\xce\xa0\xb5\x1b\x97\x9c\xb4\xfb\x6d\xfe\xb1\xec\x3b\xe9\x11\xa5\x5b\x6b\x3e\x24\xd4\xda\x48\x27\xa7\xd1\x05\x3d\x4f\x4b\x20\x80\x49\xc4\x1a\x80\x61\x39\x59\x1f\x0b\x7a\x1e\x63\x3d\x6b\x4d\x4e\xd9\x99\xb1\xc7\xc9\x5c\x58\xd0\xd3\x27\x90\x40\x84\x1a\xbd\x72\x48\x85\xa0\x93\xb9\xb0\xa0\x7d\xaa\x6c\x36\x33\x4b\x66\xd6\x79\x0c\x3f\xf9\xa6\xb4\xa0\x63\x7a\x91\x95\xda\x47\x82\x96\x11\x6c\x64\xb8\xbf\xf3\x51\x75\xea\x71\x66\x00\xff\xe7\x6c\x57\x38\x5f\x46\x39\xbc\xbc\x48\x15\x51\xc2\xb2\xfa\x7d\xe7\xe2\x4c\x9a\x9d\x0f\x8a\x91\x0e\xc6\x84\xda\x46\xd5\xb5\x24\x9d\xd1\x73\x26\x71\xed\xae\x9a\xa2\x18\x97\x58\xad\xa1\x19\x3d\x65\x92\x00\xd5\xc3\xb8\xdf\x65\x00\x96\x32\xee\xad\x07\x70\x6d\xb5\xd0\x99\x4e\xbb\xdf\x5e\x3d\x19\x0e\x54\x2f\x9a\x87\xd5\x41\x50\xe0\x36\xe7\x5b\xf6\x9d\x2b\x5b\x5b\x53\x4a\xa9\x08\x37\xaf\x69\x2d\x49\x66\xf4\x5c\x9b\x4d\x07\xf8\xe8\xd1\x3c\x17\x8b\xbb\x33\xca\xd2\x38\x0b\x9c\xf9\x0c\xf0\xcf\xd9\xae\xc8\x70\x3a\xbc\x5f\x44\xe3\xa6\x12\x89\x7d\x7d\x0a\xb9\x94\xd1\xef\x66\xbb\xe0\x8e\x46\xd6\x6b\xb4\xc6\xa7\x67\x89\x05\x3d\x65\xb8\x29\x5b\xea\x92\xb5\x90\x68\xe2\x9f\x75\xbb\x95\xd9\x2e\x48\x54\x52\x22\xab\x94\x35\xe5\xb3\x6a\xa6\xa9\x57\x62\xe7\x54\x98\x25\x73\x37\x4e\x92\xcf\xaa\xb9\x5c\xfd\xfb\xd3\xc7\xf5\xd7\x4f\xff\x0d\x00\x00\xff\xff\x2e\x40\xf2\x0c\x3e\x18\x00\x00")

func configJsonBytes() ([]byte, error) {
	return bindataRead(
		_configJson,
		"config.json",
	)
}

func configJson() (*asset, error) {
	bytes, err := configJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.json", size: 6206, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
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
	"config.json": configJson,
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
	"config.json": {configJson, map[string]*bintree{}},
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
