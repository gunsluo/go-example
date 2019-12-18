package postgres

import "github.com/gunsluo/go-example/mxo/storage"

type PostgresStorage struct {
}

func NewPostgresStorage() storage.Storage {
	return &PostgresStorage{}
}
