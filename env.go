package env // import "ztaylor.me/env"

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func init() {
	bootstrap()
}

// Cache holds the env values
var Cache = make(map[string]string)

// Name is a macro for Get("ENV")
func Name() string {
	return Get("ENV")
}

// Get uses best effort to find a value for k
//
// Uses Cache, else, use os.Getenv and save to Cache
func Get(k string) string {
	if v, ok := Cache[k]; ok {
		return v
	}
	osenv := os.Getenv(k)
	Cache[k] = osenv
	return osenv
}

// GetI uses Get with int type casting
func GetI(k string) int {
	i, _ := strconv.ParseInt(Get(k), 0, 64)
	return int(i)
}

// GetI uses Get with bool type casting
func GetB(k string) bool {
	b, _ := strconv.ParseBool(Get(k))
	return b
}

// Set overwrites a value in the Cache
func Set(k, v string) {
	Cache[k] = v
}

// Default underwrites a value in the Cache
//
// If k is already set in Cache, this operation does nothing
func Default(k, v string) string {
	if w, ok := Cache[k]; ok {
		return w
	} else {
		Set(k, v)
		return v
	}
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
		if line = strings.Trim(strings.Split(line, "#")[0], " ;\r"); line == "" {
		} else {
			SourceLine(line)
		}
	}

	return nil
}

// SourceLine imports a setting with "x=y" format
func SourceLine(line string) {
	if setting := strings.Split(line, "="); len(setting) == 1 {
		Set(setting[0], "true")
	} else if len(setting) == 2 {
		Set(setting[0], strings.Trim(setting[1], ` '";`))
	}
}

func bootstrap() {
	if err := Source(".env"); err != nil {
		fmt.Println(err)
	}
	Flags()
}
