package store

import (
	"fmt"
	"os"
	"testing"

	"github.com/blevesearch/bleve"
)

type fileInfo struct {
	Path   string `json:"path"`
	Length int    `json:"length"`
}

type bitTorrent struct {
	InfoHash string     `json:"infohash"`
	Name     string     `json:"name"`
	Keyword  string     `json:"keyword"`
	Files    []fileInfo `json:"files,omitempty"`
	Length   int        `json:"length,omitempty"`
	IsDir    bool       `json:"isdir"`
}

func testdata(docsNum, batchSize int) []map[string]interface{} {
	var (
		all  []map[string]interface{}
		docs map[string]interface{}
	)

	for i := 0; i < docsNum; i++ {
		if i%batchSize == 0 {
			docs = make(map[string]interface{})
			all = append(all, docs)
		}

		bt := &bitTorrent{
			Name:    "Jesses.Girls.XXX.DVDRip.x264-Fapulous[rarbg]",
			Keyword: "Jesses Girls XXX DVDRip x264-Fapulous[rarbg] fap-jesgir mp4",
			IsDir:   true,
			Length:  1175926458,
			Files: []fileInfo{
				fileInfo{
					Path:   "fap-jesgir.mp4",
					Length: 1175926458,
				},
			},
		}

		bt.InfoHash = fmt.Sprintf("%08X", i)
		docs[bt.InfoHash] = bt
	}

	nbt := &bitTorrent{
		InfoHash: "4ded5abc1746602e5ebbab08707ddbead59a1b2e",
		Name:     "luoji.Girls.XXX.DVDRip.x264-Fapulous[rarbg]",
		Keyword:  "luoji Girls XXX DVDRip x264-Fapulous[rarbg] fap-jesgir mp4",
		IsDir:    true,
		Length:   1175926458,
		Files: []fileInfo{
			fileInfo{
				Path:   "fap-jesgir.mp4",
				Length: 1175926458,
			},
		},
	}
	docs[nbt.InfoHash] = nbt

	return all
}

func TestIndexAlias(t *testing.T) {
	indexPath := "./data"
	i := NewShardingIndex(indexPath)
	if err := i.Open(); err != nil {
		t.Errorf("failed to open indexer: %v", err)
		return
	}

	all := testdata(1000, 100)
	for _, docs := range all {
		if err := i.Batch(docs); err != nil {
			t.Errorf("failed to index documents: %v", err)
			return
		}
	}

	query := bleve.NewMatchQuery("luoji")
	search := bleve.NewSearchRequest(query)
	res, err := i.Search(search)
	if err != nil {
		t.Errorf("failed to search documents: %v", err)
		return
	}

	if res.Total != 1 {
		t.Errorf("failed to search documents, total[%d] invalid.", res.Total)
		return
	}

	if err := os.RemoveAll(indexPath); err != nil {
		t.Errorf("failed to remove %s.", indexPath)
		return
	}
}

//go test -bench . -cpuprofile=cpu.prof
func BenchmarkIndexAlias(b *testing.B) {
	indexPath := "./data"
	i := NewShardingIndex(indexPath)
	if err := i.Open(); err != nil {
		b.Errorf("failed to open indexer: %v", err)
		return
	}

	all := testdata(1000, 100)
	for _, docs := range all {
		if err := i.Batch(docs); err != nil {
			b.Errorf("failed to index documents: %v", err)
			return
		}
	}

	query := bleve.NewMatchQuery("luoji")
	search := bleve.NewSearchRequest(query)
	res, err := i.Search(search)
	if err != nil {
		b.Errorf("failed to search documents: %v", err)
		return
	}

	if res.Total != 1 {
		b.Errorf("failed to search documents, total[%d] invalid.", res.Total)
		return
	}

	if err := os.RemoveAll(indexPath); err != nil {
		b.Errorf("failed to remove %s.", indexPath)
		return
	}
}
