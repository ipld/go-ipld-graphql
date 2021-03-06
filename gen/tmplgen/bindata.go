// Code generated for package tmplgen by go-bindata DO NOT EDIT. (@generated)
// sources:
// tmpl/common.tmpl
// tmpl/list.tmpl
// tmpl/map.tmpl
// tmpl/scalar.tmpl
// tmpl/scalar_parse.tmpl
// tmpl/scalar_parseliteral.tmpl
// tmpl/scalar_serialize.tmpl
// tmpl/struct.tmpl
// tmpl/struct_empty.tmpl
// tmpl/struct_field.tmpl
// tmpl/union.tmpl
// tmpl/union_degenerate.tmpl
package tmplgen

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

var _tmplCommonTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xcb\x6e\xdb\x30\x10\xbc\xf3\x2b\xb6\x04\x02\x48\x80\x42\xdd\x03\xf8\xd0\x3c\x0e\x41\x0d\xc7\x68\xdc\x53\x51\x04\x0c\xb9\xa2\x09\x49\xa4\x4a\xae\x1c\x1b\x86\xff\xbd\xa0\x24\x3f\xd2\x14\x6d\x7d\x30\xa9\x9d\x19\x2e\x67\x96\x65\x09\x77\x5e\x23\x18\x74\x18\x24\xa1\x86\xd7\x1d\x18\x4b\xeb\xfe\x55\x28\xdf\x96\xb6\x6b\x74\x69\xfc\x75\x5a\xaf\xbb\x60\x5b\x2c\xa3\x5a\x63\x2b\x4b\x83\xae\x34\x41\x76\xeb\x9f\x4d\x19\x31\x6c\x30\x14\x70\xff\x04\x8b\xa7\x15\x3c\xdc\x3f\xae\x04\x63\x9d\x54\xb5\x34\x08\xfb\x3d\x88\xe5\xb4\x3f\x1c\x18\xb3\x6d\xe7\x03\x41\xc6\x00\x00\xb8\xf2\x8e\x70\x4b\x7c\xfc\xaa\x5a\xe2\x6c\xdc\x5e\xdc\x62\xea\x73\x6d\xfc\x71\xcb\xff\x83\x53\x36\xd2\x99\x5e\x1a\x2c\x65\x9c\xce\x4f\x36\xde\xa9\x3e\xfa\x1b\x89\xca\xea\xc6\xba\xfa\x1f\xdc\x32\x71\xac\x33\xa5\xb2\x7a\xba\x51\x32\xbb\x7a\x3e\xdb\xe5\x2c\x67\x8c\x76\x1d\x82\xf3\x1a\xe7\x5e\x6a\x0c\x50\xf5\x4e\x65\x8a\xb6\x30\x99\x17\x77\xe3\x5a\xa4\xbe\xc7\xde\x62\x6e\x5d\x5d\xc0\x6b\x6f\x9b\xa4\x49\x5d\xc5\xc2\x6b\xbc\x1d\x0b\x39\x64\xa7\x52\x01\x18\x82\x0f\x39\x53\xde\x45\xba\xe8\x74\x47\xdb\x2f\xb8\x83\x19\xf0\xc5\xa9\xc6\x19\xdb\xc8\x90\x14\x0b\x4f\xa9\x0c\x33\xa8\x5a\x12\x0f\xe9\x88\x2a\xe3\x0b\x4f\xf0\xb8\x9c\xdf\x43\xc2\x78\x7e\x24\x3f\xba\x8d\x6c\xac\x9e\x1c\xbc\x97\x4c\x18\x4c\xe0\x32\xf8\x8d\xd5\xa8\x3f\x8a\x53\xa4\x7f\x96\x26\xc7\x67\xfe\x37\x87\xdb\x0e\x15\xa1\x5e\xa5\xe8\x66\xc0\xcf\x15\x18\xd2\xbc\x5a\x09\x38\x55\xae\x22\x67\x2c\x85\x0a\x01\xa3\x6f\x36\xf8\xd2\xca\xee\x45\x52\xd6\xc1\xf4\x16\xc4\xd7\x11\x58\xca\x20\xdb\x98\xb2\x73\x84\xa1\x92\x0a\xf7\x87\x63\x7a\xb0\x1f\x46\x48\xb1\x00\x5f\xc3\xcd\x0c\x3a\xf1\xec\xfb\xa0\x50\x9c\x93\xce\xc7\x67\x54\xc1\x27\x5f\x4f\xfc\xf4\x0b\x48\x7d\x70\xe0\x6c\x53\x5c\x04\x3b\xc0\x87\xe1\x5f\x06\x33\x9e\xf8\x39\x98\xf8\x9d\xd7\xb8\xe3\x3f\xc6\x77\x1e\xdf\x2c\xa9\x35\x90\x4c\xb8\x0c\x46\x64\xc9\xdf\xf1\x32\x4a\x46\x3c\x8f\xfe\xe6\xf7\x86\x14\xc5\xdc\xfb\xba\xef\x6e\x77\x09\xcf\x48\xe6\x67\x59\xa4\x60\x9d\xf9\x9b\xe6\x79\x60\x9c\x54\x1a\x2b\xd9\x37\xf4\x41\x31\xd8\xba\x1c\x5a\xef\x6a\xe7\xdf\x1c\xd4\xb8\x1b\xa6\x71\x03\x57\x2b\x5e\xa4\xdb\xe7\x93\xe7\x03\xfb\x15\x00\x00\xff\xff\x3e\x1f\x82\xd2\x5a\x04\x00\x00")

func tmplCommonTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplCommonTmpl,
		"tmpl/common.tmpl",
	)
}

func tmplCommonTmpl() (*asset, error) {
	bytes, err := tmplCommonTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/common.tmpl", size: 1114, mode: os.FileMode(420), modTime: time.Unix(1623174493, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplListTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x4d\x6f\xe3\x36\x10\x3d\xc7\xbf\x62\x22\xa0\x05\x05\x08\x42\x0f\x45\x0f\x02\x7c\x70\x93\x16\x30\x6a\xa8\x81\x13\xf4\x12\x04\x05\x23\x8d\x64\x56\x34\xa9\x50\xa3\x24\x86\xea\xff\xbe\xa0\x64\x3b\xb2\xc5\xec\x3a\x0b\x64\xb3\xd8\x84\x87\x20\xfc\x98\x37\x8f\x8f\x33\x2f\xca\x3d\x37\xd0\x34\x10\xc2\xff\x30\xd3\x09\x97\x31\x5f\x22\xac\xd7\x30\x86\xdc\xf0\x72\x71\x27\xc3\x18\x1f\xfe\xbe\xfd\x0f\x13\x62\xdb\x95\x6e\x7a\xa6\x55\x26\xf2\x66\x04\x00\x60\xa3\x22\xf0\x2c\xd0\x06\xc0\x0b\xda\x8d\x3f\x05\xca\xb4\x8a\x76\x60\xdd\xbc\x0b\xb2\xc3\x9b\x90\x17\xc1\xcf\x7b\xdb\x4f\xbb\x76\x5c\xad\x4a\x8c\x5a\x8a\xff\x70\x59\xa3\x9d\x1e\x70\x0d\xf6\xce\x4f\x4c\x7e\x98\xaf\x63\x3a\x31\x79\xbd\x44\x45\xfb\xf0\x2d\x89\x02\x57\x7d\x16\xdb\x93\xfd\x1b\x1e\x8e\x8e\x56\x4f\xa3\x58\xab\xb8\x96\x72\x27\xd2\x54\x91\x1f\x0c\x42\x0f\xc8\x1e\x4c\xe7\x58\x69\x79\x8f\x11\x64\xb5\x4a\x58\xb9\x83\xdf\xac\x5f\x70\xc3\x97\x95\x0f\x4c\x28\x42\x93\xf1\x04\x9b\x75\x00\x68\x8c\x36\x3e\x0c\x69\x52\x15\x80\x2e\x20\x1a\x43\x19\x5e\xea\xda\x24\x18\xb2\xa6\x69\x99\x5f\xf0\xa4\xe0\xb9\x15\x2f\xdc\x3c\xbe\x5d\xbd\x5c\x2d\x6f\xb5\x8c\xf5\x1c\x93\xda\x54\x76\xd7\x1f\x80\x8a\x0c\x4e\x75\xe1\xc8\x66\x87\x41\xaa\x8d\x02\x25\x64\x4b\x2b\xd6\x14\xeb\x14\x87\x22\x8c\x06\x4b\xdc\xe4\x1d\x51\xfb\x7c\xd7\xed\x83\xdc\x0c\x0e\xd9\x5a\xd5\x35\x81\x28\x65\x1a\x3a\x91\xed\x09\x34\xa6\xd3\x64\xb0\x5b\x3d\x08\x4a\x16\x40\xdc\xa6\xe2\x26\x0f\x19\xad\x4a\x74\x49\x97\xf0\x0a\x9f\xd2\x44\xce\xcb\xea\x9a\xda\x5b\xc2\x18\xa8\x0a\x67\x5a\x17\x75\xf9\xfb\xca\x9e\x67\xc4\x87\xc2\x75\x90\x8a\x7e\xfb\xf5\x25\x70\x53\x95\xe2\xa3\x13\x2f\xc5\x8c\xd7\x92\xdc\x60\xfd\x87\xc8\x96\x14\xfe\x61\xf5\xc8\x98\x57\xab\x42\xe9\x07\x05\x05\xae\x80\xda\x0a\xfe\xe9\xca\x0b\xac\x16\xc3\x04\xeb\xc1\x4a\xd3\xd8\xe7\x67\x5c\xa5\xc0\xf0\xae\xd7\x90\xa1\xfd\xf1\x97\x50\x69\x78\x49\x46\xa8\x1c\xbc\x99\x50\x85\xe7\x03\x53\x08\x6c\xbf\x71\x85\x2a\xae\xb8\xc9\x91\x7c\xf0\x76\xbd\x72\xee\xf9\x3e\xac\x87\x19\x45\xd6\x4a\x72\x3a\xb6\x57\x39\xae\xe8\x8e\xb8\x07\xb5\xf9\xcf\x44\xda\x09\x1e\x8d\xad\xf8\xe1\xa4\xb2\xdc\x98\xb3\xe4\x5f\x81\xc4\xa1\xa3\xcd\x5b\x84\x4e\x1a\xf0\x76\x14\x3d\x97\x2c\x4d\x03\x28\xdb\x06\x1d\x6c\x6d\x88\x6c\xab\xc9\x19\xaa\xd2\xc3\xc8\x9e\x0f\xf5\x7e\xf5\x26\x52\x1e\xe5\xce\x3d\x1b\x9c\x89\x8a\xd8\xe7\xdc\xda\x7f\xa7\x96\x37\x40\x25\x4b\xd3\xb6\xba\xa8\x68\x4a\x68\x38\x69\xe3\xa8\xbe\x64\x21\x64\x6a\x50\xd9\xd3\x4b\x5e\x20\xbb\xbe\xd9\x19\x53\x00\xbf\x0c\x03\x32\x6d\xe0\x54\x50\x78\xae\x15\x32\x97\x48\x76\xfc\x1b\x80\x6a\x01\x36\x0d\x20\x28\x8c\xf1\x91\x1c\xf9\xe1\xc8\x0e\x70\x88\xe1\x3c\x37\x54\x02\xde\xc4\x5a\xc0\x69\x04\x56\x94\xe7\x9d\xe0\x9b\x69\xb1\x77\xaf\xa9\x22\xed\xf0\x05\xcf\x72\x75\xda\x03\x3c\xdb\xe7\xdb\xb1\x2b\xaa\x31\xf0\xb2\x44\x95\xb2\xed\x4a\x57\x17\xc7\xfc\x35\xd8\x5c\xb0\x17\x28\xe4\xc9\x31\xae\x32\xe7\x2a\xc7\xd7\xf6\x95\xaf\xf9\x0c\xac\x0a\x51\xbe\xcd\x77\x60\x9b\x9e\x78\x81\x1f\x9f\xa1\x3f\x84\x27\x7f\x98\xf2\x87\x29\xbb\x41\xbe\x57\x53\x3e\xd3\xb5\x3a\xee\x5f\xf1\x17\x99\xcd\xbb\xb5\x92\x4d\xa0\xf5\x12\x54\x39\x2d\x98\xdf\xbe\xc5\x17\x9e\x62\x1d\x8c\xd6\xfe\xc9\xe8\x53\x00\x00\x00\xff\xff\x36\xc4\xaa\xf7\x95\x11\x00\x00")

func tmplListTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplListTmpl,
		"tmpl/list.tmpl",
	)
}

func tmplListTmpl() (*asset, error) {
	bytes, err := tmplListTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/list.tmpl", size: 4501, mode: os.FileMode(420), modTime: time.Unix(1623174368, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplMapTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x51\x4f\xdb\x30\x10\x7e\x1e\xbf\xe2\xf0\xc3\xe4\x48\x51\xc4\x5e\x2b\xf5\xa1\x82\x4d\x42\x65\xd9\x34\xd0\x5e\x10\x8a\xbc\xe6\x5a\xbc\xb8\x76\x70\x9c\x6e\x55\x96\xff\x3e\x39\x49\x43\x12\xa7\x88\x31\x18\x68\xe2\x1e\x10\xb6\xef\x3e\x7f\x77\xfd\x7c\xbd\x6e\x98\x86\xa2\x80\x00\x7e\xc1\x99\x5a\x30\x11\xb2\x35\x42\x59\xc2\x14\x56\x9a\xa5\xd7\x37\x22\x08\xf1\xc7\xa7\x6f\xdf\x71\x61\xe8\x6e\xa7\x5e\x1e\x2b\xb9\xe4\xab\xe2\x00\x00\xc0\x46\x4d\x80\x58\xa0\x06\x80\xf8\xd5\xc1\x07\x8e\x22\xce\x26\x2d\x58\xbd\xae\x83\xac\x91\x99\x21\x13\x78\xdb\x3b\xbe\x3d\xb5\x76\xb1\x4d\x71\x52\x51\xfc\xca\x44\x8e\x76\x39\xe0\xea\xf7\xfc\x67\x7a\x35\xbc\xaf\x66\x3a\xd3\xab\x7c\x8d\xd2\xf4\xe1\x2b\x12\x09\x6e\xbb\x2c\x76\x9e\xdd\x0c\x87\x56\xd3\xea\xd4\x28\x54\x32\xcc\x85\xa0\x96\xe9\x1c\xb7\x23\x3c\x3d\xdf\x41\x1a\x70\x2f\xfd\x37\xbd\xf5\x17\xcc\x94\xd8\xe0\x04\x74\xfd\x4f\xb4\x66\x69\xc4\xcc\x6d\x50\x27\x9e\xcc\x71\x9b\xdd\xab\x96\x1d\xd2\x67\x3c\x33\xfb\x18\x3b\x84\x5b\x36\xcb\x5c\x2e\x68\xda\xe2\x34\xfb\x9f\x99\x66\xeb\xcc\x03\xca\xa5\x41\xbd\x64\x0b\x2c\x4a\x1f\x50\x6b\xa5\x3d\x70\x8b\x68\x32\x1f\x54\x02\x93\x29\xa4\xc1\xb9\xca\xf5\x02\x03\xca\x53\x11\x07\xa1\x8a\xd1\x73\xdc\xf9\x12\x0e\x55\x32\x82\x63\x4d\xa3\xc9\xb5\x04\xc9\x45\x75\x61\xa8\x8c\x05\x71\xab\xed\xa2\x1a\x4b\xc0\x64\xc1\x47\x96\x9e\x1a\xd4\xcc\x28\x4d\xdd\xcb\x17\xd7\x5c\xc4\x1a\xa5\x75\x5e\xb3\x04\xe9\xe5\x55\x4b\xd5\x87\x23\xef\xc0\x89\x58\x2a\x0d\x87\xdc\x04\x27\x4a\x22\x1d\x4b\xdf\x9a\xac\xc2\xa3\x8a\xb3\x85\xe6\x26\x08\xf1\xa7\x19\x21\xd0\x54\xc0\xfa\x1d\x4e\x6d\x9e\x7b\x10\x47\x8a\x31\xea\xe7\x56\xa2\x97\xe6\x14\x58\x9a\xa2\x8c\xe9\x6e\xc7\xaf\xb8\xba\xbc\x5c\x9c\xe6\xf6\x4e\x20\x17\x03\x8d\x8f\xaa\xb7\x7a\xdc\x0f\xd7\xef\x68\x6f\x78\x55\x70\x6b\x4f\xa4\xe0\xc8\x6f\x44\xfc\x42\x14\x5c\x14\xf6\x12\xca\x64\x0c\x14\x6f\x3a\xb2\x08\xec\x9f\x39\x97\x71\x70\x6e\x34\x97\x2b\x20\x67\x5c\x26\xc4\x03\x2a\x11\x68\x5f\x3e\x5c\x26\x17\x4c\xaf\xd0\x78\x40\x76\x0a\x39\x3d\x21\x9e\x07\xe5\xf8\xad\xa6\xf2\x3e\xe6\x71\x5b\x07\x5b\x94\x60\x96\x59\xa8\xe7\xac\x45\x2f\xaf\x53\x69\x54\x9d\x17\x90\x96\x31\x01\x62\xb9\x92\x7d\xa9\x15\x05\xa0\x8c\xf7\x9d\x3e\x67\xc3\x98\x09\xf1\xe0\x6e\x31\xf8\x62\x8e\x22\x94\x46\x6f\x5f\x9b\x45\x63\x23\xcd\xe2\x71\xda\x45\xe2\xc3\xe6\xc5\xf4\x8a\xbb\xc4\xdb\x49\xb7\xb0\xa4\xcb\x27\x52\xb2\x1d\xf8\x4a\xef\x60\xcf\x10\xde\xe8\xf2\xaf\x67\xf1\xe8\xbd\x85\xb9\xef\x44\x3e\xef\x0f\xc3\x77\x8e\xe4\xe3\x83\xee\x53\xbe\xa3\x64\xe3\xbe\xa3\xce\xa7\xf5\xaf\x5e\x52\x13\x98\x6c\x2e\x8f\xae\xfe\x6c\xc6\x79\xe4\x9f\x3b\xff\x79\x75\xdf\xdd\xab\xba\xa5\x6f\x5f\xd1\xef\x00\x00\x00\xff\xff\xe5\x2f\xe9\x2b\xcb\x0e\x00\x00")

func tmplMapTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplMapTmpl,
		"tmpl/map.tmpl",
	)
}

func tmplMapTmpl() (*asset, error) {
	bytes, err := tmplMapTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/map.tmpl", size: 3787, mode: os.FileMode(420), modTime: time.Unix(1623174368, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplScalarTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\xc1\xca\x82\x50\x10\x85\xf7\x3e\xc5\xc1\xd5\xff\x83\xf8\x00\x42\xab\x5a\x8a\x04\x42\x5b\x19\x64\xb2\x81\x49\x6d\xb4\x82\xec\xbe\x7b\xdc\xbc\x42\x90\xcd\x6e\xe6\xfb\xe6\x70\x6e\x64\x98\x26\xa4\x78\x22\xef\x6a\xd2\x82\xce\x0c\xe7\xb0\x41\x63\xd4\x9f\x2e\x9a\x16\x7c\x2f\x6b\x52\xb2\xbf\xe5\x32\xaf\xdb\xae\x3d\x4a\x33\x45\x00\xe0\xbf\x32\x84\x89\x7d\x5e\xc8\x89\x93\x37\xdf\xf1\x50\x9b\xf4\xa3\x74\x6d\xb6\xc6\x4b\x36\x21\x95\x07\x67\x5f\x65\x9c\xab\xaa\x61\xc1\xb3\xbd\x27\x1b\xf8\x40\x7a\xfd\xa1\xf7\x9e\x7f\xa8\xb9\x8c\x6c\xa4\x5e\x5e\x77\x83\x90\x44\xee\x3f\x7a\x05\x00\x00\xff\xff\xc1\x29\x39\x00\x11\x01\x00\x00")

func tmplScalarTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplScalarTmpl,
		"tmpl/scalar.tmpl",
	)
}

func tmplScalarTmpl() (*asset, error) {
	bytes, err := tmplScalarTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/scalar.tmpl", size: 273, mode: os.FileMode(420), modTime: time.Unix(1623174368, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplScalar_parseTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\xbf\x4e\xc7\x20\x14\x85\x77\x9e\xe2\x8e\xf0\x1b\x18\x3a\x36\x71\xd0\xd9\x34\x46\xdd\x9b\x2b\xbd\xad\x44\x44\xc2\x9f\x36\x0d\xf2\xee\x06\xb4\x69\x47\xc7\x73\xe0\xfb\x4e\xee\x9c\xac\x82\x9c\x41\xc2\x37\x3c\x7e\x29\x34\x03\x7e\x52\x29\xe3\xe8\xd0\x07\xe2\x2b\x9a\x44\xa0\x6d\x24\x3f\xa3\xa2\x5c\xc4\x35\x40\x66\x00\x00\x6f\x49\x9b\x89\x3c\xf4\x77\xd5\xf4\xba\x3b\x7a\x42\xf5\x81\x0b\x95\x22\x6b\x92\xd5\x5f\xbd\x50\xc5\xcf\xe4\xbc\x1c\x68\x7b\xf8\xa5\xb8\x68\x8e\xb0\xe9\xa8\xde\x61\xed\xaa\xa5\xad\x4a\x1e\x77\x47\xe2\x6f\x43\x61\x20\x08\xd1\x6b\xbb\xf4\xad\xb8\x0c\xcb\xfb\x10\xf4\x62\x5f\xda\x2b\x5f\x3b\x71\x12\xb7\x7f\x21\xb7\x83\x99\x68\xc6\x64\xe2\xf9\xdd\x53\x4c\xde\x82\xd5\xa6\x55\x85\x5d\xca\x43\xd5\x0e\xe1\x82\x15\xf6\x13\x00\x00\xff\xff\x8c\x00\x0b\x27\x4d\x01\x00\x00")

func tmplScalar_parseTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplScalar_parseTmpl,
		"tmpl/scalar_parse.tmpl",
	)
}

func tmplScalar_parseTmpl() (*asset, error) {
	bytes, err := tmplScalar_parseTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/scalar_parse.tmpl", size: 333, mode: os.FileMode(420), modTime: time.Unix(1623174368, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplScalar_parseliteralTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\xcd\x4e\xc3\x30\x10\x84\xef\x79\x8a\x39\x26\x1c\xfc\x00\x91\x7a\x28\xe7\xaa\x42\xb4\xe2\x5a\x2d\xee\x26\x58\x18\x63\xf9\x87\xaa\x5a\xfc\xee\xc8\x6e\x49\x73\xf3\xae\x77\x66\xbe\x99\xb2\xd3\x10\x51\xf8\xdd\x7d\x6b\xb2\x7b\xfa\xe2\x52\x4e\x27\x4f\x21\xf2\xce\x24\x0e\x64\xfb\x1f\xb2\x99\xb7\x87\x23\x28\x26\xf5\x56\x87\x01\xc6\x25\x0e\x13\x69\x96\x02\xe9\x00\xe0\x3d\x1b\x7b\xe6\x80\x71\x03\x11\x1c\xaf\x9e\x5f\x48\x7f\xd2\xcc\xa5\xa8\x3a\x29\x11\xa8\xea\x8f\x1a\xf0\xca\x3e\xa8\x3d\x5f\x9e\x6f\xaa\x7e\x68\x1e\xf1\x62\x92\xfe\xc0\x12\x38\x6e\x96\xb7\xea\xd3\xd5\xf3\x70\x0f\xd3\x14\x19\x4f\x95\xe7\x90\x82\x71\x73\xa3\x1a\xdb\xd7\x8a\x45\x6d\x63\x34\xb3\xbb\x9d\x2c\x35\xee\x15\xda\xf1\x99\x27\xca\x36\x3d\x94\x81\x53\x0e\x0e\xce\xd8\xb6\x2a\xdd\x6a\xf9\xef\xda\x98\xfb\xa1\x2b\xdd\x5f\x00\x00\x00\xff\xff\x59\x1d\xf2\x3c\x3e\x01\x00\x00")

func tmplScalar_parseliteralTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplScalar_parseliteralTmpl,
		"tmpl/scalar_parseliteral.tmpl",
	)
}

func tmplScalar_parseliteralTmpl() (*asset, error) {
	bytes, err := tmplScalar_parseliteralTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/scalar_parseliteral.tmpl", size: 318, mode: os.FileMode(420), modTime: time.Unix(1623174368, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplScalar_serializeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x31\x4f\xc3\x30\x14\x84\xf7\xfc\x8a\xa3\x53\x22\xa1\xfc\x80\x48\x1d\x60\xab\x40\x5d\x60\xaf\xdc\xe4\x05\x9e\x64\xdc\x60\x3b\xa0\x60\xfc\xdf\x91\xed\xa4\x49\x89\x90\x18\xfa\x96\xc4\x77\xe7\x77\xdf\xe0\xb6\x57\x35\x9c\x43\x89\x6f\x3c\x9e\x6a\x21\xf7\xe2\x8d\xbc\x3f\x1c\x0c\x69\x16\x92\xbf\x28\xff\x10\xb2\x27\xb0\xb2\xa4\x5b\x51\x93\xf3\xc5\xf2\x00\x97\x01\x80\xf9\x64\x5b\xbf\x22\x65\xab\x6d\xfa\x29\x73\x3b\x74\x54\x8c\x91\x5a\x18\x02\x77\xb2\x29\xf7\xa7\x86\xaa\xa8\x85\x71\x0e\xdc\x22\xa7\x77\x94\xcf\x43\x47\x0f\xac\x9a\xf2\xc9\x6a\x56\x2f\xd8\xec\x94\xdd\x14\xf0\xfe\x1c\xe6\x5b\x90\xd6\x73\xc3\x9d\xd9\x29\x9b\x17\xb3\xdf\x46\xff\x66\x0b\xc5\x72\x2c\x9e\x46\x93\xed\xb5\x0a\xfe\x59\x9e\x17\x8f\x26\x2f\xb1\x48\x06\xe2\xbf\xd8\xee\x07\x4b\xe6\x92\xee\xb8\xa2\x8b\xa1\x2b\xf2\x1d\xff\xcf\x97\xbe\x97\x80\x66\x05\x98\x52\x57\x24\x34\x2b\x42\xbf\xca\x4c\xe5\xe1\xe1\x4d\xdc\xf0\x7e\x41\x11\xee\x46\x2d\x2a\x0d\xb5\xa2\x97\xb6\xfa\xbd\x47\xb1\xcc\x12\x84\xcf\x7e\x02\x00\x00\xff\xff\x81\x1c\x1d\x93\xca\x02\x00\x00")

func tmplScalar_serializeTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplScalar_serializeTmpl,
		"tmpl/scalar_serialize.tmpl",
	)
}

func tmplScalar_serializeTmpl() (*asset, error) {
	bytes, err := tmplScalar_serializeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/scalar_serialize.tmpl", size: 714, mode: os.FileMode(420), modTime: time.Unix(1623174368, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplStructTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xd0\xdd\x4a\xc3\x30\x14\x07\xf0\xeb\xf6\x29\xfe\x8c\x21\x1b\xc4\x3e\x40\x60\x57\x82\x20\x68\x04\xf1\x7e\x64\xeb\x69\xad\xc4\x64\x26\x3a\x19\x31\xef\x2e\x69\xba\x7e\x49\x73\xd5\xf3\xf5\xa3\xe7\x9c\xa5\x85\xf7\x28\xf0\x8b\x47\x73\x94\x4a\xc8\x0f\x42\x08\xd8\xa1\xb6\xf2\xf4\xf6\xa9\x0a\x41\x3f\xcf\x87\x77\x3a\x7e\x6d\xae\x99\x14\xde\x19\x5d\x35\xb5\xcf\x01\x20\x4e\x71\xac\x22\xd4\x01\x2b\xd6\x16\xee\x1b\x52\xa5\xe3\x3d\x96\xe2\x34\x14\x9f\xf7\xb7\xb0\x52\xd7\x84\x75\x15\x4b\xe0\x3b\x74\x4d\x08\xa1\x6f\x9b\xc8\x1c\x37\x13\x6e\xd0\x92\x88\xa6\x42\xf1\xe0\x9e\xe4\xe5\x40\x63\x24\xbe\xd7\xcb\x89\x78\xbb\x70\xfc\x9a\x2d\xcd\xe6\x10\x29\xb7\x20\x8c\x8e\x23\x8c\x16\xdf\x4a\x6d\x16\xd0\xed\x7f\x55\x97\x73\xf4\x85\x9c\x51\xe7\xf4\x67\xeb\xeb\xa2\xfb\x7d\x8c\xda\xb3\x0c\x29\x9b\x3a\x07\x34\xb0\x2c\xcb\x26\xf7\x24\x5d\x76\x7c\x60\x79\xd8\xe6\x7f\x01\x00\x00\xff\xff\x9c\x53\x0d\x0e\xe3\x01\x00\x00")

func tmplStructTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplStructTmpl,
		"tmpl/struct.tmpl",
	)
}

func tmplStructTmpl() (*asset, error) {
	bytes, err := tmplStructTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/struct.tmpl", size: 483, mode: os.FileMode(420), modTime: time.Unix(1623174368, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplStruct_emptyTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8d\xc1\x4a\xc4\x30\x10\x86\xef\x7d\x8a\x9f\x1e\xa4\x85\xb0\x0f\x50\xf0\xa2\xe8\x49\x56\x11\xef\xcb\x58\xa7\x6b\x24\x9b\xd4\x49\xba\x2a\x31\xef\x2e\xb1\xdd\x96\xfc\xa7\xcc\xe4\x9b\xef\x3f\x93\x20\x46\xec\xf0\x8b\x07\xd7\x93\xd9\xd3\x89\x91\x12\xae\x71\x14\x1a\xdf\x3f\xcd\x6e\xcf\x5f\x8f\xaf\x1f\xdc\x87\xe6\xb2\x99\xc7\x5b\x67\x07\x7d\x8c\x15\x00\xe4\xab\x0e\x75\x16\x2d\x82\x5a\xfd\x7f\xdc\x6b\x36\x6f\xbe\x5b\x65\xf3\x3c\x1f\xe5\xd4\x87\xc3\xdd\xb7\xf6\xc1\xd7\x1d\xae\x0a\x68\x63\x72\x5e\x7e\x46\xde\x2c\x37\xce\x19\x26\xab\x0a\xe4\x99\xbd\x33\x67\xee\x30\x4c\xb6\x6f\xc6\x15\x5e\xf6\x4f\x24\x74\xf2\x2d\x1a\x6d\x03\xcb\x40\x3d\xc7\xa4\xc0\x22\x4e\x5a\x94\x65\x39\xc2\x61\x12\x8b\x20\x13\x2b\x58\x6d\x0a\x20\x6d\xcd\xcb\x33\xa9\x2a\xb5\xd5\x5f\x00\x00\x00\xff\xff\xd2\xb1\xdc\x60\x4d\x01\x00\x00")

func tmplStruct_emptyTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplStruct_emptyTmpl,
		"tmpl/struct_empty.tmpl",
	)
}

func tmplStruct_emptyTmpl() (*asset, error) {
	bytes, err := tmplStruct_emptyTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/struct_empty.tmpl", size: 333, mode: os.FileMode(420), modTime: time.Unix(1623174368, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplStruct_fieldTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\x4d\x6e\xdb\x30\x10\x85\xf7\x39\xc5\x94\x2b\x12\x10\x78\x80\x02\x5d\xf4\x27\x05\x8c\x36\x45\x10\xa7\xeb\x80\x91\x86\x2e\x61\x89\x52\x48\xaa\x88\xc1\xf0\xee\xc5\x50\xff\x45\x5b\xbb\x9b\x78\x61\xc3\x43\xce\xfb\x66\xde\xa3\xee\x6d\x09\x31\x82\x0c\xf2\x9b\x6a\x10\x52\x7a\x78\xa0\xbf\xda\x60\x5d\xe5\x12\x55\x1c\xfa\xb6\xfe\x89\xbc\x83\x83\x53\xdd\x8f\xa7\x5a\xde\x0d\x95\x5b\xe5\x54\xe3\x05\x70\x63\x03\x3a\xad\x4a\x8c\xa9\x00\x74\xae\x75\x02\xe2\x15\x00\x40\xf0\x05\xb4\x47\x78\xfb\x0e\x3a\xb9\x6f\x7b\x57\xa2\xe4\x31\xc2\xfd\xa9\xc3\x5b\x55\x1e\xd5\x81\xa0\x32\x8f\x00\x2f\xb9\xbc\x3f\x35\x8f\x6d\x0d\x29\x89\x2c\x60\x34\xbc\x69\x8f\xa3\x1a\x7d\x1c\x86\xde\x59\xb0\xa6\x2e\x40\x37\x41\x5e\x13\x4e\x73\x74\xee\xbb\xc5\xe7\x0e\xcb\x80\x15\xe9\x14\x33\xb1\x00\x76\x21\x92\x0d\xcc\x94\xbf\x63\x24\xf8\xe8\xc5\xce\xdf\xa8\xd3\x23\x75\xe6\x33\x4d\x1b\x05\x2f\x3f\xd3\xe1\xd6\x31\x48\x89\xcf\xa3\x6b\x79\xfd\x6c\x7c\xf0\x5c\xac\x36\xd8\x08\x13\x1f\x5e\x60\xe7\x3f\xf4\xa6\x0e\x3b\x3b\x11\x56\xab\x6a\x79\xd3\xfb\xc0\x85\x7c\xef\x17\xd4\xd8\x47\x3f\x5f\x8c\xad\x16\xea\x48\xc0\xda\x23\x61\xb8\xb2\x15\x70\x7c\x5a\xf7\xc9\xa9\x4b\xee\x83\x33\xf6\x00\xec\xab\xb1\x47\x26\x80\x5b\x04\xbe\x05\xd0\xc9\xbd\x72\x07\x0c\x02\xd8\x94\xff\xee\x13\x13\x62\x3d\x69\xc8\x37\x3e\x9a\xea\x8c\x2f\xd3\x26\xa4\xba\x9d\x77\x4b\xbd\xcb\x9b\x0f\x5c\x60\xb3\x3a\x5b\x33\xa7\x25\xff\x6e\x58\x41\xaf\x64\x73\x3f\x1b\x35\x64\x3c\x34\xff\xf9\x5d\x4d\x6d\xf3\x43\x98\xdc\x3c\x13\xda\xa8\xf1\x2f\x03\x2e\xca\xf0\x55\xf3\xbb\x38\xbb\x55\x68\xff\x1d\xd8\x6f\x61\x9d\x37\x6a\x09\x61\xc9\x2d\x5d\xfd\x0a\x00\x00\xff\xff\xbf\xaf\x1e\x80\xb1\x04\x00\x00")

func tmplStruct_fieldTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplStruct_fieldTmpl,
		"tmpl/struct_field.tmpl",
	)
}

func tmplStruct_fieldTmpl() (*asset, error) {
	bytes, err := tmplStruct_fieldTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/struct_field.tmpl", size: 1201, mode: os.FileMode(420), modTime: time.Unix(1623176313, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplUnionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x55\x5d\x6f\xd3\x3c\x14\xbe\x6e\x7f\xc5\x79\xa3\xed\x55\x82\x3a\xff\x80\xa0\x5d\x0c\x10\xd2\x04\x94\x69\x1d\xdc\x00\xaa\x3c\xf7\xa4\x35\x73\xed\x60\x3b\x9b\xaa\xcc\xff\x1d\x39\xdf\x4e\x56\x36\x71\x87\xaf\x6c\x9f\x8f\x9c\xf3\xe4\x39\x8f\xef\xa9\x86\xb2\x04\x02\x8f\xf0\x51\x31\x2a\x96\x74\x8f\xe0\x1c\x9c\xc3\x56\xd3\x7c\xf7\x4b\x90\x25\x3e\x7c\x91\x5c\xc9\xb8\xbd\xa8\x4e\x6f\x95\xcc\xf8\xb6\x9c\x03\x00\xf8\x98\x14\x22\x9f\xa6\x09\x8f\x16\x95\xe1\xe6\x90\xa3\x49\xe1\xdb\x8f\x57\x6d\xec\xe7\xdb\x9f\xc8\x6c\x1d\xe6\x57\x59\x9e\x81\xa6\x72\x8b\x70\x72\xc7\xe5\x06\xd2\x73\x20\x9f\x70\x7f\x8b\xda\x38\x37\xf0\xe2\x59\xe3\xf0\x08\x97\x66\x65\x75\xc1\x6c\x60\xef\xac\xc3\x26\x16\x03\x07\x14\x06\x07\x11\x85\xef\x61\xbd\x2e\xcb\x93\xaa\x64\xe7\xaa\xbd\xcf\xd1\x9c\x83\x58\xb9\x09\x3e\x76\x06\xfd\x4d\xe3\x78\x8d\x46\x89\x7b\xf4\x0d\xa7\x90\x15\x92\xc5\x79\x07\xe0\xc0\x76\x45\x35\xdd\x9b\x04\x46\x78\x40\x0f\x08\xcf\x40\xaa\x0d\x2e\x40\xdd\x79\x30\x72\xf2\x95\x8a\x02\x49\xcc\x73\xb1\x21\x4b\xb5\xc1\xe4\xb5\x37\xf5\x01\x7e\x99\x07\x6e\xd9\xae\x0a\x24\x57\x5a\x59\x65\x0f\x39\xc6\xc9\xc8\xeb\xe5\x78\xbf\x04\x73\xbf\x18\x35\xe8\xc1\xaf\x5b\x63\x77\x74\x8b\xce\x11\x7f\x22\xed\x2f\x69\x19\x91\x4e\x4a\xc9\xa8\x10\x76\xa7\x55\xb1\xdd\xfd\x6d\xd2\xf5\xfa\x1a\x73\x3d\x4d\xad\xd1\x16\x5a\x1e\xa3\xc5\xa8\xcd\x11\x33\xfe\x81\xbe\x9e\x25\xef\xb8\xc3\x80\xbf\x30\xe1\x70\xbb\xdc\x6c\x36\x9b\x75\x37\xbd\x2d\xdb\x5b\x72\xa5\xb9\xb4\x59\x1c\x5d\x30\x5b\x50\x01\x9e\x60\x70\x7a\x93\xc2\xe9\x3d\x48\x65\x81\x37\x55\x7d\x97\xd1\xa2\x25\x6d\xb7\x49\xe6\xa3\x06\x24\x17\xed\xf0\xb8\x64\x3e\xff\x33\x29\x8f\x11\xb1\xfb\x73\x5e\xc2\x9e\x85\x24\x14\xb4\x7a\xec\xe2\x70\x0a\x9f\x94\xb4\x36\x1f\x09\xd3\x35\x02\xf7\x0e\x0d\xd3\x3c\xb7\x5c\xc9\x14\xa2\xd5\x41\xda\x1d\x5a\xce\xea\x72\x60\x5f\xf5\x00\x0f\x9a\xe6\x39\xea\x26\xe4\x3d\x47\xb1\x31\x69\x57\x4d\x7d\x1e\x2a\x22\x04\xfd\xbe\x29\xb8\xb0\x97\x72\xc8\xdb\x28\x4a\xe1\xff\x20\x3e\x1c\xf4\x5a\x85\x9e\x62\xff\x50\xd7\xa0\x97\xad\x63\x92\xd5\xca\x55\xcc\xa5\x45\x9d\x51\x86\xa5\x5b\x00\x6a\xad\xf4\x53\xe2\x62\x4d\x2f\x5b\x2b\x55\x68\x86\x24\x0e\x09\x0f\x15\x90\x5d\x5d\xde\xb2\x3a\xec\x6f\x95\x58\xaa\x6b\x64\x85\x36\xde\x23\x99\x24\xe6\x19\xfc\x37\x11\xbd\x29\xa7\xaa\xd2\x96\xca\x7a\x99\x9c\xb8\xba\x63\xd3\x64\x0d\xb9\x30\xa3\xa2\x3e\xf8\xad\x73\x71\x58\xca\x00\xbc\xe0\x7d\xf0\xa3\x04\xfd\x7b\xe0\x92\x79\x3b\x73\xfd\x9c\xfd\x0e\x00\x00\xff\xff\x01\x1b\x4d\x3e\x68\x07\x00\x00")

func tmplUnionTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplUnionTmpl,
		"tmpl/union.tmpl",
	)
}

func tmplUnionTmpl() (*asset, error) {
	bytes, err := tmplUnionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/union.tmpl", size: 1896, mode: os.FileMode(420), modTime: time.Unix(1623174368, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplUnion_degenerateTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x54\x41\x6b\xdb\x30\x14\x3e\x2b\xbf\xe2\x35\x94\xe1\x80\xa7\xee\x1c\xc8\xa1\xac\x0c\xc6\x9a\xb4\x34\xbd\x85\x10\x14\xfb\x39\x13\x96\x25\x55\x52\xc2\x8c\xe6\xff\x3e\x9e\xe4\x74\xa1\x87\x31\xd6\xeb\x2e\x46\x7e\xef\xd3\xf7\x7d\xef\xf9\xc3\x31\x7e\x04\x27\xf4\x01\xe1\xba\x95\xba\x86\xf9\x02\xf8\x12\xbb\x3d\x3a\x3f\x0c\x13\x16\xa3\x6c\xa0\xc0\x17\x98\x2e\x85\x9d\x66\x0c\x7f\xee\x2d\x7e\xa3\xc3\x3a\x38\xa9\x0f\x33\x20\xe4\x51\x4b\xa3\x77\xbb\x18\xaf\xf9\x4a\x74\x38\x0c\xe9\x9c\xf0\xf9\x9d\xdf\xd6\xf5\x17\x89\xaa\xfe\x6c\x74\x23\x0f\xc5\x74\x5a\xc2\x87\x83\x13\xf6\xfb\x8b\xe2\xa9\x11\x27\x8c\x11\xf7\x1c\x62\x1c\xed\xfc\x84\x7b\x53\x09\x95\x19\xca\x09\x63\x4f\xe8\x8d\x3a\xe1\x1c\x9a\xa3\xae\x0a\x0b\x67\x82\xb1\xfe\x28\x9c\xe8\xfc\x0c\x0a\xa9\x03\xba\x46\x54\x18\x87\x12\xd0\x39\xe3\x66\x40\xfc\x2c\xf8\x12\x4c\x4b\x73\x5a\xbe\x36\x47\x57\x21\x2f\x62\x04\xd2\x7d\x14\x55\x2b\x0e\x08\xc3\xc0\x2f\x0c\x50\x67\xdd\x77\x7b\xa3\x56\xe6\x09\xab\xa3\xf3\x84\x98\x11\x97\x6c\xe0\xca\xb4\x99\x97\x39\x0c\x47\xa7\x41\x4b\x95\x04\x57\x26\xac\x4c\x8d\xd4\x1a\xe8\xd1\x49\xd2\x0c\x9e\x2f\x85\xfd\x1a\xd0\x89\x60\x5c\x91\x59\x02\x76\x9e\x9a\x9d\x68\xb1\xe8\x84\xdd\xf8\xb4\xd7\xed\xc5\x0c\x09\xd8\x18\x07\x57\x9d\xe4\x77\x46\x63\x31\x8e\xc3\xda\x12\x4e\x49\x30\x31\x48\xbe\xc2\x1f\x21\xf3\x92\x3d\xaa\x5f\x2d\xc8\xd4\x08\x7f\x6b\x33\x15\x93\x41\x76\x73\x03\xcf\x0f\x77\x0f\x73\x68\xb1\x87\xd0\x5b\x84\x4e\xf4\xa0\x4d\x80\x3d\x42\xb6\xc4\xb3\xa4\x7f\x15\x6c\xf9\xad\xcf\x29\xf8\x37\xcd\x34\xfb\xa6\xf5\x5b\x58\xc0\xe9\x75\x57\x23\x3e\x35\x4b\xba\x36\x61\x8c\xbe\x3e\xad\x21\x46\x54\x1e\xe1\x1c\xcb\x7b\xe9\xc3\xff\x5c\xbe\x23\x97\xea\x9c\x4b\xda\xe4\x9f\x82\xb9\xd9\x4a\xab\x6a\x4e\x97\x4b\xf8\xf4\x3b\x90\xea\x4d\x20\x77\x97\x81\x54\xef\x09\x64\xd6\x5f\x80\xb0\x16\x75\x5d\x8c\x71\x38\xcd\xfe\x36\x27\xba\x1e\x86\x09\xfd\xe0\xf2\xe9\x57\x00\x00\x00\xff\xff\x04\x31\x63\x5d\xeb\x04\x00\x00")

func tmplUnion_degenerateTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplUnion_degenerateTmpl,
		"tmpl/union_degenerate.tmpl",
	)
}

func tmplUnion_degenerateTmpl() (*asset, error) {
	bytes, err := tmplUnion_degenerateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/union_degenerate.tmpl", size: 1259, mode: os.FileMode(420), modTime: time.Unix(1623174368, 0)}
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
	"tmpl/common.tmpl":              tmplCommonTmpl,
	"tmpl/list.tmpl":                tmplListTmpl,
	"tmpl/map.tmpl":                 tmplMapTmpl,
	"tmpl/scalar.tmpl":              tmplScalarTmpl,
	"tmpl/scalar_parse.tmpl":        tmplScalar_parseTmpl,
	"tmpl/scalar_parseliteral.tmpl": tmplScalar_parseliteralTmpl,
	"tmpl/scalar_serialize.tmpl":    tmplScalar_serializeTmpl,
	"tmpl/struct.tmpl":              tmplStructTmpl,
	"tmpl/struct_empty.tmpl":        tmplStruct_emptyTmpl,
	"tmpl/struct_field.tmpl":        tmplStruct_fieldTmpl,
	"tmpl/union.tmpl":               tmplUnionTmpl,
	"tmpl/union_degenerate.tmpl":    tmplUnion_degenerateTmpl,
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
	"tmpl": {nil, map[string]*bintree{
		"common.tmpl":              {tmplCommonTmpl, map[string]*bintree{}},
		"list.tmpl":                {tmplListTmpl, map[string]*bintree{}},
		"map.tmpl":                 {tmplMapTmpl, map[string]*bintree{}},
		"scalar.tmpl":              {tmplScalarTmpl, map[string]*bintree{}},
		"scalar_parse.tmpl":        {tmplScalar_parseTmpl, map[string]*bintree{}},
		"scalar_parseliteral.tmpl": {tmplScalar_parseliteralTmpl, map[string]*bintree{}},
		"scalar_serialize.tmpl":    {tmplScalar_serializeTmpl, map[string]*bintree{}},
		"struct.tmpl":              {tmplStructTmpl, map[string]*bintree{}},
		"struct_empty.tmpl":        {tmplStruct_emptyTmpl, map[string]*bintree{}},
		"struct_field.tmpl":        {tmplStruct_fieldTmpl, map[string]*bintree{}},
		"union.tmpl":               {tmplUnionTmpl, map[string]*bintree{}},
		"union_degenerate.tmpl":    {tmplUnion_degenerateTmpl, map[string]*bintree{}},
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
