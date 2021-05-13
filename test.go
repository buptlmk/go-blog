package main

import "fmt"

func main() {
	a := 1<<63 - 1
	var b int
	b = a + 1
	fmt.Println(a, b)
	nums := []int{1, 2, -1, 1, 3, 5, -2}

	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
		fmt.Println(nums)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			//return i + 1
			fmt.Println(i + 1)
			break
		}
	}

}
