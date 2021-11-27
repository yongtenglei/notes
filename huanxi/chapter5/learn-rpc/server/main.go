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
