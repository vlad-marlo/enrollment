package pgx

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/vlad-marlo/enrollment/internal/model"
	"github.com/vlad-marlo/enrollment/internal/store"
)

const (
	createRecordQuery = `INSERT INTO records(msg_type, "user")
VALUES ($1, $2)
RETURNING id, created_at;`
	getRecordByID = `SELECT msg_type, "user", created_at
FROM records
WHERE id = $1;`
	getUserRecords = `SELECT id, msg_type, created_at
FROM records
WHERE "user" = $1
ORDER BY created_at DESC;`
	getCountOfRecordsQuery = `SELECT COUNT(*) FROM records;`
	getUsersQuery          = `SELECT DISTINCT "user"
FROM records;`
)

func (str *Store) CreateRecord(ctx context.Context, record *model.Record) error {
	if record == nil {
		return store.ErrBadData
	}
	if err := str.pool.QueryRow(
		ctx,
		createRecordQuery,
		record.MsgType,
		record.User,
	).Scan(
		&record.ID,
		&record.CreatedAt,
	); err != nil {
		return err
	}
	return nil
}

func (str *Store) GetRecord(ctx context.Context, id int64) (rec *model.Record, err error) {
	rec = new(model.Record)
	rec.ID = id
	if err = str.pool.QueryRow(ctx, getRecordByID, id).Scan(&rec.MsgType, &rec.User, &rec.CreatedAt); err != nil {
		return nil, err
	}
	return
}

func (str *Store) GetUserRecords(ctx context.Context, user string) ([]model.UsersRecord, error) {
	rows, err := str.pool.Query(ctx, getUserRecords, user)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []model.UsersRecord{}, nil
		}

		return nil, fmt.Errorf("unable to get records: %w", err)
	}

	var res []model.UsersRecord

	for rows.Next() {
		var record model.UsersRecord

		if err = rows.Scan(&record.ID, &record.MsgType, &record.CreatedAt); err != nil {
			return nil, fmt.Errorf("unable to scan into record: %w", err)
		}

		res = append(res, record)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error from rows.Err(): %w", err)
	}
	return res, nil
}

func (str *Store) GetUsers(ctx context.Context) ([]string, error) {
	rows, err := str.pool.Query(ctx, getUsersQuery)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []string{}, nil
		}

		return nil, fmt.Errorf("unable to get users: %w", err)
	}

	var res []string

	for rows.Next() {
		var user string

		if err = rows.Scan(&user); err != nil {
			return nil, fmt.Errorf("unable to scan into user: %w", err)
		}

		res = append(res, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error from rows.Err(): %w", err)
	}
	return res, nil
}

func (str *Store) GetCountOfRecords(ctx context.Context) (res int64, err error) {
	if err = str.pool.QueryRow(ctx, getCountOfRecordsQuery).Scan(&res); err != nil {
		return 0, fmt.Errorf("unable to get count of records: error while scanning result: %w", err)
	}
	return
}
