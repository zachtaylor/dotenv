package env

import (
	"io/ioutil"
	"os"
	"strings"
)

// ParseEnv uses `os.Getenv` to update Service
func (s Service) ParseEnv() {
	for _, k := range s.Keys() {
		if v := os.Getenv(k); len(v) > 1 {
			s[k] = v
		}
	}
}

// ParseFile uses ParseLine to write file contents to Service
//
// Files can have comments (`#` style)
func (s Service) ParseFile(path string) error {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}

	for _, line := range strings.Split(string(file), "\n") {
		line = strings.Trim(strings.Split(line, "#")[0], " ;\r")
		if line != "" {
			s.ParseItem(line)
		}
	}

	return nil
}

// ParseItem uses "x=y" format to write values to Service, where "=y" is optional, and defaults to "=true"
func (s Service) ParseItem(item string) {
	if setting := strings.Split(item, "="); len(setting) == 1 {
		s.Set(setting[0], "true")
	} else if len(setting) == 2 {
		s.Set(setting[0], strings.Trim(setting[1], ` 	"`))
	}
}

// ParseFlags uses ParseArgs to write os.Args[1:] to Service
func (s Service) ParseFlags() {
	for _, arg := range os.Args[1:] {
		if len(arg) > 1 && arg[0] == '-' {
			s.ParseItem(arg[1:])
		}
	}
}
