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
