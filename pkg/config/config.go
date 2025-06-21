package config

import "1337b04rd/internal/ports/inbound"

type config struct {
	server serverCfg

	db dbConfig

	session sessionConfig

	redis redisConfig

	s3 minio
}

func Load() inbound.Config {
	conf := &config{}
	conf.db = conf.initDBConfig()
	conf.session = conf.initSessionConf()
	conf.server = conf.initServerCfg()
	conf.redis = conf.initRedisConf()
	conf.s3 = conf.initMinio()
	return conf
}

func (conf *config) GetServerCfg() inbound.ServerCfg {
	return &conf.server
}

func (conf *config) GetDBConfig() inbound.DBConfig {
	return &conf.db
}

func (conf *config) GetSessionConfig() inbound.SessionConfig {
	return &conf.session
}

func (conf *config) GetRedisConfig() inbound.RedisConfig {
	return &conf.redis
}

func (cfg *config) GetMinioCfg() inbound.MinioCfg {
	return &cfg.s3
}

func (cfg *config) GetMinIoConfig() inbound.MinioCfg {
	return &cfg.s3
}
