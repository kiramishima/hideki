package handlers

import (
	"github.com/unrolled/render"
	httpErrors "hideki/pkg/errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// NewHealthHandlers creates a instance of health handlers
func NewHealthHandlers(logger *zap.SugaredLogger, r *chi.Mux, render *render.Render) {
	handler := &HealthHandlers{
		logger:   logger,
		response: render,
	}

	r.Route("/v1/health", func(r chi.Router) {
		r.Get("/", handler.HealthHandler)
	})
}

type HealthHandlers struct {
	logger   *zap.SugaredLogger
	response *render.Render
}

func (h *HealthHandlers) HealthHandler(w http.ResponseWriter, req *http.Request) {

	if err := h.response.JSON(w, http.StatusAccepted, map[string]string{"status": "OK", "version": "1.0"}); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": httpErrors.InternalServerError.Error()})
		// http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}
