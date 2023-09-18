package cacher

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tiptophelmet/nomess-core/v4/errs"
	"github.com/tiptophelmet/nomess-core/v4/logger"
)

type RedisCacher struct {
	client *redis.Client
}

func InitRedisCacher() *RedisCacher {
	cache := &RedisCacher{}
	return cache
}

func (rc *RedisCacher) Connect(url string) error {
	options, err := redis.ParseURL(url)
	if err != nil {
		logger.Error("failed to connect to redis: %v", err.Error())
		return errs.ErrCacheStoreConnectionFailed
	}

	rc.client = redis.NewClient(options)
	return nil
}

func (rc *RedisCacher) Set(key string, val []byte, exp time.Duration) error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	return rc.client.Set(ctx, key, val, exp).Err()
}

func (rc *RedisCacher) Has(key string) (bool, error) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	intCmd := rc.client.Exists(ctx)
	if err := intCmd.Err(); err != nil {
		return false, err
	}

	return intCmd.Val() == 1, nil
}

func (rc *RedisCacher) Get(key string) ([]byte, error) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	stringCmd := rc.client.Get(ctx, key)
	if err := stringCmd.Err(); err != nil {
		return nil, err
	}

	return []byte(stringCmd.Val()), nil
}

func (rc *RedisCacher) Expire(key string, exp time.Duration) error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	boolCmd := rc.client.Expire(ctx, key, exp)
	if err := boolCmd.Err(); err != nil {
		return err
	}

	return nil
}

func (rc *RedisCacher) ExpireTime(key string) (time.Duration, error) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	durationCmd := rc.client.ExpireTime(ctx, key)
	if err := durationCmd.Err(); err != nil {
		return -1, err
	}

	return durationCmd.Val(), nil
}

func (rc *RedisCacher) Delete(key string) (bool, error) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	intCmd := rc.client.Del(ctx, key)

	if err := intCmd.Err(); err != nil {
		return false, err
	}

	return intCmd.Val() == 1, nil
}

func (rc *RedisCacher) Flush() error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	return rc.client.FlushAll(ctx).Err()
}
