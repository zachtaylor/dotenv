package main

import (
	"fmt"

	"ztaylor.me/env"
)

func main() {
	p := env.DefaultProvider{}
	env.ProviderSource(p, ".env")
	for k, v := range p {
		fmt.Printf("%s=%s\n", k, v)
	}
}
