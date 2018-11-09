package env // import "ztaylor.me/env"

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Provider is environment controller
type Provider interface {
	// Get returns the value by key
	Get(string) string
	// Default underwrites a value, causing no-op if a value has been previously Set()
	Default(string, string) string
	// Set overwrites a value
	Set(string, string)
}

// global is used as cache by Global()
var global Provider

// Global returns the global Provider, and bootstraps if necessary
func Global() Provider {
	if global == nil {
		global = bootstrap()
	}
	return global
}

// bootstrap used to create and init new DefaultProvider
//
// uses ProviderSource() with local ".env" file
// uses ProviderFlags() with os.Args[1:]
func bootstrap() Provider {
	global := NewDefaultProvider()
	if err := ProviderSource(global, ".env"); err != nil {
		fmt.Println(err)
	}
	ProviderFlags(global, os.Args[1:])
	return global
}

// CacheProvider is a basic k/v map
type CacheProvider map[string]string

//Get returns the value by key
func (m CacheProvider) Get(k string) string {
	return m[k]
}

// Default underwrites a value, causing no-op if a value has been previously Set()
func (m CacheProvider) Default(k, v string) string {
	if w, ok := m[k]; ok {
		return w
	}
	m.Set(k, v)
	return v
}

// Set overwrites a value
func (m CacheProvider) Set(k string, v string) {
	m[k] = v
}

// DefaultProvider wraps CacheProvider, adds os.Getenv to Provider.Get
type DefaultProvider struct {
	CacheProvider
}

// NewDefaultProvider creates an empty DefaultProvider
func NewDefaultProvider() *DefaultProvider {
	return &DefaultProvider{CacheProvider{}}
}

// Get uses best effort to find a value for k
//
// Uses cache, else caches and returns value os.Getenv
func (m *DefaultProvider) Get(k string) string {
	if v, ok := m.CacheProvider[k]; ok {
		return v
	}
	osenv := os.Getenv(k)
	m.CacheProvider[k] = osenv
	return osenv
}

// ProviderFlags uses ProviderSourceLine to read args into Provider
//
// The dash char ('-') is optional, but the equals char ('=') is required
func ProviderFlags(p Provider, args []string) {
	for _, line := range args {
		ProviderSourceLine(p, strings.Trim(line, `-`))
	}
}

// ProviderSource reads a file with the given path into Provider
//
// Comments are removed, and empty lines are skipped
func ProviderSource(p Provider, path string) error {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}

	for _, line := range strings.Split(string(file), "\n") {
		if line = strings.Trim(strings.Split(line, "#")[0], " ;\r"); line == "" {
		} else {
			ProviderSourceLine(p, line)
		}
	}

	return nil
}

// ProviderSourceLine imports a setting with "x=y" format into Provider
func ProviderSourceLine(p Provider, line string) {
	if setting := strings.Split(line, "="); len(setting) == 1 {
		p.Set(setting[0], "true")
	} else if len(setting) == 2 {
		p.Set(setting[0], strings.Trim(setting[1], ` '";`))
	}
}
