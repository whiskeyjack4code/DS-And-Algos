package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	s := []int{2, 4, 6, 10, 12, 15, 45, 60, 88}
	e := 15

	res, _ := BinarySearch(s, e)

	expected := 5

	if res != expected {
		t.Errorf("expected %d but got %d", expected, res)
	}
}
