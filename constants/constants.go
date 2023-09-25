package constants

const (
	// WEATHER_URL is the base URL for the OpenWeatherMap API.
	WEATHER_URL = "https://api.openweathermap.org/data/2.5/weather"
	// API V1 Path is the path version 1 path for this API.
	API_V1_PATH = "/v1/json"
)

// Temperature is an enum for temperature.
const (
	TEMP_COLD     = 40
	TEMP_MODERATE = 60
	TEMP_HOT      = 80
)

// Alert is an enum for weather alerts.
const (
	ALERT_THUNDERSTORM_HEAVY = 212
	ALERT_RAIN_EXTREME       = 504
	ALERT_RAIN_FREEZING      = 511
	ALERT_TORNADO            = 781
)
