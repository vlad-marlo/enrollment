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
	GetRecord(ctx context.Context, id string) (*model.GetRecordResponse, error)
	GetRecords(ctx context.Context) (*model.GetAllRecordsResponse, error)
	GetUser(ctx context.Context, user string) (*model.GetUserRecordsResponse, error)
}

type Config interface {
	BindAddr() string
}
