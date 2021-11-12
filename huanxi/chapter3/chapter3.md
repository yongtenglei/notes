# Chapter 3 Gin

## Quick Start

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

  // init A gin Engine
	r := gin.Default()

  // prepare A router
  // func(c *gin.Context) is HandlerFunc Type
	r.GET("/ping", func(c *gin.Context) {
    // response
		c.String(http.StatusOK, "pong\n")
	})

  // start service in port 8080 by default
	r.Run()

}

```

测试:

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNgy1gwbjlhz2b1j307o01laa2.jpg"> </div>

## Deference between gin.New() and gin.Default()

gin.Default() 默认使用 logger 与 recoverey 中间件

gin.New() 则是只有基本功能的 Gin.Engine

## Gin 的 分组

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	pingPongGroup := r.Group("/pingpong")
	{
		pingPongGroup.GET("/ping", Pong)
	}

	dingdongGroup := r.Group("/dingdong")
	{
		dingdongGroup.GET("/ding", Dong)
	}

	r.Run()

}

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong\n")
}

func Dong(c *gin.Context) {
	c.String(http.StatusOK, "dong\n")
}

```

测试:

<div align=center><img src="https://tva1.sinaimg.cn/large/006cK6rNgy1gwbjy37pbqj30ab02nt9c.jpg"></div>

## Gin 获取 url 参数 Param ( : and \* )

```go

func main() {
	r := gin.Default()

	pingPongGroup := r.Group("/pingpong")
	{
		pingPongGroup.GET("/ping", Pong)
	}

	dingdongGroup := r.Group("/dingdong")
	{
		dingdongGroup.GET("/ding", Dong)

		// : 后的参数
		dingdongGroup.GET("/ding/:times", Dong2)

		// * 参数
		dingdongGroup.POST("/file/*all", File)
	}

	r.Run()

}
  func Dong2(c *gin.Context) {
	times := c.Param("times")

	if times == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
    return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("Ding %s Dong", times),
	})
}

func File(c *gin.Context) {
	all := c.Param("all")
	c.JSON(http.StatusOK, gin.H{
		"msg": all,
	})
}

```

测试:

<div align=center><img src="https://tvax1.sinaimg.cn/large/006cK6rNgy1gwbnx8x1chj30ad01nt91.jpg"></div>

<div align=center><img src="https://tva3.sinaimg.cn/large/006cK6rNgy1gwbo127c0jj30i70kmacb.jpg"></div>

## Gin 的参数校验 Binding

```go
type DingDong struct {
	Times int    `uri:"times" binding:"required"`
	Name  string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()

	pingPongGroup := r.Group("/pingpong")
	{
		pingPongGroup.GET("/ping", Pong)
	}

	dingdongGroup := r.Group("/dingdong")
	{
		dingdongGroup.GET("/ding", Dong)

		// : 后的参数
		dingdongGroup.GET("/ding/:times", Dong2)
		dingdongGroup.GET("/ding/:times/:name", Dong3)

		// * 参数
		dingdongGroup.POST("/file/*all", File)
	}

	r.Run()

}

func Dong3(c *gin.Context) {
	var dingdong DingDong
	if err := c.ShouldBindUri(&dingdong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
    return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("Ding %d-%s Dong", dingdong.Times, dingdong.Name),
	})
}

```

测试:

<div align=center><img src="https://tva3.sinaimg.cn/large/006cK6rNgy1gwbog5rkoxj30p50ki0vh.jpg"></div>

参数格式正确, 可以完成隐式转换.

<div align=center><img src="https://tvax3.sinaimg.cn/large/006cK6rNgy1gwboiwcg12j30p80kj41d.jpg"></div>

参数格式不正确, 不能完成隐式转换.

required 目前没用, 代表必须有此参数.

## 获得查询参数 query(key) DefaultQuery(key, defaultValue)

```go
func main() {
	r := gin.Default()

	pingPongGroup := r.Group("/pingpong")
	{
		pingPongGroup.GET("/ping", Pong)
	}

	dingdongGroup := r.Group("/dingdong")
	{
		dingdongGroup.GET("/ding", Dong)

		// : 后的参数
		dingdongGroup.GET("/ding/:times", Dong2)
		dingdongGroup.GET("/ding/:times/:name", Dong3)

		// query 参数
		dingdongGroup.GET("/ding/query", Dong4)

		// * 参数
		dingdongGroup.POST("/file/*all", File)
	}

	r.Run()

}

func Dong4(c *gin.Context) {
	times := c.DefaultQuery("times", 1)

	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("Ding %s Dong", times),
	})
}


```

测试:

<div align=center><img src="https://tva2.sinaimg.cn/large/006cK6rNgy1gwboo2nnosj30oz0lhdir.jpg"></div>

以上存在 query 参数

<div align=center><img src="https://tva1.sinaimg.cn/large/006cK6rNgy1gwbovse98cj30oz0kxgoe.jpg"></div>

不存在 query 参数使用默认值

## Gin Post 表单数据 PostForm DefaultPostForm

```go
func main() {
	r := gin.Default()

	pingPongGroup := r.Group("/pingpong")
	{
		pingPongGroup.GET("/ping", Pong)
	}

	dingdongGroup := r.Group("/dingdong")
	{
		dingdongGroup.GET("/ding", Dong)

		// : 后的参数
		dingdongGroup.GET("/ding/:times", Dong2)
		dingdongGroup.GET("/ding/:times/:name", Dong3)

		// query 参数
		dingdongGroup.GET("/ding/query", Dong4)

		// * 参数
		dingdongGroup.POST("/file/*all", File)
		// post 表单数据
		dingdongGroup.POST("/ding/add", Dong5)
	}

	r.Run()

}
func Dong5(c *gin.Context) {
	id := rand.Intn(10000)
	name := c.DefaultPostForm("name", "defaultName")

	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

```

测试:

<div align=center><img src="https://tva2.sinaimg.cn/large/006cK6rNgy1gwbp3up5quj30p70kz77k.jpg"></div>

以上存在 name 表单数据

<div align=center><img src="https://tvax2.sinaimg.cn/large/006cK6rNgy1gwbp9iqivwj30ox0ksdj6.jpg"></div>

无 name 表单数据时, 使用默认值

## 中间件 MiddleWares

```go
package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Ding() {
	time.Sleep(time.Second * 3)
	fmt.Println("Ding")
}

func Dong() {
	time.Sleep(time.Second * 2)
	fmt.Println("Dong")
}

func DingDong(c *gin.Context) {
	Ding()
	Dong()
}

// MiddleWare
func TimerMiddleWares(c *gin.Context) {
	start := time.Now()
	c.Next()
	end := time.Now().Sub(start)
	fmt.Println("time elapsed ", end)
}

func main() {
	r := gin.Default()

	r.Use(TimerMiddleWares)

	r.GET("/ding", DingDong)

	r.Run()
}

```

测试:

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNgy1gwcgdd066uj307e014weg.jpg"></div>

<div align=center><img src="https://tva2.sinaimg.cn/large/006cK6rNgy1gwcgdlwhazj30lz04bwg0.jpg"></div>

成功得到执行结果

### 中间件调用的另一种方式

```go
func TimerMiddleWares2() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now().Sub(start)
		fmt.Println("time elapsed ", end)

	}
}

func main() {
  r := gin.Default()

  r.Use(TimerMiddleWares2())

  r.GET("/ding", DingDong)

  r.Run()
}

```

or

```go
func main() {
	r := gin.Default()

	//r.Use(TimerMiddleWares2())

	r.GET("/ding", TimerMiddleWares2(), DingDong)

	r.Run()
}

```

只对一个 uri 使用中间件, 中间件写在处理业务的 HandlerFunc 之前.

### 中间件调用的顺序

```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Middlewares1(c *gin.Context) {
	fmt.Println("=====start 1=====")
	c.Next()
	fmt.Println("=====end 1=====")

}

func Middlewares2(c *gin.Context) {
	fmt.Println("=====start 2=====")
	c.Next()
	fmt.Println("=====end 2=====")

}
func Middlewares3(c *gin.Context) {
	fmt.Println("=====start 3=====")
	c.Next()
	fmt.Println("=====end 3=====")

}
func Middlewares4(c *gin.Context) {
	fmt.Println("=====start 4=====")
	c.Next()
	fmt.Println("=====end 4=====")

}
func Middlewares5(c *gin.Context) {
	fmt.Println("=====start 5=====")
	c.Next()
	fmt.Println("=====end 5=====")

}

func main() {
	r := gin.Default()

	r.GET("/ding", Middlewares1, Middlewares2, Middlewares3, Middlewares4, Middlewares5, Dong)

	r.Run()
}

func Dong(c *gin.Context) {
	fmt.Println("dong")
}

```

测试:

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNgy1gwcgdd066uj307e014weg.jpg"></div>

<div align=center><img src="https://tva3.sinaimg.cn/large/006cK6rNgy1gwchdccc7ij30me07v406.jpg"></div>

### 中间件的退出 c.Abort()

```go
func Middlewares2(c *gin.Context) {
	fmt.Println("=====start 2=====")
	// Middlewares Chain exit

	_, ok := c.Get("id")
	if !ok {
		c.Abort()
	}

	c.Next()
	fmt.Println("=====end 2=====")

}

```

以上代码必然会执行 c.Abort()

测试:

<div align=center><img src="https://tva2.sinaimg.cn/large/006cK6rNgy1gwchun2ztdj30m0051q4g.jpg"></div>

Abort()仅仅是将 MiddleWares Chain 中的 index 增加到 abortIndex, 故不影响此中间件之后代码的执行. 只是执行此中间件后, 不执行之后的中间件.

```go
func (c *Context) Abort() {
	c.index = abortIndex
}

const abortIndex int8 = math.MaxInt8 / 2

MaxInt8  = 1 << 7 - 1

```

## ShutDown graceful

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Dong(c *gin.Context) {
	fmt.Println("dong")
}

func main() {
	r := gin.Default()
	r.GET("/ding", Dong)

	// 优雅关机
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	fmt.Println("Ready for Server...")

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Shutdown error: ", err)
		return
	}
	fmt.Println("Shutdown succuessfully")
}

```

测试:

<div align=center><img src="https://tvax2.sinaimg.cn/large/006cK6rNgy1gwciimv5fpj30ls05e75r.jpg"></div>

按下 Ctrl-c 后, quit channel 接收到 syscall.SIGINT, 优雅退出.
