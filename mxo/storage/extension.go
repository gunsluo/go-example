package storage

import (
	"reflect"

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

func NewStorageExtension(driver string, c Config) (StorageExtension, error) {
	var logger XOLogger
	if c.Logger != nil && !(reflect.ValueOf(c.Logger).Kind() == reflect.Ptr && reflect.ValueOf(c.Logger).IsNil()) {
		logger = c.Logger
	}

	var s StorageExtension
	switch driver {
	case "postgres":
		s = &PostgresStorageExtension{PostgresStorage: PostgresStorage{Logger: logger}}
	case "mssql":
		s = &MssqlStorageExtension{MssqlStorage: MssqlStorage{Logger: logger}}
	case "godror":
		s = &GodrorStorageExtension{GodrorStorage: GodrorStorage{Logger: logger}}
	default:
		return nil, errors.New("driver " + driver + " not support")
	}

	return s, nil
}
