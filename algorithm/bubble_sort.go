package main

import "fmt"

func bubblesort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Println(arr)
		}
	}
	return arr
}

func main() {
	arr := []int{1, 2, 34, 12, 32, 34, 22, 1}
	fmt.Println(bubblesort(arr))
}
