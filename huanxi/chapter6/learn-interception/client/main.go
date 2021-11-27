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
