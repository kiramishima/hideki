package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// NewHealthHandlers creates a instance of health handlers
func NewHealthHandlers(logger *zap.SugaredLogger, r *chi.Mux) {
	handler := &HealthHandlers{
		logger: logger,
	}

	r.Route("/v1/health", func(r chi.Router) {
		r.Get("/", handler.HealthHandler)
	})
}

type HealthHandlers struct {
	logger *zap.SugaredLogger
}

func (h *HealthHandlers) HealthHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]string{"status": "OK", "version": "1.0"}); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}
