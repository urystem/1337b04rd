package config

import (
	"1337b04rd/internal/core/ports/inbound"
)

type redisConfig struct {
	hostName string
	port     string
	pass     string
}

func (c *config) initRedisConf() inbound.RedisConfig {
	conf := &redisConfig{}
	conf.hostName = mustGetEnvString("REDIS_HOST")
	conf.port = mustGetEnvString("REDIS_PORT")
	conf.pass = mustGetEnvString("REDIS_PASS")
	return conf
}

func (r *redisConfig) GetHostName() string {
	return r.hostName
}

func (r *redisConfig) GetAddr() string {
	return r.port
}

func (r *redisConfig) GetPass() string {
	return r.pass
}
