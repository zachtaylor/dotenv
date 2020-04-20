// Package env provides runtime environment
package env // import "ztaylor.me/env"

import "fmt"

// Global is used as cache by Global()
var global Service

// Default returns the global Service, and bootstraps if necessary
func Default() Service {
	if global == nil {
		global = NewService()
		if err := global.ParseFile(".env"); err != nil {
			fmt.Println("env: " + err.Error())
		}
		global.ParseFlags()
	}
	return global
}
