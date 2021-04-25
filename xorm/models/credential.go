package models

import (
	"time"
)

type Credential struct {
	Id          int       `xorm:"not null pk autoincr INTEGER"`
	PhoneNumber string    `xorm:"not null default '' unique unique VARCHAR(32)"`
	Verified    bool      `xorm:"not null default false BOOL"`
	ChangedPwd  bool      `xorm:"not null default false BOOL"`
	Password    string    `xorm:"not null default '' VARCHAR(256)"`
	Created     time.Time `xorm:"not null default now() DATETIME"`
}

func (m *Credential) TableName() string {
	return "credential"
}
