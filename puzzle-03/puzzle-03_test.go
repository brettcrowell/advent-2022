package main

import (
	"reflect"
	"testing"
)

func TestGetCommonItems(t *testing.T) {
	ans := GetCommonItems([]string{"a", "b", "c"}, []string{"b", "c", "d"})
	if !reflect.DeepEqual(ans, []string{"b", "c"}) {
		t.Errorf("getCommonItems")
	}
}
