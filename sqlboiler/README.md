go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
go get -u -t github.com/volatiletech/sqlboiler
go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

sqlboiler psql
