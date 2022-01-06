package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var res int
	strArr := make([]string, 0)
	strArr = append(strArr, s)
	res = len(strArr)

	for _, v := range strArr {
		fmt.Println(string(v[2]))
	}
	return res
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
