package service

import (
	"github.com/vlad-marlo/enrollment/internal/model"
	"github.com/vlad-marlo/enrollment/internal/pkg/fielderr"
)

var (
	ErrBadRequest = fielderr.New("bad request", model.BadRequestResponse{}, fielderr.CodeBadRequest)
	ErrNotFound   = fielderr.New("not found", model.BadRequestResponse{}, fielderr.CodeNotFound)
	ErrInternal   = fielderr.New("internal error", model.BadRequestResponse{}, fielderr.CodeInternal)
)
