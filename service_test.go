package env_test

import (
	"testing"

	"ztaylor.me/env"
)

func TestServiceMatch(t *testing.T) {
	// Service default values
	s := env.Service{
		"a":   "1",
		"b":   "2",
		"c_a": "3",
		"c_b": "4",
	}

	// isolate subset with match
	sc := s.Match("c_")

	if sc.Get("a") != "3" {
		t.Fail()
	}
	if sc.Get("b") != "4" {
		t.Fail()
	}
}
