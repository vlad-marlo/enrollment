package service

import (
	"errors"
	"fmt"
	"github.com/vlad-marlo/enrollment/internal/pkg/logger"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
}

var ErrNilReference = errors.New("nil reference")

// New initialize service with logger and storage.
func New(log *zap.Logger) (*Service, error) {
	if log == nil {
		return nil, fmt.Errorf("%w: logger", ErrNilReference)
	}
	srv := new(Service)
	srv.logger = log.With(zap.String(logger.EntityField, "service"))
	return srv, nil
}
