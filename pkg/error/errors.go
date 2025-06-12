package error

import "net/http"

type ErrorResponse struct {
	Code            int                    `json:"code,omitempty"`
	CodeDescription string                 `json:"code_description,omitempty"`
	Message         string                 `json:"message,omitempty"`
	AdditionalInfo  map[string]interface{} `json:"additional_info,omitempty"`
}

type Error struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}

// Generate returns a ErrorResponse struct with proper error code, description, message and additionalInfo if any
func (er *ErrorResponse) Generate(errCode int, message string, additionalInfo map[string]interface{}) *ErrorResponse {
	er.Code = errCode
	er.CodeDescription = ErrorText(errCode)
	er.Message = message
	er.AdditionalInfo = additionalInfo

	return er
}

func ErrorText(code int) string {
	return errorText[code]
}

var errorText = map[int]string{
	ErrorBadFormat:      "BadFormatError",
	ErrorBadRequest:     "BadRequest",
	ErrorBindingRequest: "We're experiencing difficulties binding your request at the moment. Please ensure all required information is provided and try again. If the issue persists, kindly contact our support team for assistance.",
	ErrorNoRecordsFound: "NoRecordsFound",
	ErrorForbidden:      "Forbidden",
	ErrorInternalError:  "InternalServerError",
}

func GetHttpStatusCodeForError(code int) int {
	val, ok := errorCodeToHttpStatusCodeMap[code]
	if !ok {
		return http.StatusInternalServerError
	}
	return val
}
