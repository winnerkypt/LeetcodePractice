package main

import (
	"fmt"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	list1 := []int{1, 2, 4}
	list2 := []int{1, 3, 4}

	fmt.Println(mergeTwoLists(list1, list2))
}

func mergeTwoLists(list1 []int, list2 []int) []int {
	slices.Sort(list1)
	slices.Sort(list2)

	result := []int{}
	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] <= list2[j] {
				result = append(result, list1[i])
				result = append(result, list2[j])

			} else {
				result = append(result, list2[j])
				result = append(result, list1[i])

			}
		}
	}
	return result
}
