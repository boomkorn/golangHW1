package golanghw1

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type FileStore struct {
	FileName string
}

// File management
func (fs FileStore) ReadFile() []byte {
	f, err := os.Open(fs.FileName)
	if err != nil {
		os.Create(fs.FileName)
	} else {
		defer f.Close()
		byteValue, _ := ioutil.ReadAll(f)
		return byteValue
	}
	return nil
}
func (fs FileStore) WriteFile(ul UserList) {
	jsonByte, _ := json.Marshal(ul.Users)
	ioutil.WriteFile(fs.FileName, jsonByte, os.ModePerm)
}
func (fs FileStore) ClearFile() {
	ioutil.WriteFile(fs.FileName, nil, os.ModePerm)
}
func (fs FileStore) RemoveFile() {
	os.Remove(fs.FileName)
}
