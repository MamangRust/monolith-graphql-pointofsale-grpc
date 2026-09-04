package mencache

import (
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/cache"
)

type Mencache interface {
	ProductQueryCache
	ProductCommandCache
}

type mencache struct {
	ProductQueryCache
	ProductCommandCache
}

func NewMencache(cacheStore *cache.CacheStore) Mencache {
	return &mencache{
		ProductQueryCache:   NewProductQueryCache(cacheStore),
		ProductCommandCache: NewProductCommandCache(cacheStore),
	}
}
