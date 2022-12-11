package gops

import (
	"fmt"
	"os"
)

const (
	defaultFile = "default"
	configDir   = "gops"
)

var storeFileName = defaultFile

// getStoreFile resolves the file to store todo items on file system.
func getStoreFile() (*os.File, error) {
	cfgPath := getConfigPath()
	_, err := os.Stat(cfgPath)
	if err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(cfgPath, 0755); err != nil {
			return nil, err
		}
	}

	fpath := makeStoreFilePath(storeFileName)
	fp, err := os.OpenFile(fpath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return fp, nil
}

// createStoreFile create store file if not exists.
func createStoreFile() (*os.File, error) {
	var fp *os.File
	fpath := makeStoreFilePath(storeFileName)
	_, err := os.Stat(fpath)
	if err != nil && os.IsNotExist(err) {
		fp, err := os.Create(fpath)
		if err != nil {
			return nil, err
		}
		defer fp.Close()
	}
	return fp, nil
}

// getConfigPath gets config directory, on Linux usually $HOME/.congig
func getConfigPath() string {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v%c%v", cfgDir, os.PathSeparator, configDir)
}

// makeStoreFilePath makes path to store file.
func makeStoreFilePath(name string) string {
	cfgPath := getConfigPath()
	return fmt.Sprintf("%v%c%v", cfgPath, os.PathSeparator, name)
}

// getDefaultStoreFilePath gets default file to store items.
func getDefaultStoreFilePath() string {
	return makeStoreFilePath(defaultFile)
}
