package usecase

import (
	"context"
	"errors"

	"1337b04rd/internal/domain"

	"github.com/google/uuid"
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

// kerek emes au
func (u *usecase) DeleteUserOnDB(ctx context.Context, id uuid.UUID) error {
	return u.db.DeleteUser(ctx, id)
}
