package main

import (
	"fmt"
)

func merge(L, R []int) []int {
	result := make([]int, 0)

	l, r := 0, 0

	// 对比每个切片的前面的元素, 更小的先出队, 直到一个切片的元素为空
	for l < len(L) && r < len(R) {
		if L[l] < R[r] {
			result = append(result, L[l])
			l++
		} else {
			result = append(result, R[r])
			r++
		}
	}

	// 将剩下的元素依次出队
	result = append(result, L[l:]...)
	result = append(result, R[r:]...)

	return result
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	i := len(arr) / 2

	L := MergeSort(arr[:i])
	R := MergeSort(arr[i:])
	return merge(L, R)

}

func main() {
	// testMerge
	L := []int{5, 3, 5, 7, 1}
	R := []int{2, 9, 10, 8, 4}
	fmt.Println(merge(L, R))

	arr := []int{1, 4, 5, 9, 3, 5, 7, 6}
	fmt.Println(MergeSort(arr))

}
