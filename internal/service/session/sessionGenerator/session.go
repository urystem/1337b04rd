package session

import (
	"context"
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
