// Package env provides runtime environment
package env // import "ztaylor.me/env"

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Service is environment controller
type Service interface {
	// Get returns the value by key
	Get(string) string
	// Default underwrites a value, returning the final value
	Default(string, string) string
	// Set overwrites a value
	Set(string, string)
}

// global is used as cache by Global()
var global Service

// Global returns the global Service, and bootstraps if necessary
func Global() Service {
	if global == nil {
		global = NewDefaultService()
		if err := ParseFile(global, ".env"); err != nil {
			fmt.Println(err)
		}
		ParseFlags(global)
	}
	return global
}

// ParseFile uses ParseLine to write file contents to Service
//
// SQL inline comments are trimmed ("--comment"), and empty lines are skipped
func ParseFile(s Service, path string) error {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}

	for _, line := range strings.Split(string(file), "\n") {
		if line = strings.Trim(strings.Split(line, "--")[0], " ;\r"); line == "" {
		} else {
			ParseLine(s, line)
		}
	}

	return nil
}

// ParseFlags uses ParseLine to write os.Args to Service
func ParseFlags(s Service) {
	for _, line := range os.Args[1:] {
		ParseLine(s, strings.Trim(line, `-`))
	}
}

// ParseLine uses "x=y" format to write values to Service
//
// "=y" is optional; if missing, parsed as "=true"
func ParseLine(s Service, line string) {
	if setting := strings.Split(line, "="); len(setting) == 1 {
		s.Set(setting[0], "true")
	} else if len(setting) == 2 {
		s.Set(setting[0], strings.Trim(setting[1], ` '";`))
	}
}
