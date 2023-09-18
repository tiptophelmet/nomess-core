package cacher

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/tiptophelmet/nomess-core/v4/errs"
	"github.com/tiptophelmet/nomess-core/v4/logger"
)

type MemcachedCacher struct {
	client *memcache.Client
}

func InitMemcachedCacher() *MemcachedCacher {
	cache := &MemcachedCacher{}
	return cache
}

func (mc *MemcachedCacher) Connect(url string) error {
	mc.client = memcache.New(url)

	err := mc.client.Ping()
	if err != nil {
		logger.Error("failed to connect to memcached: %v", err.Error())
		return errs.ErrCacheStoreConnectionFailed
	}

	return nil
}

func (mc *MemcachedCacher) Set(key string, val []byte, exp time.Duration) error {
	return mc.client.Set(&memcache.Item{
		Key:        key,
		Value:      val,
		Expiration: int32(exp.Seconds()),
	})
}

func (mc *MemcachedCacher) Has(key string) (bool, error) {
	_, err := mc.Get(key)

	if err == nil {
		return true, nil
	}

	if err == memcache.ErrCacheMiss {
		return false, nil
	}

	return false, err
}

func (mc *MemcachedCacher) Get(key string) ([]byte, error) {
	item, err := mc.client.Get(key)
	if err != nil {
		return nil, err
	}

	return item.Value, nil
}

func (mc *MemcachedCacher) Expire(key string, exp time.Duration) error {
	return mc.client.Touch(key, int32(exp.Seconds()))
}

func (mc *MemcachedCacher) ExpireTime(key string) (time.Duration, error) {
	item, err := mc.client.Get(key)
	if err != nil {
		return -1, err
	}

	return time.Duration(item.Expiration), nil
}

func (mc *MemcachedCacher) Delete(key string) (bool, error) {
	err := mc.client.Delete(key)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (mc *MemcachedCacher) Flush() error {
	return mc.client.FlushAll()
}
