// Code generated for package configs by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.yaml
package configs

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x52\x4d\x6f\xd3\x40\x10\xbd\xfb\x57\x8c\x94\x73\x1d\x27\x69\x3e\xb4\x27\xda\x34\x15\x41\x0d\x44\xd8\x55\x8f\x68\x13\x4f\x1c\x97\xb5\x77\xb3\x3b\x4e\xb7\xdc\xb8\x70\x43\x1c\x10\xf7\x9e\x38\x02\x07\x24\x24\x8e\xfc\x97\x14\x7e\x06\x5a\x3b\xb5\x03\xcd\x69\xf5\xe6\x65\xde\xbc\xf7\x1c\xa2\xde\xa2\x66\x1e\xc0\xcb\x22\x9f\xc9\x18\x19\xc4\xb8\x28\x12\x0f\xe0\x29\x91\x9a\x4b\x4d\x0c\x46\x41\x10\x38\x06\xf2\x38\x4a\x33\x94\x05\x31\x18\x38\xe4\x4a\xa7\x84\x87\xd0\x89\x52\x6e\xd7\x19\xae\x78\x21\x68\xce\x13\x0c\xd3\x37\xc8\xa0\xe3\xd8\x33\x6e\x0f\x11\x07\x5d\xc8\x24\xe4\x5b\x9c\x73\x5a\x33\x30\x24\x35\x4f\xb0\x2d\x64\x62\xaa\xd9\x79\x2a\xf0\x39\xcf\x90\x01\x57\xaa\x81\x26\x96\x18\xf8\x42\xba\x2b\x2f\x95\x90\x3c\x7e\xbc\xa4\x28\x71\xd3\x30\x4a\xa3\x97\x5a\x30\x58\x13\x29\xd6\x6e\x77\xba\x43\x3f\xf0\x03\xbf\xc3\x9c\xbf\xb6\x21\x4e\xe9\xb2\xe6\x4f\x33\x9e\xe0\x8c\xdb\xea\xda\x3e\xb4\x60\x76\xfa\xef\xf0\x44\x08\x79\x33\xb1\x64\x9c\x63\x80\x23\xf0\xaf\x55\xd2\x3c\xb1\x7e\xab\x3c\x69\x32\x19\xcb\x9c\xd0\xd2\x61\x68\x67\x9c\xf8\x82\x1b\x2c\x93\x3b\x8d\x6e\x15\x32\xc8\x6e\xcd\x46\x38\x3d\x83\x3a\x2f\x13\xd0\x52\x12\xb4\x7e\xdd\xdd\x7f\xfa\x76\xff\xfe\xcb\xee\xe7\xc7\x3f\xdf\x3f\xef\x3e\xfc\xf0\x00\xe6\xdc\x98\x1b\xa9\x63\x06\x9d\x6e\xef\xb8\x3f\x18\x8e\xa0\x55\xb3\x76\x5f\xdf\xfd\xbe\x7b\xeb\xea\x94\x86\x1c\xe3\xc1\x74\xaf\x17\x0c\x4a\xc1\x2a\xe0\x85\x90\xc9\x2b\x83\x7a\x9b\x2e\xd1\x03\x88\xf8\x42\xe0\x5c\xe3\x2a\xb5\xfb\x99\x07\x30\x5e\x73\x6d\x90\x18\x14\xb4\x1a\x95\xc2\xda\x94\xf5\x33\x88\x74\x81\x55\xc5\xd3\x58\xe0\x58\xe6\xb9\x69\x5a\x7f\xa1\x30\xdf\x43\xbd\xc0\x7b\x76\x15\x39\xa7\x21\x2e\xb5\xdb\xb5\x5c\x63\x7e\x9d\xe2\xeb\x22\xf7\x00\xa6\xc6\x14\xa8\x2b\xc5\xa3\xe6\x9a\x89\x55\xa9\x46\x06\xc3\x6e\x10\x78\x93\x8c\xa7\x82\xd5\x8e\x4c\x46\xca\xdf\x6c\xfc\xa5\xcc\xdc\x49\xe5\x07\x7b\x3c\xe8\xef\xb3\xab\xcc\x59\x6b\xed\x93\x86\x53\xe7\x65\xab\x5f\x29\x1c\x86\x17\x0c\xa8\xb2\x71\xae\x65\xf6\xff\xbf\x22\xf9\xd0\xf3\x01\xfe\x37\x00\x00\xff\xff\x41\xc0\xfb\x28\x3f\x03\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 831, mode: os.FileMode(420), modTime: time.Unix(1612796952, 0)}
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
	"configs/config.yaml": configsConfigYaml,
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
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
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
