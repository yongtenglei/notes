# 排序算法

## 冒泡算法

```go
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

```

## 选择排序

```go
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

```

## 插入排序

```go
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

```

## 堆排序

```go
package main

import "fmt"

// heapify 对节点为 i, 有 n 个元素的数组, 进行堆化
func heapify(arr []int, n, i int) {
	// 节点不能大于等于总节点数
	if i >= n {
		return
	}

	c1 := 2*i + 1
	c2 := 2*i + 2

	// 进行由小到大排列, 建造大顶堆,
	// 排序时, 先将最后一个节点与第一个节点交换, 然后砍断最后一个节点.
	max := i

	if c1 < n && arr[c1] > arr[max] {
		max = c1
	}

	if c2 < n && arr[c2] > arr[max] {
		max = c2
	}

	// 需要交换, 交换后, 需要堆交换后的节点进行heapify
	if max != i {
		arr[max], arr[i] = arr[i], arr[max]
		heapify(arr, n, max)
	}
}

// buildHeap 对整个数组进行堆化, 只需要从最后一个节点的parent及之前的依次向前堆化
// 最后一个节点 n - 1
// 最后一个节点的parent (最后一个节点) - 1 / 2
func buildHeap(arr []int) {
	n := len(arr)

	lastNode := n - 1
	parent := int((lastNode - 1) / 2)

	for i := parent; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func HeapSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	buildHeap(arr)

	// 排序时, 先将最后一个节点与第一个节点交换, 然后砍断最后一个节点.
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}

}

func main() {
	// 测试 heapify
	testHeapify := []int{2, 4, 5, 3}
	heapify(testHeapify, len(testHeapify), 0)
	fmt.Println(testHeapify)

	// 测试 buildHeap
	testBuildHeap := []int{1, 4, 5, 9, 3, 5, 7, 6}
	buildHeap(testBuildHeap)
	fmt.Println(testBuildHeap)

	// HeapSort
	arr := []int{1, 4, 5, 9, 3, 5, 7, 6}
	HeapSort(arr)
	fmt.Println(arr)

}

```

## 归并排序

```go
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

```

## 快速排序

```go
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

```
