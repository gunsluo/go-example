## README

### Postgres
xo 'postgres://postgres:password@127.0.0.1:5432/dbname?sslmode=disable' -o postgres

### Mssql
xo 'sqlserver://SA:Tes9ting@localhost:1433/instance?database=dbname' -o sqlserver 

### Postgres & Mssql 
xo --escape-all 'pgsql://postgres:password@127.0.0.1:5432/xo?sslmode=disable,sqlserver://SA:Tes9ting@localhost:1433/instance?database=xo' -o /Users/luoji/gopath/src/github.com/gunsluo/go-example/mxo/storage --template-path=/Users/luoji/gopath/src/github.com/xo/xo/templates
