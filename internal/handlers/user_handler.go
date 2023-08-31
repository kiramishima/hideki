package handlers

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/unrolled/render"
	"go.uber.org/zap"
	ports "hideki/internal/core/ports/service"
	httpErrors "hideki/pkg/errors"
	"hideki/pkg/utils"
	"net/http"
)

// NewUserHandlers creates a instance of auth handlers
func NewUserHandlers(r *chi.Mux, logger *zap.SugaredLogger, s ports.UserService, render *render.Render) {
	handler := &UserHandlers{
		logger:   logger,
		service:  s,
		response: render,
	}

	r.Route("/v1/me", func(r chi.Router) {
		r.With(jwtauth.Verifier(utils.TokenAuth), jwtauth.Authenticator).Get("/", handler.GetProfileHandler)
	})
}

type UserHandlers struct {
	logger   *zap.SugaredLogger
	service  ports.UserService
	response *render.Render
}

func (h *UserHandlers) GetProfileHandler(w http.ResponseWriter, req *http.Request) {
	_, uid := utils.GetUserIDInJWTHeader(req)
	h.logger.Info("UID ->", uid)

	ctx := req.Context()

	resp, err := h.service.GetProfile(ctx, uid)
	if err != nil {
		h.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			// http.Error(w, ErrTimeout, http.StatusGatewayTimeout)
			h.response.JSON(w, http.StatusGatewayTimeout, map[string]string{"error": httpErrors.ErrTimeout.Error()})
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				// http.Error(w, ErrBadEmailOrPassword, http.StatusBadRequest)
				h.response.JSON(w, http.StatusBadRequest, map[string]string{"error": httpErrors.ErrBadEmailOrPassword.Error()})
			} else {
				// http.Error(w, ErrBadEmailOrPassword, http.StatusInternalServerError)
				h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.ErrBadEmailOrPassword.Error()})
			}
		}
		return
	}
	// writeJSON(w, http.StatusAccepted, resp, nil); err != nil
	if err := h.response.JSON(w, http.StatusAccepted, resp); err != nil {
		h.logger.Error(err.Error())
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.InternalServerError.Error()})
		// http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandlers) GetProfilePremiumHandler(w http.ResponseWriter, req *http.Request) {
}
