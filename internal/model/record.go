package model

import "time"

type Record struct {
	ID        int64
	User      string
	MsgType   string
	CreatedAt time.Time
}

type UsersRecord struct {
	ID        int64     `json:"id"`
	MsgType   string    `json:"msg_type"`
	CreatedAt time.Time `json:"created_at"`
}
