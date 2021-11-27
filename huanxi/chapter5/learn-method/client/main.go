package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/learn-method/pb/person"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())

	// normal method
	fmt.Println("=================Normal==============")
	client := person.NewSearchServiceClient(conn)
	res, err := client.Search(context.Background(), &person.PersonReq{Name: "rey"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// stream in
	fmt.Println("===================StreamIn=============")
	streamSearchClient, err := client.StreamSearch(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	i := 0
	for i < 5 {
		streamSearchClient.Send(&person.PersonReq{Name: fmt.Sprintf("%s%d", "stream in ", i)})
		time.Sleep(time.Second)
		i++
	}

	res, err = streamSearchClient.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	// stream out
	fmt.Println("===================StreamOut=============")
	searchStreamClient, err := client.SearchStream(context.Background(), &person.PersonReq{Name: "rey"})
	if err != nil {
		fmt.Println(err)
	}

	for {
		req, err := searchStreamClient.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(req)
	}
	// stream both
	fmt.Println("===================StreamBoth=============")
	StreamSearchStreamClient, err := client.StreamSearchStream(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			req, err := StreamSearchStreamClient.Recv()
			if err != nil {
				fmt.Println(err)
				wg.Done()
				break
			}
			fmt.Println(req)
		}
	}()

	go func() {
		i := 0
		for {
			err := StreamSearchStreamClient.Send(&person.PersonReq{Name: fmt.Sprintf("%s%d", "rey", i)})
			if err != nil {
				fmt.Println(err)
				wg.Done()
				break
			}
			i++
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
