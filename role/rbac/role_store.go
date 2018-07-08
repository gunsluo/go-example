package rbac

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// XODB This should work with database/sql.DB and database/sql.Tx.
type XODB interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	DriverName() string
	Rebind(query string) string
}

// Role the role statement in system, a group of user with the same permissions
type Role struct {
	ID          int64
	Name        string
	Description string
}

type roleStore struct{}

func newRoleStore() *roleStore {
	return &roleStore{}
}

// Insert adds a role to the storage.
func (s *roleStore) Insert(db XODB, role *Role) error {
	if role == nil {
		return nil
	}

	query := "INSERT INTO role (name, description) VALUES(?,?) RETURNING id"
	if err := db.QueryRow(db.Rebind(query), role.Name, role.Description).Scan(&role.ID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Delete remove a role FROM the storage.
func (s *roleStore) Delete(db XODB, name string) error {
	_, err := db.Exec(db.Rebind("DELETE FROM role WHERE name=?"), name)
	return errors.WithStack(err)
}

// Exist get a role FROM the storage if it exist.
func (s *roleStore) Exist(db XODB, name string) (bool, error) {
	var total int64

	err := db.QueryRow(db.Rebind("SELECT COUNT(1) as total FROM role WHERE name=?"), name).Scan(&total)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return total > 0, nil
}

// Get gets a role by name FROM the storage.
func (s *roleStore) Get(db XODB, name string) (*Role, error) {
	var (
		id   int64
		desc sql.NullString
	)

	err := db.QueryRow(db.Rebind("SELECT id, description FROM role WHERE name = ?"), name).Scan(&id, &desc)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	role := &Role{ID: id, Name: name}
	if desc.Valid {
		role.Description = desc.String
	}

	return role, nil
}

// GetAll gets roles FROM the storage.
func (s *roleStore) GetAll(db XODB, limit, offset int64, conditions ...string) ([]*Role, error) {
	var query string
	var args []interface{}
	if len(conditions) == 0 || conditions[0] == "" {
		query = db.Rebind("SELECT id, name, description FROM role limit ? offset ?")
	} else {
		query = db.Rebind("SELECT id, name, description FROM role WHERE name like ? limit ? offset ?")
		args = append(args, "%"+conditions[0]+"%")
	}
	args = append(args, limit, offset)

	rows, err := db.Query(query, args...)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	roles := []*Role{}
	for rows.Next() {
		var id int64
		var s sql.NullString
		var d sql.NullString
		if err := rows.Scan(&id, &s, &d); err != nil {
			return nil, errors.WithStack(err)
		}

		r := &Role{ID: id}
		if s.Valid {
			r.Name = s.String
		}
		if d.Valid {
			r.Description = d.String
		}
		roles = append(roles, r)
	}

	return roles, nil
}

// Count returns count of actions in the database by a condition
func (s *roleStore) Count(db XODB, conditions ...string) (int64, error) {
	var query string
	var args []interface{}

	if len(conditions) == 0 || conditions[0] == "" {
		query = db.Rebind("SELECT COUNT(1) as count FROM role")
	} else {
		query = db.Rebind("SELECT COUNT(1) as count FROM role WHERE name like ?")
		args = append(args, "%"+conditions[0]+"%")
	}

	var count int64
	err := db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return count, nil
}
