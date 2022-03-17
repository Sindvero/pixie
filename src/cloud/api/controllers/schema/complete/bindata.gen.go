// Code generated for package complete by go-bindata DO NOT EDIT. (@generated)
// sources:
// 01_base_schema.graphql
// 02_unauth_schema.graphql
// 03_auth_schema.graphql
package complete

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

var __01_base_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\xc5\xa0\x14\x54\x5c\x2a\x10\x4a\x49\x4f\x81\xe0\x07\x1c\x7b\x38\x47\x72\xbc\x3e\xef\x46\x47\x84\xf8\x77\x94\xcb\x5d\x77\xd5\x6c\x31\xf3\xb4\x4f\x43\xe2\xec\xf1\xeb\x80\xd3\xc2\xb6\x0e\xf8\xd8\xc2\x01\xf3\x62\xde\x26\x29\x03\xde\xaf\x97\xfb\x73\xae\xc3\x57\x22\xb4\x32\x20\x0a\xb5\x3c\x1a\x7c\xce\x72\x06\xe7\x6a\x2b\x6c\xad\xd4\x83\xeb\xf0\x29\x38\x13\xa1\xd1\x1b\x51\x7d\x0e\x4c\x92\x23\x9b\x22\xb1\x11\xbe\xc4\xeb\xce\x12\x95\xfb\x0e\x26\x18\xe9\x3a\xf0\xc7\x58\x22\x23\xc6\x15\x62\x89\x0d\xdf\x53\xde\xb9\xc9\xac\xea\xd0\xf7\xc7\xc9\xd2\x32\x1e\x82\xcc\xfd\xb1\xf9\x9a\x4e\xf9\x96\x4f\xdb\x73\xfd\xa4\xba\x50\xfb\xe7\x97\x57\xe7\x36\xf8\xae\x75\xf1\x2c\x22\x75\xc0\x9b\x48\xa6\x2f\x0f\x9b\xd4\xa5\x70\xb3\xbc\xdf\xf9\x0f\x00\x00\xff\xff\x6f\xc4\xb8\xef\x28\x01\x00\x00")

func _01_base_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__01_base_schemaGraphql,
		"01_base_schema.graphql",
	)
}

func _01_base_schemaGraphql() (*asset, error) {
	bytes, err := _01_base_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_base_schema.graphql", size: 296, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_unauth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\xb1\x0a\xc2\x40\x0c\x86\xf7\x7b\x8a\xdc\x56\x5f\xe1\x36\x1d\x84\x0e\x0a\xa2\x9b\x38\x84\x9a\xd6\x60\x2f\x29\x77\xb1\x58\xc4\x77\x17\x0b\xa5\x15\xb7\x9f\xe4\xfb\xbf\x9f\x9e\x46\x72\x05\x1b\x3a\x82\xc3\x83\xd2\x00\x2f\x07\x80\xc9\xb8\xc6\xca\x72\x31\xa5\x3d\x46\x0a\x70\xb4\xc4\xd2\xf8\x55\x80\xf5\x44\x94\x52\xab\x77\x00\x3d\x25\xae\x87\x52\x7a\x36\x3a\xe9\x9d\xa4\xe0\x39\x2f\x9b\x1b\xd5\x96\x50\xbc\x7b\x3b\x37\xce\xfe\xa8\xc6\x79\x36\x8a\x39\xc0\x79\xfa\xf8\xcb\x3f\x3d\x82\x3d\xa5\xcc\x3a\xeb\x1d\x40\x75\x43\x69\xa8\xd5\x66\x79\x34\x8e\x94\x0d\x63\xb7\xcb\x01\xb6\xad\xa2\x7d\x85\x9f\x00\x00\x00\xff\xff\xef\x77\x02\x34\xfc\x00\x00\x00")

func _02_unauth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__02_unauth_schemaGraphql,
		"02_unauth_schema.graphql",
	)
}

func _02_unauth_schemaGraphql() (*asset, error) {
	bytes, err := _02_unauth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_unauth_schema.graphql", size: 252, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_auth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\xdf\x8f\xdb\x36\xf2\x7f\xf7\x5f\x31\x69\x1e\xea\x05\xf6\x1b\x14\x5f\x5c\x8b\x83\x9f\x4e\xb5\x95\x46\x97\x5d\xaf\xcf\x76\xd2\x2b\x82\x20\xa0\xa5\xb1\x45\x58\x26\x55\x92\xf2\xae\xef\xd0\xff\xfd\x30\x24\x25\x91\xb2\xb6\xe9\xf6\xee\x4d\xfc\x35\xf3\x99\x9f\x9c\xa1\xf0\xc9\xa0\x28\xc0\x5c\x6a\x84\x7f\x34\xa8\x2e\xf0\xef\x09\x40\xa3\x51\xcd\xe0\x83\x46\x95\x89\xbd\x7c\x35\x01\x90\xea\x30\x83\x07\x75\x68\xc7\xb4\x63\x83\xc6\x70\x71\xd0\x6e\x67\x3b\x6a\x57\x13\x63\x14\xdf\x35\x06\xfd\x7a\x3f\xf6\xf4\x68\x52\xcf\xe0\x53\xc7\xe6\x33\x2d\xe4\x55\xa3\x0d\xaa\x29\x2f\x66\x90\x2d\x5e\xdd\xcc\x60\xee\x66\x5a\xce\x7e\xc3\x8f\x97\x25\x3b\xe1\x54\xb0\x13\xce\x60\x63\x14\x17\x87\xe7\x37\x13\x9b\x70\x25\xe4\x34\x97\x42\x60\x6e\xb8\x14\xd7\x3c\xfb\xb5\x9e\x20\x4f\x94\xe1\x7b\x96\x9b\x29\xf3\x1f\xdb\x4b\x8d\x33\x48\x82\x91\x25\x71\x97\xb5\x53\x74\x90\x35\x46\xe6\xf2\x54\x57\x68\x70\xca\x45\xdd\x98\x16\xf6\x2d\xe4\x8d\xd2\x52\xad\xa4\x9e\x41\x26\xcc\x2d\x30\xcb\x72\x06\x49\x70\x26\xb1\x73\x44\xfc\xb6\x45\xfe\x21\x5b\xb4\x34\x6e\xe2\xcd\x6b\xd4\x4d\x75\xc5\xf6\x2d\xc7\xaa\x18\xf2\xde\xd3\xa4\x97\x20\xd8\x9b\x0a\xc3\xcd\xe5\x3d\x17\xc5\xed\x04\x00\x40\xe1\xaf\x0d\x57\x58\x24\xea\x40\x9b\x49\xa1\xe3\xdb\x3f\xff\x01\x78\x16\x48\x8b\x71\x02\xf0\x1a\x36\xb9\xe2\xb5\x39\x1d\x14\xa0\x28\x6a\xc9\x85\xd1\xb7\xa0\x70\x8f\x0a\x8c\x84\x42\xe6\x1a\xb8\x80\xbc\x92\x4d\xc1\x6a\xfe\xa6\x56\xd2\xc8\x09\x40\xc5\xcf\xf8\x91\xe3\x23\xc1\xb9\xf3\xdf\xf7\x68\x58\xc1\x0c\x73\x46\x6e\x77\xcc\xa5\x30\x28\x8c\x0e\x6c\x7c\x37\x58\xa2\xed\xda\xe2\x20\x72\x0e\x51\x4c\xcc\xad\x8e\x90\xda\x44\x0b\x5e\xa6\x05\xd6\x95\xbc\xc0\x11\x2f\x7a\x02\x50\xd8\xd1\x09\x85\x79\x8f\x17\x62\xb0\x08\x27\x62\x3e\xd1\xde\x80\x4d\x74\xc4\x73\x49\x56\x59\xcb\x82\xd5\xdc\xd3\x4e\x56\xd9\x15\x51\xb7\x1a\x50\x73\x9b\x3c\x99\x55\xd5\x1c\xb8\x98\x00\xd4\xf6\x43\x4f\x8f\x5c\x14\x33\x3f\x4d\x76\xbd\x99\xc1\x27\x37\x72\xe4\x14\x92\xac\x5c\x0a\x37\x39\x97\x62\xcf\x0f\x96\x7a\x1f\x8c\x9f\xc2\x35\x3a\xf6\xdb\x64\x12\x26\x9c\xfb\xc6\x30\x22\x61\x73\xce\x5c\x21\x33\xe8\x03\x2f\x0a\x64\xf8\x5b\x81\xb5\xc2\x9c\x19\x2c\xa6\x0a\x99\xa6\xd8\xf8\xc6\x6f\xd0\xc0\x14\x82\x90\x8f\x90\x5b\x02\x05\x9c\x39\x83\xfa\xc9\x2b\xf1\x9b\x9b\x09\xc0\x87\xba\x60\x06\x3f\xf2\x7f\x71\x1b\xd2\x04\xd4\xfb\x28\xb9\x68\xb6\x78\x75\x0b\xe7\x60\x71\x06\x69\xc1\x0d\xdb\x55\xd1\x91\x91\xec\xe2\x20\x47\x56\xb9\x32\x12\xc0\x02\xc9\xe7\x17\xcf\xd8\xf4\x47\x29\x2b\x64\xa2\x27\xe7\xcc\xd2\x9b\xa7\x25\xe0\xc6\xe3\x27\x9d\x80\x61\x16\x9e\xea\x2e\x39\xb7\xc2\x44\x49\xfa\xe6\x3a\x69\x6f\xd0\xc4\x79\x7a\xca\x82\x14\x1e\x52\x09\x52\xf9\xcd\x58\x72\xcf\xc4\x99\x3b\x38\x53\x3c\x31\x5e\x75\x3e\x41\xe9\x46\x69\xb3\x0c\x93\xf6\x2d\x54\x6c\x30\x75\xd3\xde\x3d\x44\x26\x96\x6f\x85\xea\xc4\xb5\xe6\x52\xe8\x29\xdd\x32\x9d\x01\x9b\x78\x31\x06\x1c\x2c\xf4\xc4\x43\x1b\x3e\xa8\xc3\x54\xaa\xc3\x10\x45\xb6\xe8\xb9\x3f\xa8\x43\xa7\x5c\xa9\x0e\x1d\x63\xd9\xcf\xf7\x4c\x83\xcd\x44\x27\xb8\x38\x1d\x3f\x27\xda\x56\x1e\x51\x04\xc4\x6e\x3a\xde\x13\x80\x35\x9e\xe5\x11\x93\xaa\x0a\xf6\xea\x78\x73\xe0\x01\x6b\x3c\xc9\xb3\x95\xf5\xad\x92\x27\x12\x27\xd0\x4e\xb8\x35\x8e\x76\x27\xda\xfa\x6b\xb1\x7c\x0b\x28\x48\xac\xa2\x27\x74\x0b\xb9\xdd\x16\x08\x1d\x9e\xd6\x11\xd3\xdf\x26\x13\x1b\xf0\xad\xe2\x6d\xc0\x7b\x4f\x9e\x00\x44\x97\xf8\x04\x20\x76\x1a\x4a\x4a\x3c\x37\x8d\x8a\xf6\x0c\xad\xe5\xa6\xfa\x2b\x87\x26\xb8\x4e\xea\x5a\xc9\x73\x88\xbb\xc3\x92\x2d\xd2\x15\x33\xa5\x85\x92\x2d\xd2\x21\xb1\x9a\x99\xb2\x1f\xb7\x87\xbc\x21\xbf\x82\xbf\x90\x27\xc6\xc5\x90\xa2\x53\xa1\x43\xc4\x2a\x1d\x99\x8f\x17\x48\x60\x28\x77\x7b\x5c\x3e\x5d\x76\x6a\x6b\xdd\xc9\xb2\x66\x82\x55\x17\xc3\x73\xfd\x50\x1b\x49\xb7\x79\x44\xca\xc1\x0a\x0f\xf7\xe1\x69\x8f\x1b\xd9\xa8\x0d\xa2\x78\xee\x9c\x2d\x11\x9e\x89\xf8\x71\x02\xe3\xa7\xfe\x10\xe6\x0e\x68\x7c\x69\x0d\x54\xec\x93\x7b\x62\xee\xf5\x0c\xde\x56\x92\x19\x77\x51\xea\xfc\xda\x48\x8e\xd0\x80\xc0\x91\x72\x6a\x6f\x8c\x97\xd0\x1b\xbd\xa9\xff\x0b\x7c\x11\xbd\xff\x09\x4c\x14\xcd\x69\xa4\x7c\xdb\x18\x66\xd0\x32\x48\xd2\xcd\x97\x0f\xcb\xf7\xcb\x87\x9f\x97\x7e\xb4\x4a\x97\x8b\x6c\xf9\x93\x1f\xad\x3f\x2c\x97\xfd\xe8\x6d\x92\xdd\xa5\x0b\x3f\xd8\xa6\xeb\xfb\x6c\x99\x6c\xd3\xc5\x28\xa7\xbe\x2e\x75\x8c\x92\x6d\xc0\xe8\x35\x24\x02\xb0\xe0\xc6\x97\xb4\x20\x73\xaa\x75\x81\xef\x81\xd9\x8c\x0d\x25\xd3\x70\x92\x05\xdf\x73\x2c\xc0\x94\x08\xce\x8b\x0c\x3e\x19\xd8\x5d\x80\x0b\x8d\x8a\x7c\x08\xa4\x82\x82\xee\x41\xfa\xce\x4b\xa6\x58\x4e\x97\xff\x1b\xcb\x64\x5b\x72\xaa\x0f\xf3\xaa\x29\x50\x53\x69\x61\x0f\x08\x4b\xef\x88\x97\x9d\x64\xaa\x00\x26\x0a\xa8\x99\x76\x04\xe4\xe9\xc4\x44\x61\x8f\x13\xe2\x74\x91\x6d\x1d\x5c\xd0\x58\x61\xde\xe3\x15\xd5\x65\x1c\x74\x5e\x4a\x8d\x02\x98\x88\x4a\x6c\xd0\xcd\xe1\x80\x9a\xce\xbe\x69\x61\x15\x9c\x2a\x17\x0d\xb6\x62\x7d\x6d\x41\x45\x47\x4c\xc9\x0c\x70\x03\xba\x94\x4d\x55\x00\xe5\x71\xbb\x89\x58\x7d\xab\x7d\x73\x40\x65\x30\x4d\x0a\x52\x0c\xa3\x1c\x52\x2b\x4e\xd6\x35\x6c\xd7\x4a\xb1\x49\xef\xd2\xf9\xf6\x77\xfc\x81\xea\x38\xef\x0e\xef\x23\x77\x78\xff\x65\xf5\xb0\xf0\x5f\x9b\x8f\xf3\xf6\x6b\xbe\xce\x56\x5b\x3f\x58\x26\xf7\xe9\x66\x95\xcc\xd3\x76\xfc\xb0\x48\xfb\x88\x0b\x58\x6d\x3a\x0d\x58\x56\xae\x8e\x1c\xc7\x32\x48\x9d\xde\xb3\xa9\x94\xb6\xed\x4f\x37\x7b\x62\x26\x2f\xb1\xc8\x44\x81\x4f\xb6\xf5\xc8\x84\xf9\x4c\xf5\x38\xf9\xf7\x18\x71\xeb\xf8\x1d\xba\x2d\xdb\x0d\x40\x91\xcb\x90\xab\x15\xf8\x04\x72\x6f\x15\x6b\xd8\xce\x59\xc2\x94\xa8\x43\x3b\xba\x02\x73\x2f\x15\xa9\xd9\xb0\x9d\x45\x61\x1b\x35\x4b\xe8\xe7\x12\x4d\x89\xca\xfb\x0d\x39\x17\x0b\x0e\xd3\x39\x30\xe4\x07\x44\xdf\x31\x7c\xe4\x55\x05\x27\x76\x74\x56\xf6\xae\x08\xf8\x84\x79\x63\x33\x27\xf1\xe9\x47\xc9\xde\x50\x22\x25\xe2\x7d\xca\x84\x10\xdf\xa0\x15\xeb\x45\xfd\x3c\x6a\x1f\xd7\x77\x05\x6a\xd8\x4b\x75\x62\x86\x2a\x67\x17\x7b\x04\xb6\x0b\x44\xed\xdb\xc4\xc7\x92\xe7\xa5\x75\xfc\x1d\xa2\x80\x9a\x29\x8d\x05\x45\xe8\xb5\x3b\xcb\xce\xe7\x9d\xbf\xb3\xdd\xc6\xc8\x1a\x6a\xa9\xb9\xc5\x4b\xf2\x75\x3c\xb3\xb0\x1b\x8d\x14\x3a\xc4\x40\xb8\x18\x9c\x59\xc5\x8b\xdb\x40\x3f\xad\x02\xdf\xd8\xfb\x3e\xed\xe6\x43\x65\xbd\x86\xa4\xaa\x22\x93\x92\x59\x90\xe5\x65\x60\x7d\x02\xa9\xbd\x8d\x37\x91\x76\x23\xff\x19\x57\x6a\xd0\xd1\x06\x9a\x7d\x26\x33\x68\xef\x15\xad\x7c\x54\x10\xf0\x02\x8b\x3f\x6a\xd6\x57\x91\x9e\xa4\x02\x21\xad\xdb\x52\x4f\xd6\x28\x81\x05\x28\x8b\xc4\x79\x6e\xcd\x94\xe1\xac\x82\xa9\x51\x0d\xde\xd0\xf6\x0e\xd2\x74\xcf\x2a\x8d\xd4\x21\x95\x4c\x27\x45\x61\xed\xc3\xaa\x7b\x1b\x6e\x7a\xa4\x66\xa2\x0e\x97\x71\x81\x8a\x02\xac\x71\xf7\xfa\xb0\xf8\x19\xbf\xb2\x7c\xa8\xf6\xdb\x4e\xa8\x35\x3b\x44\x53\x6d\x6b\x17\xce\x68\xc3\x94\x99\xcb\x46\x18\x1b\x72\x3d\x94\xf7\x7f\xd5\xe9\x19\x85\x53\xf7\x08\x31\xdb\x68\x6c\xf9\x09\x23\x18\xd4\x6a\x0c\x26\x5b\x82\x2b\x59\xfc\x29\xa9\x1a\xfd\x62\xb1\xf2\x56\x8d\xf6\x5d\x2a\xd6\xa9\xeb\xae\x91\x44\xa3\xd5\x56\xcc\xb6\xe9\x1e\xd3\x87\xcd\xf6\xbe\x35\x0d\x44\x70\x3e\x58\xe0\x9e\x91\x57\x5a\x03\xd0\x1d\x26\xa4\x29\x7d\x38\x1d\x85\x7c\x14\xe4\xf2\xf3\x4d\x74\x69\xd3\x39\xbf\x5f\x43\x89\xac\x32\xe5\x85\x8e\x96\xc8\x94\xd9\x21\xf3\x9e\xa5\x30\x47\x7e\xc6\x82\xae\x5a\x85\x87\xa6\x62\x0a\xb8\x30\xa8\xa8\xbc\xb5\xf7\xad\x29\x5d\x0e\xf0\xdd\x36\x91\x53\xa8\x6b\x29\x0a\x42\x60\xa4\x7d\x56\x42\x6d\xb4\x07\xf1\x2e\x4d\xee\xb6\xef\x7e\xb9\x06\xd1\x88\x00\x86\x4d\x9b\x3d\xc5\xdc\x3d\xd2\x51\xfd\x20\x61\xc5\x9f\x38\xc2\xbc\x92\x8d\xbb\xf1\xb9\xf6\xe1\xd5\xa6\x97\x5e\x86\x5b\xd8\xd9\x6c\x27\xbe\x35\xf0\x6b\x83\xea\x62\xd3\x09\x85\xa6\x96\x27\xf4\x66\xf3\xb7\xb8\x42\x8d\xa7\x5d\x85\x1a\xde\x6d\xb7\xab\x6f\x35\x7c\xff\xdd\x77\xde\xfa\x9d\xfe\xc6\xc1\xdb\x6c\x7f\x90\xf6\x19\x8b\xeb\x1e\xab\x97\xe3\xa7\xf5\x6a\xde\x4a\x40\xf7\xc5\x4e\x21\x3b\xea\x37\x96\x40\x29\x6b\x74\xd9\x98\x99\xae\x74\x68\x05\xb7\x74\x73\x02\xba\x63\xf9\x91\x0a\x15\x2e\xd0\x8a\x4c\xc1\x7f\xa2\xdc\x02\x1e\x91\x43\xe2\x71\x2e\xb2\xcd\xfc\x61\xb9\x4c\xe7\x5b\x5b\xe1\x0d\xf5\x4c\xfd\x20\xd9\xe6\xb1\x44\x31\x54\x34\x77\x33\xb5\x92\x39\x6a\x4d\xa9\xb3\xdd\xde\xea\x60\xb5\x48\xb6\xae\x8c\x74\x74\xdd\xab\x8a\xab\x97\x5a\xc9\x9d\xda\x69\x8a\xd2\x96\xa6\x10\x66\xe2\x02\xd2\x26\xb3\x7d\xa3\xdc\x6d\xea\xdc\xd8\xd2\x47\x0d\x6c\x27\x1b\xa7\x82\x47\x9f\xf5\xb8\x09\x7d\x53\xaa\x21\x94\x6b\x19\x3d\x96\x47\xa6\xc1\xa8\x8b\xf7\x3f\xc7\xc0\x41\xda\x33\x5e\x61\xe7\x35\x42\x3e\x92\xc0\x0c\x76\xac\x88\x14\x68\x85\x4c\xfb\x1a\x79\xa0\xc1\x02\x0f\x8a\x15\xbd\x81\x03\xfd\x55\xfc\x88\xd5\x85\xd8\xee\x30\xf0\x38\xe2\x7d\xe2\x87\xd2\xd0\xb4\x7d\xa6\xf0\xae\x4a\x6d\x46\x6b\xb5\xf4\xa7\x75\xb2\x70\x25\xb8\xcd\x56\xe1\xcb\x94\x8d\xf6\x9a\x69\x6d\x4a\x25\x9b\x43\x99\x0e\x7b\xf5\x3e\x7f\x07\x8f\x6a\x71\xdf\xd1\x66\xb2\x28\x8d\xb4\x19\xf3\x5d\x1b\x33\x51\xf2\x8b\x9f\xcc\xa2\xa7\xb2\x6e\xf5\x23\x2a\xcd\x07\xc9\xcf\x71\x78\x7e\xe5\xaa\x13\x57\x68\xcc\x65\x3e\xbe\x78\xfd\xd6\xdc\x26\x58\x25\xab\x55\xc5\x04\x76\x79\xdd\x56\x8e\xdd\xc8\x25\xd4\x2e\xaf\x2c\x98\x61\x5f\xdf\x2e\x9a\xd3\x52\x16\xa8\x7d\xee\xb5\x13\x99\xd0\x46\x35\xd4\xcd\x61\x11\x2f\x3a\x9d\xde\x5f\xdf\x08\xb5\xc2\x33\x97\x8d\xde\x8c\x29\xfd\x6a\x3d\xba\xaf\x86\xa6\x8c\xff\x50\x38\xa3\xd6\x49\x51\x28\xd4\xd1\xbd\x64\xe4\x11\xc5\x75\x2b\xda\x3f\xb3\xd9\xa3\x57\x0f\x2f\xdc\xae\xdd\x71\x71\x8c\xce\xbe\x86\xf5\x57\xde\xe6\x2d\xf5\xe1\x93\xfc\xd7\x9e\x4d\x86\x6d\xed\x0b\xd9\xb4\xef\xef\xbe\x24\x70\x3c\x67\x57\x28\xac\x05\x9e\xaa\x76\x77\x88\xe0\xcc\xf5\xdf\x37\x0f\xcb\x3f\x03\x22\xfe\x5f\xf0\x22\x49\x6d\xf9\xd5\xa2\x8c\xa3\xf6\x45\xcc\x9f\x91\x7f\xf0\x27\xc3\x87\x47\x2c\x7a\xd7\x33\x06\x3f\xb1\x2c\x19\x80\xa8\xa1\xb7\xc3\xbb\x6c\xf9\xe1\x9f\x5f\x92\xfb\xc5\x0f\x7f\x69\xa7\x16\xc9\xfa\xe7\x6c\x19\xcf\xcd\x1f\x96\xdb\x24\x5b\xa6\xeb\x2f\x9b\x74\xfb\xe5\x97\xe4\xfe\x6e\x33\xbe\x34\x42\x2f\xde\xb0\x4d\xef\x57\x77\x94\x74\x1d\x91\x2e\x04\xfa\x3f\x6c\xee\xaf\xa5\x8a\x7c\x57\x97\xec\xff\xbf\xff\x21\x92\x31\x7e\xa2\x7a\x49\x0e\x1d\x7f\xe0\x0a\xde\x95\x9d\xc5\xaf\x1f\x1c\xaf\x0f\x06\x6f\xc3\x2e\xe8\x9e\x79\x17\x74\xf6\x77\x8f\xaa\xff\xa7\xb0\xb2\x3f\x37\x48\x70\xfd\xa6\x2d\x59\xed\xda\x68\xbd\x1a\x3c\xde\x8e\xb7\xd5\x36\xb7\xcb\x83\x0c\x7a\x2f\xe2\xa0\xcd\x48\x62\xd6\x4d\x5d\x4b\x65\x74\xf7\x4e\x1c\x3d\x1c\x76\x7f\x82\xae\xef\x9d\x56\xb8\x62\x40\xb4\xf3\xb7\xfe\xff\x92\x95\x62\x15\x3e\x4b\xac\xde\x7f\x59\xa7\xdb\x74\xb9\xcd\x1e\x96\x7d\x95\x1e\xbc\x31\x8f\x0a\x7e\x66\x55\x73\x15\x6b\x57\xb2\x5f\x19\xe5\xa5\x74\x7f\x97\x80\xb3\x6b\xf7\x3e\xfe\x69\x6c\x8f\xaf\xf3\xaf\xdd\xec\x3f\x01\x00\x00\xff\xff\x6c\x6e\x3a\x7d\x92\x1f\x00\x00")

func _03_auth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__03_auth_schemaGraphql,
		"03_auth_schema.graphql",
	)
}

func _03_auth_schemaGraphql() (*asset, error) {
	bytes, err := _03_auth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_auth_schema.graphql", size: 8082, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
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
	"01_base_schema.graphql":   _01_base_schemaGraphql,
	"02_unauth_schema.graphql": _02_unauth_schemaGraphql,
	"03_auth_schema.graphql":   _03_auth_schemaGraphql,
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
	"01_base_schema.graphql":   &bintree{_01_base_schemaGraphql, map[string]*bintree{}},
	"02_unauth_schema.graphql": &bintree{_02_unauth_schemaGraphql, map[string]*bintree{}},
	"03_auth_schema.graphql":   &bintree{_03_auth_schemaGraphql, map[string]*bintree{}},
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
