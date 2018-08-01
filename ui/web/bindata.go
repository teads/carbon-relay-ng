// Code generated by go-bindata.
// sources:
// admin_http_assets/app.css
// admin_http_assets/app.js
// admin_http_assets/index.html
// DO NOT EDIT!

package web

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

var _admin_http_assetsAppCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xce\x4b\x8f\xb1\x4a\xce\xc9\x4f\xcc\x8e\xd5\x51\x88\xce\x4b\xd7\x85\xb3\x53\x12\x4b\x12\x75\x91\x05\x2a\x90\x79\x7a\x30\xb6\x8e\x82\x1e\x42\x02\x22\x9e\x91\x99\x92\xaa\x50\xcd\xa5\xa0\xa0\xa0\x90\x92\x59\x5c\x90\x93\x58\x69\xa5\x90\x97\x9f\x97\xaa\xa0\x98\x99\x5b\x90\x5f\x54\x92\x98\x57\x62\xcd\x55\x0b\x08\x00\x00\xff\xff\x9c\x9e\x21\xb0\x7a\x00\x00\x00")

func admin_http_assetsAppCssBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppCss,
		"admin_http_assets/app.css",
	)
}

func admin_http_assetsAppCss() (*asset, error) {
	bytes, err := admin_http_assetsAppCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.css", size: 122, mode: os.FileMode(436), modTime: time.Unix(1521629563, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _admin_http_assetsAppJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x4d\x8f\xdb\x36\x13\xbe\xfb\x57\x0c\x84\x05\x2c\x21\x8a\x9c\xf7\x3d\x14\xa8\x16\x46\xea\xa6\x0d\x9a\xc3\x02\xc5\x36\x6d\x0f\x5b\x07\xe0\x4a\x63\x2d\x61\x9a\x14\x48\xca\x6b\xc3\xd1\x7f\x2f\x48\x7d\x51\xb4\xb3\xd9\x78\xdd\xd6\x17\x5b\xd4\x7c\x3c\xf3\x70\x66\x38\xf4\x96\x48\x20\x65\x09\x73\xe0\xf8\x08\x84\x17\x15\x23\x32\xd9\x88\xbc\x62\x18\x06\x19\x91\xf7\x82\xbf\x96\xc8\xc8\xfe\x35\x2f\x82\x18\xee\x02\x5e\xdc\xa2\x12\x95\xcc\x30\x88\x21\xa8\x68\x72\x2f\x84\x56\x5a\x92\x32\x58\x46\xd7\x93\x09\x29\xcb\x24\x13\x5c\x4b\xc1\x18\xca\x30\xb8\x21\x94\xbf\xd3\xcc\xea\x5e\xa9\x4c\x94\x56\xef\x4a\x3a\x46\xae\x36\x22\x27\x46\x62\x55\xf1\x4c\x53\xc1\xc3\x46\x30\x86\x5e\x2c\x86\x46\x28\x3a\x4c\x00\x9a\xb7\x09\x61\x28\xb5\x82\x39\xdc\x2d\xaf\x27\x00\x26\x94\x77\x82\xaf\x68\x01\xf3\x41\x31\x0c\x66\x99\x5d\x9c\x05\x51\x27\xf5\x91\xdc\x33\xf4\x84\xb4\x59\x73\x64\x6e\xf1\x51\x52\x8d\xd2\x13\x93\xed\xb2\x9a\xa5\x94\xe7\xb8\x1b\x14\x7e\x64\x24\x5b\x33\xaa\xb4\xa7\x71\xdf\xad\x1f\xab\x2c\x8a\x42\x62\x41\xb4\xf0\xbd\x90\xfe\xc5\xb1\xd2\xad\xa8\xb4\x0f\x5e\x9a\x35\x35\x4b\xd7\xb8\x0f\x62\x38\xac\x71\x9f\xc2\xf4\x87\x35\xee\xa7\x75\x0c\x87\xba\xd7\xfd\x09\x95\xa6\x9c\x18\x8a\xbf\x6c\x61\x96\x0f\x52\xae\xfb\xc9\xc0\xfc\x96\x30\x9a\x2f\xf2\x5c\xa2\x32\xfc\xcf\x3e\xdd\x7d\x4a\x97\xaf\xfe\x4a\xef\xde\xbc\xfe\x7e\xf9\x2a\x4c\xed\x63\xf4\xf6\x6a\x76\xed\xe9\x58\xf0\x1f\xf7\x25\x5a\xad\x50\x21\xcf\xc3\x05\x63\x9f\xdf\x53\xa9\x74\x74\x43\x74\xf6\x10\x7d\x0e\x33\xc1\x15\x55\x1a\xb9\xfe\x85\xa8\x07\xca\x8b\x68\xe6\xdb\xc1\x02\x77\x30\x87\xb0\xcf\x98\x08\x4c\x66\x98\x8f\x44\x5d\x49\xde\x3f\x9a\x8f\x46\xa5\xd3\x21\xbb\xb6\x84\x55\x18\x8d\x24\xcc\xc7\x30\x44\xd5\x1f\xc6\x3e\xcc\x41\xcb\x0a\xaf\x3d\x09\x2d\xf7\x47\x5a\x60\x2b\xe7\x16\x8b\x9f\x77\x65\x6b\xd9\x57\xab\x21\x33\x91\x85\x27\x7c\x82\xe3\x71\x45\x98\x3a\x72\x59\x7b\xcf\x6d\x74\xad\x96\x2b\xdd\x49\xd6\x66\xb1\x8e\x42\xb3\x65\xde\x8e\x15\xc5\xfb\x8a\x67\x97\xe6\x8d\xae\xa0\x79\x05\xf3\x39\x04\x64\x5b\x04\xa7\xe2\xec\x1d\x9c\x22\xd6\x8f\x72\x6c\x32\x47\xa6\xc9\xe5\x8d\x4a\xba\xc5\x4b\x5b\x65\x44\xe9\x4b\xdb\xdc\x90\xdd\xc5\x4d\x52\x7e\x69\x93\x4a\xe7\xb8\xbd\xb8\xd1\x6a\x73\x8e\xc9\xd3\x35\x73\x54\x5f\x4f\x56\x4c\xdb\xca\xfb\xe4\xa7\xf9\x2e\x6a\x80\xd8\xe3\x23\x29\x50\x0f\x55\x94\x13\x4d\xa2\x0e\x66\x6b\x40\xb7\xa7\x8c\x79\xd7\x78\x6d\x1a\x71\xed\x7b\x09\xed\x72\x73\x74\x8d\xcd\x66\xab\xa2\x8b\xbe\x95\xcf\xba\x03\x2e\x5b\x15\x0d\x6a\xd7\x1a\xc7\x47\xe7\xdc\x6a\xfa\x52\xf3\xd8\xf8\xe8\x0e\xce\x3c\x77\xc4\x8e\x5a\xc1\xe9\xe3\xf5\x94\x93\xe4\x4a\x91\x2d\x86\x51\xa2\x1f\x90\x0f\xb0\x25\xaa\x72\xd8\xb5\x67\x62\x73\x44\x07\x52\x00\xea\xb8\x37\x8b\x52\x1e\x59\x1d\x40\x1e\x36\xaa\x48\x01\xa5\x4c\x0c\xe1\x09\x4a\x29\x64\xbd\xf4\x88\xef\xf5\x24\x6e\xc4\x16\x4f\xb1\x30\xec\xb3\x49\x45\xcb\xb7\xdc\x84\xd3\x85\x44\xd8\x8b\x0a\x54\xd5\xfe\x78\x24\x5c\x83\x16\x90\x23\x43\x8d\xd0\x0d\x06\x80\xdc\x9c\x12\x5c\x24\x30\x85\x57\x60\xac\x0d\xa0\x7b\xde\x1a\xa5\xf0\x30\xb5\xa7\xeb\x34\xa5\xf9\xae\x7e\x92\x05\x3f\x6f\x38\x3e\x2e\x8a\xa2\x25\x72\x98\x23\xc2\x83\x39\x5c\x53\x08\x48\x61\xa6\xb5\x0f\x5c\xa3\xdc\x12\x96\x7e\xf7\x26\x86\x3f\x09\xd5\x29\xfc\xef\xff\x6f\x6a\x2f\x17\x46\x63\xc8\x19\xd9\xb0\x28\x8a\x6f\x4c\x84\x17\x42\xff\x2f\x52\xe5\x34\x49\x67\x26\xcb\x30\xdf\x3d\x99\x2e\x83\xcf\x8b\x24\x4c\x37\x39\xda\xda\x33\xbf\xc3\xc3\xaf\x34\x5b\x33\x4c\x9b\xd6\x18\xc3\x6f\xa5\x10\x2c\xb5\x4d\x35\x06\xef\x5d\xbb\x3b\x66\x64\x5b\x30\x66\x47\xb5\xc0\xcf\xa4\xce\xc3\x39\x2d\xc5\xa8\x7e\x6b\x3f\xf9\x87\x02\xfa\x57\xf2\x6b\x20\xce\xc4\xec\x92\x66\x67\xf1\xaf\x31\xd7\xf0\x65\xe9\x3a\xd4\x31\x58\x9d\x78\x44\xfc\x38\x02\xa8\xe3\x06\xae\x17\xc3\xb3\xd1\x9f\x42\xde\x94\x86\x7b\xf3\x79\x69\x65\xf4\xb7\xa5\x27\x0b\xa3\xf7\xf8\xf2\xba\x68\x0f\x02\x3f\x71\xd7\xb8\xff\x1a\xfc\xb7\x53\xb7\xb3\xdb\xed\xe8\xd1\x98\xab\x97\xb9\x4b\x9d\x83\x65\x7c\x49\x73\x11\xc5\xf0\x0c\x56\x47\xb0\xdc\x1b\xdf\x31\xba\x18\xce\xa6\x4d\x94\xc8\x7d\xd2\xc0\x6e\x7a\xeb\xdb\x5c\xa6\xec\x8d\xfd\x03\x57\x9a\xf0\xcc\x5e\x59\xed\x82\xd5\x0d\x3b\x84\x1a\x37\x25\x23\x1a\x7f\x97\x2c\x85\x69\x55\xe6\x44\x37\xbb\x71\x63\x65\x1f\xf4\x86\x4d\xdb\xcc\x85\x35\xee\xef\x05\x91\x79\x57\xc2\xed\xf2\xf0\x6f\x43\xea\x60\xe9\xff\x42\x18\xa1\x68\x2b\xc5\x9d\x29\x3b\xf2\xdb\x68\xec\xf7\xb5\xff\x56\xac\x47\x81\x8e\x67\xd2\xb1\x8b\x24\x63\x42\x61\xe8\x9a\x75\x6e\x84\xf5\x91\xe9\xcc\x28\xb1\xe7\x9b\xcf\xa9\xda\x50\xa5\xc2\x69\xa3\x38\x3d\x65\xbc\x2b\x76\xb0\x2e\xd2\xd6\x55\xdc\xdf\xf4\x94\x60\x5b\x4c\x1d\x2f\x16\x66\xfa\x45\x08\xed\x04\xdd\xfd\x3f\x94\x89\x72\x3f\x0a\x50\xdd\xd1\x7c\xb7\x74\x91\x4c\xdc\xef\x76\x4a\x85\x71\x46\x24\x12\x55\xc5\xf4\xb8\xd9\x43\xe8\x6d\x90\xd3\x24\x43\x97\xcc\xa1\x1d\xd5\xc6\xf1\xdf\x01\x00\x00\xff\xff\xb3\xc7\xa9\xbc\xc9\x12\x00\x00")

func admin_http_assetsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppJs,
		"admin_http_assets/app.js",
	)
}

func admin_http_assetsAppJs() (*asset, error) {
	bytes, err := admin_http_assetsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.js", size: 4809, mode: os.FileMode(436), modTime: time.Unix(1524840913, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _admin_http_assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x3b\x7b\x6f\xdb\xb6\xf6\xff\xf7\x53\x30\xc2\x80\xb4\xf8\x45\x52\xe2\x6e\xd8\x16\xd8\x06\xb2\xac\xfd\xad\xf7\xae\x6b\x91\x6e\xf7\x81\x61\x77\xa0\xc5\x63\x89\x09\x45\xaa\x24\xe5\xc4\x37\xe8\x77\xbf\x20\x25\x5a\x92\xad\x97\x9b\x57\xd7\x3f\x1a\x51\x3a\x3c\x3c\xef\x07\x69\x4e\x0f\x7e\x7c\x77\xfe\xeb\xbf\xdf\xbf\x42\x89\x4e\xd9\xfc\xd9\xd4\xfc\x41\x0c\xf3\x78\xe6\x01\xf7\x10\x8f\x7d\x9c\x65\x33\x2f\xc2\x72\x21\xb8\x2f\x81\xe1\xb5\xcf\x63\xcf\x40\x02\x26\xf3\x67\x08\x4d\x53\xd0\x18\x45\x09\x96\x0a\xf4\xcc\xcb\xf5\xd2\xff\xce\xab\x3e\x24\x5a\x67\x3e\x7c\xcc\xe9\x6a\xe6\xfd\xcb\xff\xed\xcc\x3f\x17\x69\x86\x35\x5d\x30\xf0\x50\x24\xb8\x06\xae\x67\xde\x9b\x57\x33\x20\x31\xd4\xe6\x71\x9c\xc2\xcc\x5b\x51\xb8\xce\x84\xd4\x35\xd0\x6b\x4a\x74\x32\x23\xb0\xa2\x11\xf8\x76\x70\x84\x28\xa7\x9a\x62\xe6\xab\x08\x33\x98\x9d\xec\xa0\x21\xa0\x22\x49\x33\x4d\x05\xaf\x61\xda\x01\xc3\xb9\x4e\x84\xdc\x81\x60\x94\x5f\x21\x09\x6c\xe6\xd1\xc8\x20\x48\x24\x2c\x67\x5e\x10\x84\x41\x10\x2e\xf1\xca\xbc\x0c\x68\x24\xbc\xf9\x33\x03\xad\xa9\x66\x30\x3f\x2f\x04\x76\x61\x05\xf6\xcb\xff\x23\x4c\x52\xca\xa7\x61\xf1\xd1\xc2\x1d\xf8\x3e\xfa\x19\x6b\x50\x1a\x45\x22\xcd\x28\x03\x82\x30\x27\x28\xa5\x9c\x2e\x29\x10\x74\xfe\xe1\x03\xf2\xfd\x2d\x0a\x94\x5e\x33\x50\x09\x80\x76\x74\x84\x61\x8a\x6f\x22\xc2\x83\x85\x10\x5a\x69\x89\x33\x33\x88\x44\x1a\x6e\x5e\x84\x2f\x83\x49\x70\x1c\x46\x4a\x55\xef\x82\x94\xf2\x20\x52\xca\xab\xa8\x79\x67\x05\x84\x19\xd2\x09\xa4\xf0\x80\x6b\xfb\x76\x81\x2d\x0a\x7a\xd7\xc1\x59\xb6\x45\xec\x4f\xbf\xbe\xfd\xf9\x1b\xa4\x12\x9a\x5a\xa9\x5d\x80\xca\x04\x27\xc1\xa5\x42\x6f\x5e\x7d\x87\x54\x9e\x19\xb3\x41\x62\x59\x02\x02\x83\x14\xb8\x56\x85\x88\x81\x50\x8c\x3e\xe6\x20\x29\x28\xc7\xe7\x81\xef\xff\x4e\x97\x88\x69\xf4\xe6\x15\xfa\xfe\x0f\xfb\xae\xb0\x1a\xa4\x64\x34\xf3\x8c\x21\xab\xd3\x30\x14\x4a\x05\x25\xd7\x86\x51\xe3\x30\xdf\xa8\x84\xae\xc2\x97\xc1\xb7\xc1\xa4\x1a\x5b\xf6\x2e\x95\x37\x9f\x86\x05\x9a\xb1\x18\x65\xc1\x4a\x78\x12\x7c\x1d\x4c\xdc\xa8\x1d\xdb\xc1\xef\xc0\x09\x5d\xfe\x61\x58\x98\x86\x85\x47\x3e\x9b\x2e\x04\x59\x23\x29\x98\x31\x7c\x11\xe5\x86\x6f\x0f\x55\x92\x7b\x4d\x6f\x80\x20\x8e\x57\x0b\x2c\x1d\xf3\x84\xae\x50\xc4\xb0\x52\x33\xaf\xfc\x50\xfc\xf1\x29\x5f\x81\x54\xe0\x95\xf8\x38\x5e\xd1\x18\x5b\x3f\x32\xf3\x9a\x33\x8d\xdb\x60\xca\x41\x96\xdf\xda\xf0\xfa\x86\xc8\x1a\x04\x42\x53\xbc\x05\xb1\x90\x98\x93\x8d\xe6\xbd\x6d\x57\x9a\x86\x78\x83\x3e\x24\x74\xd5\xb3\x56\x24\x18\xc3\x99\x02\xe4\x1e\xea\xcb\xe6\xac\x06\xed\xd8\xe5\x78\x55\x83\xb1\x56\xe9\xa0\x70\xa4\xe9\x0a\x1a\x5f\x2d\xf1\x1b\x3a\x7f\x12\x29\xd4\x88\x2b\x08\x64\x74\x0b\x5d\xd7\xfc\x05\x26\x6f\x41\x4b\x1a\xa9\x70\xf2\x75\x12\x5c\x2a\x23\xe2\x1f\xb0\x31\x56\xfb\xb6\x17\xf3\x34\xcc\x59\xbb\x50\x0e\x7c\x3f\x0c\x38\x5e\x55\xb2\xf0\xfd\x79\x05\x53\x3e\x3c\xeb\x52\x64\xa9\xf6\x14\xd3\x22\x19\x98\x2f\x52\x30\x06\x72\xe6\xbd\xc5\x94\x9f\x6b\x56\x33\x84\x36\x55\x48\x71\x5d\xcc\x64\x02\x5f\xd5\xb5\xce\x40\x6a\xf3\x41\x42\x06\x58\xcf\xbc\xe2\x05\xe5\xc8\x3e\x28\x6f\x7e\x7b\x6b\x9f\x82\x54\xc5\x9f\x3e\x4d\x43\x3b\xa8\x21\x68\xd0\xcb\xfc\x94\xf8\x27\x93\xa6\xee\x92\xc9\xfc\x1f\x98\x51\x62\xed\x75\x1a\x26\x93\x2d\xd9\x6b\xbc\x60\xe0\x70\x14\x03\xfb\xbf\xe1\x92\x00\x57\x40\xb6\xb4\x6d\xe6\xb8\xb4\xb7\xfd\x5e\xee\xbe\xb4\xe0\xa5\xf9\x4e\x43\x9d\x74\x41\x94\x7a\x47\x93\xe0\xb8\x0f\xac\xe4\x05\x90\x90\x04\x64\x3b\xe4\x34\xdc\x25\xc4\x40\xb6\x10\x3d\xd5\x26\x4e\xec\xc1\x0a\x99\xdf\xde\x46\x82\x2f\x69\x1c\x54\x62\xfd\x93\xc1\x0a\xd8\x9f\x0c\x62\x1c\xad\x8d\x9a\x74\x8b\x74\x86\x66\xa7\x93\xe3\xfd\xa6\xc2\x9f\x56\x06\x5d\x93\x3a\xa4\xb0\xcd\xef\x34\xb4\xea\xae\xbb\x51\xcd\x88\xc7\x98\x58\x61\x64\x3f\x30\x1c\x5d\x31\xaa\xf4\x93\xd9\xd8\x5b\xac\xa3\x04\x65\x12\x96\xf4\xa6\xd7\xd2\x2c\x9c\xca\x17\x4a\x4b\xca\xe3\x61\x50\x09\x31\xf4\x62\x3c\x8b\x8c\x1e\xd5\x3d\x59\x63\x3d\x1e\x2c\x4c\x2c\xb0\xb2\x0a\x16\x4e\xc0\x3b\xb2\xea\x94\x8a\xb3\x9b\x45\x50\x48\xa5\xdb\xbe\x2a\xc8\x8d\x5c\xc6\x00\x5b\xc9\x0c\x00\x6e\xf4\x0e\x37\xda\x8f\x80\x6b\x93\xf8\xa6\xb8\x08\x87\x34\xba\x9a\x79\x12\x52\xb1\x82\x8d\x01\x3d\xff\x8a\x72\x02\x37\x2f\xbc\xf9\x74\x93\x77\x62\xb6\xce\x12\x53\x63\xa2\xcd\x93\x5f\x4c\xf3\x23\x2a\x23\x06\x5e\x38\x37\xc9\xe1\x4e\xae\xd0\xe2\x0c\xa8\xb4\xee\x0b\xb8\x96\x54\x83\x54\x2d\xd6\xbd\x14\x32\x2d\x6b\x67\x59\x82\xbd\x16\x32\xf5\x1c\xed\xe6\xbb\x8f\x09\x41\xf6\x21\x96\x22\xcf\x50\x82\x95\xbf\x04\x20\x0b\x1c\x5d\xb9\xfc\xb2\xb4\x93\x78\xec\xab\x7c\x91\x52\x93\x0c\x08\x71\xeb\x3e\x7f\xe1\x21\x2e\x56\xa5\xe3\x3f\x89\x7b\xbd\x63\xa4\xcf\x07\x7e\x81\xeb\x7e\x4f\x7a\x22\x0f\x92\x95\x07\x39\xed\xa8\xfd\x3d\x48\x06\x82\x91\x61\x8f\x90\x01\x87\xeb\x31\x60\x29\xbe\x2f\xb7\xd9\x58\xc8\xd3\x79\x0d\xda\x3f\x91\x36\x58\xec\xf2\x8b\xf6\x89\x76\x32\xe5\x59\x6e\x0b\xa7\x54\x10\xd3\x2d\x71\xb8\x76\x82\x08\xde\x31\xe2\x95\xfe\x28\xcc\x63\x7d\x95\xb2\x74\xf3\x50\xc6\x70\x04\x89\x60\xc4\x14\x71\x16\x4c\x9a\xfe\x5c\x42\x87\x46\x90\x4b\x85\xc6\x3f\x13\x71\xdd\xf4\x75\x63\x1d\xc1\x57\x94\x5b\x07\xed\xa1\xdb\x62\x51\x19\xe6\x7d\x68\x40\x4a\x21\x83\x0c\x6b\x0d\x92\x7b\xf3\x57\x37\x19\x44\xda\x34\x2b\x82\xfb\x90\x66\x7a\x8d\x5c\xe2\x32\x98\x7a\xc8\x6d\x26\xf2\xe6\xa7\x11\x96\x77\xcf\x6a\xf9\x05\xae\x9d\x5a\xb8\x79\x1c\x56\x8b\x01\x7b\x0a\xfa\xf5\x3a\x33\x44\xe6\xe9\xc2\x14\xff\xed\xdc\xbc\xc5\x37\x8e\x9b\xd4\x3c\xb6\x72\x73\x27\x93\x4a\xf1\xcd\x7d\x98\x94\x45\xd3\x65\x52\x18\x15\x5c\x22\x2d\x10\x01\x2e\x34\xa0\x14\xdf\x20\x09\x56\x0f\xc5\x5e\xc1\x73\x8c\x32\xa1\xa8\xe9\xfb\x4a\xe8\x23\x24\x24\xf2\x4f\x4c\x3e\x43\x5c\x20\x46\x53\xaa\x5f\x3c\x94\x35\x76\x7c\x58\xe4\x5a\x0b\xee\xc4\xbe\xd0\x1c\x2d\x34\xf7\x55\x6a\xff\x64\x92\xa6\x58\xae\xed\xf3\x82\x09\x93\x62\x0b\x9d\x16\x99\xd5\xea\x94\x50\x65\x92\x02\xd9\x12\x57\x25\xf1\x33\x42\xa6\x61\xb1\xcc\x3e\x64\xb7\x45\xcd\xbd\xaa\x8d\xd0\xd8\xd0\x6e\x05\x72\x16\xc7\x12\x62\xac\xc5\x50\x0d\x82\xe3\xf8\xde\xca\x8f\x6a\xd1\x2f\xa0\x00\x79\x9d\xf3\xa8\x68\x62\xbb\xeb\x88\x8b\xa1\x52\xfd\x5d\xae\x8d\x8b\x1b\x66\xb1\xee\x03\x3c\xc7\x51\x02\x7d\x00\x6f\x4c\x4a\x5e\x61\x86\x9e\x2b\x88\x5e\xf4\x41\xfe\x13\x53\x3d\x0c\xf5\xa3\x14\xd9\x05\xee\x2d\xa2\x1e\xae\x4a\xc2\x55\x95\x84\x2b\x43\xdb\xbf\x4e\xc2\xc1\x32\xe7\xc3\x05\x10\x1e\xd5\x39\x58\xc0\x77\xb9\x7e\x9d\xea\x31\x90\x56\x63\x63\x00\x9d\xe6\xc6\xc0\x1a\xdd\x8d\x81\x2b\xb5\x77\x4f\x45\x5d\xcd\xef\x1e\xa5\xac\xfb\xa2\x0b\xbb\xb3\x38\x0e\x5e\xe7\xdc\xa5\xdb\xa5\x79\x1c\x2e\x1e\x5c\xb4\xa8\xb2\xb0\xc1\x5a\xa6\xc0\x99\x67\x23\xd9\x59\x1c\x1b\xb0\x3e\x7a\x1a\xf9\xb9\x0c\xad\xc6\xc6\x3f\x33\x35\x37\x30\x74\x66\xe5\x52\xfb\x54\x70\xb4\x2c\xd9\x38\x45\x78\x15\x1f\x21\x02\x4c\x63\xf3\x47\xd2\x15\x1c\x21\x86\x95\x3e\x32\x29\xfb\x08\xa5\x94\x1f\x21\xa5\x09\xac\x4c\x76\x56\x79\xfa\xd7\x28\x0f\x8d\x72\x6d\xd0\xf6\x36\x2d\xb4\x1d\x0c\x2b\xb8\x04\xdc\x51\x6a\x81\x6d\x5f\x95\x5a\x6c\x77\x54\x6a\x89\xa3\x4b\xad\x16\x33\x92\x10\xe7\x0c\x4b\x04\x37\x99\x04\xa5\x6c\x3e\xfb\xab\x28\xaa\x88\xc5\x9b\xe6\x2a\xd7\x4b\x33\x1a\xd1\x5f\xe5\x7a\xa8\x18\x7e\x10\x4e\x1a\xf5\x7c\x94\x40\x74\xb5\x10\x37\xde\x2e\x5f\x36\x73\x38\xb6\xca\x41\x1b\x57\x5f\x58\x33\x62\x48\x77\xb9\xcc\x51\x4f\x37\xe3\x56\xb5\xd4\xbd\x25\xfc\xcf\xef\xc7\xfe\xf7\x7f\xfc\xdf\x57\xe1\xde\xde\xe2\x56\xb9\xa3\xc3\x54\x68\x06\x1b\x94\xe7\x94\x23\x05\xa6\x9e\x54\x2f\x6a\xdd\xca\xc7\x1c\x73\x4d\xff\x4b\x79\x8c\x1c\xb2\x2f\xd6\x9d\x06\x34\x69\x2a\x0d\xa7\xc5\x6b\xfb\xfc\x90\x1a\x34\x2b\xdc\x51\x7b\x05\x8a\xcf\xd5\x9c\xe9\x33\x0d\x86\x2f\x4d\x5d\x63\x63\x46\x59\xf0\x39\x8d\x99\x21\xb2\xe3\x7b\x0a\x1c\x0f\xdf\xf7\x3a\x3d\x7e\xb1\x2d\xef\x85\xc8\xb5\xdd\xe7\xea\xdd\x72\x17\xb9\x86\xfb\xdb\x6f\x37\xd8\x9e\xb8\xd7\x2d\x7a\x3d\xdb\xe4\x15\xbf\x55\xb8\x86\x43\xc6\x90\xb1\x4a\x53\xea\x2b\x94\x80\x74\x3f\x4c\x69\x9b\x69\x79\x40\x57\xb0\xee\x6d\x94\x2d\x90\xb1\x8e\xe1\xf3\xaf\x27\x39\x52\x23\xc4\x94\x47\x7d\x20\x1f\x32\x21\x58\x1f\xc0\x7b\x1a\x5d\xb1\x3e\xfe\x50\x24\x98\x89\x3e\xb3\x09\x7a\xa4\x13\x08\x23\xf5\x3d\xda\x6a\x13\x07\x8a\xe8\xdb\xd3\xfa\x65\x0c\xaf\x3d\x84\x25\xc5\x7e\x42\x09\x01\x3e\xf3\xb4\xcc\xc1\xfe\x4a\xc6\x84\xd6\x9e\x33\x65\x87\x96\xf2\xa5\xf0\xec\xc1\xc4\x15\xf4\x1f\x60\xef\xce\x30\x36\xb4\xe7\x94\xd4\xd8\x00\xc8\xc1\x53\xc9\xde\xc9\x23\x0e\x2a\x7b\xe7\x0f\xec\x40\x6c\xcf\xdd\xd8\x8a\xf7\xd2\x1b\x2d\xd2\x6a\xd2\xa4\xfd\xf0\xc6\x86\x1b\x2b\xf5\x87\xef\xf0\xb5\xac\x5b\x25\x31\x56\x29\x03\x02\x4a\x53\x5e\xff\x31\x53\x0b\x4f\x73\x13\x86\x06\xed\x30\x4a\x60\x25\x0d\xa1\x34\x4e\x74\x9f\x41\xfa\x7e\xa7\x00\xcb\xe5\x3e\xff\x6b\x21\x62\x4b\xe3\xed\x21\xc1\x3c\x06\x79\x78\x8a\x0e\x48\x20\x38\xa3\x1c\x8e\xd0\xa1\x51\xcc\xe1\x29\x72\x6f\x3e\x19\xab\x20\xa3\x4d\xf2\x5e\x16\x19\x7b\xc6\x7e\xd7\x75\x46\x1d\xcf\x7f\xe6\x1a\xb8\x08\xd0\xf7\x8f\xbd\x75\x8f\xac\x1d\xbf\x35\xc0\x4d\x5d\x4a\x02\x65\x32\x82\xd7\x67\xa3\x09\x21\x5e\xd8\xea\xb9\x3d\x1e\x5d\xe3\x01\x39\x26\xd0\xc3\x71\x91\xd9\xb4\xd5\xcb\x86\xb8\xf2\x15\x8d\xf9\x63\xb1\xd2\x1f\xc4\x7e\xac\x42\x48\x11\xca\x8e\x9e\xf2\x28\xfa\x33\x7e\xd1\xf5\x18\x19\x76\xbf\x36\xa4\xf5\x20\xd5\xe4\x8a\xe0\xef\xb0\x76\x2d\xc7\xaf\xeb\xac\x63\x9f\xa2\xda\xf2\x6c\xee\x98\xb9\xc2\xb0\x75\xd5\x47\x64\xa2\xa0\x7c\x90\x8b\x06\xf1\x0a\x38\x39\x63\xcc\x96\x8f\x7d\x7b\xba\x76\x05\x8b\xb4\x83\xaa\xe6\x89\xab\xeb\x1e\x2c\x4d\xc3\x0d\xf1\xf6\x49\xeb\xd6\xf4\xae\x66\xb8\x4e\xfc\x91\x1d\xbd\xa6\x52\xe9\x72\x2c\x24\x32\x75\x3d\x55\x1a\xb8\xfe\x09\xab\x64\xe0\x70\xbf\xb3\x25\x7e\x44\x15\xbe\xb7\x79\xd2\x29\xd1\x8d\x86\xd5\x58\x42\x3e\x31\xf5\x1f\x5c\x02\x76\x0c\xd4\x5e\x0c\xf3\x50\x01\x3f\x31\x1b\x8d\x8d\xf3\x8b\xb1\x1b\xe7\x17\x7b\x6f\x9c\x3f\x22\x4b\x65\xdf\xe7\x98\xda\x0c\xf7\x08\x73\x15\x8a\x9d\xf3\x9e\xf2\xcb\x63\xf2\xd8\xbb\x91\x54\x5a\x63\x51\xb6\x94\x96\xd8\xa8\x61\x06\x77\x91\x9e\x90\xec\xf7\x65\x9d\x52\x86\x80\x66\xd5\x32\x92\xf0\x8e\x8c\xfc\xe0\x7b\x5f\x55\xe0\x1e\xbd\xfb\x75\xff\x3f\x2d\xdd\xde\xf7\x6a\x5e\x5a\xa8\x06\xb5\xfb\x09\x9b\x2b\x0b\x07\xbe\x8f\xc2\xcd\x15\x05\xbb\x1b\xe4\x5e\xff\xe0\x6e\x18\xa1\x48\x48\x40\x7f\xc3\x2b\xfc\xc1\xde\x96\xb1\xc8\x66\x7b\xff\xab\xdd\x0d\x42\xef\x8d\xa3\x11\x84\x35\xd2\x09\x20\xe0\x04\x89\xa5\x7d\x74\x77\x6c\x90\x12\x76\x9c\xe1\x18\x14\x62\x02\x13\xb4\xc4\x4a\xc3\xe6\x92\x4d\xdb\xdd\x1f\x7c\x89\x6f\x82\x58\x88\x98\x01\xce\xa8\xb2\x17\x80\xcc\xbb\x90\xd1\x85\x0a\x2f\x3f\xe6\x20\xd7\xe1\x49\x70\x72\x12\x9c\x94\xa3\xe1\x7b\x45\xe3\x6f\x65\x5d\x6e\x5f\x08\x6b\xe2\xed\x22\x3a\x12\x04\x02\xcc\xed\x69\xde\xa5\x0a\x84\x8c\xc3\x93\x60\x12\x9c\x1c\x87\xe5\xcb\xf1\x77\x9f\x06\x51\xf9\x12\x94\xc8\x65\x04\x63\xf8\x8e\x08\xbf\x54\x41\xc4\x44\x4e\x96\x0c\x4b\xd8\x12\xa7\x43\x99\x53\xbf\x92\xc4\xb1\x11\xee\x71\x58\x7f\xe7\xeb\x8c\xa9\x01\x59\xd8\x2b\x69\x5d\xe4\x14\xee\x67\x1a\xa0\x90\xc7\xbe\x86\x34\x63\x58\x83\x87\x28\x99\x79\x79\x46\xb0\x2e\xb6\x3f\xde\x0a\x82\x59\x90\xe8\x94\xb5\xdc\xa6\x4a\xcd\xc7\xed\xeb\x52\xd3\xe4\x65\xf3\xbb\xbd\x50\xe8\xcd\x7f\xb3\x48\x91\xf5\xed\x53\x74\x7b\x6b\x1f\xdc\x8e\x56\xf2\xb2\xe1\x4c\x68\x6b\x1b\x79\x79\x87\x1d\xe4\x9d\xad\xe2\x5d\x0e\x4c\x3c\xa8\xdf\xbb\xaa\x01\x54\x6b\x6c\x5d\xba\xc2\x0b\x60\x86\x82\x99\x67\xb2\x98\x37\x7f\x5f\xe4\xb2\x69\x68\xbf\x34\x60\xb7\xd3\x69\xc1\xb9\x99\xe0\xc2\xb3\x45\x31\xa6\x3e\x2b\x6b\xd7\x31\xe5\x41\xb3\x94\x36\x48\x6d\xe9\xdb\x55\x45\x6f\xd5\xce\x35\xf8\xfe\x13\xf3\xbe\x93\xf2\xed\xcb\x2a\xdd\x77\x57\x46\x88\x19\x13\x22\xbd\x6a\xe7\x79\xac\x98\xcd\x04\x27\x66\x8b\x62\x84\x98\x4f\x26\xdf\x06\xc7\xc1\x71\x70\x72\x3a\x39\x3e\xfe\xba\xf7\x37\x2a\x2d\x35\x4b\x8b\xe0\xcd\xc2\xfb\x08\xbe\x80\xef\x10\xfc\x29\x3a\x4c\x84\xd2\xa7\x99\x90\xfa\x70\x3f\xa1\x77\xde\x46\x2c\xfc\x60\x29\x44\x73\x27\xa4\x2d\xa9\x9b\x14\x4e\x60\x89\x73\xa6\xbd\xda\x6e\x43\x84\x79\x04\xec\xf9\x0b\x6f\x7e\xce\x84\x82\xdd\x5c\xdd\x51\x20\x94\x95\xc1\x56\x05\xb0\x6c\x24\xff\xda\x32\xe2\xca\x2c\x51\xc4\x92\xed\x35\x1a\x49\xd9\x25\xef\x2a\xf6\x4d\xc3\x22\xef\x4f\xc3\xe2\x22\xf9\xff\x02\x00\x00\xff\xff\xf8\x76\x31\xcb\x59\x3e\x00\x00")

func admin_http_assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsIndexHtml,
		"admin_http_assets/index.html",
	)
}

func admin_http_assetsIndexHtml() (*asset, error) {
	bytes, err := admin_http_assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/index.html", size: 15961, mode: os.FileMode(436), modTime: time.Unix(1524840913, 0)}
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
	"admin_http_assets/app.css":    admin_http_assetsAppCss,
	"admin_http_assets/app.js":     admin_http_assetsAppJs,
	"admin_http_assets/index.html": admin_http_assetsIndexHtml,
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
	"admin_http_assets": {nil, map[string]*bintree{
		"app.css":    {admin_http_assetsAppCss, map[string]*bintree{}},
		"app.js":     {admin_http_assetsAppJs, map[string]*bintree{}},
		"index.html": {admin_http_assetsIndexHtml, map[string]*bintree{}},
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
