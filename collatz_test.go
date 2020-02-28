package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	expected := []int{7, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}
	actual := collatz(7)
	if !reflect.DeepEqual(expected, *actual) {
		t.Errorf("Bad collatz series: %v", actual)
	}
	if maxList(actual) != 52 {
		t.Errorf("Bad maximum: %d", maxList(actual))
	}
	if listToString(actual) != "7 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1" {
		t.Errorf("Bad string list: %s", listToString(actual))
	}
}
