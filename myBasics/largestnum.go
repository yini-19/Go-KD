package main

import "fmt"

func largestNum(nums []int) int {
	largest := nums[0]

	for _, v := range nums {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	numbr := []int{23, 12, 34, 54, 23, 5}
	fmt.Println(largestNum(numbr))
}
