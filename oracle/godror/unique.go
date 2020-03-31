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

func main() {
	db, err := sql.Open("godror", `c##admin/password@127.0.0.1:1521/ORCLCDB`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	a := &Account{
		Subject: "xxx",
		Email:   "luoji@gmail.com",
	}
	//err = InsertAccount(db, a)
	err = InsertAccountByFields(db, a)
	if err != nil {
		panic(err)
	}
	log.Println("ok:", a.ID, a.CreatedDate, a.ChangedDate, a.DeletedDate)
}

// Account represents a row from '"C##ADMIN"."account"'.
type Account struct {
	ID          int          `json:"id"`           // id
	Subject     string       `json:"subject"`      // subject
	Email       string       `json:"email"`        // email
	CreatedDate sql.NullTime `json:"created_date"` // created_date
	ChangedDate sql.NullTime `json:"changed_date"` // changed_date
	DeletedDate sql.NullTime `json:"deleted_date"` // deleted_date

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
	const sqlstr = `INSERT INTO "C##ADMIN"."account" (` +
		`"subject", "email", "created_date", "changed_date", "deleted_date"` +
		`) VALUES (` +
		`:1, :2, :3, :4, :5` +
		`) RETURNING "id" INTO :6`

	// run query
	log.Println(sqlstr, a.Subject, a.Email, a.CreatedDate, a.ChangedDate, a.DeletedDate)
	_, err = db.Exec(sqlstr, a.Subject, a.Email, a.CreatedDate, a.ChangedDate, a.DeletedDate, sql.Out{Dest: &a.ID})
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

	sqlstr := `INSERT INTO "C##ADMIN"."account" (` +
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
