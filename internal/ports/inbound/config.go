package inbound

import "time"

type ServerCfg interface {
	GetPort() int
}

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

type MinioCfg interface {
	GetEndpoint() string
	GetAccessKey() string
	GetSecretKey() string
	GetRegion() string
	GetSecure() bool
	GetPostBucketName() string
	GetCommentBucketName() string
}

type Config interface {
	GetServerCfg() ServerCfg
	GetDBConfig() DBConfig
	GetSessionConfig() SessionConfig
	GetRedisConfig() RedisConfig
}
