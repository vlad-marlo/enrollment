package model

type CreateRecordRequest struct {
	User    string `json:"user" xml:"user" form:"user"`
	MsgType string `json:"msg_type" xml:"msg_type" form:"msg_type"`
}
