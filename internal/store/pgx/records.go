package pgx

import (
	"context"
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

func (str *Store) GetRecordByID(ctx context.Context, id int64) (rec *model.Record, err error) {
	rec = new(model.Record)
	rec.ID = id
	if err = str.pool.QueryRow(ctx, getRecordByID, id).Scan(&rec.MsgType, &rec.User, &rec.CreatedAt); err != nil {
		return nil, err
	}
	return
}
