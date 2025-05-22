package config

type minio struct {
	endpoint       string
	accessKey      string
	secretKey      string
	bucketComments string
	bucketThreads  string
	region         string
	useSSL         bool
}

