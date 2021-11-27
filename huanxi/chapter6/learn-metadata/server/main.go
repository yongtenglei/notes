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
