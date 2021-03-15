package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
	"strings"
	"time"
)

type Session struct {
	expireDuration time.Duration
	repo           domain.SessionRepository
}

func NewSessionUsecase(sessionRepo domain.SessionRepository, expireDuration time.Duration) domain.SessionUsecase {
	return &Session{repo: sessionRepo, expireDuration: expireDuration}
}

func (s *Session) Create(ctx context.Context, userId uint) (*domain.Session, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	sid := strings.ReplaceAll(uid.String(), "-", "")
	return s.repo.Create(ctx, &domain.Session{Id: sid, UserId: userId, ExpireDuration: s.expireDuration})
}

func (s *Session) Get(ctx context.Context, id string) (*domain.Session, error) {
	return s.repo.Get(ctx, id)
}

func (s *Session) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
