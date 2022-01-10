package main

import "fmt"

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

func main() {
	arr := []int{1, 4, 5, 9, 3, 5, 7, 6}
	SelectSort(arr)
	fmt.Println(arr)
}
