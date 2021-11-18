# Three Way to implement Fibonacci

```go
package main

import (
	"fmt"
	"time"
)

// Fibonacci implemented with definition
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// FibonacciWithNB implemented with definition as well as having a map to store values which are calculated before (Notebook)
func FibonacciWithNB(n int, nb map[int]int) int {
	if v, ok := nb[n]; ok {
		return v
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	v := FibonacciWithNB(n-1, nb) + FibonacciWithNB(n-2, nb)

	nb[n] = v

	return v
}

// FibonacciIter implemented in iteration way
func FibonacciIter(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}

func main() {
	fmt.Println("======Fibonacci=======")
	s := time.Now()
	a := Fibonacci(0)
	e := time.Now().Sub(s)
	fmt.Println(a, e)

	fmt.Println("======FibonacciWithNB=======")
	s = time.Now()
	b := FibonacciWithNB(0, map[int]int{})
	e = time.Now().Sub(s)
	fmt.Println(b, e)

	fmt.Println("======FibonacciIter=======")
	s = time.Now()
	c := FibonacciIter(0)
	e = time.Now().Sub(s)
	fmt.Println(c, e)
}
```

## Speed to find "0"

<div align=center><img src="https://tvax1.sinaimg.cn/large/006cK6rNgy1gwgwq60q0kj308x069mxm.jpg"></div>

最快的是迭代, map 的使用消耗了一些资源.

## Speed to find "50"

<div align=center><img src="https://tva2.sinaimg.cn/large/006cK6rNgy1gwgww5uyu8j30cp0eeae0.jpg"></div>

较大的数列, 迭代依旧表现出了较好的性能. 纯递归的方法已经十分吃力.

## More

<div align=center><img src="https://tvax4.sinaimg.cn/large/006cK6rNgy1gwgx18jx5uj30eg0d6djj.jpg"></div>

迭代方法依旧有效.

## 迭代法的另一个应用

前面的算法只是算出了 n 对应的 fibonacci 数列中的数, 但没有保存数列, 由表可以看出 fibonacci 数列是 a 组成的数列, 所以我们可以使用一个切片 or 数组来保存并返回.

<div align=center>

| a   | b   | n   |
| --- | --- | --- |
| 0   | 1   | 0   |
| 1   | 1   | 1   |
| 1   | 2   | 2   |
| 2   | 3   | 3   |
| 3   | 5   | 4   |
| 5   | 8   | 5   |
| 8   | 13  | 6   |
| ..  | ..  | ..  |

</div>

`一个例子:`

```go
// FibonacciIter implemented in iteration way and return Fibonacci Array
func FibonacciIterArray(n int) []int {
	array := []int{0}
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
		array = append(array, a)
	}

	return array
}

func main() {
	fmt.Println("======FibonacciIterArray=======")
	c1 := FibonacciIterArray(0)
	fmt.Println(c1)
	c2 := FibonacciIterArray(1)
	fmt.Println(c2)
	c3 := FibonacciIterArray(2)
	fmt.Println(c3)
}

```

<div align=center><img src="https://tva2.sinaimg.cn/large/006cK6rNgy1gwgxb4q7xwj30et0b3gnw.jpg"></div>

<div align=center><img src="https://tva2.sinaimg.cn/large/006cK6rNgy1gwgxdy0jvhj30kp05wgnp.jpg"></div>
