package client

import (
	"context"
	"fmt"
	pgxzap "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/stretchr/testify/assert"
	"github.com/vlad-marlo/enrollment/internal/pkg/pgx/migrator"
	"github.com/vlad-marlo/enrollment/internal/pkg/retryer"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"os"
	"strings"
	"testing"
	"time"
)

type (
	Config interface {
		URI() string
	}
	Client struct {
		pool *pgxpool.Pool
		log  *zap.Logger
	}
)

const (
	RetryAttempts = 4
	RetryDelay    = 2 * time.Second
)

// New opens new postgres connection, configures it and return prepared client.
func New(lc fx.Lifecycle, cfg Config, log *zap.Logger) (*Client, error) {
	var pool *pgxpool.Pool
	log.Info("initializing postgres client with config", zap.Any("cfg", cfg))

	c, err := pgxpool.ParseConfig(
		cfg.URI(),
	)
	if err != nil {
		return nil, fmt.Errorf("error while parsing db uri: %w", err)
	}

	var lvl = tracelog.LogLevelError
	c.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxzap.NewLogger(log),
		LogLevel: lvl,
	}

	pool, err = pgxpool.NewWithConfig(context.Background(), c)
	if err != nil {
		return nil, fmt.Errorf("postgres: init pgxpool: %w", err)
	}

	cli := &Client{
		pool: pool,
		log:  log,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return retryer.TryWithAttemptsCtx(ctx, pool.Ping, RetryAttempts, RetryDelay)
		},
		OnStop: func(ctx context.Context) error {
			pool.Close()
			return nil
		},
	})
	log.Info("created postgres client")
	return cli, nil
}

// NewTest prepares test client.
//
// If error occurred while creating connection then test will be skipped.
// Second argument cleanup function to close connection and rollback all changes.
func NewTest(t testing.TB) (*Client, func()) {
	t.Helper()

	pool, err := pgxpool.New(context.Background(), os.Getenv("TEST_DB_URI"))
	if err != nil {
		t.Skipf("can not create pool: %v", err)
	}

	cli := &Client{
		pool: pool,
		log:  zap.L(),
	}

	if err = retryer.TryWithAttemptsCtx(context.Background(), pool.Ping, 5, 200*time.Millisecond); err != nil {
		t.Skipf("can not get access to db: %v", err)
	}
	if _, err = migrator.MigrateUp(cli); err != nil {
		t.Skipf("can not migrate database: %v", err)
	}
	return cli, func() {
		_, err = migrator.MigrateDown(cli)
		assert.NoError(t, err)
		teardown(cli.pool)()
	}
}

func BadCli(t testing.TB) *Client {
	t.Helper()

	pool, err := pgxpool.New(context.Background(), "postgresql://postgres:postgres@localhost:4321/unknown_db")
	if err != nil {
		t.Skipf("can not create pool: %v", err)
	}

	cli := &Client{
		pool: pool,
		log:  zap.L(),
	}

	if err = retryer.TryWithAttemptsCtx(context.Background(), pool.Ping, 5, 200*time.Millisecond); err == nil {
		t.Skip("must have no connection to database")
	}
	return cli
}

// teardown return func for defer it to clear tables.
//
// Always pass one or more tables in it.
func teardown(pool *pgxpool.Pool, tables ...string) func() {
	return func() {
		_, _ = pool.Exec(context.Background(), fmt.Sprintf("TRUNCATE %s CASCADE;", strings.Join(tables, ", ")))
		pool.Close()
	}
}

// L return global client logger.
//
// If client is nil object then global logger will be returned.
func (cli *Client) L() *zap.Logger {
	if cli == nil {
		zap.L().Error("unexpectedly got nil client dereference")
		return zap.L()
	}
	return cli.log
}

// P returns client's configured logger.
//
// If client is nil object then will be returned nil pool.
func (cli *Client) P() *pgxpool.Pool {
	if cli == nil {
		zap.L().Error("unexpectedly got nil client dereference")
		return nil
	}
	return cli.pool
}
