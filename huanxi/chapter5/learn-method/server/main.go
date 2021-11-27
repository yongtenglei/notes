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
