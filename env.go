package env

import (
	"io/ioutil"
	"os"
	"strings"
)

var Cache = make(map[string]string)

type Option interface {
	String() string
}

func Get(k string) string {
	if v, ok := Cache[k]; ok {
		return v
	}
	osenv := os.Getenv(k)
	Cache[k] = osenv
	return osenv
}

func Set(k, v string) {
	Cache[k] = v
}

func SourceLine(line string) {
	if setting := strings.Split(line, "="); len(setting) != 2 {
	} else {
		Set(setting[0], strings.Trim(setting[1], ` '"`))
	}
}

func Default(k, v string) {
	if _, ok := Cache[k]; !ok {
		Set(k, v)
	}
}

func Bootstrap() {
	Source(".env")
	Flags()
}

func Flags() {
	for _, line := range os.Args[1:] {
		SourceLine(strings.Trim(line, `-`))
	}
}

func Source(name string) error {
	file, e := ioutil.ReadFile(name)
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
