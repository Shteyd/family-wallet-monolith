package customerHandler

import (
	"context"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/delivery/http/handler/internal/request"
	"monolith/pkg/indigo"
	"monolith/pkg/json"
	"monolith/pkg/serializer"

	"github.com/go-playground/validator/v10"
	"github.com/indigo-web/indigo/http"
	"github.com/indigo-web/indigo/http/status"
	"github.com/indigo-web/indigo/router/inbuilt/types"
)

type CustomerCreator interface {
	Save(context.Context, core.Customer) (core.Customer, error)
}

func NewSignUp(usecase CustomerCreator) types.Handler {
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

		customer, err := usecase.Save(r.Ctx, requestBody.ToEntity())
		if err != nil {
			jsonBody := serializer.ErrorResponse(int(status.InternalServerError), err, nil)
			return indigo.ErrorResponse(r, status.InternalServerError, jsonBody)
		}

		jsonBody := serializer.SuccessResponse(customer)
		return indigo.SuccessResponse(r, jsonBody)
	}
}
