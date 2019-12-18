package storage

type Storage interface {
	UserByID(db XODB, id int) (*User, error)
}
