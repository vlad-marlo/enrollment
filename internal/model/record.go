package model

import "time"

type Record struct {
	ID        int64
	User      string
	MsgType   string
	CreatedAt time.Time
}
