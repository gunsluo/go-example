package store

import "github.com/blevesearch/bleve"

type Indexer interface {
	Open() error
	Index(id string, data interface{}) error
	Batch(ms map[string]interface{}) error
	Count() (uint64, error)
	Search(req *bleve.SearchRequest) (*bleve.SearchResult, error)
	SetShardingDirStrategy(ShardingDirStrategyFn)
	SetIndexMapping(imfn IndexMappingFn)
	GetInternal(id string, v interface{}) error
	EnableKVStore()
	Close() error
	Clear() error
}
