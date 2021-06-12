package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func main() {
	query()
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

func demo1() {
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

	//ctx := context.Background()
	var indexName = "ac.auditlog"

	res, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		panic(err)
	}
	res.Body.Close()

	if res.StatusCode == 404 {
		indexSettings := `{
  "mappings": {
    "properties": {
      "@timestamp": {
        "type": "date"
      },
      "subject": {
        "type": "keyword"
      },
      "content": {
        "type": "text"
      },
      "format": {
        "type": "keyword"
      },
      "event": {
        "type": "keyword"
      },
      "labels": {
        "type": "keyword"
      },
      "orgId": {
        "type": "integer"
      },
      "reason": {
        "type": "text"
      }
    }
  }
}`

		res, err := es.Indices.Create(indexName,
			es.Indices.Create.WithBody(strings.NewReader(indexSettings)),
		)
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		if res.IsError() {
			var e map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
				panic(err)
			}

			err := fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
			panic(err)
		}

		fmt.Println("created", res)
	} else {
		fmt.Println("exists")
	}
}

func bulk() {
	/*
		cfg := elasticsearch.Config{
			Addresses: []string{
				"http://localhost:9200",
			},
			// Username: "foo",
			// Password: "bar",
		}
		es, err := elasticsearch.NewClient(cfg)
		if err != nil {
			panic(err)
		}

		res, err = es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex(indexName))
	*/
}

func query() {
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

	ctx := context.Background()
	var indexName = "ac.auditlog"

	pitRes, err := es.OpenPointInTime(
		es.OpenPointInTime.WithContext(ctx),
		es.OpenPointInTime.WithIndex(indexName),
		es.OpenPointInTime.WithKeepAlive("1m"),
	)
	if err != nil {
		panic(err)
	}
	defer pitRes.Body.Close()

	if pitRes.IsError() {
		var msg esMessage
		if err := json.NewDecoder(pitRes.Body).Decode(&msg); err != nil {
			fmt.Printf("parsing error %w\n", err)
			return
		}
		fmt.Printf("Error response status: %s type: %s Reason: %s\n",
			pitRes.Status(), msg.Error.Type, msg.Error.Reason)
		return
	}

	var pitResp PointInTimeResp
	if err := json.NewDecoder(pitRes.Body).Decode(&pitResp); err != nil {
		fmt.Printf("parsing error %w\n", err)
		return
	}
	fmt.Printf("-->%+v\n", pitResp.Id)

	var query = `{"query":{"match_all":{}},"sort":[{"@timestamp":{"order":"desc"}},{"id":{"order":"desc"}}],"size":100,"pit":{"id":"` + pitResp.Id + `","keep_alive":"1m"}}`
	res, err := es.Search(
		es.Search.WithContext(ctx),
		//es.Search.WithIndex(indexName),
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var msg esMessage
		if err := json.NewDecoder(res.Body).Decode(&msg); err != nil {
			fmt.Printf("parsing error %w\n", err)
			return
		}
		fmt.Printf("Error response status: %s type: %s Reason: %s\n",
			res.Status(), msg.Error.Type, msg.Error.Reason)
		return
	}

	var resp esResp
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		fmt.Printf("parsing error %w\n", err)
		return
	}

	fmt.Printf("%d\n", resp.Hits.Total.Value)

	var hits []logHit
	if err := resp.Hits.Hits.Marshal(&hits); err != nil {
		fmt.Printf("parsing sources error %v\n", err)
		return
	}

	var logs []*AuditloggerEntity
	for _, hit := range hits {
		source := hit.Source
		logs = append(logs, &source)
		fmt.Printf("%+v, %#+v\n", source, hit.Sort)
		v, err := parseSortValue2String(hit.Sort)
		if err != nil {
			fmt.Printf("parsing sort %v\n", err)
			return
		}
		fmt.Printf("%s\n", v)
	}
}

type esMessage struct {
	Error esError `json:"error"`
}

type esError struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}

type esResp struct {
	Took int    `json:"took"`
	Hits esHits `json:"hits"`
}

type esHits struct {
	Total    esTotalHits `json:"total"`
	MaxScore float64     `json:"max_score"`
	Hits     Hits        `json:"hits"`
}

type esTotalHits struct {
	Value    int64  `json:"value"`
	Relation string `json:"relation"`
}

type Hits []json.RawMessage

func (h Hits) Marshal(sourcesSlicePtr interface{}) error {
	if len(h) == 0 {
		return nil
	}

	sliceValue := reflect.Indirect(reflect.ValueOf(sourcesSlicePtr))
	if sliceValue.Kind() != reflect.Slice {
		return fmt.Errorf("parameters must be a slice")
	}
	sliceType := sliceValue.Type()
	sliceElementType := sliceType.Elem()

	for _, item := range h {
		buffer, err := item.MarshalJSON()
		if err != nil {
			return err
		}
		pv := reflect.New(sliceElementType)
		v := pv.Interface()
		if err := json.Unmarshal(buffer, v); err != nil {
			return err
		}

		if sliceElementType.Kind() == reflect.Ptr {
			sliceValue.Set(reflect.Append(sliceValue, reflect.ValueOf(v)))
		} else if sliceElementType.Kind() == reflect.Struct {
			sliceValue.Set(reflect.Append(sliceValue, reflect.Indirect(reflect.ValueOf(v))))
		}
	}

	return nil
}

type PointInTimeResp struct {
	Id string `json:"id,omitempty"`
}

type Hit struct {
	Index string    `json:"_index"`
	Type  string    `json:"_type"`
	Id    string    `json:"_id"`
	Score float64   `json:"_score"`
	Sort  SortValue `json:"sort"`
}

type SortValue []json.RawMessage

func parseSortValue2String(v SortValue) (string, error) {
	if len(v) == 0 {
		return "[]", nil
	}

	buf, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

type logHit struct {
	Hit
	Source AuditloggerEntity `json:"_source"`
}

type nameHit struct {
	Hit
	Source nameSource `json:"_source"`
}

type nameSource struct {
	Timestamp string `json:"@timestamp"`
	Name      string `json:"name"`
}

// AuditloggerEntity is information of audit log
type AuditloggerEntity struct {
	Id        string   `json:"id"`
	Timestamp string   `json:"@timestamp"`
	Subject   string   `json:"subject"`
	Content   string   `json:"content"`
	Format    string   `json:"format"`
	Event     string   `json:"event"`
	Labels    []string `json:"labels"`
	OrgId     int64    `json:"orgId"`
	Where     string   `json:"where"`
	Reason    string   `json:"reason"`
}
