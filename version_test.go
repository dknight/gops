package gops

import (
	"reflect"
	"testing"
)

// shit
func TestVersion(t *testing.T) {
	v := Version()
	typ := reflect.TypeOf(v)
	if typ.String() != "string" {
		t.Error("Version should return string", "got", typ)
	}
}
