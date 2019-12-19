package storage

type PostgresStorage struct {
}

func NewPostgresStorage() Storage {
	return &PostgresStorage{}
}
