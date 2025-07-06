package usecase

import "context"

func (u *usecase) Archiver(ctx context.Context) error {
	return u.db.Archiver(ctx)
}
