package config

import (
	"time"

	"1337b04rd/internal/ports/inbound"
)

type sessionConfig struct {
	cookieName string
	duration   time.Duration
}

func (c *config) initSessionConf() inbound.SessionConfig {
	conf := &sessionConfig{}
	conf.cookieName = mustGetEnvString("")
	conf.duration = time.Duration(mustGetEnvInt(""))
	return conf
}

func (sc *sessionConfig) GetCookieName() string {
	return sc.cookieName
}

func (sc *sessionConfig) GetDuration() time.Duration {
	return sc.duration
}
