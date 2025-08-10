package usecase

import (
	"context"
	"errors"
	"strconv"

	"1337b04rd/internal/domain"
)

func (u *usecase) imageSaver(ctx context.Context, file *domain.InPutObject, id uint64) error {
	if file == nil {
		return nil
	}

	file.ObjName = strconv.FormatUint(id, 10)
	err := u.s3.PutComment(ctx, file)
	if err != nil {
		return errors.Join(err, u.db.DeleteComment(ctx, id))
	}
	return nil
}
