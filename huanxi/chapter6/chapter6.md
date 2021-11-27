[TOC]

# Chapter6

## Message id 顺序重要！！！

```go
message PersonRes {
  string name = 1;
  int32 age = 2;
}

message PersonRes {
int32 age = 2;
string name = 1;
}
```

业务完成后, message 属性顺序不可随意调换, 可能影响业务的处理.

## import 外界的 proto 包

例如, 导入 "google/protobuf/empty.proto"

```proto
syntax = "proto3"; // 使用proto3 编译

package person; // proto 方法的包

option go_package="github/learn-proto/pb/person;person"; // 完整路径 ; 别名

import "google/protobuf/empty.proto";

message PersonReq {
  string name = 1;
  int32 age = 2;
}

message PersonRes {
  string name = 1;
  int32 age = 2;
}

service SearchService {
  rpc alive(google.protobuf.Empty) returns (PersonRes);

}

```

直接编译可能会导致错误 or 导入错误的包.

解决方法:

1. [protobuf_released](https://github.com/protocolbuffers/protobuf/releases/tag/v3.19.1) 下载源码

2. 将源码中 google 目录 copy 到指定目录

例如:

<div align=center><img src="https://tva3.sinaimg.cn/large/006cK6rNgy1gwu59m3gpij30do04k756.jpg">

</div>

3. 再次编译

## GRPC 中的数据传输 metadata

```go
// server
package main

import (
	"context"
	"fmt"
	"net"

	"github.com/learn-metadata/pb/person"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type PersonServer struct {
	person.UnimplementedSearchServiceServer
}

func (p *PersonServer) TestMeta(ctx context.Context, req *person.PersonReq) (res *person.PersonRes, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("no metadata there")
	}

	for k, v := range md {
		fmt.Printf("%s-%s\n", k, v)
	}

	return &person.PersonRes{Name: "finished"}, nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &PersonServer{})
	s.Serve(l)

}

```

```go
// client
package main

import (
	"context"

	"github.com/learn-metadata/pb/person"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())
	client := person.NewSearchServiceClient(conn)

	md := metadata.New(map[string]string{
		"rey":       "30",
		"charlotte": "22",
	},
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	client.TestMeta(ctx, &person.PersonReq{Name: "home"})
}
```

<div align=center><img src="https://tva1.sinaimg.cn/large/006cK6rNgy1gwu5c3tlrej30r40s9aip.jpg">

</div>

## GRPC 中的拦截器 interceptor

```go
// server
package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/learn-interception/pb/person"
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

func main() {
	l, _ := net.Listen("tcp", ":8888")

	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		elapsed := time.Since(start)
		fmt.Println(elapsed)
		return resp, nil
	}

	opt := grpc.UnaryInterceptor(interceptor)
	s := grpc.NewServer(opt)
	person.RegisterSearchServiceServer(s, &PersonServer{})
	s.Serve(l)

}


```

```go
//client
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/learn-interception/pb/person"
	"google.golang.org/grpc"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		elapsed := time.Since(start)
		fmt.Println(elapsed)
		return err

	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	conn, _ := grpc.Dial("localhost:8888", grpc.WithInsecure(), opt)
	client := person.NewSearchServiceClient(conn)
	pr, err := client.Search(context.Background(), &person.PersonReq{Name: "testInterceptor"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pr)

}

```

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNgy1gwu5ep8hj3j31810u8qcm.jpg">

</div>
