package service

// Cache is a basic k/v map implementing Service
type Cache map[string]string

// Get returns the value by key
func (s Cache) Get(k string) string {
	return s[k]
}

// Default underwrites a value, returning the final value
func (s Cache) Default(k, v string) string {
	if val, ok := s[k]; ok {
		return val
	}
	s.Set(k, v)
	return v
}

// Set overwrites a value
func (s Cache) Set(k string, v string) {
	s[k] = v
}
