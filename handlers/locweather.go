package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"weatherapi.com/weather/weatherclient"
)

// CurrentLocWeather is the handler for the /v1/json/weather/loc/ endpoint.
// It returns the current weather for a given location.
func CurrentLocWeather(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		apiKey := r.URL.Query().Get("api_key")
		latStr := r.URL.Query().Get("lat")
		lonStr := r.URL.Query().Get("lon")

		if apiKey == "" || latStr == "" || lonStr == "" {
			http.Error(w, "Missing required parameters.", http.StatusBadRequest)
			return
		}

		lat, err := strconv.ParseFloat(latStr, 64)
		if err != nil {
			http.Error(w, "Invalid latitude.", http.StatusBadRequest)
			return
		}
		lon, err := strconv.ParseFloat(lonStr, 64)
		if err != nil {
			http.Error(w, "Invalid longitude.", http.StatusBadRequest)
			return
		}

		currentWeather, err := weatherclient.NewCurrentWeatherData(lat, lon, apiKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		currentWeatherResponse, err := weatherclient.NewCurrentWeatherResponse(currentWeather.Weather, currentWeather.Main.Temp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(currentWeatherResponse)
	}
}
