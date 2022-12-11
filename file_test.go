package gops

import (
	"fmt"
	"os"
	"testing"
)

func Test_getStoreFile(t *testing.T) {
	fp, err := getStoreFile()
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

func Test_getDefaultStoreFilePath(t *testing.T) {
	exp := fmt.Sprintf("%v%c%v", getConfigPath(),
		os.PathSeparator, defaultFile)
	if exp != getDefaultStoreFilePath() {
		t.Error("Expected", exp, "got", getDefaultStoreFilePath())
	}
}
