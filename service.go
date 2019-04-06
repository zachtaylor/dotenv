package env

import "ztaylor.me/env/internal/service"

// NewCacheService creates an empty basic Service
func NewCacheService() Service {
	return service.Cache{}
}

// NewDefaultService creates an empty basic Service
func NewDefaultService() Service {
	return &service.Default{
		Cache: service.Cache{},
	}
}
