package main

import "fmt"

// O(n)
func FindDuplicate2(s []int) bool {
  
  buffer := make([]int, len(s), len(s))

  for i := 0; i < len(s); i++ {
    if buffer[s[i]] == 1 {
      return true
    } else {
      buffer[s[i]] = 1
    }
  }
  return false
}

// O(n2)
func FindDuplicate(s []int) bool {
  sLen := len(s)
  for i := 0; i < sLen; i++ {
    for j := 0; j < sLen; j++ {
      if(i != j && s[i] == s[j]){
        return true
      }
    }
  }
  return true
}

//  O(n2)
func BubbleSort(list []int) []int {
	listLen := len(list) - 1
	var sorted bool = false

	for !sorted {
		sorted = true // Tentatively set to true until a swap is encountered.

		for i := 0; i < listLen; i++ {
			var temp int

			if list[i] > list[i+1] {
				sorted = false    // Oh boy, we had a swap. Setting this back to false.
				temp = list[i]
				list[i] = list[i+1]
				list[i+1] = temp
			} // Next largest number has bubbled up at this point
		} // End pass through loop

		listLen-- // Minus array length by 1 because top most value(s) are in order
	} // End main loop
	return list
}

func main() {

	s := []int{2, 1, 4, 3, 8, 4, 9, 4, 5, 8, 1, 4}
	fmt.Println(BubbleSort(s))

  s2 := []int {1,2,3,4,5,6,7,4}
  fmt.Println(FindDuplicate(s2))
  
  s3 := []int {1,2,3,4,5,6,7,4}
  fmt.Println(FindDuplicate2(s3))

}
