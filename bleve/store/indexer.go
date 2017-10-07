package store

import "github.com/blevesearch/bleve"

type Indexer interface {
	Open() error
	Index(id string, data interface{}) error
	Batch(ms map[string]interface{}) error
	Count() (uint64, error)
	Search(req *bleve.SearchRequest) (*bleve.SearchResult, error)
	RegisterShardingDirStrategy(ShardingDirStrategyFn)
	Close() error
	Clear() error
}
