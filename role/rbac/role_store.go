package rbac

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

/*
CREATE TABLE IF NOT EXISTS role (
    id BIGSERIAL PRIMARY KEY,
    name varchar(127) NOT NULL UNIQUE,
	description varchar(511) NOT NULL,
    created_date timestamp DEFAULT now(),
    changed_date timestamp DEFAULT now(),
    deleted_date timestamp
);
*/

// Role the role statement in system, a group of user with the same permissions
type Role struct {
	ID          int64
	Name        string
	Description string
}

type roleStore struct {
	db       *sqlx.DB
	database string
}

func newRoleStore(db *sqlx.DB) *roleStore {
	database := db.DriverName()
	switch database {
	case "pgx", "pq":
		database = "postgres"
	}

	return &roleStore{
		db:       db,
		database: database,
	}
}

// Create adds a role to the storage.
func (s *roleStore) Create(role, desc string) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return errors.WithStack(err)
	}

	if err = s.create(tx, role, desc); err != nil {
		return errors.WithStack(err)
	}

	if err = tx.Commit(); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, rollErr.Error())
		}
		return errors.WithStack(err)
	}

	return nil
}

func (s *roleStore) create(tx *sqlx.Tx, role, desc string) error {
	//query := Migrations[s.database].QueryInsertPolicyActions
	query := "INSERT INTO role (name, description) VALUES(?,?)"
	if _, err := tx.Exec(tx.Rebind(query), role, desc); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// BatchCreate add multiple roles to the storage in the transrole.
func (s *roleStore) BatchCreate(roles []*Role) error {
	if len(roles) == 0 {
		return nil
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return errors.WithStack(err)
	}

	for _, role := range roles {
		if err = s.create(tx, role.Name, role.Description); err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return errors.Wrap(err, rollErr.Error())
			}
			return errors.WithStack(err)
		}
	}

	if err = tx.Commit(); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, rollErr.Error())
		}
		return errors.WithStack(err)
	}

	return nil
}

// Delete remove a role from the storage.
func (s *roleStore) Delete(role string) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return errors.WithStack(err)
	}

	if err = s.delete(tx, role); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, rollErr.Error())
		}
		return errors.WithStack(err)
	}

	if err = tx.Commit(); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, rollErr.Error())
		}
		return errors.WithStack(err)
	}

	return nil
}

func (s *roleStore) delete(tx *sqlx.Tx, role string) error {
	_, err := tx.Exec(tx.Rebind("DELETE FROM role WHERE name=?"), role)
	return errors.WithStack(err)
}

// Exist get a role from the storage if it exist.
func (s *roleStore) Exist(role string) (bool, error) {
	query := s.db.Rebind("select count(1) as total from role where name=?")
	rows, err := s.db.Query(query, role)
	if err != nil {
		return false, errors.WithStack(err)
	}
	defer rows.Close()

	var total int
	if rows.Next() {
		if err := rows.Scan(&total); err != nil {
			return false, errors.WithStack(err)
		}
	}

	return total > 0, nil
}

// Get gets a role from the storage.
func (s *roleStore) Get(id int64) (*Role, error) {
	query := s.db.Rebind("select name, description from role where id = ?")

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	if rows.Next() {
		var s sql.NullString
		var d sql.NullString
		if err := rows.Scan(&s, &d); err != nil {
			return nil, errors.WithStack(err)
		}

		role := &Role{ID: id}
		if s.Valid {
			role.Name = s.String
		}
		if d.Valid {
			role.Description = d.String
		}

		return role, nil
	}

	return nil, sql.ErrNoRows
}

// GetByName gets a role by role' name from the storage.
func (s *roleStore) GetByName(name string) (*Role, error) {
	query := s.db.Rebind("select id, description from role where name = ?")

	rows, err := s.db.Query(query, name)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	if rows.Next() {
		var id int64
		var d sql.NullString
		if err := rows.Scan(&id, &d); err != nil {
			return nil, errors.WithStack(err)
		}

		role := &Role{ID: id, Name: name}
		if d.Valid {
			role.Description = d.String
		}

		return role, nil
	}

	return nil, sql.ErrNoRows
}

// GetAll gets roles from the storage.
func (s *roleStore) GetAll(limit, offset int64, conditions ...string) ([]Role, error) {
	var query string
	var args []interface{}
	if len(conditions) == 0 || conditions[0] == "" {
		query = s.db.Rebind("select id, name, description from role limit ? offset ?")
	} else {
		query = s.db.Rebind("select id, name, description from role where name like ? limit ? offset ?")
		args = append(args, "%"+conditions[0]+"%")
	}
	args = append(args, limit, offset)

	rows, err := s.db.Query(query, args...)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	roles := []Role{}
	for rows.Next() {
		var id int64
		var s sql.NullString
		var d sql.NullString
		if err := rows.Scan(&id, &s, &d); err != nil {
			return nil, errors.WithStack(err)
		}

		r := Role{ID: id}
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
func (s *roleStore) Count(conditions ...string) (int64, error) {
	var rows *sql.Rows
	var err error
	if len(conditions) == 0 || conditions[0] == "" {
		rows, err = s.db.Query(s.db.Rebind("select count(1) as count from role"))
	} else {
		rows, err = s.db.Query(s.db.Rebind("select count(1) as count from role where name like ?"), "%"+conditions[0]+"%")
	}

	if err != nil {
		return 0, errors.WithStack(err)
	}
	defer rows.Close()

	var count sql.NullInt64
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, errors.WithStack(err)
		}
	}

	if !count.Valid {
		return 0, nil
	}

	return count.Int64, nil
}
