// Package storage contains the types for schema.
package storage

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

// Storage is interface structure for database operation that can be called
type Storage interface {
	// InsertAccount inserts the Account to the database.
	InsertAccount(db XODB, a *Account) error
	// InsertAccountByFields inserts the Account to the database.
	InsertAccountByFields(db XODB, a *Account) error
	// DeleteAccount deletes the Account from the database.
	DeleteAccount(db XODB, a *Account) error
	// DeleteAccounts deletes the Account from the database.
	DeleteAccounts(db XODB, a []*Account) error
	// Update updates the Account in the database.
	UpdateAccount(db XODB, a *Account) error
	// UpdateAccountByFields updates the Account in the database.
	UpdateAccountByFields(db XODB, a *Account, fields, retCols []string, params, retVars []interface{}) error
	// Save saves the Account to the database.
	SaveAccount(db XODB, a *Account) error
	// Upsert performs an upsert for Account.
	UpsertAccount(db XODB, a *Account) error
	// GetMostRecentAccount returns n most recent rows from 'account',
	// ordered by "created_date" in descending order.
	GetMostRecentAccount(db XODB, n int) ([]*Account, error)
	// GetMostRecentChangedAccount returns n most recent rows from 'account',
	// ordered by "changed_date" in descending order.
	GetMostRecentChangedAccount(db XODB, n int) ([]*Account, error)
	// GetAllAccount returns all rows from 'account', based on the AccountQueryArguments.
	// If the AccountQueryArguments is nil, it will use the default AccountQueryArguments instead.
	GetAllAccount(db XODB, queryArgs *AccountQueryArguments) ([]*Account, error)
	// CountAllAccount returns a count of all rows from 'account'
	CountAllAccount(db XODB, queryArgs *AccountQueryArguments) (int, error)
	// InsertUser inserts the User to the database.
	InsertUser(db XODB, u *User) error
	// InsertUserByFields inserts the User to the database.
	InsertUserByFields(db XODB, u *User) error
	// DeleteUser deletes the User from the database.
	DeleteUser(db XODB, u *User) error
	// DeleteUsers deletes the User from the database.
	DeleteUsers(db XODB, u []*User) error
	// Update updates the User in the database.
	UpdateUser(db XODB, u *User) error
	// UpdateUserByFields updates the User in the database.
	UpdateUserByFields(db XODB, u *User, fields, retCols []string, params, retVars []interface{}) error
	// Save saves the User to the database.
	SaveUser(db XODB, u *User) error
	// Upsert performs an upsert for User.
	UpsertUser(db XODB, u *User) error
	// GetMostRecentUser returns n most recent rows from 'user',
	// ordered by "created_date" in descending order.
	GetMostRecentUser(db XODB, n int) ([]*User, error)
	// GetMostRecentChangedUser returns n most recent rows from 'user',
	// ordered by "changed_date" in descending order.
	GetMostRecentChangedUser(db XODB, n int) ([]*User, error)
	// GetAllUser returns all rows from 'user', based on the UserQueryArguments.
	// If the UserQueryArguments is nil, it will use the default UserQueryArguments instead.
	GetAllUser(db XODB, queryArgs *UserQueryArguments) ([]*User, error)
	// CountAllUser returns a count of all rows from 'user'
	CountAllUser(db XODB, queryArgs *UserQueryArguments) (int, error)
	// UsersBySubjectFK retrieves rows from user by foreign key Subject.
	// Generated from foreign key Account.
	UsersBySubjectFK(db XODB, subject string, queryArgs *UserQueryArguments) ([]*User, error)
	// CountUsersBySubjectFK count rows from user by foreign key Subject.
	// Generated from foreign key Account.
	CountUsersBySubjectFK(db XODB, subject string, queryArgs *UserQueryArguments) (int, error)
	// AccountInUser returns the Account associated with the User's Subject (subject).
	// Generated from foreign key 'user_account_subject_fk'.
	AccountInUser(db XODB, u *User) (*Account, error)
	// AccountByID retrieves a row from '"public"."account"' as a Account.
	// Generated from index 'account_pk'.
	AccountByID(db XODB, id int) (*Account, error)
	// AccountBySubject retrieves a row from '"public"."account"' as a Account.
	// Generated from index 'account_subject_unique_index'.
	AccountBySubject(db XODB, subject string) (*Account, error)
	// UserByID retrieves a row from '"public"."user"' as a User.
	// Generated from index 'user_pk'.
	UserByID(db XODB, id int) (*User, error)
}

// PostgresStorage is Postgres for the database.
type PostgresStorage struct {
	logger XOLogger
}

func (s *PostgresStorage) info(format string, args ...interface{}) {
	if len(args) == 0 {
		xoLog(s.logger, logrus.InfoLevel, format)
	} else {
		var params []interface{}
		params = append(params, format)
		params = append(params, args...)
		xoLogf(s.logger, logrus.InfoLevel, "%s %v", params...)
	}
}

// MssqlStorage is Mssql for the database.
type MssqlStorage struct {
	logger XOLogger
}

func (s *MssqlStorage) info(format string, args ...interface{}) {
	if len(args) == 0 {
		xoLog(s.logger, logrus.InfoLevel, format)
	} else {
		var params []interface{}
		params = append(params, format)
		params = append(params, args...)
		xoLogf(s.logger, logrus.InfoLevel, "%s %v", params...)
	}
}

// New is a construction method that return a new Storage
func New(driver string, c Config) (Storage, error) {
	// fix bug which interface type is not nil and interface value is nil
	var logger XOLogger
	if c.Logger != nil && !(reflect.ValueOf(c.Logger).Kind() == reflect.Ptr && reflect.ValueOf(c.Logger).IsNil()) {
		logger = c.Logger
	}

	var s Storage
	switch driver {
	case "postgres":
		s = &PostgresStorage{logger: logger}
	case "mssql":
		s = &MssqlStorage{logger: logger}
	default:
		return nil, errors.New("driver " + driver + " not support")
	}

	return s, nil
}

// Account represents a row from '"public"."account"'.
type Account struct {
	ID          int      `json:"id"`           // id
	Subject     string   `json:"subject"`      // subject
	Email       string   `json:"email"`        // email
	CreatedDate NullTime `json:"created_date"` // created_date
	ChangedDate NullTime `json:"changed_date"` // changed_date
	DeletedDate NullTime `json:"deleted_date"` // deleted_date

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Account exists in the database.
func (a *Account) Exists() bool {
	return a._exists
}

// Deleted provides information if the Account has been deleted from the database.
func (a *Account) Deleted() bool {
	return a._deleted
} // User represents a row from '"public"."user"'.
type User struct {
	ID          int            `json:"id"`           // id
	Subject     string         `json:"subject"`      // subject
	Name        sql.NullString `json:"name"`         // name
	CreatedDate NullTime       `json:"created_date"` // created_date
	ChangedDate NullTime       `json:"changed_date"` // changed_date
	DeletedDate NullTime       `json:"deleted_date"` // deleted_date

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted provides information if the User has been deleted from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// extension block
// GraphQL extension

// GraphQL related types
const GraphQLCommonTypes = `
        type PageInfo {
            hasNextPage: Boolean!
            hasPreviousPage: Boolean!
            startCursor: ID
            endCursor: ID
        }
        scalar Time
        enum FilterConjunction{
            AND
            OR
        }
    `

// PageInfoResolver defines the GraphQL PageInfo type
type PageInfoResolver struct {
	startCursor     graphql.ID
	endCursor       graphql.ID
	hasNextPage     bool
	hasPreviousPage bool
}

// StartCursor returns the start cursor (global id)
func (r *PageInfoResolver) StartCursor() *graphql.ID {
	return &r.startCursor
}

// EndCursor returns the end cursor (global id)
func (r *PageInfoResolver) EndCursor() *graphql.ID {
	return &r.endCursor
}

// HasNextPage returns if next page is available
func (r *PageInfoResolver) HasNextPage() bool {
	return r.hasNextPage
}

// HasPreviousPage returns if previous page is available
func (r *PageInfoResolver) HasPreviousPage() bool {
	return r.hasNextPage
}

// ResolverConfig is a config for Resolver
type ResolverConfig struct {
	Logger   XOLogger
	DB       XODB
	S        Storage
	Recorder EventRecorder
	Verifier Verifier
}

// resolverExtensions it's passing between root resolver and  children resolver
type resolverExtensions struct {
	logger   XOLogger
	db       XODB
	storage  Storage
	recorder EventRecorder
	verifier Verifier
}

// RootResolver is a graphql root resolver
type RootResolver struct {
	ext resolverExtensions
}

// NewRootResolver return a root resolver for ggraphql
func NewRootResolver(c *ResolverConfig) *RootResolver {
	logger := c.Logger
	if logger == nil {
		logger = logrus.New()
	}

	return &RootResolver{
		ext: resolverExtensions{
			logger:   logger,
			db:       c.DB,
			storage:  c.S,
			recorder: c.Recorder,
			verifier: c.Verifier,
		},
	}
}

// BuildSchemaString build root schema string
func (r *RootResolver) BuildSchemaString(extraQueries, extraMutations, extraTypes string) string {
	return `
        schema {
            query: Query
            mutation: Mutation
        }

        type Query {
    ` +
		r.GetAccountQueries() +
		r.GetUserQueries() + extraQueries +
		`}

    type Mutation {
    ` +
		r.GetAccountMutations() +
		r.GetUserMutations() + extraMutations +
		`}

    ` +
		r.GetAccountTypes() +
		r.GetUserTypes() +
		GraphQLCommonTypes +
		extraTypes
}

func encodeCursor(typeName string, id int) graphql.ID {
	return graphql.ID(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%d", typeName, id))))
}

// EventRecorder is event recorder
type EventRecorder interface {
	RecordEvent(ctx context.Context, resource, action string, args interface{}) error
}

// Bool returns a nullable bool.
func Bool(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}

// BoolPointer converts bool pointer to sql.NullBool
func BoolPointer(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{}
	}
	return sql.NullBool{Bool: *b, Valid: true}
}

// PointerBool converts bool to pointer to bool
func PointerBool(b sql.NullBool) *bool {
	if !b.Valid {
		return nil
	}
	return &b.Bool
}

// NullDecimalString converts decimal.NullDecimal to *string
func NullDecimalString(b decimal.NullDecimal) *string {
	if !b.Valid {
		return nil
	}
	x := b.Decimal.String()
	return &x
}

// Int64 returns a nullable int64
func Int64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}

// Int64Pointer converts a int64 pointer to sql.NullInt64
func Int64Pointer(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: *i, Valid: true}
}

// PointerInt64 converts sql.NullInt64 to pointer to int64
func PointerInt64(i sql.NullInt64) *int64 {
	if !i.Valid {
		return nil
	}
	return &i.Int64
}

// Float64 returns a nullable float64
func Float64(i float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: i, Valid: true}
}

// Float64Pointer converts a float64 pointer to sql.NullFloat64
func Float64Pointer(i *float64) sql.NullFloat64 {
	if i == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{Float64: *i, Valid: true}
}

// PointerFloat64 converts sql.NullFloat64 to pointer to float64
func PointerFloat64(i sql.NullFloat64) *float64 {
	if !i.Valid {
		return nil
	}
	return &i.Float64
}

// String returns a nullable string
func String(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

// StringPointer converts string pointer to sql.NullString
func StringPointer(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

// PointerString converts sql.NullString to pointer to string
func PointerString(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}

// Time returns a nullable Time
func Time(t time.Time) NullTime {
	return NullTime{Time: t, Valid: true}
}

// TimePointer converts time.Time pointer to NullTime
func TimePointer(t *time.Time) NullTime {
	if t == nil {
		return NullTime{}
	}
	return NullTime{Time: *t, Valid: true}
}

// TimeGqlPointer converts graphql.Time pointer to NullTime
func TimeGqlPointer(t *graphql.Time) NullTime {
	if t == nil {
		return NullTime{}
	}
	return NullTime{Time: t.Time, Valid: true}
}

// PointerTime converts NullTIme to pointer to time.Time
func PointerTime(t NullTime) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

// PointerGqlTime converts NullType to pointer to graphql.Time
func PointerGqlTime(t NullTime) *graphql.Time {
	if !t.Valid {
		return nil
	}
	return &graphql.Time{Time: t.Time}
}

// PointerStringInt64 converts Int64 pointer to string pointer
func PointerStringInt64(i *int64) *string {
	if i == nil {
		return nil
	}
	s := strconv.Itoa(int(*i))
	return &s
}

// PointerStringSqlInt64 converts sql.NullInt64 pointer to graphql.ID pointer
func PointerStringSqlInt64(i sql.NullInt64) *string {
	if !i.Valid {
		return nil
	}
	s := strconv.Itoa(int(i.Int64))
	return &s
}

// PointerStringFloat64 converts Float64 pointer to string pointer
func PointerStringFloat64(i *float64) *string {
	if i == nil {
		return nil
	}
	s := fmt.Sprintf("%.6f", *i)
	return &s
}

// PointerFloat64SqlFloat64 converts sql.NullFloat64 pointer to graphql.ID pointer
func PointerFloat64SqlFloat64(i sql.NullFloat64) *float64 {
	if !i.Valid {
		return nil
	}
	s := i.Float64
	return &s
}

// access control
// Verifier is access control verifier
type Verifier interface {
	VerifyAC(ctx context.Context, resource, action string, args interface{}) error
	VerifyRefAC(ctx context.Context, resource, action string, args interface{}) error
}

// GraphQLResource is a resource of graphql API
type GraphQLResource struct {
	Name     string
	Describe string
}

// GetResolverResources get all resource
func (r *RootResolver) GetResolverResources(includes []GraphQLResource, excludes []string) ([]GraphQLResource, error) {
	uniqueResources := make(map[string]GraphQLResource)
	for _, r := range r.getAccountGraphQLResources() {
		if v, ok := uniqueResources[r.Name]; ok {
			return nil, errors.Errorf("duplicate resource %s", v.Name)
		} else {
			uniqueResources[v.Name] = v
		}
	}
	for _, r := range r.getUserGraphQLResources() {
		if v, ok := uniqueResources[r.Name]; ok {
			return nil, errors.Errorf("duplicate resource %s", v.Name)
		} else {
			uniqueResources[v.Name] = v
		}
	}

	for _, r := range includes {
		if v, ok := uniqueResources[r.Name]; ok {
			return nil, errors.Errorf("duplicate resource %s", v.Name)
		} else {
			uniqueResources[v.Name] = v
		}
	}

	for _, k := range excludes {
		delete(uniqueResources, k)
	}

	var resources []GraphQLResource
	for _, v := range uniqueResources {
		resources = append(resources, v)
	}

	return resources, nil
}