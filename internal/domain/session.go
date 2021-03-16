package domain

import (
	"context"
	"time"
)

type Session struct {
	Id             string
	UserId         uint
	ExpireDuration time.Duration
}

type ISessionUsecase interface {
	Create(ctx context.Context, userId uint) (*Session, error)
	Get(ctx context.Context, id string) (*Session, error)
	Delete(ctx context.Context, id string) error
}

type ISessionRepository interface {
	Create(ctx context.Context, session *Session) (*Session, error)
	Get(ctx context.Context, id string) (*Session, error)
	Delete(ctx context.Context, id string) error
}
