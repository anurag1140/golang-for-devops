package errors

import "net/http"

func BadRequest(msg string) APIError {
	return APIError{
		Code:    CodeBadRequest,
		Message: msg,
		Status:  http.StatusBadRequest,
	}
}

func Unauthorized(msg string) APIError {
	return APIError{
		Code:    CodeUnauthorized,
		Message: msg,
		Status:  http.StatusUnauthorized,
	}
}

func Forbidden(msg string) APIError {
	return APIError{
		Code:    CodeForbidden,
		Message: msg,
		Status:  http.StatusForbidden,
	}
}

func NotFound(msg string) APIError {
	return APIError{
		Code:    CodeNotFound,
		Message: msg,
		Status:  http.StatusNotFound,
	}
}

func Internal(msg string) APIError {
	return APIError{
		Code:    CodeInternal,
		Message: msg,
		Status:  http.StatusInternalServerError,
	}
}
