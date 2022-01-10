package main

import "fmt"

func BubbleSort(arr []int) {
	// 执行 n - 1 次排序
	for i := 1; i <= len(arr)-1; i++ {
		// 每次排序执行 n - i 次交换
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{1, 4, 5, 9, 3, 5, 7, 6}
	BubbleSort(arr)
	fmt.Println(arr)
}
