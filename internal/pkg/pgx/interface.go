package pgx

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Client interface {
	P() *pgxpool.Pool
	L() *zap.Logger
}
