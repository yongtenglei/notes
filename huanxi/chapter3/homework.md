# Chapeter 3 Homework

1. 介绍一下 Gin

   Gin 是一个用 Go (Golang) 编写的 HTTP web 轻量级框架. 它是一个类似于 martini 但拥有更好性能的 API 框架, 优于 httprouter, 速度提高了近 40 倍. 天生支持中间件, 路由, 初级的参数验证等等.

2. Defalut 和 New 有什么不同

   gin.Default() 默认使用 logger 与 recoverey 中间件

   gin.New() 则是只有基本功能的 Gin.Engine

3. 针对表单做 restful 的处理

```go
func RouterSetup() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(conf.Conf.AppConf.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		// tags
		apiV1.GET("/tags", v1.GetTags)          // 获取多个tag
		apiV1.GET("/tags/:id", v1.GetTag)       // 获取特定id的tag
		apiV1.POST("/tags", v1.AddTag)          // 创建一个tag
		apiV1.PUT("/tags/:id", v1.EditTag)      // 修改tag
		apiV1.DELETE("/tags/:id", v1.DeleteTag) // 删除tag

		// articles
		apiV1.GET("/articles", v1.GetArticles)          // 获取多个tag
		apiV1.GET("/articles/:id", v1.GetArticle)       // 获取特定id的tag
		apiV1.POST("/articles", v1.AddArticle)          // 创建一个tag
		apiV1.PUT("/articles/:id", v1.EditArticle)      // 修改tag
		apiV1.DELETE("/articles/:id", v1.DeleteArticle) // 删除tag

	}
	return r
}

```

4. 针对 JSON 做 restful 的处理 (ShouldBind)

```go
// models/tag.go
package models

import (
	"blog/conf"
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}


```

```go
// models/article.go
package models

import (
	"blog/conf"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag" gorm:"foreignkey:TagID"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

```

```go
// controllers/v1/tag.go
package v1

import (
	"blog/conf"
	"blog/logic"
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"

	"github.com/gin-gonic/gin"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	page := util.GetPage(c)
	size, err := util.StrToInt(c.DefaultQuery("size", util.IntToStr(conf.Conf.AppConf.PageSize)))
	if err != nil {
		code := e.INVALID_PARAMS
		util.ResposeWithError(c, code)
		return
	}

	var tags []models.Tag

	// 业务处理
	tags, err = logic.GetTags(page, size)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	// 返回
	util.ResposeWithSuccessData(c, tags)
}

//新增文章标签
func AddTag(c *gin.Context) {
	var tag models.Tag

	// 获取参数, 校验参数
	if err := c.ShouldBindJSON(&tag); err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	// 业务处理
	if err := logic.AddTag(&tag); err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	// 返回
	util.ResposeWithSuccess(c)
}

//获取特定id的文章标签
func GetTag(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	var tag models.Tag
	tag, err = logic.GetTag(id)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccessData(c, tag)
}

//修改文章标签
func EditTag(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	var tag models.Tag
	err = c.ShouldBindJSON(&tag)
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	tag.ID = uint(id)

	err = logic.EditTag(&tag)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccess(c)

}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	err = logic.DeleteTag(id)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccess(c)
}

```

```go
// controllers/v1/article.go
package v1

import (
	"blog/conf"
	"blog/logic"
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"

	"github.com/gin-gonic/gin"
)

//获取多个文章
func GetArticles(c *gin.Context) {
	page := util.GetPage(c)
	size, err := util.StrToInt(c.DefaultQuery("size", util.IntToStr(conf.Conf.AppConf.PageSize)))
	if err != nil {
		code := e.INVALID_PARAMS
		util.ResposeWithError(c, code)
		return
	}

	var articles []models.Article

	// 业务处理
	articles, err = logic.GetArticles(page, size)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	// 返回
	util.ResposeWithSuccessData(c, articles)
}

//新增文章
func AddArticle(c *gin.Context) {
	var article models.Article

	// 获取参数, 校验参数
	if err := c.ShouldBindJSON(&article); err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	// 业务处理
	if err := logic.AddArticle(&article); err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	// 返回
	util.ResposeWithSuccess(c)
}

//获取特定id的文章
func GetArticle(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	var tag models.Article
	tag, err = logic.GetArticle(id)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccessData(c, tag)
}

//修改文章
func EditArticle(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	var article models.Article
	err = c.ShouldBindJSON(&article)
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	article.ID = uint(id)

	err = logic.EditArticle(&article)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccess(c)

}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, err := util.StrToInt(c.Param("id"))
	if err != nil {
		util.ResposeWithError(c, e.INVALID_PARAMS)
		return
	}

	err = logic.DeleteArticle(id)
	if err != nil {
		util.ResposeWithError(c, e.ERROR)
		return
	}

	util.ResposeWithSuccess(c)
}

```

5. 举例如何用到中间件

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

`中间件调用的另一种方式`

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

6. 如果多个中间件，调用顺序如何？（代码演示）

中间件调用的顺序

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

7. Gin 终止其中一个中间件，要如何做？

   中间件的退出 c.Abort()

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

8. 如何优雅退出 Gin 的程序
   ShutDown graceful

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

9. 利用之前的 Go 语言学到的知识，可以总结出 一次请求处理的大体流程？

   1. r.Run() 启动服务, 其是 http.ListenAndServe 的封装

   2. 调用 engine.ServeHTTP(w http.ResponseWriter, req \*http.Request)从 sync.pool 中获取 Context 对象. Gin 中使用 Context 池是一个很巧妙的设计.

   3. engine.ServeHTTP 中调用 engine.handleHTTPRequest(\*gin.Context)方法处理请求, 从 Context 对象中获得 HandlerFunc 链, 并进行处理.

   4. 从处理业务逻辑的 HandlerFunc 中返回响应.

10. gin 返回 html 的处理 (选做)

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// r.LoadHTMLGlob("templates/**/*") 加载templates文件夹中的所有HTML文件
	r.LoadHTMLFiles("index.html")

	r.GET("/index", GetIndex)

	r.Run()
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"name": "rey",
	})
}

```

```go
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width" />
    <title>Index</title>
  </head>
  <body>
    <h1>Hello, {{ .name }}</h1>
  </body>
</html>

```

测试:

<div align=center><img src="https://tvax3.sinaimg.cn/large/006cK6rNgy1gwco2ksb8dj30et0480t1.jpg"></div>

11. gin 如何处理静态文件（选做）

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// r.LoadHTMLGlob("templates/**/*") 加载templates文件夹中的所有HTML文件
	r.LoadHTMLFiles("index.html")

	// load css file
	r.Static("/static", "./static")

	r.GET("/index", GetIndex)

	r.Run()
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"name": "rey",
	})
}
```

```html
<!./index.html>
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width" />
    <title>Index</title>
    <link
      rel="stylesheet"
      href="/static/index.css"
      type="text/css"
      media="all"
    />
  </head>
  <body>
    <h1>Hello, {{ .name }}</h1>
  </body>
</html>
```

```css
/* static/index.css*/
h1 {
  background: #090;
  color: cyan;
  font-size: 50px;
}
```

测试:

<div align=center><img src="https://tvax4.sinaimg.cn/large/006cK6rNgy1gwcoef3esjj30e3055gm7.jpg"></div>
