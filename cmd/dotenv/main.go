package main

import (
	"fmt"
	"ztaylor.me/env"
)

func main() {
	if len(env.Cache) < 1 {
		fmt.Println("[env empty]")
	}
	for k, v := range env.Cache {
		fmt.Printf("%s=%s\n", k, v)
	}
}
