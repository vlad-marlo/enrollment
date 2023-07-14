package model

import (
	"fmt"
	"time"
)

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

func (m *Record) String() string {
	return fmt.Sprintf("Record:\n\tID: %d\n\tUser: %s\n\tMessage Type: %s\n\tCreated At: %s", m.ID, m.User, m.MsgType, m.CreatedAt)
}

func (m *UsersRecord) String() string {
	return fmt.Sprintf("Record:\n\tID: %d\n\tMessage Type: %s\n\tCreated At: %s", m.ID, m.MsgType, m.CreatedAt)
}
