package caching

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	cached *cache.Cache
)

//Initialize is used to Initialize the in-memory cache
func Initialize() {
	cached = cache.New(-1, 5*time.Minute)

	//Initialize Cache during startup
}

//SetCache - setting key - value in the in-mem cache
func SetCache(key string, dst interface{}) bool {
	cached.Set(key, dst, cache.NoExpiration)
	return true
}

//GetCache - getting key value from the in-mem cache
func GetCache(key string) (interface{}, bool) {
	var found bool
	data, found := cached.Get(key)
	return data, found
}

//DeleteCache - getting key value from the in-mem cache
func DeleteCache(key string) {

	cached.Delete(key)

}
