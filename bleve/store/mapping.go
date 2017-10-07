package store

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

type IndexMappingFn func() mapping.IndexMapping

func defaultIndexMappingFn() mapping.IndexMapping {
	indexMapping := bleve.NewIndexMapping()
	return indexMapping
}
