package store

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"github.com/facebookgo/errgroup"
	"github.com/go-errors/errors"
)

// ShardingIndex represents the indexing engine.
type ShardingIndex struct {
	path     string                 // Path to bleve storage
	shards   map[string]bleve.Index // Index shards i.e. bleve indexes
	shardsMu sync.RWMutex           // rw mutex
	alias    bleve.IndexAlias       // All bleve indexes as one reference, for search
}

// New returns a new indexer.
func NewShardingIndex(path string) *ShardingIndex {
	return &ShardingIndex{
		path:   path,
		shards: make(map[string]bleve.Index),
		alias:  bleve.NewIndexAlias(),
	}
}

// Open opens the indexer, preparing it for indexing.
func (i *ShardingIndex) Open() error {
	if err := os.MkdirAll(i.path, 0755); err != nil {
		return errors.Errorf("unable to create index directory %s", i.path)
	}

	/*
		for s := 0; s < cap(i.shards); s++ {
			path := filepath.Join(i.path, strconv.Itoa(s))
			b, err := bleve.New(path, indexMapping())
			if err != nil {
				return fmt.Errorf("index %d at %s: %s", s, path, err.Error())
			}

			i.shards = append(i.shards, b)
			i.alias.Add(b)
		}
	*/

	return nil
}

// Index indexes the given docs, dividing the docs evenly across the shards.
// Blocks until all documents have been indexed.
func (i *ShardingIndex) Index(id string, data interface{}) error {
	prefix := i.key(id)
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
		prefix := i.key(id)
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
	}

	for prefix, batch := range mb {
		b, _, e := i.newIndex(prefix)
		if e != nil {
			return e
		}

		if e := b.Batch(batch); e != nil {
			return errors.Wrap(e, 0)
		}
	}

	return nil
}

func (i *ShardingIndex) key(id string) string {
	if len(id) < 2 {
		return ""
	}

	return id[0:2]
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
	nb, err := bleve.New(path, indexMapping())
	if err != nil {
		return nil, false, errors.Errorf("index %s at %s: %s", prefix, path, err.Error())
	}

	i.shardsMu.Lock()
	i.shards[prefix] = nb
	i.shardsMu.Unlock()
	i.alias.Add(nb)

	return nb, true, nil
}

/*
func (i *ShardingIndex) Index(docs [][]byte) error {
	base := 0
	docsPerShard := (len(docs) / len(i.shards))
	var wg sync.WaitGroup

	wg.Add(len(i.shards))
	for _, s := range i.shards {
		go func(b bleve.Index, ds [][]byte) {
			defer wg.Done()

			batch := b.NewBatch()
			n := 0

			// Just index whole batches.
			for n = 0; n < len(ds)-(len(ds)%i.batchSz); n++ {
				data := struct {
					Body string
				}{
					Body: string(ds[n]),
				}

				if err := batch.Index(strconv.Itoa(n), data); err != nil {
					panic(fmt.Sprintf("failed to index doc: %s", err.Error()))
				}

				if batch.Size() == i.batchSz {
					if err := b.Batch(batch); err != nil {
						panic(fmt.Sprintf("failed to index batch: %s", err.Error()))
					}
					batch = b.NewBatch()
				}
			}
		}(s, docs[base:base+docsPerShard])
		base = base + docsPerShard
	}

	wg.Wait()
	return nil
}
*/

// Count returns the total number of documents indexed.
func (i *ShardingIndex) Count() (uint64, error) {
	return i.alias.DocCount()
}

func (i *ShardingIndex) Search(req *bleve.SearchRequest) (*bleve.SearchResult, error) {
	return i.alias.Search(req)
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

func indexMapping() mapping.IndexMapping {
	// a generic reusable mapping for english text
	standardJustIndexed := bleve.NewTextFieldMapping()
	standardJustIndexed.Store = false
	//standardJustIndexed.IncludeInAll = false
	standardJustIndexed.IncludeTermVectors = false
	standardJustIndexed.Analyzer = "standard"

	articleMapping := bleve.NewDocumentMapping()

	// body
	articleMapping.AddFieldMappingsAt("Body", standardJustIndexed)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.DefaultMapping = articleMapping
	indexMapping.DefaultAnalyzer = "standard"
	return indexMapping
}
