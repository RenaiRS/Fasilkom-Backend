package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var beritaCache = cache.New(5*time.Minute, 10*time.Minute)

func GetBeritaCache(key string) (interface{}, bool) {
	return beritaCache.Get(key)
}

func SetBeritaCache(key string, data interface{}) {
	beritaCache.Set(key, data, cache.DefaultExpiration)
}

func InvalidateBeritaCache() {
	beritaCache.Flush()
}
