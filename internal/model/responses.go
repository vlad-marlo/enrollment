package model

import "time"

type BadRequestResponse struct{}

type CreateRecordResponse struct {
	ID        int64     `json:"id" xml:"id"`
	User      string    `json:"user" xml:"user"`
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
	MsgType   string    `json:"msg_type" xml:"msg_type"`
}

type GetRecordResponse struct {
	ID        int64     `json:"id" xml:"id"`
	User      string    `json:"user" xml:"user"`
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
	MsgType   string    `json:"msg_type" xml:"msg_type"`
}

type GetAllRecordsResponse struct {
	RecordsStored   int64                    `json:"records_stored"`
	UsersRegistered int64                    `json:"users_registered"`
	Records         map[string][]UsersRecord `json:"users"`
}

type GetUserRecordsResponse struct {
	Count   int64         `json:"records_stored"`
	User    string        `json:"user"`
	Records []UsersRecord `json:"records"`
}
