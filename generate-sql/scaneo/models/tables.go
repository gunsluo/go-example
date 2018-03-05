package models

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        int
	Created   time.Time
	Published pq.NullTime
	Draft     bool
	Title     string
	Body      string
}
