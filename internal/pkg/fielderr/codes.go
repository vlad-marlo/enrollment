package fielderr

import "net/http"

type Code int

const (
	CodeBadRequest Code = iota
	CodeNotFound
	CodeInternal
	CodeUnauthorized
	CodeConflict
	CodeForbidden
	CodeNoContent
	CodeOK
)

var httpCodes = map[Code]int{
	CodeBadRequest:   http.StatusBadRequest,
	CodeNotFound:     http.StatusNotFound,
	CodeInternal:     http.StatusInternalServerError,
	CodeUnauthorized: http.StatusUnauthorized,
	CodeForbidden:    http.StatusForbidden,
	CodeConflict:     http.StatusConflict,
	CodeNoContent:    http.StatusNoContent,
	CodeOK:           http.StatusOK,
}
