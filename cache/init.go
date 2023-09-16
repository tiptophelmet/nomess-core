package cache

import (
	"sync"

	"github.com/tiptophelmet/nomess-core/v2/cache/cacher"
	"github.com/tiptophelmet/nomess-core/v2/logger"
)

type cacheClient struct {
	cacher cacher.Cacher
	mu     sync.Mutex
}

var client *cacheClient

func Init(driver, url string) {
	switch driver {
	case "redis":
		client = &cacheClient{cacher: cacher.InitRedisCacher()}
	case "memcached":
		client = &cacheClient{cacher: cacher.InitMemcachedCacher()}
	case "aerospike":
		client = &cacheClient{cacher: cacher.InitAerospikeCacher()}
	default:
		logger.Panic("unsupported cache.driver: %v", driver)
	}

	client.cacher.Connect(url)
}

func Connection() cacher.Cacher {
	client.mu.Lock()
	defer client.mu.Unlock()

	return client.cacher
}
