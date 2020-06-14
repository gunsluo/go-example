package storage

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"time"

	"github.com/gunsluo/go-example/opentracing/pkg/internal"
	identitypb "github.com/gunsluo/go-example/opentracing/pkg/proto/identity"
	"github.com/sirupsen/logrus"
)

// Database simulates Customer repository implemented on top of an SQL Database
type Database struct {
	//tracer    opentracing.Tracer
	//lock     sync.Mutex
	logger     logrus.FieldLogger
	accounts   map[string]*internal.Account
	identities map[string]*identitypb.Identity
}

func NewDatabase(logger logrus.FieldLogger) *Database {
	return &Database{
		//tracer: tracer,
		logger: logger,
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
	d.logger.WithField("account id", id).Info("Loading account")

	/*
		// simulate opentracing instrumentation of an SQL query
		if span := opentracing.SpanFromContext(ctx); span != nil {
			span := d.tracer.StartSpan("SQL SELECT", opentracing.ChildOf(span.Context()))
			tags.SpanKindRPCClient.Set(span)
			tags.PeerService.Set(span, "mysql")
			// #nosec
			span.SetTag("sql.query", "SELECT * FROM customer WHERE customer_id="+customerID)
			defer span.Finish()
			ctx = opentracing.ContextWithSpan(ctx, span)
		}

		if !config.MySQLMutexDisabled {
			// simulate misconfigured connection pool that only gives one connection at a time
			d.lock.Lock(ctx)
			defer d.lock.Unlock()
		}
	*/

	// simulate RPC delay
	delay := time.Duration(math.Max(1, rand.NormFloat64()*300+30) * float64(time.Millisecond))
	time.Sleep(delay)

	if a, ok := d.accounts[id]; ok {
		return a, nil
	}
	return nil, errors.New("invalid account ID")
}

func (d *Database) GetIdentity(ctx context.Context, id string) (*identitypb.Identity, error) {
	d.logger.WithField("user id", id).Info("Loading identity")

	delay := time.Duration(math.Max(1, rand.NormFloat64()*300+30) * float64(time.Millisecond))
	time.Sleep(delay)

	if a, ok := d.identities[id]; ok {
		return a, nil
	}
	return nil, errors.New("invalid user ID")
}
