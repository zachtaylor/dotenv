package main

import (
	"fmt"
	"ztaylor.me/env"
)

func main() {
	env.Bootstrap()
	for k, v := range env.Cache {
		fmt.Printf("%s=%s\n", k, v)
	}
}
