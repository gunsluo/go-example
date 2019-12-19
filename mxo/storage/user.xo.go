package storage

type User struct {
	ID          int      `json:"id"`           // id
	Subject     string   `json:"subject"`      // subject
	CreatedDate NullTime `json:"created_date"` // created_date
	ChangedDate NullTime `json:"changed_date"` // changed_date
	DeletedDate NullTime `json:"deleted_date"` // deleted_date

	// xo fields
	_exists, _deleted bool
}
