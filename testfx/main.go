package main

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	appOpts := fx.Options(
		fx.Provide(NewConfig),
		fx.Provide(NewLogger),
		fx.Provide(NewEchoHandler),
		fx.Provide(NewMux),
		fx.Provide(NewServer),
		fx.Invoke(RunServer),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)

	fx.New(appOpts).Run()
}

type Config struct {
	Address string
}

func NewConfig() Config {
	return Config{Address: ":8088"}
}

func NewLogger() *zap.Logger {
	return zap.NewExample()
}

type EchoHandler struct {
	log *zap.Logger
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, wErr := io.Copy(w, r.Body)
	if wErr != nil {
		h.log.Error("failed to write response", zap.Error(wErr))

		return
	}

	h.log.Info("handler called successfully")
}

func NewMux(echoHandler *EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo/", echoHandler)
	return mux
}

func NewServer(cfg Config, mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:    cfg.Address,
		Handler: mux,
	}
}

func RunServer(srv *http.Server) error {
	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("server failed to start or finished with error: %w", err)
	}

	return nil
}
