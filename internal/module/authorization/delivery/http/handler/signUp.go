package authHandler

import (
	"context"
	"encoding/json"
	"monolith/internal/module/authorization/delivery/http/handler/request"
	customer "monolith/internal/module/customer/core"
	token "monolith/internal/module/token/core"
	"monolith/pkg/indigo"
	"monolith/pkg/serializer"

	"github.com/go-playground/validator/v10"
	"github.com/indigo-web/indigo/http"
	"github.com/indigo-web/indigo/http/status"
	"github.com/indigo-web/indigo/router/inbuilt/types"
)

type AuthSignUp interface {
	SignUp(context.Context, customer.Customer) (token.Token, error)
}

func NewSignUp(usecase AuthSignUp) types.Handler {
	return func(r *http.Request) http.Response {
		var requestBody request.SignUpBody
		if err := json.NewDecoder(r.Body()).Decode(&requestBody); err != nil {
			jsonBody := serializer.ErrorResponse(int(status.BadRequest), err, nil)
			return indigo.ErrorResponse(r, status.BadRequest, jsonBody)
		}

		validator := validator.New()
		if err := validator.Struct(requestBody); err != nil {
			jsonBody := serializer.ErrorResponse(int(status.BadRequest), err, nil)
			return indigo.ErrorResponse(r, status.BadRequest, jsonBody)
		}

		token, err := usecase.SignUp(r.Ctx, requestBody.ToEntity())
		if err != nil {
			jsonBody := serializer.ErrorResponse(int(status.InternalServerError), err, nil)
			return indigo.ErrorResponse(r, status.BadRequest, jsonBody)
		}

		jsonBody := serializer.SuccessResponse(token)
		return indigo.SuccessResponse(r, jsonBody)
	}
}
