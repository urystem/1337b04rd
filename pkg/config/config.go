package config

import "1337b04rd/internal/ports/inbound"

type config struct {
	port int

	db inbound.DBConfig

	session inbound.SessionConfig

	redis inbound.RedisConfig

	s3 minio
}

func InitConfig() inbound.Config {
	conf := &config{}
	// conf.db = conf.initDBConfig()
	// conf.session = conf.initSessionConf()
	conf.redis = conf.initRedisConf()
	return conf
}

func (conf *config) GetPort() int {
	return conf.port
}

func (conf *config) GetDBConfig() inbound.DBConfig {
	return conf.db
}

func (conf *config) GetSessionConfig() inbound.SessionConfig {
	return conf.session
}

func (conf *config) GetRedisConfig() inbound.RedisConfig {
	return conf.redis
}
