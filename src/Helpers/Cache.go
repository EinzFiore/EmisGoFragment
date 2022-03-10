package helpers

import (
	"encoding/json"
	"time"

	"github.com/patrickmn/go-cache"
)

var myCache *AppCache

// type CacheItf interface {
// 	SetCache(key string, data interface{}, expiration time.Duration) error
// 	GetCache(key string) ([]byte, error)
// }

type AppCache struct {
	client *cache.Cache
}

func InitCache() {
	myCache = &AppCache{
		client: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func SetCache(key string, data interface{}, expiration time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	myCache.client.Set(key, b, expiration)
	return nil
}

func GetCache(key string) ([]byte, bool) {
	res, exist := myCache.client.Get(key)
	if !exist {
		return nil, false
	}

	resByte, ok := res.([]byte)
	if !ok {
		return nil, false
	}

	return resByte, true
}
