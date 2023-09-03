package http

import (
	"log"
	auth "monolith/internal/module/authorization/core"
	"monolith/pkg/slog"

	"github.com/indigo-web/indigo"
	"github.com/indigo-web/indigo/router/inbuilt"
)

type HttpServer interface{ Run() }

type _HttpServer struct {
	Engine *inbuilt.Router

	AuthModule auth.AuthorizationUsecase

	Port string
}

func NewHttpServer(authModule auth.AuthorizationUsecase, addr string) HttpServer {
	httpServer := &_HttpServer{
		Engine:     inbuilt.New(),
		AuthModule: authModule,
		Port:       addr,
	}

	httpServer.SetRouter()

	return httpServer
}

func (server *_HttpServer) Run() {
	app := indigo.NewApp(server.Port)
	log.Println("start work on addr", server.Port)
	if err := app.Serve(server.Engine); err != nil {
		slog.Fatal(err.Error(), slog.String("addr", server.Port))
	}
}
