// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package views generated by go-bindata.// sources:
// web/views/errors/404.html
// web/views/errors/500.html
// web/views/index.html
// web/views/layouts/footer.html
// web/views/layouts/header.html
// web/views/layouts/layout.html
package views

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _errors404Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xc1\x6a\xc3\x30\x0c\x86\xef\x79\x0a\xcd\xe7\x79\x69\xa0\x87\x1d\xec\xc0\xe8\x5a\xd8\x65\xdb\xa1\x85\xed\xe8\xda\x3f\xb5\xc0\xb1\xb3\x54\x4d\xd9\xdb\x8f\x34\x1d\x74\x3b\x09\x7d\xfa\x3f\x81\x64\xee\x9e\xdf\x56\xdb\xcf\xf7\x35\x45\xe9\x52\x5b\x99\xa9\x50\x72\xf9\x60\x15\xb2\x9a\x00\x5c\x68\x2b\x22\x22\xd3\x41\x1c\xf9\xe8\x86\x23\xc4\xaa\xdd\x76\xa3\x1f\xd5\xed\x28\x8a\xf4\x1a\x5f\x27\x1e\xad\xfa\xd0\xbb\x27\xbd\x2a\x5d\xef\x84\xf7\x09\x8a\x7c\xc9\x82\x2c\x56\xbd\xac\x2d\xc2\x01\x7f\xcc\xec\x3a\x58\x35\x32\xce\x7d\x19\xe4\x26\x7c\xe6\x20\xd1\x06\x8c\xec\xa1\x2f\xcd\x3d\x71\x66\x61\x97\xf4\xd1\xbb\x04\xdb\x3c\x2c\x7e\x57\x09\x4b\x42\xbb\x5c\x2c\xe9\xb5\x08\x6d\xca\x29\x07\x53\xcf\xb0\x32\xf5\x7c\x88\xd9\x97\xf0\x7d\xcd\xc7\xa6\x35\x1e\x59\x30\xfc\x97\xae\xd4\xd4\xb1\x99\xd4\xd9\x31\xf5\xe5\x47\x3f\x01\x00\x00\xff\xff\xf0\x70\x71\x97\x33\x01\x00\x00")

func errors404HtmlBytes() ([]byte, error) {
	return bindataRead(
		_errors404Html,
		"errors/404.html",
	)
}

func errors404Html() (*asset, error) {
	bytes, err := errors404HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "errors/404.html", size: 307, mode: os.FileMode(420), modTime: time.Unix(1670244046, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errors500Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x4e\xc3\x30\x0c\x86\xef\x79\x0a\x93\xf3\x42\xd9\x8d\x43\x52\x09\x8d\x21\x71\x1a\x87\x4d\x82\xa3\xd7\x5a\x8d\xa5\xd4\x29\xad\xb7\x6a\x6f\x8f\xba\x0c\x69\xdc\x38\x25\xbf\xfc\xf9\xb3\x6c\xff\xf0\xba\xdb\xec\xbf\x3e\xb6\x10\xb5\x4f\xb5\xf1\xcb\x03\x09\xa5\x0b\x96\xc4\xd6\xc6\xf8\x48\xd8\xd6\x06\x00\xc0\xf7\xa4\x08\x4d\xc4\x71\x22\x0d\xf6\xb0\x7f\x73\xcf\xf6\xbe\x14\x55\x07\x47\xdf\x27\x3e\x07\xfb\xe9\x0e\x2f\x6e\x93\xfb\x01\x95\x8f\x89\x2c\x34\x59\x94\x44\x83\x7d\xdf\x06\x6a\x3b\xfa\xd3\x29\xd8\x53\xb0\x67\xa6\x79\xc8\xa3\xde\xc1\x33\xb7\x1a\x43\x4b\x67\x6e\xc8\x5d\xc3\x0a\x58\x58\x19\x93\x9b\x1a\x4c\x14\xd6\x8f\x4f\xbf\x2a\x65\x4d\x54\xef\x4e\xc3\x04\x53\xee\x49\x23\x4b\x07\x33\x89\xc2\x3c\x66\xe9\x56\xa0\xe3\x05\xb0\x43\x16\x5f\x15\xd6\xf8\xaa\xec\x67\xfc\x31\xb7\x97\x9b\x27\xae\xcb\xe7\x1a\x1a\x12\xa5\xf1\x9f\xd6\x1b\x5c\x34\xd5\xe2\xf1\x55\x11\x2f\x93\x96\x13\xff\x04\x00\x00\xff\xff\x89\x94\x40\x58\x72\x01\x00\x00")

func errors500HtmlBytes() ([]byte, error) {
	return bindataRead(
		_errors500Html,
		"errors/500.html",
	)
}

func errors500Html() (*asset, error) {
	bytes, err := errors500HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "errors/500.html", size: 370, mode: os.FileMode(420), modTime: time.Unix(1670244046, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\x2c\x2a\xc9\x4c\xce\x49\x55\x48\xce\x49\x2c\x2e\xb6\x55\xca\x4d\x2c\xca\x4e\xc9\x2f\xcf\xd3\x4d\xca\x4f\xa9\x54\xb2\xe3\x52\x50\x50\x50\xa8\xae\xd6\x73\x84\xa8\xaa\xad\xe5\xb2\xd1\x87\xea\xb0\x03\x04\x00\x00\xff\xff\x82\x8b\x94\x11\x3b\x00\x00\x00")

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

	info := bindataFileInfo{name: "index.html", size: 59, mode: os.FileMode(420), modTime: time.Unix(1670244046, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutsFooterHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func layoutsFooterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_layoutsFooterHtml,
		"layouts/footer.html",
	)
}

func layoutsFooterHtml() (*asset, error) {
	bytes, err := layoutsFooterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layouts/footer.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(1670244046, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutsHeaderHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x41\x6a\xc4\x30\x0c\x05\xd0\x7d\x4f\x21\xb4\x17\xd9\xb7\x8e\xef\xa2\x49\x64\x47\xa9\x62\x0d\xb6\x63\x98\xdb\x97\x52\x3a\x84\x42\x61\x96\x12\xff\xff\x17\x56\x1d\xb0\x18\xb7\x36\xe3\xcd\xfd\x93\x36\xe1\x55\x2a\x82\xae\x7f\x1e\xd5\x4d\x66\x2c\x3c\x34\x73\x57\x2f\x18\xdf\x00\x00\x02\x3f\xeb\xbd\xc0\xfd\x34\x23\x93\xd4\x61\x6f\xd4\xdd\xed\xc6\x95\x78\xf9\xce\x43\xf7\x9c\x4d\xa8\x9d\xc7\xc1\xf5\x81\xb0\x55\x49\x33\xee\x3c\xb8\x2d\x55\xef\xfd\xfd\x03\x63\xd0\xdf\xb5\xc4\x90\x98\xd8\x34\x17\xda\xcf\xd6\x35\x3d\x30\x86\x49\x63\x98\xf8\x5f\xb9\x6a\xde\x7e\xe8\x4d\x0e\x79\xc2\x97\xe3\x35\xb6\x9d\x85\xfc\xc2\x85\x69\xd5\x11\xbf\x02\x00\x00\xff\xff\x39\xdf\x3f\x2c\x2c\x01\x00\x00")

func layoutsHeaderHtmlBytes() ([]byte, error) {
	return bindataRead(
		_layoutsHeaderHtml,
		"layouts/header.html",
	)
}

func layoutsHeaderHtml() (*asset, error) {
	bytes, err := layoutsHeaderHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layouts/header.html", size: 300, mode: os.FileMode(420), modTime: time.Unix(1670244046, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutsLayoutHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x57\x51\x8f\xdb\x36\x12\x7e\xd6\x02\xfb\x1f\x18\x1e\xd0\xb5\x81\x95\xd4\xed\xd3\x21\x6b\x39\x48\x36\xb9\x5e\x80\x22\xed\xa5\x29\x70\x87\x20\x38\x8c\xc5\xb1\x44\x2f\x45\xaa\x24\x65\xaf\xa2\xea\xbf\x1f\x48\x4a\x96\x9c\x78\x83\xe4\xf2\x64\x91\x33\xf3\xcd\x70\x38\x33\x1f\xbd\x7a\xf2\xf2\xd7\xbb\x77\xff\xf9\xed\x15\x29\x6d\x25\xd6\x97\x17\x2b\xf7\x4b\x04\xc8\x22\xa3\x28\xa9\xdf\x41\x60\xeb\xcb\x8b\x68\x55\xa1\x05\x92\x97\xa0\x0d\xda\x8c\xfe\xf1\xee\x1f\xf1\xdf\xe9\x24\x28\xad\xad\x63\xfc\xb3\xe1\xfb\x8c\xfe\x3b\xfe\xe3\x79\x7c\xa7\xaa\x1a\x2c\xdf\x08\xa4\x24\x57\xd2\xa2\xb4\x19\x7d\xfd\x2a\x43\x56\xe0\xcc\x4e\x42\x85\x19\xdd\x73\x3c\xd4\x4a\xdb\x99\xea\x81\x33\x5b\x66\x0c\xf7\x3c\xc7\xd8\x2f\xae\x09\x97\xdc\x72\x10\xb1\xc9\x41\x60\x76\x93\xfc\x18\x80\x2c\xb7\x02\xd7\x5d\x47\x92\x7f\x22\xb0\x77\x6e\x45\xfa\x7e\x95\x86\xfd\xcb\x8b\xa8\xeb\xf8\x96\x24\xcf\x25\x88\xf6\x23\xea\xe4\x67\xa5\x0a\x81\xcf\x59\xdf\x3b\x6b\x93\x6b\x5e\x5b\x02\xa6\x95\x39\x31\x3a\xcf\xa8\x3b\x8b\x79\x9a\xa6\x5d\x97\xfc\xa6\xd5\x43\x3b\x33\x48\x6b\x28\x10\x58\xba\x33\x29\x30\xb3\x69\x0b\x2f\x49\x76\xe6\x59\x2e\xb8\x8b\xbb\xeb\xce\xf9\xa1\x97\x17\x51\x14\xe5\x5a\x19\xa3\x34\x2f\xb8\xcc\x28\x48\x25\xdb\x4a\x35\x86\xae\x57\x69\x88\x21\x84\x8a\x72\x08\x4c\x70\x79\x4f\x34\x8a\x8c\x1a\xdb\x0a\x34\x25\xa2\xa5\x84\xb3\x8c\xda\x12\x2b\x8c\x73\x63\x28\x29\x35\x6e\x33\x9a\x1a\x0b\x96\xe7\x69\x6e\x4c\x5a\x70\x5b\x36\x9b\xb8\x02\x7d\xcf\xd4\x41\x3a\xb5\x94\x81\xbe\x4f\x9c\xfe\xfa\x71\xe0\xb3\x48\x1b\xa5\xee\x63\xef\x2e\xf5\xba\xdf\x0b\x72\xc0\x8d\xe1\xf6\xbb\x61\xf2\x12\x6a\x8b\x3a\xde\x2a\xc1\xbe\xfb\x5c\xb5\xe0\xd6\xa2\xfe\x76\x9c\x92\x17\xa5\xe0\x45\x69\x07\x24\xb8\xb9\x69\xe3\xff\x2f\xd5\x15\x70\x39\x5a\xf9\x2a\x20\xae\x62\x7f\xe6\x16\xc4\x7d\x72\xe7\x4b\xeb\xf5\x4b\xf2\xa5\xba\x38\x7b\x4c\x10\xf7\xc3\xcf\x31\xa4\x63\x81\x7d\xde\x17\x2f\x80\xb3\x26\xf8\x78\x12\xc7\xc4\x2f\x09\x38\xa9\xe5\xb9\x21\x71\xbc\x9e\xfa\xc5\x7d\x46\x7b\xd0\xe4\xbf\x65\x65\x49\x16\x7e\xfe\xfa\x8b\xbc\xff\x70\xeb\x24\x8b\x6d\x23\x73\xcb\x95\x5c\x2c\x49\xe7\x36\x08\x71\xca\x65\x45\x32\xc2\x54\xde\x54\x28\x6d\x92\x6b\x04\x8b\xaf\x04\xba\xd5\x82\x06\x60\xba\xbc\x0d\xfa\x65\x95\x18\x9d\x93\x8c\x1c\xfb\xb1\xac\x92\x8d\x8b\x29\xc9\x55\xe5\x16\x3b\xf3\x6c\xde\x6e\x43\xf8\xf4\x76\xf2\x67\xe6\xee\x0a\xb4\x83\x2f\xf3\xa2\x7d\x07\xc5\x1b\xa8\x70\xf2\xfa\xfe\xc7\x0f\x83\xa1\x49\x6a\xd0\x28\xed\x1b\xc5\x30\xe1\xd2\xa0\xb6\x2f\x70\xab\x34\x2e\xca\xea\x9a\x98\x10\x5f\xbf\x5c\xf8\x8f\xb3\xbd\xfb\xd8\xc8\x99\x72\x1b\xd6\xc4\x42\x41\x16\x85\x85\x22\xd9\x99\xe5\x69\x82\xbf\x66\x20\xb5\x7d\x9f\x3a\xeb\x74\x67\x9e\x71\x76\x66\xf6\xf4\xfd\xe9\x74\x99\x5f\xde\x81\x4b\xa6\x0e\x09\x03\x0b\xbf\x40\x8b\x9a\x64\xe4\xb3\xad\xe9\x42\xc7\xfb\x24\xce\xdf\x62\xd9\x1d\x75\x92\xba\x31\xe5\x02\x74\xe1\x93\x6c\x96\xb7\xee\x8c\x91\xd7\xba\xda\x99\xab\x6b\x22\xf1\x40\x5e\x82\xc5\xc5\xd2\x25\x6c\x12\xe6\x4a\x6e\x79\x71\x75\x4d\xae\xce\xc5\x7d\xf5\xa5\xec\xae\xd2\x81\x95\x56\x1b\xc5\x5a\x7f\x30\xc6\xf7\x24\x17\x60\x4c\x46\x5d\x6f\x93\xad\x92\x36\x36\xfc\x23\xc6\x3f\x85\xef\x2d\x54\x5c\xb4\xf1\x0d\xc9\x95\x50\x3a\x74\x6c\xfc\x93\x6f\x89\xcf\xac\x63\xd3\x54\x15\xe8\x36\x48\xa3\x55\x8e\xd2\xa2\x0e\x8b\x68\x55\x8f\xaa\x42\x15\x8a\x12\xad\x04\x0e\xdf\xeb\x15\x0c\x7d\x48\x3d\x1d\x4d\x54\x04\xeb\x55\x5a\x0f\x68\xe9\x1c\x6e\x25\x61\x3f\x40\x48\xd8\xf3\x02\x5c\x92\xe9\xe8\xaa\x11\xa3\xaf\xd3\x88\xa2\xa8\xeb\x34\xc8\x02\x49\xf2\x06\xf6\xbe\xac\xc2\x26\xb1\x58\xd5\x02\x2c\x12\x07\x67\x12\xc7\xe7\x94\x24\x33\x8d\x91\x5e\x3c\xbe\xe0\x23\x3e\xe3\x7b\xce\x50\xbb\x72\x11\x7c\x74\x9f\x36\x62\x8c\x59\xc2\x3e\xa4\x2a\x65\x7c\x7f\x3e\x69\xee\x2e\xc6\x8c\x9d\xc8\x58\x1b\x73\x29\x1d\x78\xc0\xed\x3a\xa2\x51\x32\xd4\x84\x0a\x68\x55\x63\x8d\xbf\x4f\xd4\x43\xb8\x63\x7c\x73\x14\xc7\xbb\xf1\x41\x43\x5d\xa3\xa6\xc4\xc2\x86\x4b\x86\x0f\x19\x8d\x6f\xc6\x1b\x70\x23\xf4\x98\x9e\xcf\x4c\x4f\x02\x88\x42\x13\xbe\x0d\x41\xd8\x12\x49\xde\x68\xd7\xf2\x53\xfa\x4a\xd4\x38\x34\xe4\x98\xda\x96\xa3\x60\xc7\xe0\x06\x1f\x8e\x8e\xc3\x84\x75\xa9\x1b\x73\xe3\xc5\x63\x2d\xfc\xcd\x67\x27\x9c\x90\x8e\x41\x4d\x97\x1d\x17\x2a\xb6\xaa\xa6\xeb\xd5\xf1\x32\xb6\x40\xb6\x10\x83\xd6\xea\x10\x37\x4e\x92\xf2\xb5\xab\xa1\xf1\x70\x33\x3f\xb3\xef\xd9\xd5\x4c\x5f\x5b\xa5\x5c\xad\x9d\x49\x79\x90\x1c\x53\xbe\x1a\x36\x7c\x37\x0d\xf6\xab\x74\xe8\xaf\x71\x26\xf9\x69\x34\x72\xcc\x6e\xc6\x80\xc9\xee\xf4\x1d\xf3\x98\xc1\xee\x5f\x0d\xea\xf6\x6b\xb5\x4f\x68\xfe\x2b\x6d\x8e\x74\xfe\x95\xfa\x9e\x79\x3f\xd1\xfd\x02\xfd\x3e\x06\x33\xb0\x6c\xf5\x28\xd8\x30\xb8\xfd\xe3\xec\x79\xcd\xbf\x08\x06\x3b\x78\x28\x95\x3a\x0b\x37\x1b\xe1\x50\x26\xb5\x43\x5d\x78\x86\x55\xf2\x2d\xfe\xd9\xa0\xb1\x4f\xc9\x22\xcc\xd6\xeb\x12\x24\x13\xa8\x97\x24\x5b\x07\x16\x8e\x04\x5a\x02\x35\x27\x99\x9f\xba\x9f\xc4\xd4\xf7\x57\xb7\x47\xad\x5c\x69\x73\x46\xed\x4e\x69\x73\xd4\xe3\xdb\x85\x03\x7b\x92\x91\x46\x32\xdc\x72\x89\x8c\xfc\xf0\x03\x19\xf6\x28\x1d\xb8\x3f\x8a\x42\x3c\x49\xa3\x05\xc9\xc8\xb4\x48\x34\xd6\x02\x72\x5c\x50\xa8\x79\x12\x1e\xae\x8e\xdc\xe9\xb5\xc3\x08\x44\x1b\xf5\xa3\x2b\x1f\xd1\xa7\xbe\xc6\xcd\x6f\x71\xe6\x6c\x62\x90\xed\xc1\x35\x79\x02\x1f\xab\xe4\xa0\xf4\x3d\x6a\x93\x30\xdc\xd3\x6b\x8f\x79\xe2\x7c\xc8\x63\x22\xf1\xc1\x0e\xb9\x1d\x5e\x01\x97\x17\x51\xef\x3e\x4f\xee\x7b\x18\xb1\xb3\x9b\xca\x95\x34\x96\x84\x1a\x21\x99\xa7\xc4\x50\x5c\xe1\xea\xf2\xa1\xc2\x9e\xfa\x74\x7f\x52\x76\x7d\x7f\x75\x3d\x29\xfd\x8e\xb9\x46\x7b\x46\x31\x08\x46\x65\x8d\xb5\x3a\x51\x7a\x8b\xb5\x1a\x85\xea\x20\x51\x9f\x48\x7f\x75\x3b\xa3\x18\x58\xc5\xe5\x53\xf2\xfe\x8c\xfc\x83\x57\xe0\xec\xc4\xf8\x35\x1b\x2d\x05\x6c\x50\x18\x67\x1a\x8e\x7a\xf5\xc1\xe7\xc7\xb3\xfe\xd0\x20\x61\x0a\x2d\x46\x85\xe5\x23\xa9\x4b\x87\x3f\xa4\xbe\x7d\xc2\x75\x9f\x30\x9a\x57\x9a\xd8\x6b\x18\x14\xc4\xbd\x24\x72\xcb\xf7\xfe\xe1\xe3\x06\xd9\x38\x83\xbb\x2e\xf9\x85\xcb\xfb\x61\x3b\xea\xba\xe4\xf7\x52\x1d\xdc\x0b\x30\xcc\xf3\xf0\x5c\x7b\x6d\x5e\x72\xdd\xf7\xd3\x20\xc6\x87\x3c\xb6\x9a\x17\x05\x6a\xb2\x85\x30\x88\x67\xff\xd0\xc2\x48\x0e\xb6\x77\x25\x17\x4c\xa3\x0c\x92\x89\xb8\x41\x5b\x9e\x0b\x34\xa3\xe3\x81\xb6\x4f\xd4\x1f\x27\x6e\x32\xca\x27\x9f\x81\x93\x67\xa9\xf2\x7c\x3d\xac\xff\x17\x00\x00\xff\xff\x80\x08\x6e\x4f\xd2\x0f\x00\x00")

func layoutsLayoutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_layoutsLayoutHtml,
		"layouts/layout.html",
	)
}

func layoutsLayoutHtml() (*asset, error) {
	bytes, err := layoutsLayoutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layouts/layout.html", size: 4050, mode: os.FileMode(493), modTime: time.Unix(1671626368, 0)}
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
	"errors/404.html":     errors404Html,
	"errors/500.html":     errors500Html,
	"index.html":          indexHtml,
	"layouts/footer.html": layoutsFooterHtml,
	"layouts/header.html": layoutsHeaderHtml,
	"layouts/layout.html": layoutsLayoutHtml,
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
	"errors": &bintree{nil, map[string]*bintree{
		"404.html": &bintree{errors404Html, map[string]*bintree{}},
		"500.html": &bintree{errors500Html, map[string]*bintree{}},
	}},
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"layouts": &bintree{nil, map[string]*bintree{
		"footer.html": &bintree{layoutsFooterHtml, map[string]*bintree{}},
		"header.html": &bintree{layoutsHeaderHtml, map[string]*bintree{}},
		"layout.html": &bintree{layoutsLayoutHtml, map[string]*bintree{}},
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
