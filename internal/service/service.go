package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/vlad-marlo/enrollment/internal/controller"
	"github.com/vlad-marlo/enrollment/internal/model"
	"github.com/vlad-marlo/enrollment/internal/pkg/logger"
	"go.uber.org/zap"
)

type Repository interface {
	CreateRecord(ctx context.Context, record *model.Record) error
	GetRecord(ctx context.Context, id int64) (*model.Record, error)
	GetUserRecords(ctx context.Context, user string) ([]model.UsersRecord, error)
	GetUsers(ctx context.Context) ([]string, error)
	GetCountOfRecords(ctx context.Context) (int64, error)
}

var _ controller.Service = (*Service)(nil)

type Service struct {
	logger     *zap.Logger
	repository Repository
}

var ErrNilReference = errors.New("nil reference")

// New initialize service with logger and storage.
func New(log *zap.Logger, repo Repository) (*Service, error) {
	if log == nil {
		return nil, fmt.Errorf("%w: logger", ErrNilReference)
	}
	srv := new(Service)
	srv.repository = repo
	srv.logger = log.With(zap.String(logger.EntityField, "service"))
	return srv, nil
}
