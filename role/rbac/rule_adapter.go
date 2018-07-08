package rbac

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Rule the storage structure of the rule
type Rule struct {
	ID    int64
	PType string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
	V5    string
}

// ruleAdapter represents the sqlx adapter for rule storage.
type ruleAdapter struct{}

func newRuleAdapter() *ruleAdapter {
	return &ruleAdapter{}
}

// Insert adds a rule to the storage.
func (a *ruleAdapter) Insert(db XODB, ptype string, rule []string) error {
	params := []interface{}{ptype}
	for i := range rule {
		params = append(params, rule[i])
	}

	if len(params) >= 7 {
		params = params[:7]
	} else {
		for i := len(params); i < 7; i++ {
			params = append(params, "")
		}
	}

	// TODO
	//query := Migrations[a.database].QueryInsertRule
	query := "INSERT INTO rule (p_type, v0, v1, v2, v3, v4, v5) SELECT $1::varchar, $2::varchar, $3::varchar, $4::varchar, $5::varchar, $6::varchar, $7::varchar WHERE NOT EXISTS (SELECT 1 FROM rule WHERE p_type = $1 and v0 = $2 and v1 = $3 and v2 = $4 and v3 = $5 and v4 = $6 and v5 = $7)"
	if _, err := db.Exec(db.Rebind(query), params...); err != nil {
		return errors.Wrap(err, "insert rule")
	}

	return nil
}

// Delete removes a rule from the storage.
func (a *ruleAdapter) Delete(db XODB, ptype string, rule []string) error {
	params := []interface{}{ptype}
	for i := range rule {
		params = append(params, rule[i])
	}

	if len(params) >= 7 {
		params = params[:7]
	} else {
		for i := len(params); i < 7; i++ {
			params = append(params, "")
		}
	}

	_, err := db.Exec(db.Rebind("DELETE FROM rule WHERE p_type=? and v0=? and v1=? and v2=? and v3=? and v4=? and v5=?"), params...)
	return errors.Wrap(err, "delete rule")
}

// RemoveFiltered removes rules that match the filter from the storage.
func (a *ruleAdapter) RemoveFiltered(db XODB, ptype string, fieldIndex int, fieldValues ...string) error {
	query, args := deleteSQL(ptype, fieldIndex, fieldValues...)
	if _, err := db.Exec(db.Rebind(query), args...); err != nil {
		return errors.Wrap(err, "delete filtered rule")
	}

	return nil
}

func (a *ruleAdapter) dropAllData(db XODB) error {
	_, err := db.Exec(db.Rebind("DELETE FROM rule"))
	return errors.Wrap(err, "delete all rule")
}

// Restore saves rule from model to database.
func (a *ruleAdapter) Restore(db *sqlx.DB, md model) error {
	tx, err := db.Beginx()
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}

	if err = a.dropAllData(tx); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, "rollback")
		}

		return errors.WithStack(err)
	}

	//for ptype, ast := range md["p"] {
	//	for _, rule := range ast.Rule {
	//		err = a.create(tx, ptype, rule)
	//		if err != nil {
	//			if err = tx.Rollback(); err != nil {
	//				return errors.WithStack(err)
	//			}
	//			return errors.WithStack(err)
	//		}
	//	}
	//}

	for ptype, ast := range md["g"] {
		for _, rule := range ast.Rule {
			err = a.Insert(tx, ptype, rule)
			if err != nil {
				if rollErr := tx.Rollback(); rollErr != nil {
					return errors.Wrap(err, "rollback")
				}
				return errors.WithStack(err)
			}
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

// Load loads rule from database to model.
func (a *ruleAdapter) Load(db XODB, md model) error {
	var (
		limit  int64 = 1000
		offset int64
	)

	for {
		lines, err := a.GetAll(db, limit, offset)
		if err != nil {
			return err
		}

		for _, line := range lines {
			a.loadRuleLine(line, md)
		}

		if len(lines) < int(limit) {
			break
		}

		offset += limit
	}

	return nil
}

func (a *ruleAdapter) loadRuleLine(line *Rule, md model) {
	lineText := line.PType
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}

	loadRuleLine(lineText, md)
}

// Has determines whether a model has the specified rule.
func (a *ruleAdapter) Has(db XODB, ptype string, rule []string) (bool, error) {
	var count int64

	query, args := hasSQL(ptype, 0, rule...)
	err := db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "query rule")
	}

	return count > 0, nil
}

// Get gets a rule from the storage.
func (a *ruleAdapter) Get(db XODB, id int64) (*Rule, error) {
	var p, v0, v1, v2, v3, v4, v5 sql.NullString

	query := db.Rebind("SELECT p_type, v0, v1, v2, v3, v4, v5 FROM rule WHERE id = ?")
	err := db.QueryRow(query, id).Scan(&p, &v0, &v1, &v2, &v3, &v4, &v5)
	if err != nil {
		return nil, errors.Wrap(err, "get rule")
	}

	return newRule(id, p, v0, v1, v2, v3, v4, v5), nil
}

// GetFiltered gets rules based on field filters from the database.
func (a *ruleAdapter) GetFiltered(db XODB, ptype string, fieldIndex int, fieldValues ...string) ([]*Rule, error) {
	query, args := querySQL(ptype, fieldIndex, fieldValues...)
	rows, err := db.Query(query, args...)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	rules := []*Rule{}
	for rows.Next() {
		var id int64
		var p, v0, v1, v2, v3, v4, v5 sql.NullString
		if err := rows.Scan(&id, &p, &v0, &v1, &v2, &v3, &v4, &v5); err != nil {
			return nil, errors.WithStack(err)
		}

		rule := newRule(id, p, v0, v1, v2, v3, v4, v5)
		rules = append(rules, rule)
	}

	return rules, nil
}

// GetAll gets rule from the storage.
func (a *ruleAdapter) GetAll(db XODB, limit, offset int64) ([]*Rule, error) {
	query := db.Rebind("SELECT id, p_type, v0, v1, v2, v3, v4, v5 FROM rule LIMIT ? OFFSET ?")
	rows, err := db.Query(query, limit, offset)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	rules := []*Rule{}
	for rows.Next() {
		var id int64
		var p, v0, v1, v2, v3, v4, v5 sql.NullString
		if err := rows.Scan(&id, &p, &v0, &v1, &v2, &v3, &v4, &v5); err != nil {
			return nil, errors.WithStack(err)
		}

		rule := newRule(id, p, v0, v1, v2, v3, v4, v5)
		rules = append(rules, rule)
	}

	return rules, nil
}

// Count gets total number of rule from the storage.
func (a *ruleAdapter) Count(db XODB) (int64, error) {
	var count int64

	query := db.Rebind("SELECT COUNT(id) as count FROM rule")
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return count, nil
}

func newRule(id int64, p, v0, v1, v2, v3, v4, v5 sql.NullString) *Rule {
	r := &Rule{ID: id}
	if p.Valid {
		r.PType = p.String
	}

	if v0.Valid {
		r.V0 = v0.String
	}

	if v1.Valid {
		r.V1 = v1.String
	}

	if v2.Valid {
		r.V2 = v2.String
	}

	if v3.Valid {
		r.V3 = v3.String
	}

	if v4.Valid {
		r.V4 = v4.String
	}

	if v5.Valid {
		r.V5 = v5.String
	}

	return r
}

func deleteSQL(ptype string, fieldIndex int, fieldValues ...string) (string, []interface{}) {
	where, args := whereSQL(ptype, fieldIndex, fieldValues...)
	query := "DELETE FROM rule " + where
	return query, args
}

func querySQL(ptype string, fieldIndex int, fieldValues ...string) (string, []interface{}) {
	where, args := whereSQL(ptype, fieldIndex, fieldValues...)
	query := "SELECT id, p_type, v0, v1, v2, v3, v4, v5 FROM rule " + where
	return query, args
}

func hasSQL(ptype string, fieldIndex int, fieldValues ...string) (string, []interface{}) {
	where, args := whereSQL(ptype, fieldIndex, fieldValues...)
	query := "SELECT COUNT(id) as count FROM rule " + where
	return query, args
}

func whereSQL(ptype string, fieldIndex int, fieldValues ...string) (string, []interface{}) {
	params := make([]string, 6)

	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		params[0] = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		params[1] = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		params[2] = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		params[3] = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		params[4] = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		params[5] = fieldValues[5-fieldIndex]
	}

	count := 1
	queryArgs := []interface{}{ptype}
	queryStr := " WHERE p_type = $1"
	if params[0] != "" {
		count++
		queryStr += fmt.Sprintf(" and v0 = $%d", count)
		queryArgs = append(queryArgs, params[0])
	}
	if params[1] != "" {
		count++
		queryStr += fmt.Sprintf(" and v1 = $%d", count)
		queryArgs = append(queryArgs, params[1])
	}
	if params[2] != "" {
		count++
		queryStr += fmt.Sprintf(" and v2 = $%d", count)
		queryArgs = append(queryArgs, params[2])
	}
	if params[3] != "" {
		count++
		queryStr += fmt.Sprintf(" and v3 = $%d", count)
		queryArgs = append(queryArgs, params[3])
	}
	if params[4] != "" {
		count++
		queryStr += fmt.Sprintf(" and v4 = $%d", count)
		queryArgs = append(queryArgs, params[4])
	}
	if params[5] != "" {
		count++
		queryStr += fmt.Sprintf(" and v5 = $%d", count)
		queryArgs = append(queryArgs, params[5])
	}

	return queryStr, queryArgs
}
