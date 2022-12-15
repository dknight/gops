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

func TestTruncate(t *testing.T) {
	bf := getTestBuffer()
	_, err := bf.Write([]byte{123, 32, 110})
	if err != nil {
		t.Error(err)
	}
	err = Truncate(bf)
	if err != nil {
		t.Error(err)
	}
	if bf.Len() != 0 {
		t.Error("Expected", 0, "got", bf.Len())
	}
}

func TestTruncate_File(t *testing.T) {
	fp, err := os.CreateTemp("", "tmp")
	if err != nil {
		t.Error(err)
	}
	_, err = fp.Write([]byte{123, 32, 110})
	if err != nil {
		t.Error(err)
	}
	err = Truncate(fp)
	if err != nil {
		t.Error(err)
	}

	info, err := fp.Stat()
	if err != nil {
		t.Error(err)
	}
	size := info.Size()
	if size != 0 {
		t.Error("Expected", 0, "got", size)
	}
}
