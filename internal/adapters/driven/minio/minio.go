package minio

import (
	"context"
	"time"

	"1337b04rd/internal/domain"

	"github.com/minio/minio-go/v7"
)

func (m *minIo) PutPost(ctx context.Context, obj *domain.Object) error {
	return m.putObject(ctx, m.postBuc, obj)
}

func (m *minIo) DelPost(ctx context.Context, objName string) error {
	return m.client.RemoveObject(ctx, m.postBuc, objName, minio.RemoveObjectOptions{})
}

func (m *minIo) GetPost(ctx context.Context, objName string) (string, error) {
	return m.getObject(ctx, m.postBuc, objName)
}

func (m *minIo) PutComment(ctx context.Context, obj *domain.Object) error {
	return m.putObject(ctx, m.commentBuc, obj)
}

func (m *minIo) DelComment(ctx context.Context, objName string) error {
	return m.client.RemoveObject(ctx, m.commentBuc, objName, minio.RemoveObjectOptions{})
}

func (m *minIo) GetComment(ctx context.Context, objName string) (string, error) {
	return m.getObject(ctx, m.commentBuc, objName)
}

func (m *minIo) getObject(ctx context.Context, bucName, objName string) (string, error) {
	presignedURL, err := m.client.PresignedGetObject(ctx, bucName, objName, 15*time.Minute, nil)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (m *minIo) putObject(ctx context.Context, bucName string, obj *domain.Object) error {
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
	if err != nil {
		return err
	}
	return nil
}
