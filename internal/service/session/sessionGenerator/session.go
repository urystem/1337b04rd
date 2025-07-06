package session

import (
	"context"
	"errors"
	"log/slog"

	"1337b04rd/internal/domain"
	"1337b04rd/internal/ports/inbound"
	"1337b04rd/internal/ports/outbound"

	"github.com/google/uuid"
)

type session struct {
	sessinoRedis outbound.SessionRedisInter
	morty        inbound.UseCaseRickAndMorty
}

func InitSession(sessinoRedis outbound.SessionRedisInter, morty inbound.UseCaseRickAndMorty) inbound.SessionInter {
	return &session{sessinoRedis, morty}
}

// for middleware
func (s *session) NewSession(ctx context.Context) *domain.Session {
	newCharacter, err := s.morty.GetRandomCharacterAndDel(ctx)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	newSession := &domain.Session{Uuid: uuid.New(), Name: newCharacter.Name, AvatarURL: newCharacter.Image}
	err = s.sessinoRedis.SetSession(ctx, newSession)
	if err != nil {
		//
		return nil
	}
	return newSession
}

// for middleware
func (s *session) GetSession(ctx context.Context, id string) *domain.Session {
	sesId, err := uuid.Parse(id)
	if err != nil {
		//
		return nil
	}
	ses, err := s.sessinoRedis.GetUserBySession(ctx, sesId)
	if err != nil {
		return nil
	}
	return ses
}

// for handler or service(usecase)
func (s *session) SetSavedUUID(ctx context.Context, id uuid.UUID) error {
	ses, err := s.sessinoRedis.GetUserBySession(ctx, id)
	if err != nil {
		return err
	}
	if ses.Saved {
		return errors.New("already")
	}
	ses.Saved = true
	return s.sessinoRedis.SetSession(ctx, ses)
}
