# Chapter5 RPC

## RPC 定义

RPC（Remote Procedure Call Protocol）——远程过程调用协议，它是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络技术的协议。RPC 协议假定某些传输协议的存在，如 TCP 或 UDP，为通信程序之间携带信息数据。在 OSI 网络通信模型中，RPC 跨越了传输层和应用层。RPC 使得开发包括网络分布式多程序在内的应用程序更加容易。

RPC 采用客户机/服务器模式。请求程序就是一个客户机，而服务提供程序就是一个服务器。首先，客户机调用进程发送一个有进程参数的调用信息到服务进程，然后等待应答信息。在服务器端，进程保持睡眠状态直到调用信息到达为止。当一个调用信息到达，服务器获得进程参数，计算结果，发送答复信息，然后等待下一个调用信息，最后，客户端调用进程接收答复信息，获得进程结果，然后调用执行继续进行。

### 为什么要用 RPC

RPC 协议的主要目的是做到不同服务间调用方法像同一服务间调用本地方法一样

1：服务化/微服务

2：分布式系统架构

3：服务可重用

4：系统间交互调用

### XML JSON Proto

<div align=center><img src="http://81.68.236.200:8181/uploads/lee/images/m_b333f0b235126e8e136bab0c5f0176b7_r.png">

</div>

### RPC 通信过程

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

## Go 语言原生 RPC

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

### RPC 跨语言

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

## protobuffer

Proto buffer 是 Google 的语言中立、平台中立、可扩展的机制，用于序列化结构化数据, 像 XML 一样，但更小、更快、更简单。你只需要定义一次你想要的数据结构，然后你就可以使用特殊生成的源代码来轻松地从各种数据流和各种语言中写入和读取你的结构化数据。

### protobuffer 语法

```go
syntax = "proto3"; // 使用proto3 编译

package person; // proto 方法的包

option go_package="github/learn-proto/pb/person;person"; // 完整路径 ; 别名

message Person {
  string name = 1;
  int32 age = 2;
  enum Gender{ // 枚举 必须有0值
    option allow_alias = true; // 允许枚举有相同值
    MAN = 0;
    MALE = 0;
    WOMAN = 1;
    FEMALE = 1;
    OTHER = 2;
  }
  Gender gender = 3;
  repeated string testSlice = 4; // 切片
  map <string, string> testMap = 5; // map <key:int/string, value: any type>

  //reserved "testSlice", "testMap"; // 保留字
  //reserved 4, 5;

}

// message 嵌套
message Home {
  repeated Person persons = 1;

  message Visitor {
    string name = 1;
  }

  Visitor visitor = 2;

}


service SearchPersonService {
  //  四种方法
  rpc Search(Person) returns (Person);
  rpc StreamSearch(stream Person) returns (Person);
  rpc SearchStream(Person) returns (stream Person);
  rpc StreamSearchStream(stream Person) returns (stream Person);
}

```

### 类型转换

<table width="50%" border="1">

<tbody>

<tr>

<th>.proto Type</th>

<th>Notes</th>

<th>C++ Type</th>

<th>Java Type</th>

<th>Python Type<sup>[2]</sup></th>

<th>Go Type</th>

</tr>

<tr>

<td>double</td>

<td></td>

<td>double</td>

<td>double</td>

<td>float</td>

<td>*float64</td>

</tr>

<tr>

<td>float</td>

<td></td>

<td>float</td>

<td>float</td>

<td>float</td>

<td>*float32</td>

</tr>

<tr>

<td>int32</td>

<td>Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead.</td>

<td>int32</td>

<td>int</td>

<td>int</td>

<td>*int32</td>

</tr>

<tr>

<td>int64</td>

<td>Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead.</td>

<td>int64</td>

<td>long</td>

<td>int/long<sup>[3]</sup></td>

<td>*int64</td>

</tr>

<tr>

<td>uint32</td>

<td>Uses variable-length encoding.</td>

<td>uint32</td>

<td>int<sup>[1]</sup></td>

<td>int/long<sup>[3]</sup></td>

<td>*uint32</td>

</tr>

<tr>

<td>uint64</td>

<td>Uses variable-length encoding.</td>

<td>uint64</td>

<td>long<sup>[1]</sup></td>

<td>int/long<sup>[3]</sup></td>

<td>*uint64</td>

</tr>

<tr>

<td>sint32</td>

<td>Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.</td>

<td>int32</td>

<td>int</td>

<td>int</td>

<td>*int32</td>

</tr>

<tr>

<td>sint64</td>

<td>Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.</td>

<td>int64</td>

<td>long</td>

<td>int/long<sup>[3]</sup></td>

<td>*int64</td>

</tr>

<tr>

<td>fixed32</td>

<td>Always four bytes. More efficient than uint32 if values are often greater than 2<sup>28</sup>.</td>

<td>uint32</td>

<td>int<sup>[1]</sup></td>

<td>int/long<sup>[3]</sup></td>

<td>*uint32</td>

</tr>

<tr>

<td>fixed64</td>

<td>Always eight bytes. More efficient than uint64 if values are often greater than 2<sup>56</sup>.</td>

<td>uint64</td>

<td>long<sup>[1]</sup></td>

<td>int/long<sup>[3]</sup></td>

<td>*uint64</td>

</tr>

<tr>

<td>sfixed32</td>

<td>Always four bytes.</td>

<td>int32</td>

<td>int</td>

<td>int</td>

<td>*int32</td>

</tr>

<tr>

<td>sfixed64</td>

<td>Always eight bytes.</td>

<td>int64</td>

<td>long</td>

<td>int/long<sup>[3]</sup></td>

<td>*int64</td>

</tr>

<tr>

<td>bool</td>

<td></td>

<td>bool</td>

<td>boolean</td>

<td>bool</td>

<td>*bool</td>

</tr>

<tr>

<td>string</td>

<td>A string must always contain UTF-8 encoded or 7-bit ASCII text.</td>

<td>string</td>

<td>String</td>

<td>unicode (Python 2) or str (Python 3)</td>

<td>*string</td>

</tr>

<tr>

<td>bytes</td>

<td>May contain any arbitrary sequence of bytes.</td>

<td>string</td>

<td>ByteString</td>

<td>bytes</td>

<td>[]byte</td>

</tr>

</tbody>

</table>

### 编译 protobuffer

```makefile
all:
		protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative DESTINATION

```

### 安装

[QuikeStart](https://grpc.io/docs/languages/go/quickstart/):

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

### 三种流模式

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
