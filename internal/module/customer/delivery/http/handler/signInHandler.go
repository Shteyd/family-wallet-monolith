package customerHandler

import (
	"github.com/indigo-web/indigo/http"
	"github.com/indigo-web/indigo/router/inbuilt/types"
)

func NewSignIn() types.Handler {
	return func(r *http.Request) http.Response {
		return r.Respond()
	}
}
