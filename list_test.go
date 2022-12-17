package gops

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestDisplayList(t *testing.T) {
	items := []*Item{
		NewItem(time.Unix(0, 0).UTC(), false, "first"),
		NewItem(time.Unix(0, 0).UTC(), false, "second"),
	}

	bufs := []*bytes.Buffer{
		getTestBuffer(),
		getTestBuffer(),
		getTestBuffer(),
	}

	for _, item := range items {
		item.Complete()
		item.Save(bufs[0])
	}

	for i, item := range items {
		item.Complete()
		if i == 0 {
			item.Complete()
		}
		item.Save(bufs[1])
	}

	items[0].Complete()
	items[0].Save(bufs[2])

	rds := make([]*bytes.Reader, len(bufs))
	for i := 0; i < len(bufs); i++ {
		rds[i] = bytes.NewReader(bufs[i].Bytes())
	}

	err := DisplayLists([]io.Reader{rds[0], rds[1], rds[2]})
	if err != nil {
		t.Error(err)
	}
}

// FIXME maybe use tmp dir, it is better.
func TestGetListsByPath(t *testing.T) {
	rnd := strconv.Itoa(rand.Intn(100000000))
	testFilePath := fmt.Sprintf("%v%c%v__%v", GetSystemConfigPath(),
		os.PathSeparator, "goptest", rnd)
	fp, err := os.Create(testFilePath)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(fp.Name())

	items := []*Item{
		NewItem(time.Unix(0, 0).UTC(), false, "first"),
		NewItem(time.Unix(0, 0).UTC(), false, "second"),
	}

	for _, item := range items {
		item.Save(fp)
	}

	retrieved, err := GetListsByPath(GetSystemConfigPath())
	if err != nil {
		t.Error(err)
	}

	if len(retrieved) < 1 {
		t.Error("Expected at least", 1, "got", len(retrieved))
	}
}
