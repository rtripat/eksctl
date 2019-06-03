// Code generated by go-bindata.
// sources:
// assets/aws-node-1.10.yaml
// assets/aws-node.yaml
// assets/coredns.yaml
// DO NOT EDIT!

package defaultaddons

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

var _awsNode110Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdf\x6f\xdb\x36\x10\x7e\xd7\x5f\x71\x70\x1f\x0a\x0c\x95\x1c\x75\x59\x97\xe9\xcd\x75\xbc\xac\x58\xe2\x1a\xf1\x9a\x62\x28\x02\x83\xa6\x2e\x32\x67\x8a\xe4\xc8\xa3\x1d\xf7\xaf\x1f\x28\xf9\x87\x14\xcb\xee\x36\x0c\x18\x5f\x2c\xf3\xee\xbe\xe3\x7d\xfc\xc8\x63\x1c\xc7\x11\x33\xe2\x01\xad\x13\x5a\x65\x60\xe7\x8c\x27\xcc\xd3\x42\x5b\xf1\x95\x91\xd0\x2a\x59\x5e\xb9\x44\xe8\xfe\x2a\x8d\x5e\xc1\xd2\xcf\xd1\x2a\x24\x74\xb0\xaa\x43\x1c\xcc\xf1\x49\x5b\x84\x34\xb9\x4a\x2e\xc0\x2d\xb4\x97\x39\x78\x87\x67\xa1\xe6\x48\x2c\x8d\x96\x42\xe5\x19\x0c\xa5\x77\x84\xf6\x5e\x4b\x8c\x4a\x24\x96\x33\x62\x59\x04\xa0\x58\x89\x19\xb0\xb5\x8b\x95\xce\x31\xb2\x5e\xa2\xcb\xa2\x18\x98\x11\x37\x56\x7b\xe3\x82\x53\x0c\xdc\xe6\x15\x2e\x2b\xd9\x57\xad\xd8\xda\x25\x5c\x97\x11\x80\x45\xa7\xbd\xe5\xb8\x75\xeb\x7d\xd7\xab\x7e\x03\xaa\x33\x8c\xa3\x8b\x20\xd4\x30\x6f\xd8\x9b\xd8\xf0\xa5\xd7\x7b\x3c\x86\x31\x3a\x77\x35\x8e\xce\xd1\x9d\x42\x84\x2f\x3d\x29\x1c\xf5\xde\x40\x6f\xcd\x88\x2f\xc2\x47\x81\xd4\x7b\x7c\x99\x02\x9f\x09\x55\x45\x63\x57\xb2\x9c\x61\xa9\x95\x43\x3a\x83\xfc\x18\xbd\xdc\xc2\xd5\x8e\xd8\x29\xda\x95\xe0\x38\xe0\x5c\x7b\x45\xe7\xb8\x85\x43\x11\x59\xb5\xc7\xb1\xdb\x38\xc2\xf2\x08\xfb\xff\x95\xc7\x7b\xa1\x72\xa1\x8a\xb3\x2a\xd1\x12\xef\xf1\x29\x58\x76\x44\x9f\x59\x75\x04\x70\xac\xc1\x23\x4c\xe7\xe7\x7f\x20\xa7\x4a\x7c\x9d\xcc\xfe\x33\x3e\x6b\x88\xeb\x6a\x6f\xa7\x48\x2d\x7e\x99\x31\xee\x6f\x50\xf9\x53\x9b\xca\x83\x8a\xf6\xdc\xfd\x9b\xcd\x06\x90\x6c\x8e\xb2\x12\x1f\xc0\xf2\xca\xc5\xcc\x98\x26\x0f\x06\x79\xb0\x79\x93\x33\xc2\x29\x59\x46\x58\x6c\x6a\x6f\xda\x18\xcc\xe0\x5e\x4b\x29\x54\xf1\xa9\x72\x88\x00\x1c\x4a\xe4\xa4\x6d\xed\x53\x06\xc1\xde\x36\x52\x74\x25\x01\x20\x2c\x8d\x64\x84\xdb\xa0\x46\x21\x61\xc8\x56\x7c\x37\x42\x18\x4c\x29\x4d\xd5\x5e\x37\x9c\x1d\x5f\x60\xee\x25\xda\x84\x49\xb3\x60\xc9\x81\xe4\xa0\x3b\x6e\x05\x09\xce\x64\x6c\x74\x9e\xc1\xeb\xd7\x55\xd8\xae\xe8\xea\xbb\xb5\xed\xe3\x97\xb4\x86\xb1\xd0\x8e\xc6\x48\x6b\x6d\x97\x19\x90\xf5\xbb\x79\xd2\x12\x6d\x7b\x39\x31\x68\x13\xe6\xb4\xcd\x60\xf4\x2c\x5c\x75\xca\xc3\xe0\x5a\x11\x13\x0a\x6d\xc3\x55\x94\xac\xc0\x0c\xde\x5d\xbc\xbd\xbc\x48\xd3\xcb\xef\x2f\x7f\x78\x9b\xe4\x4b\x9b\x20\xb7\x89\x77\xf1\x1a\x1d\xc5\x6f\xdb\x77\x60\xbf\xfe\x17\x07\x86\xb8\x12\xd9\x2a\x4d\x2e\x93\x74\xcf\x45\x85\x38\xf1\x52\x4e\xb4\x14\x7c\x93\xc1\x40\xae\xd9\xc6\xed\xed\x46\x5b\x6a\x50\x17\x1f\x96\x35\xd1\x96\x32\x78\x97\xbe\xfb\xf1\x6a\x6f\xde\xa9\xac\x44\xb2\x82\x1f\x50\x8e\xb4\x57\x0f\x54\xab\xac\x11\x1b\x6f\xfd\x06\x9f\xa7\xb3\x87\xc9\x70\xf6\xeb\xd5\x74\x36\x1c\x7f\x98\xdd\x7e\xbc\xb9\x1d\x3d\x8c\x6e\x1b\xae\x00\x2b\x26\x3d\x66\x70\x3d\x7a\xff\xe9\xa6\x03\xe3\xee\xf7\xd9\xf8\xe3\xf5\x68\x36\x1e\xdc\x8d\x8e\xe3\x7e\xb6\xba\xcc\x5a\xd3\x00\x4f\x02\x65\xbe\xbd\x34\x3a\x2c\x13\x46\x8b\xac\xd2\x41\x12\x6a\x08\xdb\xde\x91\xf6\xf3\xe0\xb7\xe1\x2f\x55\xd2\xe9\x64\x30\xfc\x2f\x33\xef\x4e\x40\xb2\x3f\xb6\x7b\xef\x56\xbf\x38\x4c\xfe\xe9\xd1\x91\x6b\x83\x72\xe3\x33\x48\x2f\xca\xc3\x59\x40\xee\xad\xa0\xcd\x50\x2b\xc2\x67\x6a\x7a\x1b\x2b\x56\x42\x62\x81\x79\x4b\xc3\x00\x2b\x2d\x7d\x89\x77\x41\xfd\x2d\x69\x94\x61\xa6\x5e\x6d\x3f\x9c\x80\xbe\x36\xd4\xe7\x4a\xf4\xe7\x42\x1d\x49\x84\x2b\x11\xcf\x85\x8a\x73\x61\xcf\x41\x20\xf1\x0a\x42\x21\x25\x79\x27\x88\x42\xfa\x16\xc8\x8a\xd9\xbe\xd4\xc5\x51\xb8\xd4\xc5\x99\xd0\x10\x65\xbd\xea\xe7\x9a\x2f\xd1\x26\x4e\xf3\xe5\x11\x42\x6d\x6b\x98\x6a\x6e\x1a\x47\xf6\x74\xb5\x61\x69\x55\xaa\x26\xe7\x75\xea\x63\xe2\xe2\x33\x15\x9f\x01\xea\xa2\x2f\x3e\x51\xfd\x19\x98\x36\x81\xf1\xa9\xe2\xbf\x89\xf1\x92\xce\x97\x0f\x0b\x66\xc4\xa1\x8b\x9d\x78\x08\x78\x47\xba\xbc\xdf\x4a\xfe\x1a\x9f\x84\x12\xe1\x42\xed\xe8\x75\xa8\x04\xd7\xea\x49\x14\x2e\xe9\x7e\x1e\xee\x6e\x75\xc7\x75\xe8\x5b\xdb\xf6\x1f\x01\x14\xf5\x8b\xe1\xd4\xa3\xf2\x15\xa4\x49\x7a\x11\x9a\xae\x83\xde\xb6\x2f\xf7\xde\x00\x93\x12\x14\xae\xd1\x56\xed\x78\x67\x70\xbd\xfa\xd9\xb6\x7b\x96\x55\x3d\x27\xdd\xf5\xdf\x9a\x28\x23\xbd\x65\xb2\xb9\xe2\xba\xeb\x08\x55\x78\xc9\x6c\xc3\x50\x37\xe5\x8a\x89\xd1\xf8\xc3\xb0\x9e\xfb\x2b\x00\x00\xff\xff\x9c\xa8\x6b\xec\xbf\x0b\x00\x00")

func awsNode110YamlBytes() ([]byte, error) {
	return bindataRead(
		_awsNode110Yaml,
		"aws-node-1.10.yaml",
	)
}

func awsNode110Yaml() (*asset, error) {
	bytes, err := awsNode110YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aws-node-1.10.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _awsNodeYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdf\x8f\x1a\x37\x10\x7e\xdf\xbf\x62\x44\x1f\x22\x55\xdd\xe5\x48\xaf\xe9\x75\xdf\x08\x47\xd3\xa8\x77\x04\x1d\xcd\x45\x55\x74\x42\xc6\x3b\x07\x2e\x5e\xdb\xb5\xc7\x70\xe4\xaf\xaf\xec\xe5\xc7\x2e\x2c\xa4\xad\x2a\xd5\x2f\xb0\x9e\x99\x6f\x3c\x9f\x3f\x7b\x9c\xa6\x69\xc2\x8c\x78\x44\xeb\x84\x56\x39\xd8\x19\xe3\x19\xf3\xb4\xd0\x56\x7c\x61\x24\xb4\xca\x96\x37\x2e\x13\xba\xbb\xea\x25\xdf\xc0\xd2\xcf\xd0\x2a\x24\x74\xb0\xaa\x42\x1c\xcc\xf0\x59\x5b\x84\x5e\x76\x93\x5d\x81\x5b\x68\x2f\x0b\xf0\x0e\x2f\x42\xcd\x90\x58\x2f\x59\x0a\x55\xe4\x30\x90\xde\x11\xda\x07\x2d\x31\x29\x91\x58\xc1\x88\xe5\x09\x80\x62\x25\xe6\xc0\xd6\x2e\x55\xba\xc0\xc4\x7a\x89\x2e\x4f\x52\x60\x46\xbc\xb3\xda\x1b\x17\x9c\x52\xe0\xb6\x88\xb8\xac\x64\x5f\xb4\x62\x6b\x97\x71\x5d\x26\x00\x16\x9d\xf6\x96\xe3\xd6\xad\xf3\x6d\x27\xfe\x06\x54\x67\x18\x47\x97\x40\xa8\x61\x56\xb3\xd7\xb1\xe1\x73\xa7\xf3\x74\x0a\x63\x74\xe1\x2a\x1c\x5d\xa0\x3b\x87\x08\x9f\x3b\x52\x38\xea\x7c\x07\x9d\x35\x23\xbe\x08\x7f\xe6\x48\x9d\xa7\xe3\x14\xf8\x42\xa8\x22\x8d\x6d\xc9\x0a\x86\xa5\x56\x0e\xe9\x02\xf2\x53\x72\xbc\x85\xab\x1d\xb1\x13\xb4\x2b\xc1\xb1\xcf\xb9\xf6\x8a\x2e\x71\x0b\x87\x22\xf2\xb8\xc7\xa9\xdb\x38\xc2\xf2\x04\xfb\xff\x95\xc7\x5b\xa1\x0a\xa1\xe6\x17\x55\xa2\x25\x3e\xe0\x73\xb0\xec\x88\xbe\xb0\xea\x04\xe0\x54\x83\x27\x98\xce\xcf\xfe\x40\x4e\x51\x7c\xad\xcc\xfe\x33\x3e\x2b\x88\xdb\xb8\xb7\x13\xa4\x06\xbf\xcc\x18\xf7\x37\xa8\xfc\xa9\x49\xe5\x41\x45\x7b\xee\xfe\xcd\x66\x03\x48\x36\x43\x19\xc5\x07\xb0\xbc\x71\x29\x33\xa6\xce\x83\x41\x1e\x6c\xde\x14\x8c\x70\x42\x96\x11\xce\x37\x95\x37\x6d\x0c\xe6\xf0\xa0\xa5\x14\x6a\xfe\x31\x3a\x24\x00\x0e\x25\x72\xd2\xb6\xf2\x29\x83\x60\xef\x6a\x29\xda\x92\x00\x10\x96\x46\x32\xc2\x6d\x50\xad\x90\x30\x64\x23\xbe\x1d\x21\x0c\xa6\x94\xa6\xb8\xd7\x35\x67\xc7\x17\x58\x78\x89\x36\x63\xd2\x2c\x58\x76\x20\x39\xe8\x8e\x5b\x41\x82\x33\x99\x1a\x5d\xe4\xf0\xea\x55\x0c\xdb\x15\x1d\x86\xb1\x42\x5b\x41\x9b\x81\x64\xce\x8d\x22\xab\x15\x75\x31\x71\xba\x8b\xdf\x7a\xbb\x86\x48\x46\xc7\x9b\x10\xc6\x42\x3b\x1a\x21\xad\xb5\x5d\xe6\x40\xd6\xef\xe6\x49\x4b\xb4\xcd\xc5\xa7\xa0\x4d\x98\xd3\x36\x87\xe1\x8b\x70\xf1\x4e\x08\x83\x6b\x45\x4c\x28\xb4\x35\x57\x51\xb2\x39\xe6\xf0\xe6\xea\xf5\xf5\x55\xaf\x77\xfd\xfd\xf5\x0f\xaf\xb3\x62\x69\x33\xe4\x36\xf3\x2e\x5d\xa3\xa3\xf4\x75\xf3\xc6\xec\x56\x5f\x69\xe0\x93\x2b\x91\xaf\x7a\xd9\x75\xd6\xdb\x33\x17\x11\xc7\x5e\xca\xb1\x96\x82\x6f\x72\xe8\xcb\x35\xdb\xb8\xbd\xdd\x68\x4b\x35\xa2\xd3\xc3\xb2\xc6\xda\x52\x0e\x6f\x7a\x6f\x7e\xbc\xd9\x9b\x77\x9a\x2c\x91\xac\xe0\x07\x94\x13\xa5\x56\x03\xd5\x2a\xaf\xc5\xa6\x5b\xbf\xfe\xa7\xc9\xf4\x71\x3c\x98\xfe\x7a\x33\x99\x0e\x46\xef\xa7\x77\x1f\xde\xdd\x0d\x1f\x87\x77\x35\x57\x80\x15\x93\x1e\x73\xb8\x1d\xbe\xfd\xf8\xae\x05\xe3\xfe\xf7\xe9\xe8\xc3\xed\x70\x3a\xea\xdf\x0f\x4f\xe3\x7e\xb6\xba\xcc\x1b\xd3\x00\xcf\x02\x65\xb1\xbd\x62\x5a\x2c\x63\x46\x8b\x3c\xaa\x26\x0b\x35\x84\x6d\x6f\x49\xfb\xa9\xff\xdb\xe0\x97\x98\x74\x32\xee\x0f\xfe\xcb\xcc\xbb\xf3\x92\xed\x0f\xf9\xde\xbb\xd1\x5d\x0e\x93\x7f\x7a\x74\xe4\x9a\xa0\xdc\xf8\x1c\x7a\x57\xe5\xe1\xe4\x20\xf7\x51\xfa\x5a\x11\xbe\x50\xdd\xdb\x58\xb1\x12\x12\xe7\x58\x34\x34\x0c\xb0\xd2\xd2\x97\x78\x1f\xd4\xdf\x90\x46\x19\x66\xaa\xd5\x76\xc3\x09\xe8\x6a\x43\x5d\xae\x44\x77\x26\xd4\x89\x44\xb8\x12\xe9\x4c\xa8\xb4\x10\xf6\x12\x04\x12\x8f\x10\x0a\x29\x2b\x5a\x41\x14\xd2\xd7\x40\x56\xcc\x76\xa5\x9e\x9f\x84\x4b\x3d\xbf\x10\x1a\xa2\xac\x57\xdd\x42\xf3\x25\xda\xcc\x69\xbe\x3c\x41\xa8\x6c\x35\x53\xc5\x4d\xed\xc8\x9e\xaf\x36\x2c\x2d\xa6\xaa\x73\x5e\xa5\x3e\x25\x2e\xbd\x50\xf1\x05\xa0\x36\xfa\xd2\x33\xd5\x5f\x80\x69\x12\x98\x9e\x2b\xfe\xab\x18\xc7\x74\x1e\x3f\x43\x98\x11\x87\x9e\x77\xe6\xd9\xe0\x1d\xe9\xf2\x61\x2b\xf9\x5b\x7c\x16\x4a\x84\x0b\xb5\xa5\x33\xa2\x12\x5c\xab\x67\x31\x77\x59\xfb\x63\x72\xd7\x03\x1c\xd7\xa1\xcb\x6d\x1f\x0b\x09\xc0\xbc\x7a\x5f\x9c\x7b\x82\xee\x1a\x77\x55\xe5\x8e\x8e\x55\x2f\x36\x9f\x5e\xad\x4d\x1c\x1d\x1d\x47\xda\xc6\x0b\x7c\x3b\x17\x8f\x72\x05\x62\xa4\xb7\x4c\xd6\xd7\x5c\x75\x29\xa1\xe6\x5e\x32\x5b\x33\x54\x4d\x3c\x72\x31\x1c\xbd\x1f\x54\x73\x7f\x05\x00\x00\xff\xff\xd1\x54\xb9\x74\xef\x0b\x00\x00")

func awsNodeYamlBytes() ([]byte, error) {
	return bindataRead(
		_awsNodeYaml,
		"aws-node.yaml",
	)
}

func awsNodeYaml() (*asset, error) {
	bytes, err := awsNodeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aws-node.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x5b\x6f\xdb\x3a\x12\x7e\xf7\xaf\x20\xf4\xbc\x52\xec\x26\xe9\x66\xf9\xd6\xda\xde\xae\x81\xd6\x35\xe2\xa4\x2f\x8b\x45\x30\xa6\xc6\x16\xd7\x14\xc9\xe5\xc5\xb5\xbb\xe7\xfc\xf7\x03\xea\x66\x4a\x71\x8a\x14\x2d\x70\x8e\x9e\x24\x72\xf8\xcd\x70\x2e\xdf\x8c\x40\xf3\x2f\x68\x2c\x57\x92\x92\xc3\x64\xb4\xe7\x32\xa7\x64\x8d\xe6\xc0\x19\xbe\x63\x4c\x79\xe9\x46\x25\x3a\xc8\xc1\x01\x1d\x11\x22\xa1\x44\x4a\x98\x32\x98\x4b\xdb\x7c\x5b\x0d\x0c\x29\xd9\xfb\x0d\xa6\xf6\x64\x1d\x96\x23\x42\x04\x6c\x50\xd8\x70\x84\x10\xdc\xdb\x0c\x4a\xf8\xa6\x24\x7c\xb5\x19\x53\xe5\x15\x53\xa5\x56\x12\xa5\x8b\xb1\x08\xd9\xdf\xd9\x14\xb4\x6e\xb0\xc2\x6a\x9a\xa6\xa3\xd8\x46\xb3\x01\x96\x81\x77\x85\x32\xfc\x1b\x38\xae\x64\xb6\xbf\xb3\x19\x57\x57\x87\xc9\x06\x1d\xb4\x57\x98\x0a\x6f\x1d\x9a\x7b\x25\xb0\x67\x7f\x6c\x56\x50\x62\x24\x3a\xac\xce\x6f\x94\x72\xd6\x19\xd0\x9a\xcb\x5d\xad\x28\xcd\x71\x0b\x5e\x38\xfb\xb3\xb7\x68\xfd\x56\x7b\x87\xb6\xc2\xc6\x0b\xb4\x74\x94\x12\xd0\xfc\x83\x51\x5e\x57\x86\xa5\x24\x49\x46\x84\x18\xb4\xca\x1b\x86\xcd\x1a\xca\x5c\x2b\x2e\x2b\x5b\x52\x62\xeb\x08\xd5\x1f\x5a\xe5\xf5\x4b\x17\x8c\xf0\x79\x40\xb3\x69\xce\x0a\x6e\x5d\xf5\xf2\x15\x1c\x2b\x5e\xa7\x4f\xaa\x7c\x08\xb3\x43\xf7\x2b\xe2\xf1\x9e\xcb\x9c\xcb\x5d\x2f\x2c\x20\xa5\x72\xd5\xf1\x26\x36\x97\x70\x7b\xe1\x02\xef\x94\xd7\x39\x38\xa4\x24\x71\xc6\x63\xf2\x97\x8b\xae\x12\x78\x8f\xdb\xea\x7a\x8d\xbf\xbf\xe3\xaf\x11\x21\xcf\x33\xf7\x05\x64\xeb\x37\xff\x45\xe6\xaa\xd4\xb9\x58\xb1\xaf\xae\xd3\x61\x38\x3b\x0a\x98\x2a\xb9\xe5\xbb\x4f\xa0\xff\xd4\xea\x6f\xf5\x4e\x95\xc1\x2d\x17\x48\xc9\x6f\x95\x64\x46\x6f\xaf\xc9\xff\xab\xd7\x4a\x83\x31\xca\xd8\xee\xb3\x40\x10\xae\xe8\x3e\xcf\x89\x40\x58\xed\xdb\x4c\x28\x06\x22\x02\x20\x55\x0d\x11\x2e\x2d\x32\x6f\x30\x5a\xf7\xda\x3a\x83\x50\x46\x4b\x5b\x10\xc2\x15\x46\xf9\x5d\x41\xb8\x4c\x21\xcf\x4d\x06\x46\x03\xe1\xfa\x6d\xf5\xd2\xc9\xfe\xde\xbd\x69\xa3\x4a\x74\x05\x7a\x4b\xe8\x3f\x26\xb7\xd7\xf1\xc6\xf1\x44\x32\x72\x85\x8e\x5d\x85\x12\x14\x87\x8c\x29\xb9\xed\x04\x18\xb0\x02\xc9\xf5\x78\x54\x03\x0e\x03\x86\x47\x87\x32\xbc\xda\x41\xc1\xcd\x50\x0b\x75\x2a\xf1\x57\xf0\xf7\xa5\x94\x1f\x16\x58\x0d\x9c\x84\x48\xcd\x96\xeb\xe4\xf5\x91\xb7\x1a\x19\xad\xf8\x47\x0b\xce\xc0\x52\xf2\x66\x44\x48\xa8\x55\x87\xbb\x53\x6d\x80\x3b\x69\xa4\xe4\x5e\x09\xc1\xe5\xee\xb1\xaa\xfa\x9a\x25\xe2\x15\xda\xf8\xac\x84\xe3\xa3\x84\x03\x70\x01\x9b\x90\x32\x93\x00\x87\x02\x99\x53\xa6\x96\x29\x03\x0d\x7e\x8c\x2e\xf8\xd2\x15\x5f\x9d\xbc\x0e\x4b\x2d\x3a\x1b\x62\x87\x87\x47\xf4\x54\xbd\xac\xec\x07\x6a\xa5\xf5\x5a\xf5\xde\x2b\xfe\xe5\x20\xc2\x75\x96\x71\x65\xb8\x3b\x4d\x05\x58\xbb\x8c\x28\x25\x0d\x34\x9f\x32\xc3\x1d\x67\x20\x1a\x69\xd8\x6e\xb9\xe4\xee\x74\x36\x38\x48\xbd\x7b\xb6\x1a\x62\xf6\x3f\xcf\x0d\xe6\x33\x6f\xb8\xdc\xad\x59\x81\xb9\x0f\x01\x59\xec\xa4\xea\x96\xe7\x47\x64\x3e\x30\x5d\x7c\xb2\xc6\x5c\x37\x61\x79\x40\x53\xda\xfe\x76\x5a\x47\x69\x7e\xd4\x06\xad\x3d\x37\x86\x58\x62\x8f\x27\x4a\x92\x90\xf5\x83\xe6\xa0\x6c\x32\x10\x26\x44\x69\x34\x10\x52\x80\x2c\xe4\xb3\xcd\x03\x08\x8f\xcf\x34\xd4\xbd\x53\xfa\xe3\xeb\x35\x83\x61\xc5\x2f\xd3\x0d\x65\xfe\xf6\xa6\x59\x77\x4a\x04\x8c\xd8\x11\xad\x19\xd3\x26\x7c\xef\xf2\x5c\x49\xfb\x59\x8a\xd3\xd9\x82\xb3\xe6\x64\x7e\xe4\xd6\x75\x8e\x61\x4a\x3a\xe0\x12\x4d\x04\x37\x24\x87\xfa\xe1\x25\xec\x90\x92\xb7\xe3\x37\x37\xe3\xc9\xe4\xe6\xfa\xe6\xf6\x4d\x96\xef\x4d\x86\xcc\x64\xf7\xf3\x0f\x8b\xcf\xcb\x41\xca\xe2\xde\x5e\x35\x20\xf4\x30\xc9\x26\xd9\x75\x1f\x6b\xe5\x85\x58\x29\xc1\xd9\x89\x92\xc5\x76\xa9\xdc\xca\xa0\xc5\xaa\x6d\xd5\x4f\x6f\x14\x69\x1f\xc1\x4b\xee\x06\x6e\x2a\xb1\x54\xe6\x44\xc9\xe4\xef\xe3\x4f\x7c\x90\x97\x68\x87\xd2\x4c\x7b\x4a\x26\xe3\x71\x79\x11\xa3\x07\x01\x66\x67\x29\xf9\x37\x49\xd2\x40\xc6\xc9\xdf\x48\x52\x11\x74\x73\xab\xab\xb6\x1f\x25\xe4\x3f\xdd\x91\x83\x12\xbe\xc4\x4f\xa1\x04\x23\xbd\x67\xa7\x86\x7e\x9a\xd6\x42\x91\xfe\x32\xc8\xaf\xc0\x15\x94\xc4\x1a\x7a\x77\x81\x3c\xc4\x94\x92\x30\xe5\x9c\x1b\x87\x32\x7d\x3d\x5d\x40\x57\xca\x38\x4a\xa2\x1e\xd3\xb2\x7e\x1f\x57\x1b\xe5\x14\x53\x82\x92\xc7\xd9\xea\x47\x71\x52\xc7\xf4\x45\xac\x87\xe9\x77\xb0\x7a\x9d\xaf\x45\x2b\xd1\x19\xce\x2e\x5b\x16\xa3\x55\xad\x39\x70\x98\x92\x0e\x8f\x2e\x0e\x2d\x08\xa1\xbe\xae\x0c\x3f\x70\x81\x3b\x9c\x5b\x06\xa2\xaa\x14\x1a\x7a\xb5\x8d\xdd\xcd\x40\xc3\x86\x0b\xee\xf8\xb0\xe2\x20\xcf\x87\x04\xb4\x9c\x3f\x3c\xbd\x5f\x2c\x67\x4f\xeb\xf9\xfd\x97\xc5\x74\xde\xdb\xce\x8d\xd2\xc3\x03\x20\xc4\x85\xc0\xdd\x2b\xe5\xfe\xc9\x05\x36\x43\x5c\x3f\x8c\x82\x1f\x50\xa2\xb5\x2b\xa3\x36\x18\xe3\x15\xce\xe9\x0f\xe8\xfa\x2a\x74\x9d\x28\x83\x01\x87\x34\xe9\x40\xc9\xdd\xf8\x6e\xdc\x5b\xb6\xac\xc0\xe0\xe4\x7f\x3d\x3c\xac\xa2\x8d\x40\xe4\x1c\xc4\x0c\x05\x9c\xd6\xc8\x94\xcc\x6d\x28\xf0\x48\xc2\xf1\x12\x95\x77\xdd\xe6\x6d\xb4\x67\x3d\x63\x68\xed\x43\x61\xd0\x16\x4a\xe4\x75\x8b\x6d\x9f\x2d\x70\xe1\x0d\x46\xbb\xed\xd9\x5c\xda\xb6\xec\x67\xf5\xe8\xdd\x6c\xd4\x55\xf1\x03\x55\xc3\xda\xe9\x74\xd0\x52\x2e\xf2\x57\x75\x61\x87\xcf\x1b\x4c\xc5\x9e\x6d\x29\x0f\xe8\xb7\xf6\x74\xb7\xf9\xe2\x9c\xdc\x0c\xde\x17\x66\xac\xc1\xef\xc1\xc5\x21\xeb\xd9\x6f\xcf\x79\x4e\x0c\xcd\xa4\x0e\x6a\x12\xca\x26\xb9\xb0\x6d\x99\x01\xfd\xe2\xef\xcf\x2b\x66\xb6\x66\x1c\x4e\x9b\x01\x22\x42\xfa\xe9\xe9\xae\xd3\xda\x0e\x2a\xfd\x09\xec\x92\x75\x8d\x35\x8b\x15\x25\xb3\xe5\xfa\x69\xfa\xf1\x71\xfd\x30\xbf\x7f\x5a\x84\xc4\xed\xd8\x2e\x1d\x70\x99\x8e\x49\x6a\x48\x69\xe9\x05\xc2\x7a\xe1\x40\x60\x9a\x3f\x02\x00\x00\xff\xff\x17\xfe\xaa\x4f\x0c\x11\x00\x00")

func corednsYamlBytes() ([]byte, error) {
	return bindataRead(
		_corednsYaml,
		"coredns.yaml",
	)
}

func corednsYaml() (*asset, error) {
	bytes, err := corednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"aws-node-1.10.yaml": awsNode110Yaml,
	"aws-node.yaml": awsNodeYaml,
	"coredns.yaml": corednsYaml,
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
	"aws-node-1.10.yaml": &bintree{awsNode110Yaml, map[string]*bintree{}},
	"aws-node.yaml": &bintree{awsNodeYaml, map[string]*bintree{}},
	"coredns.yaml": &bintree{corednsYaml, map[string]*bintree{}},
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

