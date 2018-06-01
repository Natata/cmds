package main

import (
	"cmds/server"
	"context"
	"fmt"
	"log"

	server "cmds/proto"

	"google.golang.org/grpc"
)

func main() {

	// server side
	set := cmds.Set{
		1: func(s string) error {
			if s == "a" {
				fmt.Println("get")
				return nil
			}
			return fmt.Errorf("wrong")
		},
		2: func(s string) error {
			if s == "b" {
				fmt.Println("get")
				return nil
			}
			return fmt.Errorf("wrong")
		},
	}

	s := cmds.InitCMDS(set)
	addr := ":8080"
	go s.Run(addr)

	// client side
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	log.Print("dial success")

	client := server.NewCommandServiceClient(conn)
	r, err := client.Send(context.Background(), &server.Request{
		Code:  1,
		Param: "a",
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
