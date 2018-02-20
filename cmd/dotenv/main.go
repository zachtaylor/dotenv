package main

import (
	"fmt"
	"os"
	"strings"
	"ztaylor.me/env"
)

func main() {
	env.Bootstrap()
	for k, v := range env.Cache {
		fmt.Printf("%s=%s\n", k, v)
	}
	for _, line := range os.Args[1:] {
		args := strings.Split(strings.Trim(line, `-`), "=")
		if len(args) == 1 {
			fmt.Printf("%s=%s\n", args[0], env.Get(args[0]))
		}
	}
}
