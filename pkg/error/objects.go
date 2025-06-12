package error

import "net/http"

const (
	EmptyMessage           = ""
	ErrorBadFormat         = 40000
	ErrorBadRequest        = 40001
	ErrorInsufficientFunds = 40002
	ErrorBindingRequest    = 40003
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
