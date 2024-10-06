package models

type Error struct {
	Message     ErrorMessage     `json:"message"`
	Code        int              `json:"statusCode"`
	Description ErrorMessageLong `json:"description"`
}

func GetError(message ErrorMessage, code int) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func GetLongError(message ErrorMessage, code int, description ErrorMessageLong) *Error {
	return &Error{
		Code:        code,
		Message:     message,
		Description: description,
	}
}

type ErrorMessage string

const (
	INTERNAL_SERVER_ERROR ErrorMessage = "Internal Server Error!"
	BAD_REQUEST           ErrorMessage = "Bad Request!"
	NOT_FOUND             ErrorMessage = "Not Found!"
)

type ErrorMessageLong string

const (
	MALFORMED_REQUEST_BODY ErrorMessageLong = "Malformed Request Body"
	MALFORMED_RISKID       ErrorMessageLong = "Malformed Risk ID"
	INVALID_REQUEST_BODY   ErrorMessageLong = "Invalid Request body"
	RISKID_NOT_EXIST       ErrorMessageLong = "Risk ID does not exist"
)
