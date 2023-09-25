package weatherclient

import (
	"errors"

	"weatherapi.com/weather/constants"
)

// CurrentWeatherResponse is the response for the current weather endpoint.
type CurrentWeatherResponse struct {
	// Condition is the weather condition.
	Condition string `json:"condition"`
	// HumanFriendlyTemp is the human friendly temperature name.
	HumanFriendlyTemp string `json:"humanFriendlyTemp"`
	// Alerts is the weather alerts.
	Alerts string `json:"alerts"`
}

// NewCurrentWeatherResponse creates a new current weather response for this challenge.
func NewCurrentWeatherResponse(weather []Weather, temp float64) (*CurrentWeatherResponse, error) {
	weatherResponse := CurrentWeatherResponse{}

	if len(weather) == 0 {
		return nil, errors.New("weather is empty")
	}

	for _, w := range weather {
		// Create a comma delimited list of weather conditions.
		if weatherResponse.Condition != "" {
			weatherResponse.Condition += ", "
		}
		weatherResponse.Condition += w.Description

		// Create a comma delimited list of weather alerts.
		if weatherResponse.Alerts != "" {
			weatherResponse.Alerts += ", "
		}
		switch w.ID {
		case constants.ALERT_RAIN_EXTREME:
			weatherResponse.Alerts += "Flood Warning is in effect, due to extreme rainfall!"
		case constants.ALERT_THUNDERSTORM_HEAVY:
			weatherResponse.Alerts += "Thunderstorm Warning is in effect!"
		case constants.ALERT_TORNADO:
			weatherResponse.Alerts += "Tornado Warning is in effect!"
		}
	}

	weatherResponse.HumanFriendlyTemp = weatherResponse.getFriendlyWeatherName(temp)

	return &weatherResponse, nil
}

func (w *CurrentWeatherResponse) getFriendlyWeatherName(temp float64) string {
	switch {
	case temp < constants.TEMP_MODERATE:
		return "Cold"
	case temp < constants.TEMP_HOT:
		return "Moderate"
	case temp > constants.TEMP_HOT:
		return "Hot"
	}

	return ""
}
