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
