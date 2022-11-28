package main

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
	exp := fmt.Sprintf("%v%c%v", cfgDir, os.PathSeparator, defaultStoreFileName)
	if fp.Name() != exp {
		t.Error("Expected", exp, "got", fp.Name())
	}
}
