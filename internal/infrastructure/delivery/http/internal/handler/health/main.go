package healthHandler

import (
	"github.com/indigo-web/indigo/http"
	"github.com/indigo-web/indigo/http/status"
	"github.com/indigo-web/indigo/router/inbuilt/types"
)

func NewGetHealth() types.Handler {
	return func(request *http.Request) http.Response {
		return request.Respond().
			WithCode(status.OK).
			WithBody("health ok")
	}
}
