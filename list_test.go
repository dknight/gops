package main

import (
	"os"
	"testing"
	"time"
)

func TestDisplayList(t *testing.T) {
	pairs := []*Item{
		NewItem(time.Unix(0, 0).UTC(), false, "first"),
		NewItem(time.Unix(0, 0).UTC(), false, "second"),
	}
	fp1 := getTempFile("todo1-")
	fp2 := getTempFile("todo2-")
	fp3 := getTempFile("todo3-")
	defer os.Remove(fp1.Name())
	defer os.Remove(fp2.Name())
	defer os.Remove(fp3.Name())

	for _, pair := range pairs {
		pair.Complete()
		pair.Save(fp1)
	}

	for i, pair := range pairs {
		pair.Complete()
		if i == 0 {
			pair.Complete()
		}
		pair.Save(fp2)
	}

	pairs[0].Complete()
	pairs[0].Save(fp3)

	lists, err := AllLists(getTempDirPath())
	if err != nil {
		t.Error(err)
	}
	if len(lists) != 3 {
		t.Error("Excepted", 3, "got", len(lists))
	}

	DisplayLists(getTempDirPath())
}
