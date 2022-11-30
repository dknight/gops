package main

import (
	"fmt"
	"os"
)

const defaultStoreFileName = ".gops"

var storeFileName = defaultStoreFileName
var storeFileResolver = getStoreFile

// getStoreFile resolves the file to store todo items on file system.
func getStoreFile() (*os.File, error) {
	home, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%v%c%v", home, os.PathSeparator, storeFileName)
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return fp, nil
}
