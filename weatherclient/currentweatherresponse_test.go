package weatherclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewCurrentWeatherResponseGood tests the NewCurrentWeatherResponse function with good data.
// Just testing for a couple of weather conditions, since this is just a challenge.
func TestNewCurrentWeatherResponseGood(t *testing.T) {
	weather := []Weather{
		{
			ID:          212,
			Main:        "Thunderstorm",
			Description: "heavy thunderstorm",
		},
		{
			ID:          504,
			Main:        "Rain",
			Description: "extreme rain",
		},
	}
	temp := 100.0

	expected := CurrentWeatherResponse{
		Condition:         "heavy thunderstorm, extreme rain",
		HumanFriendlyTemp: "Hot",
		Alerts:            "Thunderstorm Warning is in effect!, Flood Warning is in effect, due to extreme rainfall!",
	}

	actual, err := NewCurrentWeatherResponse(weather, temp)
	assert.Nil(t, err)
	assert.Equal(t, expected, *actual)
}

// TestNewCurrentWeatherResponseBad tests the NewCurrentWeatherResponse function with bad data.
func TestNewCurrentWeatherResponseBad(t *testing.T) {
	weather := []Weather{
		{
			ID:          800,
			Main:        "Clear",
			Description: "clear sky",
		},
	}
	temp := 0.0

	expected := CurrentWeatherResponse{
		Condition:         "heavy thunderstorm",
		HumanFriendlyTemp: "Hot",
		Alerts:            "Thunderstorm Warning is in effect!",
	}

	actual, err := NewCurrentWeatherResponse(weather, temp)
	assert.Nil(t, err)
	assert.NotEqual(t, expected, *actual)

	weather = []Weather{}
	temp = 0.0
	_, err = NewCurrentWeatherResponse(weather, temp)
	assert.NotNil(t, err)
	assert.Equal(t, "weather is empty", err.Error())
}
