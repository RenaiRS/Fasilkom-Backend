package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var kemahasiswaanKerjasamaCache = cache.New(5*time.Minute, 10*time.Minute)

func GetKemahasiswaanKerjasamaCache(key string) (interface{}, bool) {
	return kemahasiswaanKerjasamaCache.Get(key)
}

func SetKemahasiswaanKerjasamaCache(key string, data interface{}) {
	kemahasiswaanKerjasamaCache.Set(key, data, cache.DefaultExpiration)
}

func InvalidateKemahasiswaanKerjasamaCache() {
	kemahasiswaanKerjasamaCache.Flush()
}
