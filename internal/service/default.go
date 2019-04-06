package service

import "os"

type Default struct {
	Cache
}

func (s *Default) Get(k string) string {
	if _, ok := s.Cache[k]; !ok {
		s.Cache[k] = os.Getenv(k)
	}
	return s.Cache[k]
}

func (s *Default) Default(k, v string) string {
	if val := s.Get(k); val != "" {
		return val
	}
	s.Set(k, v)
	return v
}

func (s *Default) Set(k string, v string) {
	s.Cache[k] = v
}
