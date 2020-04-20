package env

// Service is a basic k/v map
type Service map[string]string

// NewService creates an empty basic Service
func NewService() Service {
	return Service{}
}

// Keys returns a new slice of all named Service settings
func (s Service) Keys() (keys []string) {
	keys = make([]string, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return
}

// Match returns a new Service containing values of settings prefixes matched, with the prefix removed
func (s Service) Match(pre string) Service {
	service, lpre := Service{}, len(pre)
	for _, k := range s.Keys() {
		if len(k) > lpre && k[:lpre] == pre {
			service.Set(k[lpre:], s.Get(k))
		}
	}
	return service
}

// Get returns the value by key
func (s Service) Get(k string) string {
	return s[k]
}

// Set overwrites a value
func (s Service) Set(k string, v string) {
	s[k] = v
}
