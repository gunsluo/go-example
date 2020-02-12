
### install sql-migrate
```
go get -v github.com/rubenv/sql-migrate/...
```


### run sql-migrate
```
sql-migrate new 001

sql-migrate up -limit=0
```

```
sql-migrate new -config=dbconfig.oracle.yml 001

sql-migrate up -config=dbconfig.oracle.yml -limit=0
```
