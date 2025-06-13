package minio

type minio struct {
	postBuc    string
	commentBuc string
}

func s() any {
	// // Инициализация клиента
	// minioClient, err := minio.New(endpoint, &minio.Options{
	// 	Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	// 	Secure: useSSL,
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// ctx := context.Background()
	// bucketName := "my-bucket"

	return &minio{}
}
