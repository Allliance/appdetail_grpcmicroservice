package cache

import (
	"time"

	bigcache "github.com/allegro/bigcache/v3"
)

var config bigcache.Config = bigcache.DefaultConfig(10 * time.Minute)

type AppDetailCache struct {
	cache *bigcache.BigCache
}

func NewCache() (*AppDetailCache, error) {
	cache, err := bigcache.NewBigCache(config)
	return &AppDetailCache{cache: cache}, err
}

func (cacheInstance *AppDetailCache) CacheApp(key string, data []byte) {
	cacheInstance.cache.Set(key, data)
}

func (cacheInstance *AppDetailCache) Retrieve(key string) ([]byte, error) {
	return cacheInstance.cache.Get(key)
}
