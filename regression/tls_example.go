package main

import (
	"cmds/client"
	"cmds/server"
	"flag"
	"fmt"
	"time"
)

var (
	certFile   = flag.String("cert", "", "")
	keyFile    = flag.String("key", "", "")
	serverName = flag.String("server_name", "", "")
)

func main() {
	flag.Parse()

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
	credOpt, err := cmds.CreateCred(*certFile, *keyFile)
	if err != nil {
		panic(err)
	}
	go s.Run(addr, credOpt)

	// wait server run
	time.Sleep(2 * time.Second)

	// client side
	cred, _ := client.CreateCred(*certFile, *serverName)
	cli, err := client.InitClient(addr, cred)
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
