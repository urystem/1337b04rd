package config

import "1337b04rd/internal/ports/inbound"

type serverCfg struct {
	port int
}

func (c *config) initServerCfg() inbound.ServerCfg {
	cfg := &serverCfg{}
	cfg.port = mustGetEnvInt("SERVER_PORT")
	return cfg
}

func (scfg *serverCfg) GetPort() int {
	return scfg.port
}
