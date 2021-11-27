package main

import (
	"fmt"
	"net/rpc"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c, err := rpc.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	foods := []string{"beef", "pork", "bacon"}
	reply := make([]string, len(foods))

	for i := 0; i < len(foods); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err = c.Call("FoodService.SayName", foods[i], &reply[i])
			if err != nil {
				panic(err)
			}

			fmt.Println(reply[i])

		}(i)

	}

	wg.Wait()
}
