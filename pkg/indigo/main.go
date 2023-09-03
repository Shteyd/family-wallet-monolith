package indigo

import (
	"github.com/indigo-web/indigo/http"
	"github.com/indigo-web/indigo/http/status"
)

const (
	contentTypeHeader      = "Content-Type"
	contentTypeHeaderValue = "application/json"
)

func SuccessResponse(r *http.Request, bodyJson string) http.Response {
	return r.Respond().
		WithCode(status.OK).
		WithHeader(contentTypeHeader, contentTypeHeaderValue).
		WithBody(bodyJson)
}

func ErrorResponse(r *http.Request, statusCode status.Code, bodyJson string) http.Response {
	return r.Respond().
		WithCode(statusCode).
		WithHeader(contentTypeHeader, contentTypeHeaderValue).
		WithBody(bodyJson)
}
