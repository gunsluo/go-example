package storage

type MssqlStorage struct {
}

func NewMssqlStorage() Storage {
	return &MssqlStorage{}
}
