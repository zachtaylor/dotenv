package env

import (
	"io/ioutil"
	"os"
	"strings"
)

// Cache holds all prereferenced
var Cache = make(map[string]string)

// Name is a macro for Get("ENV")
func Name() string {
	return Get("ENV")
}

// Get uses best effort to find a value for k
//
// If any value has been written to Cache, use that value
// Else, use os.Getenv and save this to Cache
func Get(k string) string {
	if v, ok := Cache[k]; ok {
		return v
	}
	osenv := os.Getenv(k)
	Cache[k] = osenv
	return osenv
}

// Set overwrites a value in the Cache
func Set(k, v string) {
	Cache[k] = v
}

// SourceLine imports a setting with "x=y" format
func SourceLine(line string) {
	if setting := strings.Split(line, "="); len(setting) != 2 {
	} else {
		Set(setting[0], strings.Trim(setting[1], ` '"`))
	}
}

// Default underwrites a value in the Cache
//
// If k is already set, this operation does nothing
func Default(k, v string) {
	if _, ok := Cache[k]; !ok {
		Set(k, v)
	}
}

// Bootstrap sources the ".env" file and parses cli flags
func Bootstrap() {
	Source(".env")
	Flags()
}

// Flags reads os.Args to source env values in Cache
//
// The dash char ('-') is optional, but the equals char ('=') is required
func Flags() {
	for _, line := range os.Args[1:] {
		SourceLine(strings.Trim(line, `-`))
	}
}

// Source reads a file with the given path
//
// Comments are removed, and empty lines are skipped
func Source(path string) error {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}

	for _, line := range strings.Split(string(file), "\n") {
		if line = strings.Trim(strings.Split(line, "#")[0], " ;"); line == "" {
		} else {
			SourceLine(line)
		}
	}

	return nil
}
