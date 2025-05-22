package config

type dbConfig struct {
	host     string
	port     uint
	user     string
	password string
	name     string
	sslMode  string
}

func initDBConfig() (any, error) {
	dbConf := &dbConfig{}
	var err error
	dbConf.host, err = mustGetEnvString("")
	
	return &dbConfig{}, err
}
