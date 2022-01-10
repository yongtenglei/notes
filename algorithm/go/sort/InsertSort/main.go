package main

import "fmt"

func InsertSort(arr []int) {
	// 指向需要排序的元素
	for i := 1; i < len(arr); i++ {
		// 与已经排序好的元素比较
		for j := i; j > 0; j-- {
			// 如果左边的比待排序元素小, 则完成该元素的排序
			if arr[j-1] < arr[j] {
				break
			}

			// 交换
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

func main() {
	arr := []int{1, 4, 5, 9, 3, 5, 7, 6}
	InsertSort(arr)
	fmt.Println(arr)
}
