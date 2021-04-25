package models

import (
	"time"
)

type Inspector struct {
	Id       int       `xorm:"not null pk autoincr INTEGER"`
	Username string    `xorm:"not null default '' VARCHAR(256)"`
	Password string    `xorm:"not null default '' VARCHAR(256)"`
	Created  time.Time `xorm:"not null default now() DATETIME"`
}

func (m *Inspector) TableName() string {
	return "inspector"
}
