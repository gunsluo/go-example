package store

import (
	"encoding/json"
)

type KVStoreMarshal func(data interface{}) ([]byte, error)
type KVStoreUnmarshal func(data []byte, v interface{}) error

func defaultKVStoreMarshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func defaultKVStoreUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
