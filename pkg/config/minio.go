package config

import "os"

type minio struct {
	endpoint       string
	accessKey      string
	secretKey      string
	region         string
	secure         bool
	bucketPosts    string
	bucketComments string
}

func (c *config) initMinio() minio {
	myMinio := minio{}
	myMinio.endpoint = mustGetEnvString("S3_ENDPOINT")
	myMinio.accessKey = mustGetEnvString("S3_ROOT_USER")
	myMinio.secretKey = mustGetEnvString("S3_ROOT_PASSWORD")
	myMinio.region = os.Getenv("S3_REGION")
	myMinio.secure = mustGetBoolean("S3_SECURE")
	myMinio.bucketPosts = mustGetEnvString("S3_BUCKETNAME_POST")
	myMinio.bucketComments = mustGetEnvString("S3_BUCKETNAME_COMMENT")
	return myMinio
}

func (m *minio) GetEndpoint() string {
	return m.endpoint
}

func (m *minio) GetAccessKey() string {
	return m.accessKey
}

func (m *minio) GetSecretKey() string {
	return m.secretKey
}

func (m *minio) GetRegion() string {
	return m.region
}

func (m *minio) GetSecure() bool {
	return m.secure
}

func (m *minio) GetPostBucketName() string {
	return m.bucketPosts
}

func (m *minio) GetCommentBucketName() string {
	return m.bucketComments
}
