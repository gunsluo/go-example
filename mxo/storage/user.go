package storage

import "time"

type User struct {
	ID          int       `json:"id"`           // id
	Subject     string    `json:"subject"`      // subject
	CreatedDate time.Time `json:"created_date"` // created_date
	ChangedDate time.Time `json:"changed_date"` // changed_date
	DeletedDate time.Time `json:"deleted_date"` // deleted_date

	// xo fields
	Exists, Deleted bool
}
