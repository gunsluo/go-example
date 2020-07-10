package storage

import (
	"github.com/pkg/errors"
)

type StorageExtension interface {
	Storager
	CustomEndpoint(db XODB, args ...interface{}) error
}

type PostgresStorageExtension struct {
	PostgresStorage
}

func (s *PostgresStorageExtension) CustomEndpoint(db XODB, args ...interface{}) error {
	// TODO:
	return nil
}

type MssqlStorageExtension struct {
	MssqlStorage
}

func (s *MssqlStorageExtension) CustomEndpoint(db XODB, args ...interface{}) error {
	// TODO:
	return nil
}

type GodrorStorageExtension struct {
	GodrorStorage
}

func (s *GodrorStorageExtension) CustomEndpoint(db XODB, args ...interface{}) error {
	// TODO:
	return nil
}

func NewStorageExtension(driver string, opts ...Option) (StorageExtension, error) {
	o := applyOptions(opts...)

	var s StorageExtension
	switch driver {
	case "postgres":
		s = &PostgresStorageExtension{PostgresStorage: PostgresStorage{Logger: o.logger}}
	case "mssql":
		s = &MssqlStorageExtension{MssqlStorage: MssqlStorage{Logger: o.logger}}
	case "godror":
		s = &GodrorStorageExtension{GodrorStorage: GodrorStorage{Logger: o.logger}}
	default:
		return nil, errors.New("driver " + driver + " not support")
	}

	return s, nil
}
