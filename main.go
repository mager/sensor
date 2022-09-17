package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mager/sensor/config"
	"github.com/mager/sensor/handler"
	"github.com/mager/sensor/logger"
	"github.com/mager/sensor/router"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			config.Options,
			logger.Options,
			router.Options,
		),
		fx.Invoke(Register),
	).Run()
}

func Register(l *zap.SugaredLogger, r *mux.Router, c config.Config) {
	l.Infow("Starting up the service")

	addr := ":8080"
	l.Info("Listening on ", addr)
	go http.ListenAndServe(addr, r)

	handler.New(l, r, c)
}
