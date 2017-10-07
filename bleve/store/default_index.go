package store

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/blevesearch/bleve"
	"github.com/facebookgo/errgroup"
	"github.com/go-errors/errors"
	"github.com/toolkits/file"
)

// ShardingIndex represents the indexing engine.
type ShardingIndex struct {
	path             string                 // Path to bleve storage
	shards           map[string]bleve.Index // Index shards i.e. bleve indexes
	shardsMu         sync.RWMutex           // rw mutex
	alias            bleve.IndexAlias       // All bleve indexes as one reference, for search
	sdsfn            ShardingDirStrategyFn  // sharding dir strategy func.
	imfn             IndexMappingFn         // index mapping func.
	kvStore          bool                   // if enable key value store option, default false
	kvStoreMarshal   KVStoreMarshal         // marshal kv store
	kvStoreUnmarshal KVStoreUnmarshal       // unmarshal kv store
}

// New returns a new indexer.
func NewShardingIndex(path string) *ShardingIndex {
	return &ShardingIndex{
		path:             path,
		shards:           make(map[string]bleve.Index),
		alias:            bleve.NewIndexAlias(),
		sdsfn:            defaultShardingDirStrategyFn,
		imfn:             defaultIndexMappingFn,
		kvStoreMarshal:   defaultKVStoreMarshal,
		kvStoreUnmarshal: defaultKVStoreUnmarshal,
	}
}

// Open opens the indexer, preparing it for indexing.
func (i *ShardingIndex) Open() error {
	if file.IsExist(i.path) == false {
		if err := os.MkdirAll(i.path, 0755); err != nil {
			return errors.Errorf("unable to create index directory %s", i.path)
		}

		return nil
	}

	// each path for load index
	subDirs, err := file.DirsUnder(i.path)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	for _, subDir := range subDirs {
		path := filepath.Join(i.path, subDir)
		nb, err := bleve.Open(path)
		if err != nil {
			return errors.Errorf("index %s at %s: %s", subDir, path, err.Error())
		}

		i.shardsMu.Lock()
		i.shards[subDir] = nb
		i.shardsMu.Unlock()
		i.alias.Add(nb)
	}

	return nil
}

// Index indexes the given docs, dividing the docs evenly across the shards.
// Blocks until all documents have been indexed.
func (i *ShardingIndex) Index(id string, data interface{}) error {
	prefix := i.sdsfn(id)
	b, _, e := i.newIndex(prefix)
	if e != nil {
		return e
	}

	// write data to index
	return b.Index(id, data)
}

func (i *ShardingIndex) Batch(ms map[string]interface{}) error {
	var (
		mb = make(map[string]*bleve.Batch)
	)

	// classified batch
	for id, data := range ms {
		prefix := i.sdsfn(id)
		b, _, e := i.newIndex(prefix)
		if e != nil {
			return e
		}

		batch, ok := mb[prefix]
		if !ok {
			batch = b.NewBatch()
			mb[prefix] = batch
		}

		// write data to batch
		if e = batch.Index(id, data); e != nil {
			return errors.Wrap(e, 0)
		}

		// kv store
		if i.kvStore {
			vBuf, e := i.kvStoreMarshal(data)
			if e != nil {
				return e
			}

			e = b.SetInternal([]byte(id), vBuf)
			if e != nil {
				return e
			}
		}
	}

	for prefix, batch := range mb {
		b := i.getIndex(prefix)
		if b == nil {
			continue
		}

		if e := b.Batch(batch); e != nil {
			return errors.Wrap(e, 0)
		}
	}

	return nil
}

// new index by sub dir
func (i *ShardingIndex) newIndex(prefix string) (bleve.Index, bool, error) {
	if prefix == "" {
		return nil, false, errors.Errorf("index prefix is nil.")
	}

	//path exist
	i.shardsMu.RLock()
	ob, ok := i.shards[prefix]
	i.shardsMu.RUnlock()
	if ok {
		return ob, false, nil
	}

	path := filepath.Join(i.path, prefix)
	nb, err := bleve.New(path, i.imfn())
	if err != nil {
		return nil, false, errors.Errorf("index %s at %s: %s", prefix, path, err.Error())
	}

	i.shardsMu.Lock()
	i.shards[prefix] = nb
	i.shardsMu.Unlock()
	i.alias.Add(nb)

	return nb, true, nil
}

func (i *ShardingIndex) getIndex(prefix string) bleve.Index {
	i.shardsMu.RLock()
	ob, ok := i.shards[prefix]
	i.shardsMu.RUnlock()
	if ok {
		return ob
	}

	return nil
}

// Count returns the total number of documents indexed.
func (i *ShardingIndex) Count() (uint64, error) {
	return i.alias.DocCount()
}

// SetShardingDirStrategy set ShardingDirStrategyFn
func (i *ShardingIndex) SetShardingDirStrategy(sdsfn ShardingDirStrategyFn) {
	i.sdsfn = sdsfn
}

// SetIndexMapping set indexMapping
func (i *ShardingIndex) SetIndexMapping(imfn IndexMappingFn) {
	i.imfn = imfn
}

// Search query form index
func (i *ShardingIndex) Search(req *bleve.SearchRequest) (*bleve.SearchResult, error) {
	return i.alias.Search(req)
}

// EnableKVStore enable kv store option
func (i *ShardingIndex) EnableKVStore() {
	i.kvStore = true
}

// GetInternal get kv store, need enable kvStore.
func (i *ShardingIndex) GetInternal(id string, v interface{}) error {
	prefix := i.sdsfn(id)
	b := i.getIndex(prefix)
	if b == nil {
		return errors.Errorf("index id:%s not exist.", id)
	}

	vBuf, err := b.GetInternal([]byte(id))
	if err != nil {
		return errors.Wrap(err, 0)
	}

	return i.kvStoreUnmarshal(vBuf, v)
}

func (i *ShardingIndex) Close() error {
	var g errgroup.Group

	i.shardsMu.Lock()
	for _, b := range i.shards {
		if e := b.Close(); e != nil {
			g.Error(errors.Wrap(e, 0))
		}
	}
	i.shardsMu.Unlock()

	return g.Wait()
}

func (i *ShardingIndex) Clear() error {
	if err := os.RemoveAll(i.path); err != nil {
		return errors.Errorf("failed to remove %s, %v.", i.path, err)
	}

	return nil
}
