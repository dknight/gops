package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

var tmpdefault = "gopsdefault"

func getTempFile(suffix string) *os.File {
	fp, err := os.CreateTemp(getTempDirPath(), suffix)
	if err != nil {
		log.Fatalln(err)
	}
	return fp
}

func getTempDirPath() string {
	path := fmt.Sprintf("%v%c%v", os.TempDir(), os.PathSeparator, "gops")
	_ = os.Mkdir(path, 0755)
	return path
}

func TestSave(t *testing.T) {
	fp := getTempFile(tmpdefault)
	defer os.Remove(fp.Name())

	item := NewItem(time.Unix(0, 0).UTC(), false, "hello")
	err := item.Save(fp)
	if err != nil {
		t.Error("Cannot save item", item, err)
	}
}

func TestAllItems(t *testing.T) {
	fp := getTempFile(tmpdefault)
	defer os.Remove(fp.Name())

	tests := []string{
		"hello",
		"busybreezy",
		"xyzzy",
	}
	for _, test := range tests {
		item := NewItem(time.Unix(0, 0).UTC(), false, test)
		item.Save(fp)
	}
	list, err := AllItems(fp)
	if err != nil {
		t.Error(err)
	}
	if len(list) != len(tests) {
		t.Error("Length expected", len(tests), "got", len(list))
	}
}

func TestFilterItemsByStatus(t *testing.T) {
	tests := map[string]bool{
		"hello":       false,
		"busy breezy": true,
		"xyzzy":       true,
	}
	var items []Item
	for k, v := range tests {
		items = append(items, *NewItem(time.Unix(0, 0).UTC(), v, k))
	}
	incomeletedItems := FilterItemsByStatus(items, itemStatusTodo)
	completedItems := FilterItemsByStatus(items, itemStatusDone)
	if len(incomeletedItems) != 1 {
		t.Error("Length expected", 1, "got", len(incomeletedItems))
	}
	if len(completedItems) != 2 {
		t.Error("Length expected", 2, "got", len(incomeletedItems))
	}
}

func TestCompleteItem(t *testing.T) {
	fp := getTempFile(tmpdefault)
	defer os.Remove(fp.Name())

	item := NewItem(time.Unix(0, 0).UTC(), false, "Complete me!")
	err := item.Save(fp)
	if err != nil {
		t.Error(err)
	}
	if item.Status {
		t.Error("Expected", false, "got", item.Status)
	}
	items, err := AllItems(fp)
	if err != nil {
		t.Error(err)
	}
	CompleteItem(1, items, fp)
	if item.Status {
		t.Error("Expected", true, "got", item.Status)
	}
}

func TestBeautifulString(t *testing.T) {
	item := NewItem(time.Unix(0, 0).UTC(), false, "Beautify me!")
	expFn := func(mark []rune) string {
		return fmt.Sprintf("  %v#%v %d %v%s%v %v",
			Color.Blue, Color.Nul,
			1,
			Color.Green, string(mark), Color.Nul,
			item.Task)
	}
	expected := expFn([]rune{itemMarkTodo})
	s := item.BeautifulString(1)
	if s != expected {
		t.Error("Expected", expected, "got", s)
	}
	item.Complete()
	expected = expFn([]rune{itemMarkDone})
	s = item.BeautifulString(1)
	if s != expected {
		t.Error("Expected", expected, "got", s)
	}
}
