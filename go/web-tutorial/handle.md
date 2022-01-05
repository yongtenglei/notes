## Handle

```go
package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("myHandler\n"))
	// fmt.Fprintln(w, "myHandler") // 本质依然是调用 w.Write
}

type WelcomeHandler struct{}

func (wh *WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WelcomeHandler") // 本质依然是调用 w.Write
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.Handle("/", &myHandler{})

	http.Handle("/welcome", &WelcomeHandler{})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}

```

```go
// Handle registers the handler for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

// Handle registers the handler for the given pattern.
// If a handler already exists for pattern, Handle panics.
func (mux *ServeMux) Handle(pattern string, handler Handler)

```

调用 http.handle 方法, 将给 DefaultServeMux 注册 handler 方法

```go
// Handler 是一个接口
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

```

<div align=center><img src="https://tvax3.sinaimg.cn/large/006cK6rNly1gx53du2qzsj31gx0il48b.jpg">

</div>

使用 http.Handle 的场景是, 如果原本一个结构体有自己的结构 or 方法, 但它想要用作 web 用途, 只需要这个结构体实现, ServeHTTP 方法即可.

## HandleFunc

```go
package main

import (
	"fmt"
	"net/http"
)

func RootHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("myHandler\n"))
	// fmt.Fprintln(w, "myHandler") // 本质依然是调用 w.Write
}

func WelcomeHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WelcomeHandler") // 本质依然是调用 w.Write
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/", RootHandleFunc)

	http.HandleFunc("/welcome", WelcomeHandleFunc)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}

```

```go
// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

// HandleFunc registers the handler function for the given pattern.
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("http: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}

```

http.HandleFunc 本质依然是向 http.DefaultServeMux 注册 handler 方法, ServeMux 的 HandleFunc 方法本质是调用 ServeMux 的 Handle 方法.

为什么, HandlerFunc(适当签名的函数) 这个整体可以成为 Handler 类型?

```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

```

HandlerFunc() 将适当函数签名的函数类型转换为 HandlerFunc 类型, HandlerFunc 实现了 ServeHTTP 方法, 使得 HandlerFunc(适当签名的函数) 这个整体成为了 Handler 类型.

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNly1gx53urkz7bj30pz0r8n3m.jpg">

</div>
