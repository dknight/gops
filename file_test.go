package gops

import (
	"fmt"
	"os"
	"testing"
)

func TestGetStoreFile(t *testing.T) {
	fp, err := GetStoreFile()
	if err != nil {
		t.Error(err)
	}

	cfgDir, err := os.UserConfigDir()
	if err != nil {
		t.Error(err)
	}

	exp := fmt.Sprintf("%v%c%v%c%v", cfgDir, os.PathSeparator,
		configDir, os.PathSeparator, storeFileName)
	if fp.Name() != exp {
		t.Error("Expected", exp, "got", fp.Name())
	}
}

func TestGetDefaultStoreFilePath(t *testing.T) {
	exp := fmt.Sprintf("%v%c%v", GetConfigPath(), os.PathSeparator,
		defaultFile)
	if exp != GetDefaultStoreFilePath() {
		t.Error("Expected", exp, "got", GetDefaultStoreFilePath())
	}
}

func TestDefaultFile(t *testing.T) {
	if DefaultFile() != defaultFile {
		t.Error("Expected", DefaultFile(), "got", defaultFile)
	}
}

func TestConfigDir(t *testing.T) {
	if ConfigDir() != configDir {
		t.Error("Expected", ConfigDir(), "got", configDir)
	}
}

func TestStoreFileName(t *testing.T) {
	if StoreFileName() != storeFileName {
		t.Error("Expected", StoreFileName(), "got", storeFileName)
	}
}

func TestSetStoreFilename(t *testing.T) {
	exp := "foo"
	SetStoreFileName(exp)
	if storeFileName != exp {
		t.Error("Expected", exp, "got", storeFileName)
	}
}
