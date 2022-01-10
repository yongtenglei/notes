package main

import (
	"context"
	"fmt"

	"github.com/rey/micro-demo/internal"
	"github.com/rey/micro-demo/proto/account"
	"google.golang.org/grpc"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
)

func main() {
	addr := fmt.Sprintf("%s:%d", internal.ConsulHost, internal.ConsulPort)
	//dialAddr := fmt.Sprintf("consul://127.0.0.1:8500/whoami?wait=14s&tag=manual")
	dialAddr := fmt.Sprintf("consul://%s/account_server?wait=14s", addr)

	conn, err := grpc.Dial(dialAddr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := account.NewAccountServiceClient(conn)
	for i := 0; i < 10; i++ {
		r, err := client.GetAccountList(context.Background(), &account.PagingReq{
			PageNo:   1,
			PageSize: 2,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(r)

	}
}
