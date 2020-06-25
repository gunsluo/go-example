package storage

import (
	"context"
	"fmt"

	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/internal"
	identitypb "github.com/gunsluo/go-example/opentelemetry/demo/pkg/proto/identity"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
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
	sqlstr := "SELECT * FROM account WHERE id=$1"

	out := &internal.Account{}
	err := d.db.QueryRowxContext(ctx, d.db.Rebind(sqlstr), id).Scan(&out.Id, &out.Name, &out.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to query account id: %s, %w", id, err)
	}

	return out, nil
}

func (d *Database) GetIdentity(ctx context.Context, id string) (*identitypb.Identity, error) {
	sqlstr := "SELECT * FROM identity WHERE id=$1"

	out := &identitypb.Identity{}
	err := d.db.QueryRowxContext(ctx, d.db.Rebind(sqlstr), id).Scan(&out.Id, &out.Name, &out.CertId)
	if err != nil {
		return nil, fmt.Errorf("failed to query identity id: %s, %w", id, err)
	}

	return out, nil
}
