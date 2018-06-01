package main

import (
	"cmds/server"
	"context"
	"fmt"
	"log"
	"net"

	server "cmds/proto"

	"google.golang.org/grpc"
)

func main() {
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
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen at: %v", addr)

	gserv := grpc.NewServer()
	server.RegisterCommandServiceServer(gserv, s)
	go gserv.Serve(lis)

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
