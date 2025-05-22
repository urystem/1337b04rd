package config

type config struct {
	port uint

	db dbConfig

	s3 minio

	session sessionConfig
}

