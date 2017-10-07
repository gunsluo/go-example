package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/gunsluo/go-example/bleve/store"
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

var batchSize = flag.Int("batchSize", 100, "batch size for indexing")
var docsNum = flag.Int("docs", 1000, "test docs num")
var maxprocs = flag.Int("maxprocs", 1, "GOMAXPROCS")
var indexPath = flag.String("index", "indexes", "index storage path")
var csv = flag.Bool("csv", false, "summary CSV output")

//./bleve -docs 1000 -maxprocs 8 -batchSize 100
func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*maxprocs)

	var i store.Indexer
	i = store.NewShardingIndex(*indexPath)
	i.RegisterShardingDirStrategy(func(id string) string {
		if len(id) < 3 {
			return ""
		}

		return id[0:3]
	})
	if err := i.Open(); err != nil {
		fmt.Println("failed to open indexer:", err)
		os.Exit(1)
	}

	var total int
	all := testdata()
	startTime := time.Now()
	for _, docs := range all {
		if err := i.Batch(docs); err != nil {
			fmt.Println("failed to index documents:", err)
			os.Exit(1)
		}

		total += len(docs)
	}
	duration := time.Now().Sub(startTime)

	count, err := i.Count()
	if err != nil {
		fmt.Println("failed to determine total document count")
		os.Exit(1)
	}
	rate := int(float64(count) / duration.Seconds())

	fmt.Printf("Commencing indexing. GOMAXPROCS: %d, batch size: %d.\n", runtime.GOMAXPROCS(-1), *batchSize)

	fmt.Println("Indexing operation took", duration)
	fmt.Printf("%d documents indexed.\n", count)
	fmt.Printf("Indexing rate: %d docs/sec.\n", rate)

	if *csv {
		fmt.Printf("csv,%d,%d,%d,%d,%d\n", total, count, runtime.GOMAXPROCS(-1), *batchSize, rate)
	}

	query := bleve.NewMatchQuery("luoji")
	//query := bleve.NewQueryStringQuery("luoji")
	search := bleve.NewSearchRequest(query)
	res, err := i.Search(search)
	fmt.Printf("Indexing result: %s %v.\n", res, err)

	// Remove any existing indexes.
	if err := i.Clear(); err != nil {
		fmt.Println("failed to remove %s.", *indexPath)
		os.Exit(1)
	}
}

func testdata() []map[string]interface{} {
	var (
		all  []map[string]interface{}
		docs map[string]interface{}
	)

	for i := 0; i < *docsNum; i++ {
		if i%*batchSize == 0 {
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
