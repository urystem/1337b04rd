package config

import "time"

type sessionConfig struct {
	cookieName string
	duration   time.Duration
}

