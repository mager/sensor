package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mager/sensor/config"
	"go.uber.org/zap"
)

// Handler struct for HTTP requests
type Handler struct {
	logger *zap.SugaredLogger
	router *mux.Router
	config config.Config
}

// New creates a Handler struct
func New(l *zap.SugaredLogger, r *mux.Router, c config.Config) *Handler {
	h := Handler{l, r, c}
	h.registerRoutes()

	return &h
}

func (h *Handler) registerRoutes() {
	h.router.HandleFunc("/forecast", h.getForecast).Methods("GET")
}

type GetForecastResp struct {
	Success bool `json:"success"`
}

func (h *Handler) getForecast(w http.ResponseWriter, r *http.Request) {
	// Call api.weather.gov /points

	// Call api.weather.gov /gridpoints

	// Return forecast
	var resp GetForecastResp

	resp.Success = true

	json.NewEncoder(w).Encode(resp)
}
