package main

import (
	"cmds/client"
	"cmds/server"
	"fmt"
	"time"
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
	cli, err := client.InitClient(addr)
	if err != nil {
		panic(err)
	}
	err = cli.Send(1, "btc")
	if err != nil {
		panic(err)
	}
	fmt.Println("success")

	err = cli.Send(1, "eth")
	fmt.Println(err)
}
