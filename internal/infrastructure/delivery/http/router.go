package http

import (
	healthHandler "monolith/internal/infrastructure/delivery/http/internal/handler/health"
	authHandler "monolith/internal/module/authorization/delivery/http/handler"
)

func (server *_HttpServer) SetRouter() {
	healthGroup := server.Engine.Group("/health")
	{
		healthGroup.Get("", healthHandler.NewGetHealth())
	}

	apiV1Group := server.Engine.Group("/api/v1")
	{
		authGroup := apiV1Group.Group("/auth")
		{
			authGroup.Post("/sign-in", // /api/v1/auth/sign-in
				authHandler.NewSignIn(server.AuthModule))
			authGroup.Post("/sign-up", // /api/v1/auth/sign-up
				authHandler.NewSignUp(server.AuthModule))
		}
	}
}
