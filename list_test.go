package main

import (
	"bytes"
	"io"
	"testing"
	"time"
)

func TestDisplayList(t *testing.T) {
	pairs := []*Item{
		NewItem(time.Unix(0, 0).UTC(), false, "first"),
		NewItem(time.Unix(0, 0).UTC(), false, "second"),
	}

	bufs := []*bytes.Buffer{
		getTestBuffer(),
		getTestBuffer(),
		getTestBuffer(),
	}

	for _, pair := range pairs {
		pair.Complete()
		pair.Save(bufs[0])
	}

	for i, pair := range pairs {
		pair.Complete()
		if i == 0 {
			pair.Complete()
		}
		pair.Save(bufs[1])
	}

	pairs[0].Complete()
	pairs[0].Save(bufs[2])

	rds := make([]*bytes.Reader, len(bufs))
	for i := 0; i < len(bufs); i++ {
		rds[i] = bytes.NewReader(bufs[i].Bytes())
	}

	DisplayLists([]io.Reader{rds[0], rds[1], rds[2]})
}
