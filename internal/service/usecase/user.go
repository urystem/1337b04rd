package usecase

import (
	"context"
	"errors"

	"1337b04rd/internal/domain"
)

func (u *usecase) AddUserToDB(ctx context.Context, ses *domain.Session) error {
	err := u.db.InsertUser(ctx, ses)
	if err != nil {
		return err
	}
	err = u.session.SetSavedUUID(ctx, ses.Uuid)
	if err != nil {
		return errors.Join(err, u.db.DeleteUser(ctx, ses.Uuid))
	}
	return nil
}
