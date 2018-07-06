// Code generated by go-bindata. DO NOT EDIT.
// sources:
// data/values.yaml
package helm

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

var _valuesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x98\x3d\x8f\xe3\x36\x13\xc7\x7b\x7f\x0a\x01\xd7\x3c\x2e\x1e\x57\x41\x8a\x74\xb2\xe2\xcb\x2e\xf6\x36\x2b\xc4\x0b\x04\x29\x69\x6a\x2c\x13\x47\x93\xca\x90\xdc\x9c\xef\x72\xdf\x3d\x90\x64\xc9\x7c\x93\xbc\x54\x52\x19\x20\xe7\xff\x9b\x11\x39\x1c\x0e\xfd\x21\x2b\x72\xb5\xda\x69\x5a\x15\xa4\x00\xd4\x3f\x65\xdf\xbe\x65\x5b\xa2\xe0\xc7\x1f\x76\x82\xca\x0a\xb2\xff\xb5\xc3\xb9\xda\x5e\x34\xa8\x6c\x53\x48\x71\x64\xf5\xe6\x26\x58\x67\x7f\x67\x7f\x1a\xa9\x21\xfb\xfe\x7d\x55\x90\x27\xb8\x44\x10\x25\xb2\x37\xa2\xe1\x09\x2e\x3e\xa8\x53\x78\x8c\x84\x38\x22\x31\x7c\x44\x29\x74\x89\xf2\xcb\x25\x89\xe4\xcb\x1c\xe6\x1e\xf0\x8d\x51\xd8\xb3\x5a\x30\x51\xa7\x7f\x65\x44\x3f\xcb\x4f\x88\x3b\x26\x8d\xb1\x0b\xa2\x09\x97\x8b\xd8\x8e\xd4\x61\xaf\x3e\x64\x54\x0a\x4d\x98\x00\xcc\xd8\x99\xd4\xa0\xb2\xa3\xc4\xac\x91\x95\x5a\x3d\x13\xa5\x01\xdb\x4c\x79\x6c\x67\x3a\x9f\x03\xd9\x9b\xb3\x99\xfd\x54\x5e\x3e\x4e\xa9\x86\xa9\x50\x54\x48\xa1\x51\x72\x0e\xa8\xa6\xc4\xbe\x89\x0d\xf9\x55\x56\x10\xea\xc6\xd1\xe9\x45\x0d\x45\x91\x79\x47\x7e\x11\x34\x22\x1a\x46\x6d\xd3\x57\x23\x04\xf0\xd0\xd8\x1a\xb7\xcd\x4b\x94\x67\xd0\x27\x30\xaa\x8d\x7b\xf7\xa5\x91\xa8\x01\x43\xf9\x8c\x9d\x8d\xfb\x0d\x6a\xa6\x34\x5e\x42\x80\x33\xe3\x48\xa4\x89\x7a\xb4\xc6\x6d\xf3\xfc\xab\x41\x28\x3e\x45\x76\xdb\x99\x89\x05\x55\x48\xa1\x24\x8f\x6c\x59\xcc\xc0\xf1\x29\x14\x3b\x70\xb8\x6e\xd2\x16\xe5\xe7\x58\xc0\x93\x56\x36\xea\x77\x38\x4c\x86\xe1\xcd\xd9\xb2\x97\xdc\xe8\x53\x57\x6d\x42\x99\x37\x17\xdf\xdd\xb9\x1d\x9d\x91\xe5\x1c\x50\x6f\xcd\xf1\x38\x9f\x13\xbe\xd9\x0c\xec\x99\x08\x52\xbf\x83\x66\xdb\x79\x25\x04\x34\xad\x32\x0a\xa8\xd9\x91\x51\xa2\xa1\xbf\x8f\xda\x75\x07\x4c\xad\xb4\x8e\xd2\xa9\x55\xb7\x99\xc4\x7b\xee\x26\x0a\x78\x25\x2c\x8b\xf0\xaa\x8b\xf2\x12\xa3\x1b\x24\x01\xab\xe0\x0c\x84\x5e\x12\xdd\xa8\x9c\x60\xa6\xf6\x09\xa3\x28\x76\x77\xa0\xe4\x59\xc3\x89\x00\x37\x03\xfa\x6a\xbd\x2c\x07\x3c\xed\x3a\xbc\x28\x92\xf3\xc0\x97\xad\xc3\x12\x9d\x1a\xe5\xa8\x8a\xb0\x12\x22\xbb\x09\x1c\x4e\x5e\xd7\x08\x35\xd1\x12\x6f\x8d\x4d\x6a\x84\x13\x8c\xbb\x7e\x12\xa2\x9f\x92\x47\xf6\xec\xc9\x1c\x80\x83\x5e\x96\xd6\x71\xc4\x3d\x2f\xc9\x09\x12\xa8\x23\x1e\xfa\x6f\xfc\x17\x5f\xe1\x02\xe6\x3d\x24\x7f\x81\xa7\x75\x4f\xac\xdb\xd5\x2c\x3b\x9c\x13\x8c\x99\x96\x35\xf9\xb0\x4e\xc9\xfd\xf2\x73\xee\xbe\xf8\xff\xb4\x53\x65\xfd\x8f\x41\x72\xe0\xa0\xba\x2e\xa6\xed\xa7\xbb\xc1\xbd\x39\x54\xf2\x4c\x98\x18\x2f\xb8\xbe\xe5\xbd\x3a\x6a\xaf\xba\xa6\x2d\x5e\xa0\x36\x2f\x48\x4f\xa0\x34\xb6\x29\x5d\xa2\x3c\x32\x0e\x9b\x97\x06\xc4\xfe\xc4\x8e\xda\xea\x84\x42\xb6\x73\xc7\x9a\x03\x67\xf4\x41\x2a\x2d\xc8\x19\xfe\x1b\xaf\x2e\xd3\xf6\xd6\xdd\xc9\xbd\xd5\x47\x89\x67\xa2\x9d\x7b\x3c\x98\xf5\x57\x91\x29\xea\x57\x72\x67\x29\xaf\xf1\xe6\x94\x4a\xd3\x1e\xf1\x21\x33\x96\xa4\xcd\x15\x32\x91\x30\x83\x8b\xee\x4b\x27\x3c\x0c\x73\x77\x1d\x6c\x46\x53\xcf\x95\x52\x4c\x8a\x3d\x50\x04\xdd\x76\x6a\xa1\x93\x1b\xd2\x33\x9d\xe4\xec\x04\x7d\x27\x66\x27\xa8\x4d\x79\xd0\x25\x51\xea\xaf\x6a\x5a\x3d\x58\x78\xbb\xf6\xd9\x1c\xa0\xdf\x26\xb5\xca\xab\x33\x13\x4f\xe3\x40\x64\xd5\xfe\x20\x67\xfe\x4c\x50\x9d\x08\xbf\x95\x6e\x57\x35\x51\x4b\x93\x90\xbe\x6c\xa6\x28\xe4\xe5\x63\x32\x7e\x86\xe0\x78\xda\x4a\xa9\xdb\xc3\xd4\xe4\x46\x4b\xd2\x34\x28\xdf\x16\x7c\xcc\x1d\x8a\x5f\x8c\x48\xfb\xda\xb9\x9e\x9c\xd5\x2b\x08\x22\xf4\xe3\xcf\xf7\x4f\x7e\xfe\x75\x38\xef\x83\xc6\x59\x33\x73\x50\x14\x59\xa3\x99\x14\x69\x34\x57\xe9\xfc\x3f\xd3\x5d\x0e\xef\xa1\x5d\x87\x4a\x64\x82\xb2\x86\xf0\x81\x3d\x10\x42\x6a\x9f\xe3\xcb\xc9\xbd\xde\x7d\x2b\x2a\x69\x90\xc2\x2f\x28\x4d\x93\xb2\x00\x8e\xd0\x26\x7e\x92\x94\xb4\xab\x32\x01\x1b\xa6\xbd\xed\x15\x00\x15\x54\xd9\xe1\x92\xb1\x73\xfb\xd6\x1e\x1f\xb1\x7b\x2d\x91\xd4\x43\xdd\x89\x3e\x63\x5d\x93\xd8\x4b\xf8\xe1\xf5\xb5\xb4\xd6\x2e\x5a\x08\x42\x5b\xa7\x8f\xb3\xde\x68\x5d\x0f\xe0\x54\x9d\x69\xea\xbc\x2e\xf0\xa0\x92\xd9\x11\x45\xfc\x35\x9a\x42\x9e\x53\xcd\xfc\x77\xc6\x4d\x5b\x9f\xac\xd4\x8f\xfc\x49\x36\xd8\x6c\xf6\x1a\x99\xa8\x63\xbb\x95\x7a\xf7\x59\xba\x75\xf4\x7f\x90\xf7\xb7\x46\xb6\x64\x1d\xfe\x6b\x93\x1c\xd9\xa0\x8a\xb0\x52\xa2\x1a\x05\x6e\x55\xec\xdf\x36\x41\x1b\x74\x7b\xf6\xc4\xba\x99\x7f\x02\x00\x00\xff\xff\x49\x2b\xc6\x89\xd2\x16\x00\x00")

func valuesYamlBytes() ([]byte, error) {
	return bindataRead(
		_valuesYaml,
		"values.yaml",
	)
}

func valuesYaml() (*asset, error) {
	bytes, err := valuesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "values.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"values.yaml": valuesYaml,
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
	"values.yaml": {valuesYaml, map[string]*bintree{}},
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
