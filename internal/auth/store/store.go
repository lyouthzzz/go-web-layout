package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Store interface {
	Write(ctx context.Context, key string, val string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type rdbStore struct {
	rdb            *redis.Client
	prefix         string
	expireDuration time.Duration
}

func NewRDBStore(rdb *redis.Client, prefix string, expDur time.Duration) Store {
	return &rdbStore{rdb: rdb, prefix: prefix, expireDuration: expDur}
}

func (r *rdbStore) Write(ctx context.Context, key string, val string) error {
	return r.rdb.Set(ctx, key, val, r.expireDuration).Err()
}

func (r *rdbStore) Get(ctx context.Context, key string) (string, error) {
	return r.rdb.Get(ctx, key).Result()
}

func (r *rdbStore) Delete(ctx context.Context, key string) error {
	return r.rdb.Del(ctx, key).Err()
}
