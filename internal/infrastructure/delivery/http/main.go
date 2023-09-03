package http

import (
	"log"
	customer "monolith/internal/module/customer/core"
	"monolith/pkg/slog"

	"github.com/indigo-web/indigo"
	"github.com/indigo-web/indigo/router/inbuilt"
)

type HttpServer interface{ Run() }

type _HttpServer struct {
	Engine *inbuilt.Router

	CustomerModule customer.CustomerUsecase

	Port string
}

func NewHttpServer(customerModule customer.CustomerUsecase, addr string) HttpServer {
	httpServer := &_HttpServer{
		Engine:         inbuilt.New(),
		CustomerModule: customerModule,
		Port:           addr,
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
