// Package dotenv is an executable that prints all values in the global env
package main

import (
	"fmt"

	"ztaylor.me/env"
	"ztaylor.me/env/internal/service"
)

func main() {
	cache := service.Cache{}
	if err := env.ParseFile(cache, ".env"); err != nil {
		fmt.Println("open .env failed")
	}
	env.ParseFlags(cache)
	if len(cache) < 1 {
		fmt.Println("env is empty")
	}
	for k, v := range cache {
		fmt.Printf("%s=%s\n", k, v)
	}
}
