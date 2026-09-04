package merchant_document_cache

import "github.com/MamangRust/monolith-graphql-pointofsale-shared/cache"

type MerchantDocumentMencache interface {
	MerchantDocumentCommandCache
	MerchantDocumentQueryCache
}

type merchantDocumentMencache struct {
	MerchantDocumentCommandCache
	MerchantDocumentQueryCache
}

func NewMerchantDocumentMencache(cacheStore *cache.CacheStore) MerchantDocumentMencache {
	return &merchantDocumentMencache{
		MerchantDocumentCommandCache: NewMerchantDocumentCommandCache(cacheStore),
		MerchantDocumentQueryCache:   NewMerchantDocumentQueryCache(cacheStore),
	}
}
