package goutil

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewSet()
	set.Add("nihao", "wohao")
	set.Traverse()

	if !set.Contains("nihao") {
		t.Error("no nihao find")
	}
}
