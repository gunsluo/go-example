package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/godror/godror"
)

const (
	ORALCE_EMPTY_STRING = "_xxxx"
)

func main() {
	db, err := sql.Open("godror", `oracle://ac:password@127.0.0.1:1521/ORCLPDB1`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	a := &Account{
		Subject: "xxx",
		Email:   "luoji@gmail.com",
		//Name:    "abc",
		//Label:   sql.NullString{String: "bcd", Valid: true},
		Name: "",
		//Label: sql.NullString{String: "", Valid: true},
		Label: sql.NullString{Valid: false},
	}
	err = InsertAccount(db, a)
	//err = InsertAccountByFields(db, a)
	if err != nil {
		panic(err)
	}
	log.Println("ok:", a.ID, a.CreatedDate, a.ChangedDate, a.DeletedDate)

	na, err := AccountByID(db, a.ID)
	if err != nil {
		panic(err)
	}
	log.Println("ok:", na.ID, na.Name, na.Label.String, na.Label.Valid, na.CreatedDate, na.ChangedDate, na.DeletedDate)
}

// Account represents a row from '"C##ADMIN"."account"'.
type Account struct {
	ID          int            `json:"id"`           // id
	Subject     string         `json:"subject"`      // subject
	Email       string         `json:"email"`        // email
	Name        string         `json:"name"`         // name
	Label       sql.NullString `json:"label"`        // label
	CreatedDate sql.NullTime   `json:"created_date"` // created_date
	ChangedDate sql.NullTime   `json:"changed_date"` // changed_date
	DeletedDate sql.NullTime   `json:"deleted_date"` // deleted_date

	// xo fields
	_exists, _deleted bool
}

// InsertAccount inserts the Account to the database.
func InsertAccount(db *sql.DB, a *Account) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO "account" (` +
		`"subject", "email", "name", "label", "created_date", "changed_date", "deleted_date"` +
		`) VALUES (` +
		`:1, :2, :3, :4, :5, :6, :7` +
		`) RETURNING "id" INTO :8`

	// run query
	log.Println(sqlstr, a.Subject, a.Email, a.Name, a.Label, a.CreatedDate, a.ChangedDate, a.DeletedDate)
	_, err = db.Exec(sqlstr, convertOracleEmptyString(a.Subject), convertOracleEmptyString(a.Email),
		convertOracleEmptyString(a.Name), convertOracleNullString(a.Label), a.CreatedDate, a.ChangedDate, a.DeletedDate, sql.Out{Dest: &a.ID})
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// InsertAccountByFields inserts the Account to the database.
func InsertAccountByFields(db *sql.DB, a *Account) error {
	var err error

	params := make([]interface{}, 0, 5)
	fields := make([]string, 0, 5)
	retFields := make([]string, 0, 5)
	retFields = append(retFields, `"id"`)
	retVars := make([]interface{}, 0, 5)
	retVars = append(retVars, sql.Out{Dest: &a.ID})

	fields = append(fields, `"subject"`)
	params = append(params, a.Subject)

	fields = append(fields, `"email"`)
	params = append(params, a.Email)
	if a.CreatedDate.Valid {
		fields = append(fields, `"created_date"`)
		params = append(params, a.CreatedDate)
	} else {
		retFields = append(retFields, `"created_date"`)
		retVars = append(retVars, sql.Out{Dest: &a.CreatedDate})
	}
	if a.ChangedDate.Valid {
		fields = append(fields, `"changed_date"`)
		params = append(params, a.ChangedDate)
	} else {
		retFields = append(retFields, `"changed_date"`)
		retVars = append(retVars, sql.Out{Dest: &a.ChangedDate})
	}
	if a.DeletedDate.Valid {
		fields = append(fields, `"deleted_date"`)
		params = append(params, a.DeletedDate)
	} else {
		retFields = append(retFields, `"deleted_date"`)
		retVars = append(retVars, sql.Out{Dest: &a.DeletedDate})
	}
	if len(params) == 0 {
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

	sqlstr := `INSERT INTO "account" (` +
		strings.Join(fields, ",") +
		`) VALUES (` + strings.Join(fieldsPlaceHolders, ",") +
		`) RETURNING ` + strings.Join(retFields, ",") +
		` INTO ` + strings.Join(retFieldsPlaceHolders, ",")

	// run query
	log.Println(sqlstr, params)
	_, err = db.Exec(sqlstr, params...)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

func AccountByID(db *sql.DB, id int) (*Account, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`"id", "subject", "email", "name", "label", "created_date", "changed_date", "deleted_date" ` +
		`FROM "AC"."account" ` +
		`WHERE "id" = :1`

	// run query
	log.Println(sqlstr, id)
	a := Account{
		_exists: true,
	}

	//var subject, email, name string
	//var label sql.NullString

	var subject, email, name OracleString
	var label OracleNullString
	err = db.QueryRow(sqlstr, id).Scan(&a.ID, &subject, &email, &name, &label, &a.CreatedDate, &a.ChangedDate, &a.DeletedDate)
	if err != nil {
		return nil, err
	}

	a.Subject = string(subject)
	a.Email = string(email)
	a.Name = string(name)
	a.Label = label.NullString()
	//a.Label.String, a.Label.Valid = label.String, label.Valid

	return &a, nil
}

type OracleString string

func (s *OracleString) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		if v == ORALCE_EMPTY_STRING {
			*s = ""
		} else {
			*s = OracleString(v)
		}
	default:
	}

	return nil
}

type OracleNullString string

func (s *OracleNullString) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*s = OracleNullString(v)
	default:
	}

	return nil
}

func (s OracleNullString) NullString() sql.NullString {
	if s == "" {
		return sql.NullString{}
	}

	if s == ORALCE_EMPTY_STRING {
		return sql.NullString{Valid: true}
	}

	return sql.NullString{String: string(s), Valid: true}
}

func convertOracleEmptyString(s string) string {
	if s == "" {
		return ORALCE_EMPTY_STRING
	}
	return s
}

func convertOracleNullString(ns sql.NullString) sql.NullString {
	if !ns.Valid {
		return sql.NullString{}
	}

	if ns.String == "" {
		return sql.NullString{String: ORALCE_EMPTY_STRING, Valid: true}
	}

	return sql.NullString{String: ns.String, Valid: true}
}
