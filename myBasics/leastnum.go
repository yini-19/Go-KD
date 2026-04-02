package main

func leastNum(nums []int) int {
	least := nums[0]
	for _, v := range nums {

		if v < least {
			least = v
		}

	}
	return least
}

// func main() {
// 	numbr := []int{23, 12, 34, 54, 23, 5}
// 	fmt.Println(leastNum(numbr))
// }
