package main

import (
	"fmt"
	"ztaylor.me/env"
)

func main() {
	for k, v := range env.Cache {
		fmt.Printf("%s=%s\n", k, v)
	}
}
