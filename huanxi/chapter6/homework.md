# Chapter 6 Homework

1. protobuf 数据类型有哪些，和 go 语言如何对应的

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

注意:

protobuf 的类型, 都对应 GO 语言的指针 or 应用类型(切片)

有负数时, 使用 sint32, sinint64

2. 说说 go_package

`option go_package="github/learn-proto/pb/person;person"; // 完整路径 ; 别名`

`;` 前指定编译后的 proto 文件的位置, 后面为别名. 一般与 package 相对应. 如 `package person`

3. 举例说明如何使用 protobuf 的 map,timestamp 和枚举

```go
syntax = "proto3";

import "google/protobuf/timestamp.proto";

option go_package="github.com/rey/444/pb/testdt;testdt";

package testdt;

enum e {
  e1 = 0;
  e2 = 1;
};

message Req {
  map<string, int32> m = 1;
  e t = 2;
  google.protobuf.Timestamp RequestTime = 3;
}

message Res {
  map<string, int32> m = 1;
  e t = 2;
  google.protobuf.Timestamp RequestTime = 3;

}

service HelloGRPC {
  rpc TestDT(Req) returns (Res);
}

```

```go
// server
package main

import (
	"context"
	"fmt"
	"net"

	"github.com/rey/testdt/pb/testdt"
	"google.golang.org/grpc"
)

type TestServer struct {
	testdt.UnimplementedHelloGRPCServer
}

func (t TestServer) TestDT(ctx context.Context, req *testdt.Req) (res *testdt.Res, err error) {
	fmt.Println(req.M)
	fmt.Println(req.RequestTime.AsTime().Local().Format("2006-01-02 15:04:05"))
	fmt.Println(req.T)

	return &testdt.Res{}, nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	server := grpc.NewServer()
	testdt.RegisterHelloGRPCServer(server, &TestServer{})
	server.Serve(l)
}

```

```go
// client
package main

import (
	"context"
	"fmt"

	"github.com/rey/testdt/pb/testdt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())
	client := testdt.NewHelloGRPCClient(conn)
	res, err := client.TestDT(context.Background(), &testdt.Req{
		M: map[string]int32{
			"rey": 30,
		},
		RequestTime: timestamppb.Now(),
		T:           testdt.E_e1,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

```

测试:

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNgy1gwv7pi2szlj31h60jdamo.jpg">

</div>

4. grpc 的拦截器是如何使用的？

```go
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

  rpc TestMeta(PersonReq) returns (PersonRes);
}

```

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
// client
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

测试:

<div align=center><img src="https://tvax2.sinaimg.cn/large/006cK6rNgy1gwv7uavbknj31h00inapz.jpg">

</div>

5. 举例说明 message 的嵌套和 import 的用法

```go
syntax = "proto3";

import "google/protobuf/timestamp.proto";

option go_package="github.com/rey/444/pb/testdt;testdt";

package testdt;

enum e {
  e1 = 0;
  e2 = 1;
};

message Req {
  map<string, int32> m = 1;
  e t = 2;
  google.protobuf.Timestamp RequestTime = 3;
  message embedded {
    string embededitem = 1;
  }
  embedded emb = 4;
}

message Res {
  map<string, int32> m = 1;
  e t = 2;
  google.protobuf.Timestamp RequestTime = 3;

}

service HelloGRPC {
  rpc TestDT(Req) returns (Res);
}

```

```go
package main

import (
	"context"
	"fmt"
	"net"

	"github.com/rey/testdt/pb/testdt"
	"google.golang.org/grpc"
)

type TestServer struct {
	testdt.UnimplementedHelloGRPCServer
}

func (t TestServer) TestDT(ctx context.Context, req *testdt.Req) (res *testdt.Res, err error) {
	fmt.Println(req.M)
	fmt.Println(req.RequestTime.AsTime().Local().Format("2006-01-02 15:04:05"))
	fmt.Println(req.T)
	fmt.Println(req.Emb.Embededitem)

	return &testdt.Res{}, nil
}

func main() {
	l, _ := net.Listen("tcp", ":8889")
	server := grpc.NewServer()
	testdt.RegisterHelloGRPCServer(server, &TestServer{})
	server.Serve(l)
}

```

```go
// client
package main

import (
	"context"
	"fmt"

	"github.com/rey/testdt/pb/testdt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, _ := grpc.Dial("localhost:8889", grpc.WithInsecure())
	client := testdt.NewHelloGRPCClient(conn)
	res, err := client.TestDT(context.Background(), &testdt.Req{
		M: map[string]int32{
			"rey": 30,
		},
		RequestTime: timestamppb.Now(),
		T:           testdt.E_e1,
		Emb:         &testdt.ReqEmbedded{Embededitem: "embededitem"},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

```

测试:

<div align=center><img src="https://tva3.sinaimg.cn/large/006cK6rNgy1gwv822qxr9j31a40k57gz.jpg">

</div>

引入别人的 proto 包

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
