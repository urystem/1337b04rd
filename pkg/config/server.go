package config

type serverCfg struct {
	port int
}

func (c *config) initServerCfg() any {
	cfg := &serverCfg{}
	cfg.port = mustGetEnvInt("")
	return &cfg
}
