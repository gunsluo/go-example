package storage

import (
	"context"
	"fmt"

	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/internal"
	identitypb "github.com/gunsluo/go-example/opentelemetry/demo/pkg/proto/identity"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

// Database simulates Customer repository implemented on top of an SQL Database
type Database struct {
	db         *sqlx.DB
	logger     *zap.Logger
	accounts   map[string]*internal.Account
	identities map[string]*identitypb.Identity
}

func NewDatabase(logger *zap.Logger, db *sqlx.DB) (*Database, error) {
	return &Database{
		db:     db,
		logger: logger,
	}, nil
}

func (d *Database) GetAccount(ctx context.Context, id string) (*internal.Account, error) {
	sqlstr := "SELECT * FROM account WHERE id=?"

	out := &internal.Account{}
	err := d.db.QueryRowxContext(ctx, d.db.Rebind(sqlstr), id).Scan(&out.Id, &out.Name, &out.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to query account id: %s, %w", id, err)
	}

	return out, nil
}

func (d *Database) GetIdentity(ctx context.Context, id string) (*identitypb.Identity, error) {
	sqlstr := "SELECT * FROM identity WHERE id=?"

	out := &identitypb.Identity{}
	err := d.db.QueryRowxContext(ctx, d.db.Rebind(sqlstr), id).Scan(&out.Id, &out.Name, &out.CertId)
	if err != nil {
		return nil, fmt.Errorf("failed to query identity id: %s, %w", id, err)
	}

	return out, nil
}

/*
// Database simulates Customer repository implemented on top of an SQL Database
type Database struct {
	db         *sqlx.DB
	logger     *zap.SugaredLogger
	accounts   map[string]*internal.Account
	identities map[string]*identitypb.Identity
}

func NewDatabase(logger *zap.Logger, db *sqlx.DB) *Database {
	//tracer := trace.Init("postgres", logger, nil)
	return &Database{
		//tracer: tracer,
		logger: logger.Sugar(),
		accounts: map[string]*internal.Account{
			"123": {
				Id:    "123",
				Name:  "Rachel's Floral Designs",
				Email: "rachel@test.com",
			},
			"567": {
				Id:    "567",
				Name:  "Amazing Coffee Roasters",
				Email: "amazing@test.com",
			},
			"392": {
				Id:    "392",
				Name:  "Trom Chocolatier",
				Email: "trom@test.com",
			},
			"731": {
				Id:    "731",
				Name:  "Japanese Desserts",
				Email: "dess@test.com",
			},
		},
		identities: map[string]*identitypb.Identity{
			"123": {
				Id:     "123",
				Name:   "Rachel's Floral Designs",
				CertId: "xxxx001",
			},
			"567": {
				Id:     "567",
				Name:   "Amazing Coffee Roasters",
				CertId: "xxxx002",
			},
			"392": {
				Id:     "392",
				Name:   "Trom Chocolatier",
				CertId: "xxxx003",
			},
			"731": {
				Id:     "731",
				Name:   "Japanese Desserts",
				CertId: "xxxx004",
			},
		},
	}
}

func (d *Database) GetAccount(ctx context.Context, id string) (*internal.Account, error) {
	sqlstr := "SELECT * FROM account WHERE id=?"
	out := &internal.Account{}
	if err := d.invokeQuery(ctx, sqlstr, id, out); err != nil {
		return nil, fmt.Errorf("failed to query account id: %s, %w", id, err)
	}

	return out, nil
}

func (d *Database) GetIdentity(ctx context.Context, id string) (*identitypb.Identity, error) {
	sqlstr := "SELECT * FROM identity WHERE id=?"
	out := &identitypb.Identity{}
	if err := d.invokeQuery(ctx, sqlstr, id, out); err != nil {
		return nil, fmt.Errorf("failed to query identity id: %s, %w", id, err)
	}

	return out, nil
}

func (d *Database) invokeQuery(ctx context.Context, sql string, in, out interface{}) error {
	//if d.tracer == nil {
	//	return d.invoke(ctx, sql, in, out)
	//}

	// simulate opentracing instrumentation of an SQL query
	/*
		if span := opentracing.SpanFromContext(ctx); span != nil {
			span := d.tracer.StartSpan("SQL SELECT", opentracing.ChildOf(span.Context()))
			ext.SpanKindRPCClient.Set(span)
			ext.PeerService.Set(span, "postgres")
			// #nosec
			span.SetTag("sql.query", sql)
			defer span.Finish()
			ctx = opentracing.ContextWithSpan(ctx, span)
		}
*/

//err := d.invoke(ctx, sql, in, out)
/*
	if err != nil {
		if span := opentracing.SpanFromContext(ctx); span != nil {
			ext.Error.Set(span, true)
		}
	}
*/

/*
	return err
}

func (d *Database) invoke(ctx context.Context, sql string, in, out interface{}) error {
	// simulate RPC delay
	delay := time.Duration(math.Max(1, rand.NormFloat64()*300+30) * float64(time.Millisecond))
	time.Sleep(delay)

	if strings.Index(sql, "account") > 0 {
		if id, ok := in.(string); ok {
			if a, ok := d.accounts[id]; ok {
				if o, ok := out.(*internal.Account); ok {
					o.Id = a.Id
					o.Name = a.Name
					o.Email = a.Email
					return nil
				}
			}
		}

		return errors.New("invalid user ID")
	}

	if strings.Index(sql, "identity") > 0 {
		if id, ok := in.(string); ok {
			if a, ok := d.identities[id]; ok {
				if o, ok := out.(*identitypb.Identity); ok {
					o.Id = a.Id
					o.Name = a.Name
					o.CertId = a.CertId
					return nil
				}
			}
		}

		return errors.New("invalid ID")
	}

	return errors.New("invalid sql")
}
*/
