# migrate

**create a new migrate file**
```
migrate create -ext sql -dir ./sql/postgres -seq account
migrate create -ext sql -dir ./sql/postgres -seq user 

migrate create -ext sql -dir ./sql/mssql -seq account
migrate create -ext sql -dir ./sql/mssql -seq user 
```


**run migrate to up/down sql**
```
migrate -database 'postgres://postgres:password@localhost:5432/migrate?sslmode=disable' -path ./sql/postgres up
migrate -database 'postgres://postgres:password@localhost:5432/migrate?sslmode=disable' -path ./sql/postgres down 


migrate -database 'sqlserver://SA:password@localhost:1433?database=migrate&encrypt=disable' -path ./sql/mssql up
migrate -database 'sqlserver://SA:password@localhost:1433?database=migrate&encrypt=disable' -path ./sql/mssql down 
```


