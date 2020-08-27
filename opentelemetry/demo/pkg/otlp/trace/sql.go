package trace

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/luna-duclos/instrumentedsql"
	"github.com/xo/dburl"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/label"
)

// OpenDB open driver
func OpenDB(tracer trace.Tracer, dsn string) (*sqlx.DB, error) {
	u, err := dburl.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse database address, %w", err)
	}

	if tracer == nil {
		return sqlx.Open(u.Driver, u.DSN)
	}

	if _, ok := tracer.(*trace.NoopTracer); ok {
		return sqlx.Open(u.Driver, u.DSN)
	}

	driver := uiqDriver()
	sql.Register(driver,
		instrumentedsql.WrapDriver(&pq.Driver{},
			instrumentedsql.WithTracer(&dbTracer{tracer: tracer}),
			instrumentedsql.WithOpsExcluded(
				instrumentedsql.OpSQLConnectorConnect,
				instrumentedsql.OpSQLRowsNext,
			)))

	db, err := sql.Open(driver, u.DSN)
	if err != nil {
		return nil, err
	}
	return sqlx.NewDb(db, u.Driver), nil
}

var uiqflag int

func uiqDriver() string {
	uiqflag++
	return fmt.Sprintf("inst-pg-%d", uiqflag)
}

type dbTracer struct {
	tracer trace.Tracer
}

func (t *dbTracer) GetSpan(ctx context.Context) instrumentedsql.Span {
	return &dbSpan{tracer: t.tracer, ctx: ctx}
}

type dbSpan struct {
	tracer trace.Tracer
	ctx    context.Context
	parent trace.Span
}

func (s *dbSpan) NewChild(name string) instrumentedsql.Span {
	s.ctx, s.parent = s.tracer.Start(s.ctx, name)
	return s
}

func (s *dbSpan) SetLabel(k, v string) {
	if s.parent == nil {
		return
	}
	s.parent.SetAttributes(label.Key(k).String(v))
}

func (s *dbSpan) SetError(err error) {
	if s.parent == nil {
		return
	}

	if err == nil || err == driver.ErrSkip {
		return
	}

	s.parent.SetStatus(codes.Internal, err.Error())
}

func (s *dbSpan) Finish() {
	if s.parent == nil {
		return
	}

	s.parent.End()
}
