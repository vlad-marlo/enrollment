package pgx

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vlad-marlo/enrollment/internal/pkg/pgx"
	"github.com/vlad-marlo/enrollment/internal/service"
	"go.uber.org/zap"
)

var (
	ErrNilReference = errors.New("unexpectedly got nil reference in storage")
)

var _ service.Repository = (*Store)(nil)

// Store is postgres storage.
type Store struct {
	log  *zap.Logger
	pool *pgxpool.Pool
}

// New return storage with provided client
func New(cli pgx.Client) (*Store, error) {
	if cli == nil {
		return nil, ErrNilReference
	}
	return &Store{
		log:  cli.L(),
		pool: cli.P(),
	}, nil
}
