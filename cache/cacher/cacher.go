package cacher

import "time"

type Cacher interface {
	Connect(url string) error

	Set(key string, val []byte, exp time.Duration) error
	Has(key string) (bool, error)
	Get(key string) ([]byte, error)

	Expire(key string, exp time.Duration) error
	ExpireTime(key string) (time.Duration, error)

	Delete(key string) (bool, error)
	Flush() error
}

