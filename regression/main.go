package main

import (
	"cmds/server"
	"context"
	"fmt"
	"log"
	"time"

	server "cmds/proto"

	"google.golang.org/grpc"
)

func main() {

	// server side
	set := cmds.Set{
		1: func(s string) error {
			if s == "btc" {
				fmt.Println("mining")
				return nil
			}
			return fmt.Errorf("wrong")
		},
		2: func(s string) error {
			fmt.Println(s)
			return nil
		},
	}

	s := cmds.InitCMDS(set)
	addr := ":8080"
	go s.Run(addr)

	// wait server run
	time.Sleep(2 * time.Second)

	// client side
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	log.Print("dial success")

	client := server.NewCommandServiceClient(conn)
	r, err := client.Send(context.Background(), &server.Request{
		Code:  1,
		Param: "btc",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	r, err = client.Send(context.Background(), &server.Request{
		Code:  1,
		Param: "eth",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	//err := s.Send("a", "a")
	//fmt.Println(err)
	//err = s.Send("a", "b")
	//fmt.Println(err)
}
