package config

import "time"

type sessionConfig struct {
	cookieName string
	duration   time.Duration
}

func (c *config) initSessionConf() sessionConfig {
	conf := sessionConfig{}
	conf.cookieName = mustGetEnvString("SESSION_NAME")
	conf.duration = time.Duration(mustGetEnvInt("SESSION_DURATION")) * 24 * time.Hour
	return conf
}

func (sc *sessionConfig) GetCookieName() string {
	return sc.cookieName
}

func (sc *sessionConfig) GetDuration() time.Duration {
	return sc.duration
}
