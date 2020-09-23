package util

import (
	"fmt"
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewSet()
	set.Add("nihao", "wohao")
	set.Traverse(func(x interface{}) {
		fmt.Println(x)
	})

	if !set.Contains("nihao") {
		t.Error("no nihao find")
	}
}
