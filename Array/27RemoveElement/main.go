package main

import "fmt"

/*
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.
*/
func main() {
	num1 := []int{3,2,2,3}
	val1 := 3
	fmt.Println(removeElement(num1, val1))

	num2 := []int{0,1,2,2,3,0,4,2}
	val2 := 2
	fmt.Println(removeElement(num2, val2))

}

func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	// fmt.Println(nums)
	return index
}
