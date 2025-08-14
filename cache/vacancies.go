package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var vacanciesCache = cache.New(5*time.Minute, 10*time.Minute)

func GetVacanciesCache(key string) (interface{}, bool) {
	return vacanciesCache.Get(key)
}

func SetVacanciesCache(key string, data interface{}) {
	vacanciesCache.Set(key, data, cache.DefaultExpiration)
}

func InvalidateVacanciesCache() {
	vacanciesCache.Flush()
}
