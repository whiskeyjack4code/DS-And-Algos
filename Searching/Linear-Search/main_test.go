package main

import "testing"

func TestLinearSearch(t *testing.T) {
  s := []int{10,4,20,100,45,87,21,0,34}
  e := 87

  res, _ := LinearSearch(s, e)
  expected := 5

  if res != expected {
    t.Errorf("expected %d and got %d", expected, res)
  }
}
