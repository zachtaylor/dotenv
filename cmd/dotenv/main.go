// Package dotenv is an executable that prints all values in the global env
package main

import (
	"fmt"

	"ztaylor.me/env"
)

func main() {
	env := env.Global()
	if len(env) < 1 {
		fmt.Println("env is empty")
	}
	for k, v := range env {
		fmt.Printf(k + "=" + v + "\n")
	}
}
