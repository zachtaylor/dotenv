// Package env provides runtime environment
package env // import "ztaylor.me/env"

import "fmt"

// Default returns a new Service, loaded with `ParseFile(".env")`, `ParseEnv()`, and `ParseAllFlags()`
func Default() Service {
	service := NewService()
	if err := service.ParseFile(".env"); err != nil {
		fmt.Println("env: " + err.Error())
	}
	service.ParseEnv()
	service.ParseAllFlags()
	return service
}

// File returns a new Service with `ParseFile(path)` loaded
func File(path string) Service {
	service := NewService()
	if err := service.ParseFile(path); err != nil {
		fmt.Println("env: " + err.Error())
	}
	return service
}

// Flags returns a new Service with `ParseAllFlags()` loaded
func Flags() Service {
	service := NewService()
	service.ParseAllFlags()
	return service
}
