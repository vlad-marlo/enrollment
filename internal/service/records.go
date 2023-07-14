package service

import (
	"context"
	"errors"
	"github.com/vlad-marlo/enrollment/internal/model"
	"github.com/vlad-marlo/enrollment/internal/store"
	"go.uber.org/zap"
	"strconv"
)

func (srv *Service) CreateRecord(ctx context.Context, req *model.CreateRecordRequest) (*model.CreateRecordResponse, error) {
	record := &model.Record{
		User:    req.User,
		MsgType: req.MsgType,
	}

	if err := srv.repository.CreateRecord(ctx, record); err != nil {
		return nil, ErrBadRequest.With(zap.Error(err))
	}

	resp := &model.CreateRecordResponse{
		ID:        record.ID,
		User:      record.User,
		CreatedAt: record.CreatedAt,
		MsgType:   record.MsgType,
	}
	return resp, nil
}

func (srv *Service) GetRecord(ctx context.Context, rawID string) (*model.GetRecordResponse, error) {
	var record *model.Record

	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return nil, ErrBadRequest.With(zap.Error(err), zap.String("raw_id", rawID))
	}

	record, err = srv.repository.GetRecord(ctx, id)
	if err != nil {
		return nil, ErrBadRequest.With(zap.Error(err))
	}

	resp := &model.GetRecordResponse{
		ID:        record.ID,
		User:      record.User,
		CreatedAt: record.CreatedAt,
		MsgType:   record.MsgType,
	}
	return resp, nil
}

func (srv *Service) GetRecords(ctx context.Context) (resp *model.GetAllRecordsResponse, err error) {
	resp = new(model.GetAllRecordsResponse)
	var users []string

	users, err = srv.repository.GetUsers(ctx)
	if err != nil {
		return nil, ErrInternal.With(zap.Error(err))
	}

	resp.RecordsStored, err = srv.repository.GetCountOfRecords(ctx)
	if err != nil {
		return nil, ErrInternal.With(zap.Error(err))
	}

	resp.UsersRegistered = int64(len(users))
	resp.Records = make(map[string][]model.UsersRecord)

	for _, user := range users {
		resp.Records[user], err = srv.repository.GetUserRecords(ctx, user)
		if err != nil {
			return nil, ErrInternal.With(zap.Error(err))
		}
	}

	return resp, nil
}

func (srv *Service) GetUser(ctx context.Context, user string) (resp *model.GetUserRecordsResponse, err error) {
	resp = new(model.GetUserRecordsResponse)
	resp.User = user

	if resp.Records, err = srv.repository.GetUserRecords(ctx, user); err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return resp, nil
		}

		return nil, ErrInternal.With(zap.Error(err))
	}

	resp.Count = int64(len(resp.Records))
	return
}
