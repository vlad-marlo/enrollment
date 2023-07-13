package controller

import (
	"context"
	"github.com/vlad-marlo/enrollment/internal/model"
)

type Interface interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Service interface {
	CreateRecord(ctx context.Context, req *model.CreateRecordRequest) (*model.CreateRecordResponse, error)
}

type Config interface {
	BindAddr() string
}
