package minio

import (
	"context"

	"1337b04rd/internal/domain"

	"github.com/minio/minio-go/v7"
)

func (m *minIo) PutPost(ctx context.Context, obj *domain.InPutObject) error {
	return m.putObject(ctx, m.postBuc, obj)
}

func (m *minIo) DelPost(ctx context.Context, objName string) error {
	return m.client.RemoveObject(ctx, m.postBuc, objName, minio.RemoveObjectOptions{})
}

func (m *minIo) GetPost(ctx context.Context, objName string) (*domain.OutputObject, error) {
	return m.getObjedctReadSeek(ctx, m.postBuc, objName)
}

func (m *minIo) PutComment(ctx context.Context, obj *domain.InPutObject) error {
	return m.putObject(ctx, m.commentBuc, obj)
}

func (m *minIo) DelComment(ctx context.Context, objName string) error {
	return m.client.RemoveObject(ctx, m.commentBuc, objName, minio.RemoveObjectOptions{})
}

func (m *minIo) GetComment(ctx context.Context, objName string) (*domain.OutputObject, error) {
	return m.getObjedctReadSeek(ctx, m.commentBuc, objName)
}

func (m *minIo) getObjedctReadSeek(ctx context.Context, bucName string, objName string) (*domain.OutputObject, error) {
	obj, err := m.client.GetObject(ctx, bucName, objName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	info, err := obj.Stat()
	if err != nil {
		return nil, err
	}

	return &domain.OutputObject{
		ReadSeekCloser: obj,
		ConType:        info.ContentType,
		Modified:       info.LastModified,
	}, nil
}

func (m *minIo) putObject(ctx context.Context, bucName string, obj *domain.InPutObject) error {
	// uploadInfo, err := m.client.PutObject(ctx, m.commentBuc, objName, reader, 0, minio.PutObjectOptions{})
	_, err := m.client.PutObject(
		ctx,
		bucName,
		obj.ObjName,
		obj,
		obj.Size,
		minio.PutObjectOptions{
			ContentType: obj.ConType,
		},
	)
	return err
}

// func (m *minIo) getObject(ctx context.Context, bucName, objName string) (string, error) {
// 	// Проверяем наличие
// 	_, err := m.client.StatObject(ctx, bucName, objName, minio.StatObjectOptions{})
// 	if err != nil {
// 		return "", fmt.Errorf("object not found: %w", err)
// 	}

// 	// Генерируем ссылку
// 	presignedURL, err := m.client.PresignedGetObject(ctx, bucName, objName, 15*time.Minute, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	return presignedURL.String(), nil
// }
