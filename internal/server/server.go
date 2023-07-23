package server

import (
	"fmt"
	"hideki/config"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type Server struct {
	http.Server
}

func NewServer(addr string, r *chi.Mux) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

// Module
var ServerModule = fx.Module("server",
	fx.Provide(func(cfg *config.Configuration, r *chi.Mux) *http.Server {
		var addr = fmt.Sprintf("%s:%d", cfg.ServerAddress, cfg.Port)
		return NewServer(addr, r)
	}),
)
