package main

import "fmt"

func findMid(arr []int, l, r int) int {
	i, j := l, r

	// 腾出左边第一个位置
	pivot := arr[l]

	for i < j {
		// 从右边寻找比pivot小的数字
		for i < j && arr[j] >= pivot {
			j--
		}

		// 将找到的数字填入预留的i中后, i向右腾出一个新位置
		if i < j {
			arr[i] = arr[j]
			i++
		}

		// 从左边找到比pivot大的数字, 相等时也需要填入右边
		for i < j && arr[i] < pivot {
			i++
		}

		if i < j {
			arr[j] = arr[i]
			j--
		}
	}

	// 最后当 l == r, 将此位置填入 pivot
	// 至此, 得到: (nums < pivot) pivot (nums > pivot)
	arr[i] = pivot

	// l 为中心位置
	return i

}
func QueckSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	i := findMid(arr, 0, len(arr)-1)

	// 向两边执行 QueckSort
	QueckSort(arr[:i])   // 不包含 i
	QueckSort(arr[i+1:]) // 包含i+1
}

func main() {
	arr := []int{1, 4, 5, 9, 3, 5, 7, 6}
	QueckSort(arr)
	fmt.Println(arr)

}
