package inbound

import "time"

type DBConfig interface {
	GetHostName() string
	GetPort() int
	GetUser() string
	GetPassword() string
	GetDBName() string
	GetSSLMode() string
}

type SessionConfig interface {
	GetCookieName() string
	GetDuration() time.Duration
}

type RedisConfig interface {
	GetAddr() string
	GetPass() string
}

type Config interface {
	GetPort() int
	GetDBConfig() DBConfig
	GetSessionConfig() SessionConfig
	GetRedisConfig() RedisConfig
}
