package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m+n == len(nums1) {
		for i := range nums2 {
			nums1[m] = nums2[i]
			m += 1
		}
	}
	for i := range nums1 {
		for j := i + 1; j < len(nums1); j++ {
			if nums1[i] > nums1[j] {
				temp := nums1[i]
				nums1[i] = nums1[j]
				nums1[j] = temp
			}
		}
	}
	fmt.Println(nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3

	merge(nums1, m, nums2, n)
}
