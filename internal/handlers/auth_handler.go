package handlers

import (
	"errors"
	"hideki/internal/core/domain"
	ports "hideki/internal/core/ports/service"
	httpErrors "hideki/pkg/errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// NewAuthHandlers creates a instance of auth handlers
func NewAuthHandlers(r *chi.Mux, logger *zap.SugaredLogger, s ports.AuthService) {
	handler := &AuthHandlers{
		logger:  logger,
		service: s,
	}

	r.Route("/v1/auth", func(r chi.Router) {
		r.Post("/sign-in", handler.SignInHandler)
	})
}

type AuthHandlers struct {
	logger  *zap.SugaredLogger
	service ports.AuthService
}

func (h *AuthHandlers) SignInHandler(w http.ResponseWriter, req *http.Request) {
	var form = &domain.AuthRequest{}

	err := readJSON(w, req, &form)

	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, ErrInvalidRequestBody, http.StatusBadRequest)
		return
	}
	h.logger.Info(form)
	ctx := req.Context()

	resp, err := h.service.Login(ctx, form)
	if err != nil {
		h.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			http.Error(w, ErrTimeout, http.StatusGatewayTimeout)
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				http.Error(w, ErrBadEmailOrPassword, http.StatusBadRequest)
			} else {
				http.Error(w, ErrBadEmailOrPassword, http.StatusInternalServerError)
			}
		}
		return
	}

	if err := writeJSON(w, http.StatusAccepted, resp, nil); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}
