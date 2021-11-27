# Chapter5

1. 介绍一下 RPC

RPC（Remote Procedure Call Protocol）——远程过程调用协议，它是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络技术的协议。RPC 协议假定某些传输协议的存在，如 TCP 或 UDP，为通信程序之间携带信息数据。在 OSI 网络通信模型中，RPC 跨越了传输层和应用层。RPC 使得开发包括网络分布式多程序在内的应用程序更加容易。

RPC 采用客户机/服务器模式。请求程序就是一个客户机，而服务提供程序就是一个服务器。首先，客户机调用进程发送一个有进程参数的调用信息到服务进程，然后等待应答信息。在服务器端，进程保持睡眠状态直到调用信息到达为止。当一个调用信息到达，服务器获得进程参数，计算结果，发送答复信息，然后等待下一个调用信息，最后，客户端调用进程接收答复信息，获得进程结果，然后调用执行继续进行。

RPC 协议的主要目的是做到不同服务间调用方法像同一服务间调用本地方法一样

1：服务化/微服务

2：分布式系统架构

3：服务可重用

4：系统间交互调用

RPC 通信过程:

<div align=center><img src="http://81.68.236.200:8181/uploads/lee/images/m_d36c644ce31f3f548b9b6fcf565f398e_r.png
">
</div>

<div align=center><img src="https://tva3.sinaimg.cn/large/006cK6rNly1gwsdyeei71j31440ny7gc.jpg">
</div>

    1. Client 向 Server 发起请求

    2. Client stub 将网络信息, 参数信息等编码, 发送给 Client sockets

    3. sockets 通过网络信息将参数信息等网络传输给 Server sockets

    4. Server sockets 将参数信息传递给 server stub

    5. Server stub 对信息进行解码

    6. 根据参数信息进行处理

    7. 将 response 传给 server stub 进行编码

    8. 网络传输给 client sockets

    9. Client sockets 将 Response 信息传递给 Client stub 对 Reponse 信息进行解码

    10. Client 收到 Response

2. POST http 客户端调用

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	// GET
	//resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
	//if err != nil {
	//// handle error
	//}

	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//// handle error
	//}

	//fmt.Println(string(body))

	// POST 1
	fmt.Println("=========method1===========")
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=rey"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	// POST 2 FORM
	fmt.Println("=========method2===========")
	resp, err = http.PostForm("http://www.01happy.com/demo/accept.php",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

```

结果:

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNly1gwsgu9r9zcj31hc0ibn1x.jpg">
</div>

3. RPC 四要素

IDL Interface Description Language: GO or other

Server: GO or other

传输协议：TCP

数据协议：protobuf

4. GO 原生的跨语言的并发高的实例

原生高并发 rpc:

```go
// Server
package main

import (
	"net"
	"net/rpc"
)

type FoodService struct {
}

func (f *FoodService) SayName(req string, res *string) error {
	*res = "ordered " + req
	return nil
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	err = rpc.RegisterName("FoodService", &FoodService{})
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go rpc.ServeConn(conn)

	}

}

```

```go
// Client
package main

import (
	"fmt"
	"net/rpc"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c, err := rpc.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	foods := []string{"beef", "pork", "bacon"}
	reply := make([]string, len(foods))

	for i := 0; i < len(foods); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err = c.Call("FoodService.SayName", foods[i], &reply[i])
			if err != nil {
				panic(err)
			}

			fmt.Println(reply[i])

		}(i)

	}

	wg.Wait()
}

```

RPC 跨语言 高并发 使用 JSON

```go
// Server
package main

import (
	"net"
	"net/rpc"
)

type FoodService struct {
}

func (f *FoodService) SayName(req string, res *string) error {
	*res = "ordered " + req
	return nil
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	err = rpc.RegisterName("FoodService", &FoodService{})
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

    //使用json解析
    go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	}

}

```

```go
// Client
package main

import (
	"net"
	"net/rpc"
)

func main() {
    c, err := net.Dial("tcp", "localhost:8888")
    if err != nil {
        fmt.Println(err)
        return
    }

    reply := ""
    client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))

    //其他语言使用json传递进行rpc-json类型 跨语言传输
    //{"method":"FoodService.SayName",params:["小龙虾"],id:1}
    err = client.Call("FoodService.SayName", "小龙虾", &reply)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(reply)
}
```

5. protobuffer 与 JSON 的优势

[reference](https://zhuanlan.zhihu.com/p/53339153)

压缩环境下 Protobuf 和 JSON 的结果非常相似。Protobuf 消息比 JSON 消息小 9％，减少了 4%的时间。
非压缩的环境下, 在与 JSON 相比，Protobuf 表现更好。 消息体积减少了 34％，快了 21％。
发送 POST 请求时, 差别不是很大, 因为 POST 请求一般不会处理大的 Message Body.

JSON 优势:

    1.数据格式比较简单, 易于读写, 格式都是压缩的, 占用带宽小

    2.易于解析这种语言, 客户端 JavaScript 可以简单的通过 eval()进行 JSON 数据的读取（浏览器解析）

    3.因为 JSON 格式能够直接为服务器端代码使用, 大大简化了服务器端和客户端的代码开发量, 但是完成的任务不变, 且易于维护。能够被大多数后端语言支持

Protobuffer 优势: 1. 二进制 效率高

    2. 代码生成机制

    3. 向前后兼容

    4. 支持多语言

6. grpc + probuffer 调用实例

```proto
syntax = "proto3";

option go_package="./;Hello_GRPC";

package Hello_GRPC;

message Req {
  string message = 1;
}

// TODO(rey): Describe this message.
message Res {
  string message = 1;

  // Next available id: 1
}

service HelloGRPC {
  rpc SayHi(Req) returns (Res);
}


```

```go
// Server
package main

import (
	"context"
	"fmt"
	"net"

	hello_grpc "github.com/hello_grpc/pb"
	"google.golang.org/grpc"
)

type Server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *Server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "from server, "}, nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &Server{})
	s.Serve(l)

}

```

```proto
// Client
package main

import (
	"context"
	"fmt"

	hello_grpc "github.com/hello_grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, e := grpc.Dial("localhost:8888", grpc.WithInsecure())
	defer conn.Close()
	fmt.Println(e)
	client := hello_grpc.NewHelloGRPCClient(conn)
	req, _ := client.SayHi(context.Background(), &hello_grpc.Req{Message: "from client"})
	fmt.Println(req.GetMessage())
}

```

结果:

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNly1gwshv0z6y7j31h4049wh2.jpg">

</div>

7. 三种流模式

```protobuf
syntax = "proto3"; // 使用proto3 编译

package person; // proto 方法的包

option go_package="github/learn-proto/pb/person;person"; // 完整路径 ; 别名

message PersonReq {
  string name = 1;
  int32 age = 2;
}

message PersonRes {
  string name = 1;
  int32 age = 2;
}

service SearchService {
  //  四种方法
  rpc Search(PersonReq) returns (PersonRes);
  rpc StreamSearch(stream PersonReq) returns (PersonRes);
  rpc SearchStream(PersonReq) returns (stream PersonRes);
  rpc StreamSearchStream(stream PersonReq) returns (stream PersonRes);
}

```

```go
// Server
package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/learn-method/pb/person"
	"google.golang.org/grpc"
)

type PersonServer struct {
	person.UnimplementedSearchServiceServer
}

func (p *PersonServer) Search(ctx context.Context, req *person.PersonReq) (res *person.PersonRes, err error) {

	name := req.GetName()
	res = &person.PersonRes{Name: "receive name: " + name}
	return res, nil
}

func (p *PersonServer) StreamSearch(server person.SearchService_StreamSearchServer) error {
	for {
		req, err := server.Recv()
		if err != nil {
			server.SendAndClose(&person.PersonRes{Name: "finished"})
			break
		}
		fmt.Println(req)
	}
	return nil
}

func (p *PersonServer) SearchStream(req *person.PersonReq, server person.SearchService_SearchStreamServer) error {
	name := req.Name
	i := 0
	for i < 5 {
		server.Send(&person.PersonRes{Name: fmt.Sprintf("%s%s", "server sent ", name)})
		time.Sleep(time.Second)
		i++
	}
	return nil
}

func (p *PersonServer) StreamSearchStream(server person.SearchService_StreamSearchStreamServer) error {

	msgChan := make(chan string)

	go func() {
		i := 0
		for i < 5 {
			req, err := server.Recv()
			if err != nil {
				fmt.Println(err)
				msgChan <- "finished"
			}

			msgChan <- req.Name
			i++
			time.Sleep(time.Second)

		}
	}()

	for {
		s := <-msgChan
		if s == "finished" {
			break
		}
		server.Send(&person.PersonRes{Name: s})
	}

	return nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &PersonServer{})
	s.Serve(l)

}

```

```go
// Client
package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/learn-method/pb/person"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())

	// normal method
	fmt.Println("=================Normal==============")
	client := person.NewSearchServiceClient(conn)
	res, err := client.Search(context.Background(), &person.PersonReq{Name: "rey"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// stream in
	fmt.Println("===================StreamIn=============")
	streamSearchClient, err := client.StreamSearch(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	i := 0
	for i < 5 {
		streamSearchClient.Send(&person.PersonReq{Name: fmt.Sprintf("%s%d", "stream in ", i)})
		time.Sleep(time.Second)
		i++
	}

	res, err = streamSearchClient.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	// stream out
	fmt.Println("===================StreamOut=============")
	searchStreamClient, err := client.SearchStream(context.Background(), &person.PersonReq{Name: "rey"})
	if err != nil {
		fmt.Println(err)
	}

	for {
		req, err := searchStreamClient.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(req)
	}
	// stream both
	fmt.Println("===================StreamBoth=============")
	StreamSearchStreamClient, err := client.StreamSearchStream(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			req, err := StreamSearchStreamClient.Recv()
			if err != nil {
				fmt.Println(err)
				wg.Done()
				break
			}
			fmt.Println(req)
		}
	}()

	go func() {
		i := 0
		for {
			err := StreamSearchStreamClient.Send(&person.PersonReq{Name: fmt.Sprintf("%s%d", "rey", i)})
			if err != nil {
				fmt.Println(err)
				wg.Done()
				break
			}
			i++
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}

```

运行结果:

<div align=center><img src="https://tvax3.sinaimg.cn/large/006cK6rNly1gwsfazoaqfj31hc0u0dz6.jpg">
</div>
