package gops

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var (
	defaultFile = "default"
	configDir   = "gops"
)

var storeFileName = defaultFile

// GetStoreFile resolves the file to store todo items on file system.
func GetStoreFile() (*os.File, error) {
	cfgPath := GetSystemConfigPath()
	_, err := os.Stat(cfgPath)
	if err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(cfgPath, 0755); err != nil {
			return nil, err
		}
	}

	fpath := MakeStoreFilePath(StoreFileName())
	fp, err := os.OpenFile(fpath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return fp, nil
}

// CreateStoreFile create store file if not exists.
func CreateStoreFile() (*os.File, error) {
	var fp *os.File
	fpath := MakeStoreFilePath(StoreFileName())
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

// GetSystemConfigPath gets config directory, on Linux usually $HOME/.config
func GetSystemConfigPath() string {
	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v%c%v", userCfgDir, os.PathSeparator, configDir)
}

// MakeStoreFilePath makes path to store file.
func MakeStoreFilePath(name string) string {
	return fmt.Sprintf("%v%c%v", GetSystemConfigPath(), os.PathSeparator, name)
}

// GetDefaultStoreFilePath gets default file to store items.
func GetDefaultStoreFilePath() string {
	return MakeStoreFilePath(defaultFile)
}

// Truncate truncates writer, currently implemeted only for file and buffer.
func Truncate(wr io.Writer) error {
	var err error
	switch wr.(type) {
	case *os.File:
		err = wr.(*os.File).Truncate(0)
		if err != nil {
			return err
		}
	case *bytes.Buffer:
		wr.(*bytes.Buffer).Truncate(0)
	}
	return err
}

// For export -------------------------------------------------------------

// DefaultFile gets default file name in config directory.
func DefaultFile() string {
	return defaultFile
}

// ConfigDir gets the gops name in config directory.
func ConfigDir() string {
	return configDir
}

// StoreFileName gets the file name where items are written.
func StoreFileName() string {
	return storeFileName
}

// SetStoreFileName set store file name where items are written.
func SetStoreFileName(name string) {
	storeFileName = name
}
