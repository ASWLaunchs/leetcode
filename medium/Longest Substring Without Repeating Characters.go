package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	var index int //下标
	var res int   //结果
	//初始化结果
	res = 0
	//创建切片
	strArr := make([]string, 0)
	//转化为数组
	for i := 0; i < len(s); i++ {
		strArr = append(strArr, string(s[i]))
	}

	//排序
	for i, v := range strArr {
		index = i
		for j := index + 1; j < len(strArr); j++ {
			if v == strArr[j] {
				if len(strArr[index:j]) > res {
					res = len(strArr[index:j])
				}
				fmt.Printf("从下标%d开始的切片为%s \n", index, strArr[index:j])
				index = j
			} else if v != strArr[j] && j == len(strArr)-1 {
				fmt.Printf("从下标%d开始的切片为%s \n", index, strArr[0:j])
			}
		}
	}
	return res
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
	s = " "
	fmt.Println(lengthOfLongestSubstring(s))
	s = "au"
	fmt.Println(lengthOfLongestSubstring(s))
}
