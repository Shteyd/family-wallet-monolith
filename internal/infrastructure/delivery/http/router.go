package http

import (
	healthHandler "monolith/internal/infrastructure/delivery/http/internal/handler/health"
	customerHandler "monolith/internal/module/customer/delivery/http/handler"
)

func (server *_HttpServer) SetRouter() {
	healthGroup := server.Engine.Group("/health")
	{
		healthGroup.Get("", healthHandler.NewGetHealth())
	}

	apiV1Group := server.Engine.Group("/api/v1")
	{
		customerGroup := apiV1Group.Group("/customer")
		{
			customerGroup.Post("/sign-in", // /api/v1/customer/sign-in
				customerHandler.NewSignIn())
			customerGroup.Post("/sign-up", // /api/v1/customer/sign-up
				customerHandler.NewSignUp(server.CustomerModule))
		}
	}
}
