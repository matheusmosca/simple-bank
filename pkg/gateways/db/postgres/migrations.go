// Code generated by go-bindata. DO NOT EDIT.
// sources:
// migrations/000001_account_schema.up copy.sql (277B)
// migrations/000001_initial.down.sql (38B)
// migrations/000001_initial.up.sql (277B)
// migrations/000002_account_schema.down.sql (38B)

package postgres

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __000001_account_schemaUpCopySql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\x31\x4f\xc3\x30\x10\x46\xe7\xf8\x57\x7c\x63\x2b\x75\x08\x88\x4e\x9d\xdc\x72\x05\x8b\xd8\x29\xce\x59\x50\x96\xea\x70\x5c\x11\xa9\x04\x14\x1c\x7e\x3f\xa2\x43\xa3\x8e\xa7\x77\xef\x9e\x74\x6b\x7a\x30\x6e\xa5\xd4\xc6\x93\x66\x02\xeb\x75\x45\x30\x5b\xb8\x9a\x41\xaf\xa6\xe1\x06\x12\xe3\xd7\xd8\x67\x35\x53\x45\xd7\x62\x1c\xbb\x16\x3b\x6f\xac\xf6\x7b\x3c\xd1\x7e\xa1\x8a\x5e\x3e\x13\x7e\x65\x88\x1f\x32\xcc\x6e\x97\xcb\xf9\xd9\x76\xa1\xaa\x16\xaa\x88\xdf\xc7\x0b\xbb\xb9\x9b\x10\x82\x33\xcf\x81\x16\xaa\xf8\x49\x71\x48\x79\x5a\x2a\xcb\xab\x03\xef\x72\x92\x3e\x26\x74\x7d\x46\x9b\x8e\x32\x9e\x32\xca\xab\xc2\x90\x24\xa7\xf6\x20\x19\x6c\x2c\x35\xac\xed\x0e\x2f\x86\x1f\xcf\x23\xde\x6a\x47\x53\xf5\x9e\xb6\x3a\x54\x8c\x4d\xf0\x9e\x1c\x1f\x2e\x86\x9a\xff\x7f\xa1\xb6\xd6\xf0\xea\x2f\x00\x00\xff\xff\x77\x35\xfc\x1e\x15\x01\x00\x00")

func _000001_account_schemaUpCopySqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_account_schemaUpCopySql,
		"000001_account_schema.up copy.sql",
	)
}

func _000001_account_schemaUpCopySql() (*asset, error) {
	bytes, err := _000001_account_schemaUpCopySqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_account_schema.up copy.sql", size: 277, mode: os.FileMode(0644), modTime: time.Unix(1619986388, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x99, 0xc0, 0x2e, 0x38, 0x1, 0xed, 0xe5, 0x38, 0x50, 0x85, 0xe8, 0x3b, 0x0, 0x71, 0x66, 0x68, 0xd3, 0x10, 0x2d, 0x88, 0x29, 0x26, 0x4a, 0xd6, 0x8e, 0xb, 0x17, 0xbd, 0x38, 0x3e, 0xfd, 0x7c}}
	return a, nil
}

var __000001_initialDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x4c\x4e\xce\x2f\xcd\x2b\x51\x70\x76\x0c\x76\x76\x74\x71\xb5\xe6\x02\x04\x00\x00\xff\xff\x28\x19\x2b\xd0\x26\x00\x00\x00")

func _000001_initialDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initialDownSql,
		"000001_initial.down.sql",
	)
}

func _000001_initialDownSql() (*asset, error) {
	bytes, err := _000001_initialDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_initial.down.sql", size: 38, mode: os.FileMode(0644), modTime: time.Unix(1620163736, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xac, 0x90, 0x9f, 0xce, 0x7e, 0x76, 0xaa, 0xad, 0x5d, 0xd6, 0x12, 0xc7, 0x31, 0xae, 0xb1, 0xdb, 0x4, 0xa8, 0xff, 0x98, 0xb0, 0x97, 0x19, 0x93, 0x4a, 0x29, 0x44, 0xaa, 0x3b, 0x67, 0x26, 0x9e}}
	return a, nil
}

var __000001_initialUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\x31\x4f\xc3\x30\x10\x46\xe7\xf8\x57\x7c\x63\x2b\x75\x08\x88\x4e\x9d\xdc\x72\x05\x8b\xd8\x29\xce\x59\x50\x96\xea\x70\x5c\x11\xa9\x04\x14\x1c\x7e\x3f\xa2\x43\xa3\x8e\xa7\x77\xef\x9e\x74\x6b\x7a\x30\x6e\xa5\xd4\xc6\x93\x66\x02\xeb\x75\x45\x30\x5b\xb8\x9a\x41\xaf\xa6\xe1\x06\x12\xe3\xd7\xd8\x67\x35\x53\x45\xd7\x62\x1c\xbb\x16\x3b\x6f\xac\xf6\x7b\x3c\xd1\x7e\xa1\x8a\x5e\x3e\x13\x7e\x65\x88\x1f\x32\xcc\x6e\x97\xcb\xf9\xd9\x76\xa1\xaa\x16\xaa\x88\xdf\xc7\x0b\xbb\xb9\x9b\x10\x82\x33\xcf\x81\x16\xaa\xf8\x49\x71\x48\x79\x5a\x2a\xcb\xab\x03\xef\x72\x92\x3e\x26\x74\x7d\x46\x9b\x8e\x32\x9e\x32\xca\xab\xc2\x90\x24\xa7\xf6\x20\x19\x6c\x2c\x35\xac\xed\x0e\x2f\x86\x1f\xcf\x23\xde\x6a\x47\x53\xf5\x9e\xb6\x3a\x54\x8c\x4d\xf0\x9e\x1c\x1f\x2e\x86\x9a\xff\x7f\xa1\xb6\xd6\xf0\xea\x2f\x00\x00\xff\xff\x77\x35\xfc\x1e\x15\x01\x00\x00")

func _000001_initialUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initialUpSql,
		"000001_initial.up.sql",
	)
}

func _000001_initialUpSql() (*asset, error) {
	bytes, err := _000001_initialUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_initial.up.sql", size: 277, mode: os.FileMode(0644), modTime: time.Unix(1620163738, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x99, 0xc0, 0x2e, 0x38, 0x1, 0xed, 0xe5, 0x38, 0x50, 0x85, 0xe8, 0x3b, 0x0, 0x71, 0x66, 0x68, 0xd3, 0x10, 0x2d, 0x88, 0x29, 0x26, 0x4a, 0xd6, 0x8e, 0xb, 0x17, 0xbd, 0x38, 0x3e, 0xfd, 0x7c}}
	return a, nil
}

var __000002_account_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x4c\x4e\xce\x2f\xcd\x2b\x51\x70\x76\x0c\x76\x76\x74\x71\xb5\xe6\x02\x04\x00\x00\xff\xff\x28\x19\x2b\xd0\x26\x00\x00\x00")

func _000002_account_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_account_schemaDownSql,
		"000002_account_schema.down.sql",
	)
}

func _000002_account_schemaDownSql() (*asset, error) {
	bytes, err := _000002_account_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_account_schema.down.sql", size: 38, mode: os.FileMode(0644), modTime: time.Unix(1619901595, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xac, 0x90, 0x9f, 0xce, 0x7e, 0x76, 0xaa, 0xad, 0x5d, 0xd6, 0x12, 0xc7, 0x31, 0xae, 0xb1, 0xdb, 0x4, 0xa8, 0xff, 0x98, 0xb0, 0x97, 0x19, 0x93, 0x4a, 0x29, 0x44, 0xaa, 0x3b, 0x67, 0x26, 0x9e}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"000001_account_schema.up copy.sql": _000001_account_schemaUpCopySql,
	"000001_initial.down.sql":           _000001_initialDownSql,
	"000001_initial.up.sql":             _000001_initialUpSql,
	"000002_account_schema.down.sql":    _000002_account_schemaDownSql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"000001_account_schema.up copy.sql": {_000001_account_schemaUpCopySql, map[string]*bintree{}},
	"000001_initial.down.sql": {_000001_initialDownSql, map[string]*bintree{}},
	"000001_initial.up.sql": {_000001_initialUpSql, map[string]*bintree{}},
	"000002_account_schema.down.sql": {_000002_account_schemaDownSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
