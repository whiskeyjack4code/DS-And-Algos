package main

import "fmt"

func LinearSearch(s []int, e int) (int, bool) {
	for i, v := range s {
		if v == e {
			return i, true
		}
	}
	return 0, false
}

func main() {

	s := []int{10, 50, 30, 25, 45, 100, 24}
	e := 25

	i, ok := LinearSearch(s, e)

	if ok {
		fmt.Println(e, "found at index:", i)
	} else {
		fmt.Println(e, "not in slice")
	}
}
