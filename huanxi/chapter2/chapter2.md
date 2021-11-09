# Chapter 2

## 接口

一种"契约", go 语言的接口是非侵入式接口, 即实现了某个接口的所有方法, 就实现了这个接口.

鸭子 🦆 类型: 如果某个事物能够做鸭子 🦆 可以做的事情, 那它就是一只鸭子 🦆.

### 接口的实现

```go
package main

import "fmt"

type DuckyType interface {
	Swimming()
}

type Ducky struct {
	Name string
}

func (d Ducky) Swimming() {
	fmt.Println(d.Name + "is swimming")
}

func main() {
	var ducky DuckyType
	ducky = Ducky{Name: "TangDuck"}
	ducky.Swimming()
}
```

### 接口的组合

```go
package main

import "fmt"

type DuckyType interface {
	Swimming()
}

type Ducky struct {
	Name string
}

func (d Ducky) Swimming() {
	fmt.Println(d.Name + "is swimming")
}

type DoggyType interface {
	Woofing()
}

type Doggy struct {
	Name string
}

func (d Doggy) Woofing() {
	fmt.Println(d.Name + "is woofing")
}

func (d Doggy) Swimming() {
	fmt.Println(d.Name + "is swimming")
}

type DockyType interface {
	DuckyType
	DoggyType
}

func main() {
	var ducky DuckyType
	ducky = Ducky{Name: "TangDuck"}
	ducky.Swimming()

	var doggy DoggyType
	doggy = Doggy{Name: "pluto"}
	doggy.Woofing()

	var docky DockyType
	docky = doggy.(DockyType)
	docky.Swimming()
	docky.Woofing()

}

```

### 接口的断言

```go
package main

import "fmt"

func ShowType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int type")
	case float32, float64:
		fmt.Println("float type")

	default:
		fmt.Println("not a number type")
	}

}

func main() {
	i := 1
	f32 := float32(1.1)
	f64 := float64(1.1)
	s := "non-number"
	ShowType(i)
	ShowType(f32)
	ShowType(f64)
	ShowType(s)
}
```

## defer 和 recover

```go
package main

import "fmt"

func Test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("==========1=========")
	fmt.Println("==========2=========")
	panic("==========3=========")
	fmt.Println("==========4=========")

}

func main() {

	fmt.Println("=======start======")
	Test()
	fmt.Println("=======end========")
}
```

## test

1. 需要测试的文件 main.go, 则测试文件命名为 main_test.go.

2. 需要测试的函数为 Show(), 则测试函数命名为 TestShow().

```go
// main.go
package main

func Show() string {
	return "test successfully"
}

func main() {
	Show()
}

// main_test.go
package main

import "testing"

func TestShow(t *testing.T) {
	s := Show()
	want := "test successfully"
	if s != want {
		t.Errorf("Show() = %v, want = %v", s, want)
	}
	t.Logf("Show() = %v, want = %v", s, want)

}
```

## channel

1. 声明

`var ChanName chan Type`

`ChanName := make(chan Type[, buffer_size])`

2. channel 读写

`variable := <- chan` 读

`chan <- variable` 写

3. 限制 channel 的权限

创建只能创建双向通道, 之后可以限权

`func(ch chan <- )` 在函数中只写

`func(ch <- chan)` 在函数中只读

在读操作中, 返回两个值, value 与 ok (可省略), ok 可判断 channel 是否关闭.

    1. channel 关闭后, ok 必为false

    2. ok为false, 不一定channel关闭, 可能为空, 产生的为默认值

```go
package main

import "fmt"

func main() {
	c := make(chan int, 10)
	c <- 1
	v, ok := <-c
	fmt.Println(v, ok)
	close(c)
	v, ok = <-c
	fmt.Println(v, ok)

}
// output:
// 1 true
// 0 false
```

4.  阻塞

    1. 同步 channel

    ```go
    package main

    import "fmt"

    func main() {
        c := make(chan int)
        c <- 1
        fmt.Println(<-c)
    }
    // deadlock
    ```

    没有缓冲区的 channel, 不管是读还是写都会陷入阻塞

    执行 `c <- 1` 时, 陷入阻塞, `<- c` 永远没有机会执行, 因此 deadlock.

    ```go
    package main

     import (
       "fmt"
       "time"
     )

     func main() {
       c := make(chan int)
       go func() { c <- 1 }()
       fmt.Println(<-c)
       time.Sleep(time.Second)
     }
     // outpute: 1
    ```

    2. 有缓冲的 channel

    写: 当缓冲区满时, 写操作会陷入阻塞

    读: 当缓冲区为 size 为 0 时, 读操作会陷入阻塞

    读写值为 nil 的 channel 时, 会 panic

    同步 channel 需要同时准备好, 协调运作

    channel 需要 close 后, 才可以进行 range 操作, 否则会 panic

    ```go
    package main

    import "fmt"

    func main() {
      c := make(chan int, 10)
      for i := 0; i < 9; i++ {
        c <- i
      }

      close(c)

      for i := range c {
        fmt.Println(i)
      }
    }
    ```

## WaitGroup

```go
 package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintCow(wg sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("🦬")
		time.Sleep(time.Second)
	}
}

func PrintSheep(wg sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("🐏")
		time.Sleep(time.Second)
	}
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go PrintCow(wg)
	go PrintSheep(wg)

	wg.Wait()

}
```

`wg.Add(num)` 内部计数器加上 num 的值

`wg.Done()` 内部计数器减 1

`wg.Wait()` 阻塞至等待内部计数器为 0

## select

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 5)
	t := time.After(time.Second * 4)

	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- i
			time.Sleep(time.Second)
		}
	}()

Loop:
	for {
		select {
		case v := <-ch1:
			fmt.Println("ch1 = ", v)

		case v := <-ch2:
			fmt.Println("ch2 = ", v)

		case <-t:
			fmt.Println("time out")
			close(ch1)
			close(ch2)
			break Loop
		default:
		}
	}

}

```

## 互斥锁 Mutex, RWMutex

### 互斥锁

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type NewsFactory struct {
	m    sync.Mutex
	news map[string]int
}

func (nf *NewsFactory) WNews() {
	nf.m.Lock()
	defer nf.m.Unlock()
	nf.news["big thing"]++
}

func (nf *NewsFactory) RNews() {
	nf.m.Lock()
	defer nf.m.Unlock()
	fmt.Println(nf.news["big thing"])
}

func main() {
	newsfactory := &NewsFactory{
		m: sync.Mutex{},
		news: map[string]int{
			"big thing": 0,
		},
	}

	for i := 0; i < 10; i++ {
		go func() {
			newsfactory.WNews()
		}()
	}

	time.Sleep(time.Second * 3)

	newsfactory.RNews()

}

```

### 读写锁 RWMutex

1. 写操作时唯一抢占资源

2. 读操作时共享资源
