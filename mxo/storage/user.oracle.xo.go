// Package storage contains the types for schema.
package storage

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// InsertUser inserts the User to the database.
func (s *GodrorStorage) InsertUser(db XODB, u *User) error {
	var err error

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO "AC"."user" (` +
		`"subject", "name", "created_date", "changed_date", "deleted_date"` +
		`) VALUES (` +
		`:1, :2, :3, :4, :5` +
		`) RETURNING "id" INTO :6`

	// run query
	s.Logger.Info(sqlstr, u.Subject, u.Name, u.CreatedDate, u.ChangedDate, u.DeletedDate)
	_, err = db.Exec(sqlstr, RealOracleEmptyString(u.Subject), RealOracleNullString(u.Name), u.CreatedDate, u.ChangedDate, u.DeletedDate, sql.Out{Dest: &u.ID})
	if err != nil {
		return err
	}

	return nil
}

// InsertUserByFields inserts the User to the database.
func (s *GodrorStorage) InsertUserByFields(db XODB, u *User) error {
	var err error

	params := make([]interface{}, 0, 5)
	fields := make([]string, 0, 5)
	retFields := make([]string, 0, 5)
	retFields = append(retFields, `"id"`)
	retVars := make([]interface{}, 0, 5)
	retVars = append(retVars, sql.Out{Dest: &u.ID})

	fields = append(fields, `"subject"`)
	params = append(params, RealOracleEmptyString(u.Subject))

	if u.Name.Valid {
		fields = append(fields, `"name"`)
		params = append(params, RealOracleNullString(u.Name))

	} else {
		retFields = append(retFields, `"name"`)
		retVars = append(retVars, sql.Out{Dest: &u.Name.String})

	}
	if u.CreatedDate.Valid {
		fields = append(fields, `"created_date"`)
		params = append(params, u.CreatedDate)

	} else {
		retFields = append(retFields, `"created_date"`)
		retVars = append(retVars, sql.Out{Dest: &u.CreatedDate})

	}
	if u.ChangedDate.Valid {
		fields = append(fields, `"changed_date"`)
		params = append(params, u.ChangedDate)

	} else {
		retFields = append(retFields, `"changed_date"`)
		retVars = append(retVars, sql.Out{Dest: &u.ChangedDate})

	}
	if u.DeletedDate.Valid {
		fields = append(fields, `"deleted_date"`)
		params = append(params, u.DeletedDate)

	} else {
		retFields = append(retFields, `"deleted_date"`)
		retVars = append(retVars, sql.Out{Dest: &u.DeletedDate})

	}

	if len(params) == 0 {
		// FIXME(jackie): maybe we should allow this?
		return errors.New("all fields are empty, unable to insert")
	}
	params = append(params, retVars...)

	var fieldsPlaceHolders []string
	for i := range fields {
		fieldsPlaceHolders = append(fieldsPlaceHolders, ":"+strconv.Itoa(i+1))
	}
	var retFieldsPlaceHolders []string
	for i := range retFields {
		retFieldsPlaceHolders = append(retFieldsPlaceHolders, ":"+strconv.Itoa(len(fieldsPlaceHolders)+i+1))
	}

	sqlstr := `INSERT INTO "AC"."user" (` +
		strings.Join(fields, ",") +
		`) VALUES (` + strings.Join(fieldsPlaceHolders, ",") +
		`) RETURNING ` + strings.Join(retFields, ",") +
		` INTO ` + strings.Join(retFieldsPlaceHolders, ",")

	// run query
	s.Logger.Info(sqlstr, params)
	_, err = db.Exec(sqlstr, params...)
	if err != nil {
		return err
	}
	FixRealOracleEmptyString(&u.Subject)

	if !u.Name.Valid {
		FixRealOracleNullString(&u.Name)
	}

	return nil
}

// UpdateUser updates the User in the database.
func (s *GodrorStorage) UpdateUser(db XODB, u *User) error {
	var err error

	// sql query
	const sqlstr = `UPDATE "AC"."user" SET ` +
		`"subject" = :1, "name" = :2, "created_date" = :3, "changed_date" = :4, "deleted_date" = :5` +
		` WHERE "id" = :6`

	// run query
	s.Logger.Info(sqlstr, u.Subject, u.Name, u.CreatedDate, u.ChangedDate, u.DeletedDate, u.ID)
	_, err = db.Exec(sqlstr, RealOracleEmptyString(u.Subject), RealOracleNullString(u.Name), u.CreatedDate, u.ChangedDate, u.DeletedDate, u.ID)
	return err
}

// UpdateUserByFields updates the User in the database.
func (s *GodrorStorage) UpdateUserByFields(db XODB, u *User, fields, retCols []string, params, retVars []interface{}) error {
	if len(fields) == 0 {
		return nil
	}

	if len(fields) != len(params) {
		return errors.New("fields length is not equal params length")
	}

	if len(retCols) != len(retVars) {
		return errors.New("retCols length is not equal retVars length")
	}

	var setstr string
	var idxvals []interface{}
	var oparams []interface{}
	for i, field := range fields {
		if i != 0 {
			setstr += ", "
		}
		setstr += field + ` = :%d`
		idxvals = append(idxvals, i+1)
		switch v := (params[i]).(type) {
		case string:
			oparams = append(oparams, RealOracleEmptyString(v))
		case sql.NullString:
			oparams = append(oparams, RealOracleNullString(v))
		default:
			oparams = append(oparams, v)
		}

	}
	id := u.ID

	oparams = append(oparams, id)
	idxvals = append(idxvals, len(oparams))
	var sqlstr = fmt.Sprintf(`UPDATE "AC"."user" SET `+setstr+` WHERE "id" = :%d`, idxvals...)
	s.Logger.Info(sqlstr, params, id)
	if _, err := db.Exec(sqlstr, oparams...); err != nil {
		return err
	}

	if len(retCols) > 0 {
		err := db.QueryRow(`SELECT `+strings.Join(retCols, ",")+` from "AC"."user" WHERE "id" = :1`, id).Scan(retVars...)
		if err != nil {
			return err
		}
		for _, val := range retVars {
			switch v := val.(type) {
			case *string:
				FixRealOracleEmptyString(v)
			case *sql.NullString:
				FixRealOracleNullString(v)
			}
		}
	}

	return nil
}

// SaveUser saves the User to the database.
func (s *GodrorStorage) SaveUser(db XODB, u *User) error {

	return s.InsertUser(db, u)
}

// UpsertUser performs an upsert for User.
func (s *GodrorStorage) UpsertUser(db XODB, u *User) error {
	var err error

	// sql query

	const sqlstr = `MERGE INTO "AC"."user" t ` +
		`USING (SELECT :1 AS "id", :2 AS "subject", :3 AS "name", :4 AS "created_date", :5 AS "changed_date", :6 AS "deleted_date" FROM dual) s ` +
		`ON (t."id" = s."id") ` +
		`WHEN MATCHED THEN UPDATE SET "subject" = s."subject", "name" = s."name", "created_date" = s."created_date", "changed_date" = s."changed_date", "deleted_date" = s."deleted_date" ` +
		`WHEN NOT MATCHED THEN INSERT ("subject", "name", "created_date", "changed_date", "deleted_date") VALUES (s."subject", s."name", s."created_date", s."changed_date", s."deleted_date")`

	// run query
	s.Logger.Info(sqlstr, u.ID, u.Subject, u.Name, u.CreatedDate, u.ChangedDate, u.DeletedDate)
	_, err = db.Exec(sqlstr, u.ID, RealOracleEmptyString(u.Subject), RealOracleNullString(u.Name), u.CreatedDate, u.ChangedDate, u.DeletedDate)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes the User from the database.
func (s *GodrorStorage) DeleteUser(db XODB, u *User) error {
	var err error

	// sql query
	const sqlstr = `DELETE FROM "AC"."user" WHERE "id" = :1`

	// run query
	s.Logger.Info(sqlstr, u.ID)
	_, err = db.Exec(sqlstr, u.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUsers deletes the User from the database.
func (s *GodrorStorage) DeleteUsers(db XODB, us []*User) error {
	var err error

	if len(us) == 0 {
		return nil
	}

	var args []interface{}
	var placeholder string
	for i, u := range us {
		args = append(args, u.ID)
		if i != 0 {
			placeholder = placeholder + ", "
		}
		placeholder += fmt.Sprintf(":%d", i+1)
	}

	// sql query
	var sqlstr = `DELETE FROM "AC"."user" WHERE "id" in (` + placeholder + `)`

	// run query
	s.Logger.Info(sqlstr, args)
	_, err = db.Exec(sqlstr, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetMostRecentUser returns n most recent rows from 'user',
// ordered by "created_date" in descending order.
func (s *GodrorStorage) GetMostRecentUser(db XODB, n int) ([]*User, error) {
	const sqlstr = `SELECT ` +
		`"id", "subject", "name", "created_date", "changed_date", "deleted_date" ` +
		`FROM "AC"."user" ` +
		`ORDER BY "created_date" DESC FETCH NEXT :1 ROWS ONLY`

	s.Logger.Info(sqlstr, n)
	q, err := db.Query(sqlstr, n)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*User
	for q.Next() {
		u := User{}

		// scan
		err = q.Scan(&u.ID, &u.Subject, &u.Name, &u.CreatedDate, &u.ChangedDate, &u.DeletedDate)
		if err != nil {
			return nil, err
		}
		FixRealOracleEmptyString(&u.Subject)

		FixRealOracleNullString(&u.Name)

		res = append(res, &u)
	}

	return res, nil
}

// GetMostRecentChangedUser returns n most recent rows from 'user',
// ordered by "changed_date" in descending order.
func (s *GodrorStorage) GetMostRecentChangedUser(db XODB, n int) ([]*User, error) {
	const sqlstr = `SELECT ` +
		`"id", "subject", "name", "created_date", "changed_date", "deleted_date" ` +
		`FROM "AC"."user" ` +
		`ORDER BY "changed_date" DESC FETCH NEXT :1 ROWS ONLY`

	s.Logger.Info(sqlstr, n)
	q, err := db.Query(sqlstr, n)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*User
	for q.Next() {
		u := User{}

		// scan
		err = q.Scan(&u.ID, &u.Subject, &u.Name, &u.CreatedDate, &u.ChangedDate, &u.DeletedDate)
		if err != nil {
			return nil, err
		}
		FixRealOracleEmptyString(&u.Subject)

		FixRealOracleNullString(&u.Name)

		res = append(res, &u)
	}

	return res, nil
}

// GetAllUser returns all rows from 'user', based on the UserQueryArguments.
// If the UserQueryArguments is nil, it will use the default UserQueryArguments instead.
func (s *GodrorStorage) GetAllUser(db XODB, queryArgs *UserQueryArguments) ([]*User, error) { // nolint: gocyclo
	queryArgs = ApplyUserQueryArgsDefaults(queryArgs)

	desc := ""
	if *queryArgs.Desc {
		desc = "DESC"
	}

	dead := "NULL"
	if *queryArgs.Dead {
		dead = "NOT NULL"
	}

	orderBy := "id"
	foundIndex := false
	dbFields := map[string]bool{
		"id":           true,
		"subject":      true,
		"name":         true,
		"created_date": true,
		"changed_date": true,
		"deleted_date": true,
	}

	if *queryArgs.OrderBy != "" && *queryArgs.OrderBy != defaultOrderBy {
		foundIndex = dbFields[*queryArgs.OrderBy]
		if !foundIndex {
			return nil, fmt.Errorf("unable to order by %s, field not found", *queryArgs.OrderBy)
		}
		orderBy = *queryArgs.OrderBy
	}

	var params []interface{}
	placeHolders := ""
	params = append(params, *queryArgs.Offset)
	offsetPos := len(params)

	params = append(params, *queryArgs.Limit)
	limitPos := len(params)

	var sqlstr = fmt.Sprintf(`SELECT %s FROM %s WHERE %s "deleted_date" IS %s ORDER BY "%s" %s OFFSET :%d ROWS FETCH NEXT :%d ROWS ONLY`,
		`"id", "subject", "name", "created_date", "changed_date", "deleted_date" `,
		`"AC"."user"`,
		placeHolders,
		dead,
		orderBy,
		desc,
		offsetPos,
		limitPos)
	s.Logger.Info(sqlstr, params)

	q, err := db.Query(sqlstr, params...)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*User
	for q.Next() {
		u := User{}

		// scan
		err = q.Scan(&u.ID, &u.Subject, &u.Name, &u.CreatedDate, &u.ChangedDate, &u.DeletedDate)
		if err != nil {
			return nil, err
		}
		FixRealOracleEmptyString(&u.Subject)

		FixRealOracleNullString(&u.Name)

		res = append(res, &u)
	}

	return res, nil
}

// CountAllUser returns a count of all rows from 'user'
func (s *GodrorStorage) CountAllUser(db XODB, queryArgs *UserQueryArguments) (int, error) {
	queryArgs = ApplyUserQueryArgsDefaults(queryArgs)

	dead := "NULL"
	if *queryArgs.Dead {
		dead = "NOT NULL"
	}

	var params []interface{}
	placeHolders := ""

	var err error
	var sqlstr = fmt.Sprintf(`SELECT count(*) from "AC"."user" WHERE %s "deleted_date" IS %s`, placeHolders, dead)
	s.Logger.Info(sqlstr)

	var count int
	err = db.QueryRow(sqlstr, params...).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

// UsersBySubjectFK retrieves rows from "AC"."user" by foreign key Subject.
// Generated from foreign key Account.
func (s *GodrorStorage) UsersBySubjectFK(db XODB, subject string, queryArgs *UserQueryArguments) ([]*User, error) {
	queryArgs = ApplyUserQueryArgsDefaults(queryArgs)

	desc := ""
	if *queryArgs.Desc {
		desc = "DESC"
	}

	dead := "NULL"
	if *queryArgs.Dead {
		dead = "NOT NULL"
	}

	var params []interface{}
	placeHolders := ""
	params = append(params, subject)
	placeHolders = fmt.Sprintf(`%s "subject" = :%d AND `, placeHolders, len(params))

	params = append(params, *queryArgs.Offset)
	offsetPos := len(params)

	params = append(params, *queryArgs.Limit)
	limitPos := len(params)

	var sqlstr = fmt.Sprintf(
		`SELECT %s FROM %s WHERE %s "deleted_date" IS %s ORDER BY "%s" %s OFFSET :%d ROWS FETCH NEXT :%d ROWS ONLY`,
		`"id", "subject", "name", "created_date", "changed_date", "deleted_date" `,
		`"AC"."user"`,
		placeHolders,
		dead,
		"id",
		desc,
		offsetPos,
		limitPos)

	s.Logger.Info(sqlstr, params)
	q, err := db.Query(sqlstr, params...)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*User
	for q.Next() {
		u := User{}

		// scan
		err = q.Scan(&u.ID, &u.Subject, &u.Name, &u.CreatedDate, &u.ChangedDate, &u.DeletedDate)
		if err != nil {
			return nil, err
		}
		FixRealOracleEmptyString(&u.Subject)

		FixRealOracleNullString(&u.Name)

		res = append(res, &u)
	}

	return res, nil
}

// CountUsersBySubjectFK count rows from "AC"."user" by foreign key Subject.
// Generated from foreign key Account.
func (s *GodrorStorage) CountUsersBySubjectFK(db XODB, subject string, queryArgs *UserQueryArguments) (int, error) {
	queryArgs = ApplyUserQueryArgsDefaults(queryArgs)

	dead := "NULL"
	if *queryArgs.Dead {
		dead = "NOT NULL"
	}

	var params []interface{}
	placeHolders := ""
	params = append(params, subject)
	placeHolders = fmt.Sprintf(`%s "subject" = :%d AND `, placeHolders, len(params))

	var err error
	var sqlstr = fmt.Sprintf(`SELECT count(*) from "AC"."user" WHERE %s "deleted_date" IS %s`, placeHolders, dead)
	s.Logger.Info(sqlstr)

	var count int
	err = db.QueryRow(sqlstr, params...).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

// AccountInUser returns the Account associated with the User's Subject (subject).
//
// Generated from foreign key 'user_account_subject_fk'.
func (s *GodrorStorage) AccountInUser(db XODB, u *User) (*Account, error) {
	return s.AccountBySubject(db, u.Subject)
}

// UserByID retrieves a row from '"AC"."user"' as a User.
//
// Generated from index 'USER_PK'.
func (s *GodrorStorage) UserByID(db XODB, id int) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`"id", "subject", "name", "created_date", "changed_date", "deleted_date" ` +
		`FROM "AC"."user" ` +
		`WHERE "id" = :1`

	// run query
	s.Logger.Info(sqlstr, id)
	u := User{}

	err = db.QueryRow(sqlstr, id).Scan(&u.ID, &u.Subject, &u.Name, &u.CreatedDate, &u.ChangedDate, &u.DeletedDate)
	if err != nil {
		return nil, err
	}
	FixRealOracleEmptyString(&u.Subject)

	FixRealOracleNullString(&u.Name)

	return &u, nil
}
