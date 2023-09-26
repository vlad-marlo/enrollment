package migrator

import (
	"context"
	"github.com/vlad-marlo/enrollment/internal/pkg/pgx"
	"github.com/vlad-marlo/enrollment/internal/pkg/retryer"
	"time"
)

const (
	migrationRetryAttempts = 2
	migrationsRetryDelay   = time.Second
)

var (
	migrations = []string{
		`CREATE TABLE IF NOT EXISTS records
(
    "id"         BIGSERIAL PRIMARY KEY UNIQUE NOT NULL,
    "msg_type"   TEXT                         NOT NULL,
    "user"       TEXT                         NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`,
	}
	migrateDown = []string{
		`DROP TABLE IF EXISTS records;`,
	}
	Migrations = len(migrations)
)

func MigrateUp(cli pgx.Client) (int, error) {
	i := 0
	for _, migration := range migrations {
		if err := retryer.TryWithAttempts(
			func() error {
				_, err := cli.P().Exec(context.Background(), migration)
				return err
			},
			migrationRetryAttempts,
			migrationsRetryDelay,
		); err != nil {
			return i, err
		}
		i++
	}
	return i, nil
}

func MigrateDown(cli pgx.Client) (int, error) {
	i := 0
	for _, migration := range migrateDown {
		if err := retryer.TryWithAttempts(
			func() error {
				_, err := cli.P().Exec(context.Background(), migration)
				return err
			},
			migrationRetryAttempts,
			migrationsRetryDelay,
		); err != nil {
			return i, err
		}
		i++
	}
	return i, nil
}
