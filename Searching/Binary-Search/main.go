package main

import "fmt"

func BinarySearch(s []int, e int) (int, bool) {
	upper := len(s) - 1
	lower := s[0]

	for lower <= upper {
		mid := (upper + lower) / 2
		midValue := s[mid]

		if e == midValue {
			return mid, true
		}

		if e < midValue {
			upper = mid - 1
		}

		if e > midValue {
			lower = mid + 1
		}
	}

	return 0, false
}

func main() {

	s := []int{2, 4, 6, 10, 12, 15, 45, 60, 88}
	e := 15

	v, ok := BinarySearch(s, e)
	if ok {
		fmt.Println(e, "found at index:", v)
	} else {
		fmt.Println(e, "not found in slice.")
	}

}
