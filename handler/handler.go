package handler

import (
	"encoding/json"
	"fmt"
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
	h.router.HandleFunc("/weather", h.getCurrentWeather).Methods("GET")
}

type OpenWeatherMapAPIGetWeatherResp struct {
	Main OpenWeatherMapAPIGetWeatherRespMain `json:"main"`
}

type OpenWeatherMapAPIGetWeatherRespMain struct {
	Temp float64 `json:"temp"`
}

type GetCurrentWeatherResp struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
}

func (h *Handler) getCurrentWeather(w http.ResponseWriter, r *http.Request) {
	// Get city parameter from URL
	city := r.URL.Query().Get("city")

	h.logger.Info("City?? ", city)

	// Fetch weather data from openweathermap API
	apiKey := h.config.OpenWeatherMapKey
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	h.logger.Infow("Calling OpenWeatherMap API", "city", city)

	// Call OpenWeatherMap API
	resp, err := http.Get(url)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode response into OpenWeatherMapAPIGetWeatherResp
	var owmResp OpenWeatherMapAPIGetWeatherResp
	err = json.NewDecoder(resp.Body).Decode(&owmResp)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Format the response
	celciusTemp := adaptTempToCelcius(owmResp.Main.Temp)
	response := GetCurrentWeatherResp{
		City:  city,
		TempC: celciusTemp,
		TempF: adaptCelciusTempToFarhenheit(celciusTemp),
	}

	json.NewEncoder(w).Encode(response)
}

func adaptTempToCelcius(temp float64) float64 {
	return roundToTwoDecimals(temp - 273.15)
}

func adaptCelciusTempToFarhenheit(temp float64) float64 {
	return roundToTwoDecimals(temp*1.8 + 32)
}

func roundToTwoDecimals(num float64) float64 {
	return float64(int(num*100)) / 100
}
