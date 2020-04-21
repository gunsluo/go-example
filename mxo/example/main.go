package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/gunsluo/go-example/mxo/example/graphiql"
	"github.com/gunsluo/go-example/mxo/storage"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/xo/dburl"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/godror/godror"
	_ "github.com/lib/pq"
)

const (
	postgresMode = "postgres"
	mssqlMode    = "mssql"
	oracleMode   = "oracle"
	serverMode   = "server"
)

var dsns = map[string]string{
	postgresMode: "postgres://postgres:password@localhost:5432/xo?sslmode=disable",
	mssqlMode:    "sqlserver://SA:Tes9ting@localhost:1433/instance?database=xo&encrypt=disable",
	oracleMode:   "oracle://ac:password@127.0.0.1:1521/ORCLPDB1",
}

func main() {
	var mode string
	if len(os.Args) <= 1 {
		mode = postgresMode
	} else {
		mode = os.Args[1]
	}

	fmt.Println("run mode:", mode)

	var driver string
	var fn func(string, string)
	switch mode {
	case serverMode:
		if len(os.Args) > 2 {
			driver = os.Args[2]
		} else {
			driver = postgresMode
		}
		fn = server
	case postgresMode:
		driver = mode
		fn = testStoage
	case mssqlMode:
		driver = mode
		fn = testStoage
	case oracleMode:
		driver = mode
		fn = testStoage
	default:
		fmt.Println("invalid parameter, it should be 'postgres' & 'mssql' & 'oracle' & 'server'")
	}

	dsn, ok := dsns[driver]
	if !ok {
		fmt.Println("invalid parameter, it should be 'postgres' & 'mssql' & 'oracle'")
		return
	}

	u, err := dburl.Parse(dsn)
	if err != nil {
		panic(err)
	}

	fn(u.Driver, u.DSN)
}

func testStoage(driver, dsn string) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	s, err := storage.New(driver, storage.Config{Logger: logrus.New()})
	//s, err := storage.NewStorageExtension(driver, storage.Config{Logger: logrus.New()})
	if err != nil {
		panic(err)
	}

	testStoageAndDB(s, db)

	db.Close()
}

func testStoageAndDB(s storage.Storager, db *sqlx.DB) {
	//func testStoageAndDB(s storage.StorageExtension, db *sqlx.DB) {
	account := &storage.Account{
		Subject:     "luoji",
		Email:       "mock",
		Name:        "",
		Label:       sql.NullString{},
		CreatedDate: sql.NullTime{Time: time.Now(), Valid: true},
		ChangedDate: sql.NullTime{Time: time.Now(), Valid: true},
		//DeletedDate: time.Now(),
	}

	err := s.InsertAccount(db, account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("insert account: id=%d, err=%v\n", account.ID, err)

	account.Email = "luoji@gmail.com"
	err = s.UpdateAccount(db, account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("update account: id=%d, err=%v\n", account.ID, err)

	account.Subject = "luoji1"
	err = s.UpsertAccount(db, account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("upsert account: id=%d, err=%v\n", account.ID, err)

	account2 := &storage.Account{
		Subject:     "luoji2",
		Email:       "mock",
		Name:        "",
		Label:       sql.NullString{Valid: true},
		CreatedDate: sql.NullTime{Time: time.Now(), Valid: true},
		ChangedDate: sql.NullTime{Time: time.Now(), Valid: true},
		DeletedDate: sql.NullTime{Time: time.Now(), Valid: true},
	}
	err = s.InsertAccountByFields(db, account2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("insert account by fields: id=%d, err=%v\n", account2.ID, err)

	var fields = []string{`"email"`, `"name"`, `"label"`}
	var retCols = []string{`"name"`, `"label"`}
	var params = []interface{}{"luoji2@gamil.com", "", sql.NullString{}}
	var retVars = []interface{}{&account2.Name, &account2.Label}
	err = s.UpdateAccountByFields(db, account2, fields, retCols, params, retVars)
	if err != nil {
		panic(err)
	}
	fmt.Println("query account: ", account2.ID, account2.Name, account2.Label.String, account2.Label.Valid)

	user := &storage.User{
		Subject:     account.Subject,
		CreatedDate: sql.NullTime{Time: time.Now(), Valid: true},
		ChangedDate: sql.NullTime{Time: time.Now(), Valid: true},
	}

	err = s.InsertUser(db, user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("insert user: id=%d, err=%v\n", user.ID, err)

	user.Name = sql.NullString{Valid: true, String: "luoji"}
	err = s.UpdateUser(db, user)
	if err != nil {
		panic(err)
	}

	user.Name = sql.NullString{Valid: true, String: "luoji2"}
	err = s.UpsertUser(db, user)
	if err != nil {
		panic(err)
	}

	// query
	a, err := s.AccountByID(db, account.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("query account: ", a.ID, a.Name, a.Label.String, a.Label.Valid, a.CreatedDate, a.ChangedDate, a.DeletedDate)

	a, err = s.AccountByID(db, account2.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("query account: ", a.ID, a.Name, a.Label.String, a.Label.Valid, a.CreatedDate, a.ChangedDate, a.DeletedDate)

	a, err = s.AccountBySubject(db, account.Subject)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query account by subject: %v\n", a)

	a, err = s.AccountInUser(db, user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query account in user: %v\n", a)

	as, err := s.GetMostRecentAccount(db, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query account in user: %d %v\n", len(as), as)

	as, err = s.GetMostRecentChangedAccount(db, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query account in user: %d %v\n", len(as), as)

	as, err = s.GetAllAccount(db, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query all account in user: %d %v\n", len(as), as)

	count, err := s.CountAllAccount(db, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("count all account in user: %d\n", count)

	u, err := s.UserByID(db, user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query user: %v\n", u)

	us, err := s.UsersBySubjectFK(db, account.Subject, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query user by subject fk: %s %v\n", account.Subject, len(us))

	total, err := s.CountUsersBySubjectFK(db, account.Subject, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("total user by subject fk: %s %v\n", account.Subject, total)

	err = s.DeleteUser(db, user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("delete user: %v\n", err)

	err = s.DeleteAccount(db, account)
	if err != nil {
		panic(err)
	}

	err = s.DeleteAccount(db, account2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("delete account: %v\n", err)
}

// Server is a graphql server.
type gqlServer struct {
	address string
	engine  *gin.Engine

	db storage.XODB
	s  storage.Storager
}

// NewGQLServer is graphql server
func NewGQLServer(address string, logger *logrus.Logger, db storage.XODB, s storage.Storager) (*gqlServer, error) {

	// graphql API
	rootResolver := storage.NewRootResolver(&storage.ResolverConfig{Logger: logger, DB: db, S: s, Verifier: &Verifier{}})
	schemaString := rootResolver.BuildSchemaString("", "", "")
	fmt.Println("--->", schemaString)

	schema, err := graphql.ParseSchema(schemaString,
		rootResolver,
	)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't parse schema")
	}

	engine := gin.Default()
	graphqlRoute := engine.Group("/graphql")
	graphQLHandler := Handler{
		Schema: schema,
	}
	graphqlRoute.Any("", graphQLHandler.Serve)

	// debug
	debug := engine.Group("/")
	{
		debugPage := bytes.Replace(graphiql.GraphiQLPage, []byte("fetch('/'"), []byte("fetch('/graphql'"), -1)
		debug.GET("/debug.html", func(c *gin.Context) {
			c.Data(http.StatusOK, "text/html; charset=utf-8", debugPage)
		})
	}

	return &gqlServer{
		address: address,
		engine:  engine,
		db:      db,
		s:       s,
	}, nil
}

func (s *gqlServer) Run() error {
	httpServer := http.Server{
		Addr:    s.address,
		Handler: s.engine,
	}

	// listening http server
	fmt.Println("Starting http server listening on:", s.address)
	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func server(driver, dsn string) {
	var logger *logrus.Logger
	logger = logrus.New()

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	s, err := storage.New(driver, storage.Config{Logger: logger})
	if err != nil {
		panic(err)
	}

	server, err := NewGQLServer(":8080", logger, db, s)
	if err != nil {
		panic(err)
	}

	server.Run()
}

// Handler a graphql Handle responds to an HTTP request.
type Handler struct {
	Schema *graphql.Schema
}

// Serve implementation function of graphql handler
func (h *Handler) Serve(c *gin.Context) {
	r := c.Request
	w := c.Writer

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	schema := string(body)
	res, err := h.Query(c, schema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (h *Handler) Query(c *gin.Context, schema string) ([]byte, error) {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	if err := json.NewDecoder(strings.NewReader(schema)).Decode(&params); err != nil {
		return nil, err
	}

	response := h.Schema.Exec(c.Request.Context(), params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return responseJSON, nil
}

// Verifier ac verifier
type Verifier struct {
	Enable bool
	// TODO: ac client
}

// VerifyAC is
func (v *Verifier) VerifyAC(ctx context.Context, resource, action string, args interface{}) error {
	fmt.Printf("enable ac, resource: %s action: %s\n", resource, action)
	return nil
}

func (v *Verifier) VerifyRefAC(ctx context.Context, resource, action string, args interface{}) error {
	return nil
}
