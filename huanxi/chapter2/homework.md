1. 说明一下接口是什么，和面向对象有什么关系？
   （选做部分:如果你知道 java，那么，Go 语言的接口和 java 接口有什么不同？）

一种”约定”, 定义了对象的行为, 在 go 语言中可以使用接口实现面向对象中的多态.

Go 语言的接口是非侵入式接口, 即实现了某个接口的所有方法, 就实现了这个接口.

Java 中必须显式的指定 implement 了哪一种接口.

2. 举例说明鸭子类型

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

type Doggy struct {
    Name string
}

func (d Doggy) Swimming() {
    fmt.Println(d.Name + "is swimming")
}

func main() {
    var ducky DuckyType
    ducky = Ducky{Name: "TangDuck"}
    ducky.Swimming()

    ducky = Doggy{Name: "pluto"}
    ducky.Swimming()
}
```

3. go 语言中的标准的接口，有哪些？ 并举例说明 1-2 个接口的实现，通过接口如何实现多态？

标准接口: io 接口 Writer Reader WriterReader

接口的实现及多态:

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

4. 函数传值和传引用有何不同？ 各举一个例子

函数传值为值的拷贝, 不能更改原始数据, 只能进行读等操作. 传引用为传递指针, 可以更改原始数据.

传引用, 可以更改原数组的值.

```go
package main

import "fmt"

func Add1(s *[5]int) {
	for i := 0; i < len(s); i++ {
		s[i] += 1
	}

}
func main() {
	s := [5]int{1, 2, 3, 4, 5}
	fmt.Println(s)
	Add1(&s)
	fmt.Println(s)

}
```

传值, 不能改变原始数据.

```go
package main

import "fmt"

func Add1(s [5]int) {
	for i := 0; i < len(s); i++ {
		s[i] += 1
	}

}
func main() {
	s := [5]int{1, 2, 3, 4, 5}
	fmt.Println(s)
	Add1(s)
	fmt.Println(s)

}

```

传递切片时, 不需要传引用, 因为切片是引用类型.

5. 举例说明 函数变量

```go
package main

import "fmt"

func FuncVariable() {
	fmt.Println("function can be stored into variable")
}

func main() {
	f := FuncVariable
	f()

}
// output:
// function can be stored into variable
```

6. 举例说明 匿名函数

```go
package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("this is a anonymous function")
	}
	f()
}
// output:
// this is a anonymous function

```

7. 举例说明 闭包

```go
package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(str string) string {
	return func(str string) string {
		if !strings.HasSuffix(str, suffix) {
			return str + suffix
		}
		return str
	}
}

func main() {
	checkSuffix := makeSuffix(".txt")
	file := checkSuffix("rey")
	fmt.Println(file)
	// output: rey.txt
}

```

8.  举例说明 变长函数

```go
package main

import (
	"fmt"
)

func MutiSummation(nums ...int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return
}

func main() {
	sum := MutiSummation(1, 2, 3)
	fmt.Println(sum)
	// output: 6
}

```

9.  延长函数的调用顺序是什么？ 举例说明

先声明的 defer 函数, 后执行

```go
package main

import (
	"fmt"
)

func DeferSequence() {
	defer fmt.Println("==========1==========")
	defer fmt.Println("==========2==========")
	defer fmt.Println("==========3==========")
	defer fmt.Println("==========4==========")
	defer fmt.Println("==========5==========")
	fmt.Println("==========6==========")
}

func main() {
	DeferSequence()
}
// output:
// ==========6==========
// ==========5==========
// ==========4==========
// ==========3==========
// ==========2==========
// ==========1==========
```

10. go 语言是如何做测试的？ 举例说明

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

11. 如何理解 线程安全？

尽管使用多个线程(进程), 依然能够像一个串行程序一样工作良好. 不会发生竞争, 死锁等问题, 协同完好的完成任务.

12. 如何理解 Go 语言的并发模型？

Go 语言的并发模型时 CSP, 主张使用消息传递来共享内存, 而不是使用共享内存来传递消息.

Go 语言的 MPG 模型, 在运行时, 只有一个 M 与 P 一一绑定, P 则可以携带多个 G, 多个 G 在 P 的队列中等待执行. 有时, 因为 I/O 等原因阻塞了 M, P 则会携带除阻塞的 G, 与其他 M 绑定.

13. 缓冲通道与无缓冲通道有合不同？

无缓冲通道需要读取和写入同时准备好才能正常工作, 否则会引起死锁.

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

同步读取和写入可以正常工作.

有缓冲的 channel

写: 当缓冲区满时, 写操作会陷入阻塞

读: 当缓冲区为 size 为 0 时, 读操作会陷入阻塞

读写值为 nil 的 channel 时, 会 panic

同步 channel 需要同时准备好, 协调运作

channel 需要 close 后, 才可以进行 range 操作, 否则会 panic

14. 单向通道优势是什么？

采取权限最小原则, 可以保护一些不想被修改的数据被误操作, 减少犯错.

15. 关闭通道，会造成哪些影响？

关闭通道后, 再往通道写入数据会引发 panic. 当有缓冲区的 channel 被读取完后, 再次读取会读取到 channel 的默认值.

16. 什么场景使用 select?

从不同的携程中获取随机顺序的值时, 使用 select.

17. 举例说明 mutex 和 rwmutex

mutex, 读写均会独占资源.

rwmutex, 写操作独占资源, 读操作共享资源.

18. 举例说明 条件变量

生产者消费者模型

```go
func main() {
    cond := sync.NewCond(new(sync.Mutex))
    condition := 0

    // 消费者
    go func() {
        for {
            // 消费者开始消费时，锁住
            cond.L.Lock()
            // 如果没有可消费的值，则等待
            for condition == 0 {
                cond.Wait()
            }
            // 消费
            condition--
            fmt.Printf("Consumer: %d\n", condition)

            // 唤醒一个生产者
            cond.Signal()
            // 解锁
            cond.L.Unlock()
        }
    }()

    // 生产者
    for {
        // 生产者开始生产
        cond.L.Lock()

        // 当生产太多时，等待消费者消费
        for condition == 100 {
            cond.Wait()
        }
        // 生产
        condition++
        fmt.Printf("Producer: %d\n", condition)

        // 通知消费者可以开始消费了
        cond.Signal()
        // 解锁
        cond.L.Unlock()
    }
}
```

19. 举例说明 WaitGroup

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

20. 举例说明 context.Context

Context with timeout

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
	loop:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("timeout")
				break loop
			default:
				fmt.Println("still working")
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Wait()
}

```

21. 说说你对 GO 语言错误处理的理解？

go 语言中的错误以一种更加灵活的方式, 像参数一样可以在函数间传递, 不仅如此, 它更可以携带信息, 它是一个值.

但也有缺点, 造成代码的冗余, 但这些都不是 golang 的错. golang 有能力优雅的处理 error, 因为 error 是一个值. 详情 [Error are values](https://go.dev/blog/errors-are-values)

22. go 语言如何做依赖管理？

经过多年的迭代与官方的支持, go modules 成为新一代的 go 依赖管理工具.

vendor 实现了将项目引用的外部包引进 vendor 目录, 使 go build 的时候, 直接从 vendor 目录中拉取依赖.

23. go mod 常用命令有哪些？

`go mod init` 初始化项目
`go mod tidy` 自动拉取已引用依赖
`go get pakage_name` 拉取依赖
