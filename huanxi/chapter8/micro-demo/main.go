package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/rey/micro-demo/biz"
	_ "github.com/rey/micro-demo/dao/mysql"
	"github.com/rey/micro-demo/proto/account"
	"google.golang.org/grpc"
)

func main() {
	ip := flag.String("ip", "localhost", "specific ip")
	port := flag.Int("port", 9409, "specific port")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *ip, *port)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	account.RegisterAccountServiceServer(server, &biz.AccountServer{})
	if err = server.Serve(l); err != nil {
		panic(err)
	}

}
