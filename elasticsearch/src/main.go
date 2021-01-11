package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func main() {
	demo()
}

func demo() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		// Username: "foo",
		// Password: "bar",
	}
	es, err := elasticsearch.NewClient(cfg)
	//es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	{
		//version
		log.Println(elasticsearch.Version)
		resp, err := es.Info()
		if err != nil {
			panic(err)
		}
		log.Println(resp)
	}

	var indexName = "my-index-test"
	{
		// 2. Index documents concurrently
		var wg sync.WaitGroup
		for i, title := range []string{"Test One", "Test Two"} {
			wg.Add(1)

			go func(i int, title string) {
				defer wg.Done()

				// Build the request body.
				var b strings.Builder
				// @timestamp
				b.WriteString(`{"@timestamp": "2099-11-15T13:12:00",`)
				b.WriteString(`"title" : "`)
				b.WriteString(title)
				b.WriteString(`"}`)

				// Set up the request object.
				req := esapi.IndexRequest{
					Index: indexName,
					//DocumentID: strconv.Itoa(i + 1),
					Body:    strings.NewReader(b.String()),
					Refresh: "true",
				}

				// Perform the request with the client.
				res, err := req.Do(context.Background(), es)
				if err != nil {
					log.Fatalf("Error getting response: %s", err)
				}
				defer res.Body.Close()

				if res.IsError() {
					log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
				} else {
					// Deserialize the response into a map.
					var r map[string]interface{}
					if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
						log.Printf("Error parsing the response body: %s", err)
					} else {
						// Print the response status and indexed document version.
						log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
					}
				}
			}(i, title)
		}
		wg.Wait()
	}

	{
		// 3. Search for the indexed documents
		//
		// Build the request body.
		var buf bytes.Buffer
		query := map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"title": "test",
				},
			},
		}
		if err := json.NewEncoder(&buf).Encode(query); err != nil {
			log.Fatalf("Error encoding query: %s", err)
		}

		// Perform the search request.
		res, err := es.Search(
			es.Search.WithContext(context.Background()),
			es.Search.WithIndex(indexName),
			es.Search.WithBody(&buf),
			es.Search.WithTrackTotalHits(true),
			es.Search.WithPretty(),
		)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			var e map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
				log.Fatalf("Error parsing the response body: %s", err)
			} else {
				// Print the response status and error information.
				log.Fatalf("[%s] %s: %s",
					res.Status(),
					e["error"].(map[string]interface{})["type"],
					e["error"].(map[string]interface{})["reason"],
				)
			}
		}

		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}
		// Print the response status, number of results, and request duration.
		log.Printf(
			"[%s] %d hits; took: %dms",
			res.Status(),
			int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
			int(r["took"].(float64)),
		)
		// Print the ID and document source for each hit.
		for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
			log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		}

		log.Println(strings.Repeat("=", 37))
	}
}
