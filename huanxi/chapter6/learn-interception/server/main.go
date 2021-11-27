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
