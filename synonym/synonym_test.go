package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	// test removeDuplicates

	ans := removeDuplicates([]string{"a", "b", "a", "c", "b"})
	if !reflect.DeepEqual(ans, []string{"a", "b", "c"}) {
		t.Errorf("RemoveDuplicates failed")
	}
}
