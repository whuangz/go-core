package networking

import "net/http"

const statusUnprocessableEntity = 422

type ResponseError struct {
	reason string `json:"reason"`
	status int    `json:"status`
}

func (r ResponseError) ResponseError() string { return r.reason }

func (r ResponseError) Status() int { return r.status }

func NotFound(reason string) ResponseError {
	return ResponseError{reason, http.StatusNotFound}
}

func Forbidden(reason string) ResponseError {
	return ResponseError{reason, http.StatusForbidden}
}

func Conflict(reason string) ResponseError {
	return ResponseError{reason, http.StatusConflict}
}

func BadRequest(reason string) ResponseError {
	return ResponseError{reason, http.StatusBadRequest}
}

func UnprocessableRequest(reason string) ResponseError {
	return ResponseError{reason, statusUnprocessableEntity}
}

func UnauthorizedRequest(reason string) ResponseError {
	return ResponseError{reason, http.StatusUnauthorized}
}

func InternalServerError(reason string) ResponseError {
	return ResponseError{reason, http.StatusInternalServerError}
}

func CustomError(reason string, status int) ResponseError {
	return ResponseError{reason, status}
}
