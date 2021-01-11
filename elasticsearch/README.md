# Elasticsearch

### 查看所有index
```
localhost:9200/_cat/indices?v&pretty
```

### 创建index
```
curl -X POST 'http://localhost:9200/my-index-000001/_doc?pretty' -H 'Content-Type: application/json' -d '
{
  "@timestamp": "2099-11-15T13:12:00",
  "message": "GET /search HTTP/1.1 200 1070000",
  "user": {
    "id": "kimchy"
  }
}'

{
  "_index" : "my-index-000001",
  "_type" : "_doc",
  "_id" : "y7428HYB4EFdDc-0zTX4",
  "_version" : 1,
  "result" : "created",
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 0,
  "_primary_term" : 1
}
```

### 查询index

1. 指定匹配条件
```
$ curl -X GET 'http://localhost:9200/my-index-000001/_search?pretty=true' -H 'Content-Type: application/json' -d '
{
  "query" : {
    "match" : { "user.id": "kimchy" }
  }
}'

{
  "took" : 1065,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 1,
      "relation" : "eq"
    },
    "max_score" : 0.9808291,
    "hits" : [
      {
        "_index" : "my-index-000001",
        "_type" : "_doc",
        "_id" : "x74f8HYB4EFdDc-0JjWR",
        "_score" : 0.9808291,
        "_source" : {
          "@timestamp" : "2099-11-15T13:12:00",
          "message" : "GET /search HTTP/1.1 200 1070000",
          "user" : {
            "id" : "kimchy"
          }
        }
      }
    ]
  }
}
```

2. 匹配所有
```
curl -X GET 'http://localhost:9200/my-index-000001/_search?pretty=true' -H 'Content-Type: application/json' -d '
{
  "query" : {
    "match_all" : {}
  }
}'
```


3. 匹配时间
```
curl -X GET 'http://localhost:9200/my-index-000001/_search?pretty=true' -H 'Content-Type: application/json' -d '
{
  "query" : {
    "range" : {
      "@timestamp": {
        "from": "2099-11-15T13:00:00",
        "to": "2099-11-15T14:00:00"
      }
    }
  }
}'
```

### 更新index
```
curl -XPOST 'localhost:9200/my-index-000001/_update/y7428HYB4EFdDc-0zTX4' -H 'Content-Type: application/json' -d '
{
  "doc": {
    "user": {
       "id": "luoji"
    }
  }
}'
```

### 删除index

```
curl -XDELETE 'localhost:9200/my-index-000001?pretty'
```

