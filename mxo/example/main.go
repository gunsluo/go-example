package main

import (
	"bytes"
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

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

const (
	driverPostgres = "postgres"
	dsnPostgres    = "postgres://postgres:password@localhost:5432/xo?sslmode=disable"

	driverMssql = "mssql"
	dsnMssql    = "sqlserver://SA:Tes9ting@localhost:1433/instance?database=xo&encrypt=disable"
)

func main() {
	var driver string
	if len(os.Args) <= 1 {
		driver = "postgres"
	} else {
		driver = os.Args[1]
	}

	fmt.Println("run driver:", driver)
	switch driver {
	case "postgres":
		testStoage(driver)
	case "mssql":
		testStoage(driver)
	case "server":
		if len(os.Args) > 2 {
			driver = os.Args[2]
		} else {
			driver = "postgres"
		}

		if driver != "postgres" && driver != "mssql" {
			fmt.Println("invalid parameter, it should be 'postgres' & 'mssql'")
			return
		}

		server(driver)
	default:
		fmt.Println("invalid parameter, it should be 'postgres' & 'mssql' & 'server'")
	}
}

func getDsn(driver string) string {
	switch driver {
	case "postgres":
		return dsnPostgres
	case "mssql":
		return dsnMssql
	}

	return dsnPostgres
}

func testStoage(driver string) {
	dsn := getDsn(driver)
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	s, err := storage.New(driver, storage.Config{Logger: logrus.New()})
	if err != nil {
		panic(err)
	}

	testStoageAndDB(s, db)

	db.Close()
}

func testStoageAndDB(s storage.Storage, db *sqlx.DB) {
	account := &storage.Account{
		Subject:     "luoji",
		CreatedDate: storage.NullTime{Time: time.Now(), Valid: true},
		ChangedDate: storage.NullTime{Time: time.Now(), Valid: true},
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

	account.Subject = "luoji1"
	err = s.UpsertAccount(db, account)
	if err != nil {
		panic(err)
	}

	user := &storage.User{
		Subject:     account.Subject,
		CreatedDate: storage.NullTime{Time: time.Now(), Valid: true},
		ChangedDate: storage.NullTime{Time: time.Now(), Valid: true},
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
	fmt.Printf("query account: %v\n", a)

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
	fmt.Printf("delete account: %v\n", err)
}

// Server is a graphql server.
type gqlServer struct {
	address string
	engine  *gin.Engine

	logger logrus.FieldLogger
	db     storage.XODB
	s      storage.Storage
}

// NNewGQLServer is graphql server
func NewGQLServer(address string, logger logrus.FieldLogger, db storage.XODB, s storage.Storage) (*gqlServer, error) {

	// graphql API
	rootResolver := storage.NewRootResolver(&storage.ResolverConfig{Logger: logger, DB: db, S: s})
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
		logger:  logger,
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
	s.logger.Infoln("Starting http server listening on:", s.address)
	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func server(driver string) {
	dsn := getDsn(driver)
	logger := logrus.New()

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
