package handlers

import (
	"errors"
	"github.com/unrolled/render"
	"hideki/internal/core/domain"
	ports "hideki/internal/core/ports/service"
	httpErrors "hideki/pkg/errors"
	httpUtils "hideki/pkg/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// NewAuthHandlers creates a instance of auth handlers
func NewAuthHandlers(r *chi.Mux, logger *zap.SugaredLogger, s ports.AuthService, render *render.Render) {
	handler := &AuthHandlers{
		logger:   logger,
		service:  s,
		response: render,
	}

	r.Route("/v1/auth", func(r chi.Router) {
		r.Post("/sign-in", handler.SignInHandler)
		r.Post("/sign-up", handler.SignUpHandler)
	})
}

type AuthHandlers struct {
	logger   *zap.SugaredLogger
	service  ports.AuthService
	response *render.Render
}

func (h *AuthHandlers) SignInHandler(w http.ResponseWriter, req *http.Request) {
	var form = &domain.AuthRequest{}

	err := httpUtils.ReadJSON(w, req, &form)

	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, httpErrors.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	h.logger.Info(form)
	ctx := req.Context()

	resp, err := h.service.FindByCredentials(ctx, form)
	if err != nil {
		h.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			_ = h.response.JSON(w, http.StatusGatewayTimeout, httpErrors.ErrTimeout)
			// http.Error(w, httpErrors.ErrTimeout.Error(), http.StatusGatewayTimeout)
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				_ = h.response.JSON(w, http.StatusBadRequest, map[string]string{"error": httpErrors.BadQueryParams.Error()})
				// http.Error(w, httpErrors.ErrBadEmailOrPassword.Error(), http.StatusBadRequest)
			} else if errors.Is(err, httpErrors.ErrUserNotFound) {
				_ = h.response.JSON(w, http.StatusBadRequest, map[string]string{"error": httpErrors.ErrBadEmailOrPassword.Error()})
			} else {
				_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.InternalServerError.Error()})
				// http.Error(w, httpErrors.ErrBadEmailOrPassword.Error(), http.StatusInternalServerError)
			}
		}
		return
	}

	if err := h.response.JSON(w, http.StatusAccepted, resp); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.InternalServerError.Error()})
		// http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

// SignUpHandler for register new users
func (h *AuthHandlers) SignUpHandler(w http.ResponseWriter, req *http.Request) {
	var form = &domain.RegisterRequest{}

	err := httpUtils.ReadJSON(w, req, &form)

	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, httpErrors.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	h.logger.Info(form)
	ctx := req.Context()

	err = h.service.Register(ctx, form)
	if err != nil {
		h.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			_ = h.response.JSON(w, http.StatusGatewayTimeout, httpErrors.ErrTimeout)
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				_ = h.response.JSON(w, http.StatusBadRequest, httpErrors.ErrBadEmailOrPassword)
			} else {
				_ = h.response.JSON(w, http.StatusInternalServerError, httpErrors.ErrBadEmailOrPassword)
			}
		}
		return
	}

	if err := h.response.JSON(w, http.StatusAccepted, domain.SuccessResponse{Message: "Success. Check you email for activate your account."}); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.InternalServerError.Error()})
		// http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}
