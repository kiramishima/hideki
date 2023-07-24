package handlers

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	ports "hideki/internal/core/ports/service"
	"net/http"
)

// NewUserHandlers creates a instance of auth handlers
func NewUserHandlers(r *chi.Mux, logger *zap.SugaredLogger, s ports.AuthService) {
	handler := &UserHandlers{
		logger:  logger,
		service: s,
	}

	r.Route("/v1/me", func(r chi.Router) {
		r.Get("/", handler.GetProfileHandler)
	})
}

type UserHandlers struct {
	logger  *zap.SugaredLogger
	service ports.AuthService
}

func (h *UserHandlers) GetProfileHandler(w http.ResponseWriter, req *http.Request) {

}

func (h *UserHandlers) GetProfilePremiumHandler(w http.ResponseWriter, req *http.Request) {
}
