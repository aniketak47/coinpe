package error

import "net/http"

const (
	EmptyMessage           = ""
	ErrorInsufficientFunds = 40001
	ErrorUnauthorized      = 40101
	ErrorForbidden         = 40301
	ErrorNoRecordsFound    = 40401
	ErrorInternalError     = 50001
)

var EmptyInterface map[string]interface{}

var errorCodeToHttpStatusCodeMap = map[int]int{
	ErrorInternalError: http.StatusInternalServerError,
	ErrorForbidden:     http.StatusForbidden,
}
