# migrate

**install migrate**
```
go get -u -d github.com/golang-migrate/migrate/cmd/migrate
cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
git checkout $TAG  # e.g. v4.1.0
go build -tags 'postgres,sqlserver' -ldflags="-X main.Version=$(git describe --tags)" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/cmd/migrate
```


**create a new migrate file**
```
migrate create -ext sql -dir ./sql/postgres -seq account
migrate create -ext sql -dir ./sql/postgres -seq user 

migrate create -ext sql -dir ./sql/mssql -seq account
migrate create -ext sql -dir ./sql/mssql -seq user 
```


**run migrate to up/down sql**
```
migrate -database 'postgres://postgres:password@localhost:5432/db?sslmode=disable' -path ./sql/postgres up
migrate -database 'postgres://postgres:password@localhost:5432/db?sslmode=disable' -path ./sql/postgres down 


migrate -database 'sqlserver://SA:password@localhost:1433?database=db&encrypt=disable' -path ./sql/mssql up
migrate -database 'sqlserver://SA:password@localhost:1433?database=db&encrypt=disable' -path ./sql/mssql down 
```

