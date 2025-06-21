package minio

import (
	"context"

	"1337b04rd/internal/ports/inbound"
	"1337b04rd/internal/ports/outbound"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minIo struct {
	client     *minio.Client
	postBuc    string
	commentBuc string
	// avatarBuc  string
}

func InitMinio(ctx context.Context, cfg inbound.MinioCfg) (outbound.MinioInter, error) {
	// // Инициализация клиента
	endpoint := cfg.GetEndpoint()
	accessKey := cfg.GetAccessKey()
	secretkey := cfg.GetSecretKey()
	secure := cfg.GetSecure()

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretkey, ""),
		Secure: secure,
	})
	if err != nil {
		return nil, err
	}

	postBucName := cfg.GetPostBucketName()
	err = createBucket(ctx, minioClient, postBucName)
	if err != nil {
		return nil, err
	}

	commentBucName := cfg.GetCommentBucketName()
	err = createBucket(ctx, minioClient, commentBucName)
	if err != nil {
		return nil, err
	}

	// avaBucName := cfg.GetAvatarBucketName()
	// err = createBucket(ctx, minioClient, avaBucName)
	// if err != nil {
	// 	return nil, err
	// }

	return &minIo{
		client:     minioClient,
		postBuc:    postBucName,
		commentBuc: commentBucName,
		// avatarBuc:  avaBucName,
	}, nil
}

func createBucket(ctx context.Context, mc *minio.Client, bucName string) error {
	err := mc.MakeBucket(ctx, bucName, minio.MakeBucketOptions{})
	if err == nil {
		return nil
	}

	exists, errExists := mc.BucketExists(ctx, bucName)
	if errExists != nil {
		return errExists
	} else if !exists {
		return err
	}
	return nil
}
