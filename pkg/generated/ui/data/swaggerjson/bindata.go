// Code generated by go-bindata.
// sources:
// assets/generated/swagger/api.swagger.json
// DO NOT EDIT!

package swaggerjson

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x4b\x8f\xdb\xc8\x11\xbe\xcf\xaf\x28\x30\x01\x72\xf1\x8e\xbc\xce\x65\x31\x97\xc4\x98\x01\xd6\x83\xf5\x6c\x8c\x75\x62\x1f\x12\x43\x28\x91\x25\xb2\x2c\xb2\x9b\xee\x2e\x4a\x51\x82\xf9\xef\x41\x37\x49\x89\xa4\x48\x3d\x28\x8f\x2d\x3b\xeb\x8b\x6d\x75\x57\xd7\xeb\xeb\x7a\xb5\xf4\xdf\x2b\x80\xc0\xae\x30\x8e\xc9\x04\x37\x10\xbc\xb8\x7e\x1e\x3c\x73\x9f\xb1\x9a\xeb\xe0\x06\xdc\x3a\x40\x20\x2c\x29\xb9\xf5\xdb\xb4\xb0\x42\x06\x1e\x50\x61\x4c\x06\xde\x3d\xbc\x47\x43\xf0\x8a\xd2\x9c\x0c\xbc\x7c\x73\xef\xa9\x01\x82\x25\x19\xcb\x5a\x39\x9a\xe5\xf3\xeb\x1f\xab\x63\x01\x82\x50\x2b\xc1\x50\x36\x67\x03\x04\x0a\x33\x7f\xf8\x03\x87\x09\x52\x0a\xef\x48\xd1\x7f\x18\x2b\x0a\x80\xa0\x30\xa9\x5b\x4f\x44\x72\x7b\x33\x99\xc4\x2c\x49\x31\xbb\x0e\x75\x36\xb1\x98\xd9\x42\xc5\x3f\x84\x2a\x94\x49\x98\xe1\x0f\xcb\x6c\x85\x86\xb6\xa4\x94\x21\x7b\xe2\x6c\x59\x9e\xfa\xd7\xd8\x7d\xe2\x88\x03\xbf\xe7\xf1\x0a\xe0\xd1\xab\x6c\xc3\x84\x32\xb2\xc1\x0d\xfc\xb3\x14\xd5\xf3\xab\xe5\x76\xff\x71\x14\x1f\xfc\xde\x50\x2b\x5b\xb4\x36\x63\x9e\xa7\x1c\xa2\xb0\x56\x93\x8f\x56\xab\xed\xde\xdc\xe8\xa8\x08\x8f\xdc\x8b\x92\xd8\xad\xdd\x27\x98\xf3\x64\xf9\xe3\x24\x2c\xcd\xde\x34\x5a\x4c\x4d\x1b\x3a\xf1\x8b\x2c\x43\xb3\x76\xba\xbe\xe7\x34\x05\x43\x62\x98\x96\x04\x92\x10\x58\x41\x29\x2c\xe8\x39\x20\x54\x87\x01\xaa\x08\x58\x2c\x2c\x8a\x19\x85\x5a\xcd\x39\x86\xb9\x36\x10\x6a\xa5\x28\x14\x5e\xb2\xac\x37\x76\x04\x08\x74\x4e\xc6\x8b\x7c\x1f\x39\x1e\x3f\x93\x54\x60\x68\x6e\x32\x64\x73\xad\x2c\xd9\x96\x6c\x00\xc1\x8b\xe7\xcf\x3b\x1f\x01\x04\x11\xd9\xd0\x70\x2e\x15\x50\x1a\x07\x95\x1a\x39\x87\xe0\x0e\x19\x40\xf0\x47\x43\x73\x47\xf1\x87\x49\x44\x73\x56\xec\x4e\xb0\xce\xff\xa5\xfb\xb7\xb2\xfd\x46\x79\xba\x0e\x5a\xe4\x8f\x57\x7d\xff\x7e\x6c\x28\x91\xa3\xc1\x8c\x84\xcc\xd6\x65\xe5\x9f\x8e\xf8\x35\x6e\xfd\xdf\xcf\xf6\xaa\xf6\x2b\x66\xe4\xac\xef\x7c\x51\xdb\x5f\x34\xcc\x08\x52\xad\x17\x14\x41\x91\x5f\x77\x8f\x60\x4f\xf9\xa9\x20\xb3\xee\x2e\x19\xfa\x54\xb0\x21\xe7\x88\x39\xa6\x96\x3a\xcb\xb2\xce\xbd\x60\x56\x0c\xab\x38\xe8\x55\xf8\x43\x43\x61\xc1\xb8\xab\x6a\x7d\xd3\xb7\xc4\x1f\xae\x3a\x96\x0a\x22\x4a\x49\x68\x3f\x0a\xcb\x3d\x5b\xd4\xed\x41\xd4\x9d\xdf\x7a\xb1\xa0\x6a\x89\x77\x29\xb8\x7a\x9f\xa0\x00\xdb\x26\xae\xfe\x64\xc1\x11\x3a\x78\x45\x64\xc5\xe8\xf5\xb7\x87\xac\xbc\x38\x10\xdc\x30\xfa\x58\x58\x01\x84\xdc\xe8\x25\xbb\x4c\x73\x14\xc4\x5e\x7a\xb2\x4a\x80\x5f\x75\x44\xf6\xf2\x70\xd6\x92\xf1\x8b\xe0\x6c\xa6\xa3\x1d\x1c\x94\x10\xe9\x5b\x69\x20\x44\x4c\xd1\x05\xc8\xe7\x51\xfb\xc1\xc6\xc7\x28\x3d\x1e\x69\x57\x0d\x9b\x75\xb3\xec\x24\x65\x2b\xe3\x52\x2d\x82\xa3\x75\x81\xbe\x3a\xcb\x1e\x95\x41\x5f\x3b\x86\x17\x07\xc4\xb6\x7c\xa3\x90\xf8\x19\x9d\x52\xe4\xb1\xc1\x88\x4e\xf5\x4b\x61\x14\x54\xa4\xa0\xbd\x8d\xac\xaf\x72\x10\x62\x5e\x92\x3a\x22\x66\xfc\x4c\xf2\x8f\xf2\x80\x4a\xf2\x7b\x35\xd7\x26\xf3\x3b\x2e\xd2\x69\x83\xd2\x5e\x70\xd2\x02\x71\x9f\xad\x08\x5c\x37\xe1\x8a\x6b\x8e\xc8\x65\x18\xef\xab\xca\x7f\xdf\x63\x1a\x13\xa1\x2c\x17\x97\xab\x6b\x90\x1e\x93\xc6\xda\x1e\xbe\x3c\x10\xb6\xe5\xfb\xff\xc9\x61\x6d\xbd\xbf\x4e\x12\xdb\x36\xde\x27\xc7\xc9\x8a\x14\x78\x1b\x32\x00\x67\xba\x10\xc0\x9c\xc1\x92\x59\x1e\x0a\x94\xef\xca\x13\x2e\x3d\x42\x56\x62\x7e\xb1\x94\xb6\x19\x31\x34\x04\xda\x36\xf9\xdd\xca\xa7\x9c\xa9\xfc\x46\x99\x5e\xd2\x03\x86\x09\x2b\x7a\x9b\x53\xd8\xf4\x68\x1d\xbf\xf4\xec\x23\x85\xdb\xd2\x21\xc8\x8d\xf3\x89\x70\xc7\xc4\x41\xa2\xad\x74\x8d\xde\x89\x81\xcf\x5a\x6b\xf5\xbc\xe7\xef\x09\x81\x23\xf6\x71\xf8\xed\xdb\x57\x80\x61\x48\xd6\x6e\x35\x7d\xec\x05\x63\xc7\xc6\x3d\xb0\x38\x43\x99\x98\x65\xba\x8b\xf2\xd3\x74\x12\x8c\x41\x2b\x9f\x81\x62\x16\x30\x94\x6b\xcb\xa2\x4d\x03\x0e\x4d\xa7\x3b\x96\xa1\xce\x32\x3e\xc3\x8a\x68\x93\x7a\x02\xe0\x58\x56\xc7\x0d\xb2\x13\x43\x34\xb5\x82\x9d\x06\xfb\x58\x96\xef\x13\x92\x84\x0c\x68\x03\x4a\x8b\xe7\xea\x4e\x84\x15\x5a\x08\x53\x42\x05\xab\x84\x14\xcc\x0a\x4e\x07\x84\x70\x4b\xd1\x34\x1a\x2b\xc0\x1d\x8a\x9f\x78\xf8\x63\x06\xd4\xd4\x67\xf9\xb1\x42\x95\x63\x12\x6b\x28\x2c\x45\x2e\x8f\x86\x3a\xcb\x39\xa5\x7e\x8e\xd5\xa2\x19\xc5\xef\xb6\x22\xf6\xac\xfa\xcf\xcf\x53\x14\x87\xf1\x51\xe7\xbf\xa9\x88\x81\xa5\x74\x53\xc9\x2f\xf2\x77\x6f\x02\xa6\x50\xca\x55\x44\xe5\xb8\xb0\xe2\xdd\x7b\xfb\x86\x7b\xaa\x33\x6e\x5d\x95\x68\xc7\x20\x71\x54\xb5\xd7\x6f\x60\x8c\xa2\xa9\xf2\x7d\xfb\x80\x28\x68\x0c\xb6\xd3\x7d\xc0\x42\x59\x77\xff\xe1\x34\x51\x46\xe1\x66\xfc\x6d\xa6\x86\x7e\x55\xab\xdd\x16\x56\x09\x87\x89\x53\x70\x85\xca\x57\x77\x18\x79\x70\x36\x6c\xd0\xaf\x9f\xf1\x51\xff\x69\x55\x3c\x36\xdf\x9c\xa3\x6f\xa9\x07\xcc\x8d\xce\x06\x94\x3e\x01\xb9\x65\x9e\x3e\x03\xbb\x7a\x31\x64\xcb\x99\xd6\x2e\x18\xb6\xad\x59\x66\xa9\xc1\xe5\x2d\xb2\xd1\xa1\x9a\x2d\x20\xd8\xc2\xe7\xc5\x79\xe1\x8a\xa9\x4f\x05\x59\x39\x5a\xd5\x4a\xc9\x3b\x12\xe4\xf4\x5e\x28\x3b\x47\x53\x8e\x46\xdd\xd1\xfb\xbb\xce\x70\xba\x1f\x9d\xa3\x63\x40\xcf\xf8\xbb\x9f\x43\xf9\x4e\x71\x76\x9c\xd9\x3e\x77\x1c\xe4\xb8\x7d\xfd\x38\x9b\x6b\xe3\x21\xc5\x27\x22\xff\x8e\x32\x7c\xed\x8f\x82\xc5\xef\x80\xf8\x22\x80\x38\xe0\x0b\x43\x28\xf4\x95\xf3\xe8\xf0\x13\xd2\x66\x06\x3e\x54\x96\x2c\x7e\xb2\x63\x2a\xad\x4e\x07\xe6\xaa\xd8\xe5\xb6\xe6\xfa\xa5\x98\x91\x51\x24\x54\x0e\xd4\x56\xda\x2c\xc8\x15\x9b\x11\xd9\x6b\xb8\xd5\x4a\x8c\x4e\x21\x4f\x51\x6d\xa8\xac\xcf\xf7\x91\xeb\xea\x33\x56\x14\xc1\x6c\xed\xb5\x69\x24\x9d\xeb\x7e\x05\x12\x8e\x93\x29\x2e\x91\x53\x9c\x71\xca\xb2\x7e\x9a\x78\xbe\x5b\x33\xd7\x86\x66\x0b\xaf\x5e\x0e\xdc\x01\x12\xa7\xfb\x74\x8e\x33\xc3\xe1\xe8\xfe\xa0\x24\xaf\x3c\x3a\x5c\x61\x86\xa5\x61\xa7\xde\xb0\xdf\x4e\x29\xe4\x6a\x59\xc3\x96\x0e\x5f\xf8\x12\x47\xdf\xa3\x66\x98\xf3\x94\x54\x94\x6b\x56\x63\xdb\x48\xb6\x60\x13\x5d\xa4\x91\x03\x09\xc2\x12\xd3\x82\x20\xe5\x05\x01\xe7\x37\xb9\x36\x52\x55\xd6\x9c\xa6\xd5\x0e\x36\x52\x60\x0a\xf7\x6f\x26\x6e\xf9\x5f\xea\x0d\x5a\xd7\x28\xcd\x30\x5c\x38\xac\xd1\xbf\x85\x8c\xc2\x14\xc2\xc2\x8a\xce\xc8\xd8\x0a\x81\x38\x4b\xa9\x6a\xa7\xb2\x42\x71\xe8\xba\xb9\x63\x0a\xd8\xdc\xf0\x12\x85\xa6\x0b\x1a\xbc\xa3\xfb\x9b\xa0\x92\x1e\x16\xb4\xde\xb4\x74\xd6\x26\xc0\x4a\x34\x64\x95\xe9\xc7\x05\xee\x8b\x2d\x23\x87\xc3\x8e\x6b\x03\x1b\xd1\xbd\x0c\x98\x6c\x7b\x6a\xcc\x66\x88\xd8\xf9\xb2\xc8\x31\x57\xa1\x59\x6f\x0c\x47\xa9\xc8\x57\xa9\xdd\x44\x5a\xcb\x03\xf5\x84\xef\x68\x07\xf5\x3c\xaa\x5f\x9e\x83\x6e\xfd\x8d\x6b\xea\x3b\xa3\xfa\x6d\x7d\x28\x4e\x9f\x51\xb7\xbc\x6d\xd5\x2a\xa7\xb6\x13\x7d\xaf\x76\x97\x67\xd2\xfb\xf6\x3c\x80\xcb\x51\x9c\x5d\x5b\x07\xbf\x7d\xa8\x7e\xfa\x9c\xd0\xbc\x08\x87\xb3\xc1\xeb\xee\x7b\xef\xe9\x7e\xfa\xee\x7c\x34\x2e\xf2\x34\x1a\xe0\x13\x6c\x78\xe0\xb1\xf3\xf2\xac\x7a\x8b\xaa\x1b\x4a\xaa\x89\xd7\x40\x24\xa9\x8b\xe7\xf3\x81\xbf\xe7\xd5\x73\x4f\x07\xd5\x2a\xde\x73\x6d\x2d\xbb\xd2\xc0\x70\x9c\x08\x28\xbd\x3a\xc5\x59\xad\x77\x97\xcb\x73\xcd\xfd\x1c\x36\x4f\x5a\x3e\xf5\xfe\xed\x97\xbd\x2e\x99\x72\xef\x83\x06\x0c\x63\xfe\xf0\xbb\xc8\xfe\x51\x77\x73\xe7\xae\xdd\x37\x76\xdb\xd0\x79\x2e\xbe\x3d\xab\x1b\xb7\x56\x3a\xd9\xf1\xd1\xb6\xa7\x7b\x8d\x33\x4a\xbf\x4a\x9f\xeb\xca\x0c\x55\xf5\xba\x08\xa9\x97\xa3\xdf\x0b\xae\xfa\x1d\xcd\xa2\xac\x9d\x7b\x79\xec\xc7\xf1\xee\x23\xef\x37\x3f\x55\x3f\xe7\x15\xa6\x19\x23\x3c\xe7\x55\x1d\xd0\x1c\x6b\xd1\x7f\x19\x69\xd8\x4b\x0d\x12\x9f\x67\xe8\xbb\xdb\x78\x9e\xa1\x69\x61\x5d\x0f\x77\xc6\x75\xab\x0f\x18\x7e\xe5\x6d\xce\x62\x9e\xee\x2d\xb9\xd9\x47\x6a\x33\xc8\x85\x95\x50\xdc\xfa\x2a\x42\xcb\x7f\xac\xe4\xcf\x2f\xf6\xc8\xe0\xfb\xe4\xa3\x64\x40\x6b\x57\xda\x9c\x34\x41\xed\x99\x99\x2d\x7e\xb2\x9b\x0b\xe2\xd8\xfa\x2b\xdb\x9c\x8f\x5d\x83\xef\xec\x95\x8e\xc8\x5d\x69\xad\xd2\x35\x20\x64\x58\x96\x5d\xf3\x12\x6e\x73\xa6\x34\x72\xcb\x3e\x9f\x50\x34\x30\x29\xf3\xb1\xec\xe9\x8b\xe4\x6e\xa2\x38\x5c\x4b\x38\x4b\x94\xc2\x6d\x8c\xb0\xa7\x9d\xdf\x4d\x65\x8e\xde\xe6\x14\xf2\xbc\xfa\xe5\x44\xf5\x35\xbe\xfa\x33\x6f\xbe\xee\x4f\x3a\xea\xf1\xc6\x9d\x0e\x1b\x5f\xb8\xe8\xf8\xe8\x41\x1b\xaa\xbe\xe7\x72\xf4\xcf\x5b\x4e\xff\x45\x8a\x93\xe9\xea\xf1\xea\x7f\x01\x00\x00\xff\xff\x0c\x7e\xc3\x96\x71\x33\x00\x00")

func apiSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apiSwaggerJson,
		"api.swagger.json",
	)
}

func apiSwaggerJson() (*asset, error) {
	bytes, err := apiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.swagger.json", size: 13169, mode: os.FileMode(420), modTime: time.Unix(1537457358, 0)}
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
	"api.swagger.json": apiSwaggerJson,
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
	"api.swagger.json": &bintree{apiSwaggerJson, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
