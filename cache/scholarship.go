package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var scholarshipCache = cache.New(5*time.Minute, 10*time.Minute)

func GetScholarshipCache(key string) (interface{}, bool) {
	return scholarshipCache.Get(key)
}

func SetScholarshipCache(key string, data interface{}) {
	scholarshipCache.Set(key, data, cache.DefaultExpiration)
}

func InvalidateScholarshipCache() {
	scholarshipCache.Flush()
}
