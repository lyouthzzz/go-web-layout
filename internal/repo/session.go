package repo

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
	"github.com/spf13/cast"
)

type sessionRepo struct {
	rdb *redis.Client
}

func NewSessionRepo(rdb *redis.Client) domain.ISessionRepository {
	return &sessionRepo{rdb: rdb}
}

func (s *sessionRepo) Create(ctx context.Context, session *domain.Session) (*domain.Session, error) {
	err := s.rdb.Set(ctx, fmt.Sprintf("session:%s", session.Id), session.UserId, session.ExpireDuration).Err()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionRepo) Get(ctx context.Context, id string) (*domain.Session, error) {
	userIdVal, err := s.rdb.Get(ctx, fmt.Sprintf("session:%s", id)).Result()
	if err != nil {
		if err == redis.Nil {
			// todo
		}
		return nil, err
	}
	return &domain.Session{Id: id, UserId: cast.ToUint(userIdVal)}, nil
}

func (s *sessionRepo) Delete(ctx context.Context, id string) error {
	err := s.rdb.Del(ctx, fmt.Sprintf("session:%s", id)).Err()
	if err != nil {
		return err
	}
	return nil
}
