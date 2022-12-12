package gops

import (
	"testing"
	"time"
)

func TestResolveDoneColor(t *testing.T) {
	complete := NewItem(time.Unix(0, 0).UTC(), true, "done")
	incomplete := NewItem(time.Unix(0, 0).UTC(), false, "not done")

	doneColor := ResolveDoneColor(complete.Status)
	if doneColor != Color.Green {
		t.Error("Expected", Color.Green, "got", doneColor)
	}

	notDoneColor := ResolveDoneColor(incomplete.Status)
	if doneColor != Color.Green {
		t.Error("Expected", Color.Nul, "got", notDoneColor)
	}
}
