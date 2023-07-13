package model

type BadRequestResponse struct{}

type CreateRecordResponse struct {
	ID      int64  `json:"id" xml:"id"`
	User    string `json:"user" xml:"user" form:"user"`
	MsgType string `json:"msg_type" xml:"msg_type" form:"msg_type"`
}
