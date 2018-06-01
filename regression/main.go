package main

import (
	"cmds"
	"fmt"
)

func main() {

	set := cmds.Set{
		"a": func(s string) error {
			if s == "a" {
				fmt.Println("get")
				return nil
			}
			return fmt.Errorf("wrong")
		},
		"b": func(s string) error {
			if s == "b" {
				fmt.Println("get")
				return nil
			}
			return fmt.Errorf("wrong")
		},
	}

	s := cmds.InitCMDS(set)
	err := s.Send("a", "a")
	fmt.Println(err)
	err = s.Send("a", "b")
	fmt.Println(err)
}
