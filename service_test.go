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

	if len(sc) != 2 {
		t.Fail()
	}
	if sc["a"] != "3" {
		t.Fail()
	}
	if sc["b"] != "4" {
		t.Fail()
	}
}

func TestServiceMerge(t *testing.T) {
	// Service default values
	s := env.Service{
		"a": "1",
		"b": "2",
	}

	// merge it up with itself
	s = s.Merge("c_", s)

	if s["a"] != "1" {
		t.Fail()
	}
	if s["c_a"] != "1" {
		t.Fail()
	}
	if s["b"] != "2" {
		t.Fail()
	}
	if s["c_b"] != "2" {
		t.Fail()
	}
}
