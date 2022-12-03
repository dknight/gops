package main

import (
	"fmt"
	"os"
)

const defaultFile = "default"
const configDir = "gops"

var storeFileName = defaultFile
var storeFileResolver = getStoreFile

// getStoreFile resolves the file to store todo items on file system.
func getStoreFile() (*os.File, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	cfgPath := fmt.Sprintf("%v%c%v", cfgDir, os.PathSeparator, configDir)

	_, err = os.Stat(cfgPath)
	if os.IsNotExist(err) {
		if err := os.Mkdir(cfgPath, 0755); err != nil {
			return nil, err
		}
	}

	fpath := fmt.Sprintf("%v%c%v", cfgPath, os.PathSeparator, storeFileName)
	fp, err := os.OpenFile(fpath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return fp, nil
}
