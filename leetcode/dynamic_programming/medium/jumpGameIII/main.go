package main

import "fmt"

func canReach(arr []int, start int) bool {
	seen := make(map[int]struct{}, len(arr))
	return dfs(arr, start, seen)
}

func dfs(arr []int, start int, seen map[int]struct{}) bool {
	if start < 0 || start >= len(arr) {
		return false
	}

	if arr[start] == 0 {
		return true
	}

	if _, ok := seen[start]; ok {
		return false
	} else {
		seen[start] = struct{}{}
	}

	return dfs(arr, start-arr[start], seen) || dfs(arr, start+arr[start], seen)

}
func main() {
	//arr := []int{4, 2, 3, 0, 3, 1, 2}
	arr := []int{3, 0, 2, 1, 2}
	fmt.Println(canReach(arr, 2))
}
