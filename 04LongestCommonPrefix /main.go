package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[0 : len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}

	}
	return prefix
}

func main() {

	strs := []string{"flower", "flow", "flight"}
	strs1 := []string{"dog", "racecar", "car"}
	strs2 := []string{"a"}

	fmt.Println(longestCommonPrefix(strs))
	fmt.Println(longestCommonPrefix(strs1))
	fmt.Println(longestCommonPrefix(strs2))
}
