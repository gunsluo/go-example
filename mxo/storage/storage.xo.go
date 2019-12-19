package storage

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type Storage interface {
	UserByID(db XODB, id int) (*User, error)
	InsertUser(db XODB, u *User) error
}

type Config struct {
	Logger logrus.FieldLogger
}

func New(driver string, c Config) (Storage, error) {
	var s Storage
	switch driver {
	case "mssql":
		s = NewMssqlStorage()
	case "postgres":
		s = NewPostgresStorage()
	default:
		return nil, errors.New("driver " + driver + " not support")
	}

	logger = c.Logger
	return s, nil
}

var logger logrus.FieldLogger

func xoLog(s string, args ...interface{}) {
	if logger != nil {
		logger.Infof("%s %v", s, args)
	}
}
