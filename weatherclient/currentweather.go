/*
weatherclient is a package for getting weather data,
modifying it, and returning it.
*/
package weatherclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"weatherapi.com/weather/constants"
)

// CurrentWeatherData is a struct that holds the current weather data.
// This is the data returned from the openweathermap API.
type CurrentWeatherData struct {
	// Coordinates is the lat and lon coordinates of the location.
	Coordinates Coordinates `json:"coord"`
	// Name is the name of the location.
	Name string `json:"name"`
	// Weather is the weather data.
	Weather []Weather `json:"weather"`
	// Main is the main weather data.
	Main Main `json:"main"`
}

// Coordinates is a struct that holds the coordinates of the location.
type Coordinates struct {
	// Lat is the latitude of the location.
	Lat float64 `json:"lat"`
	// Lon is the longitude of the location.
	Lon float64 `json:"lon"`
}

// Weather is a struct that holds the weather data.
type Weather struct {
	// ID is the weather condition code.
	ID int32 `json:"id"`
	// Main is the main weather category.
	Main string `json:"main"`
	// Description is the description of the weather code.
	Description string `json:"description"`
}

// Main is a struct that holds the main weather data.
type Main struct {
	// Temp is the outside temperature.
	Temp float64 `json:"temp"`
}

// NewCurrentWeather creates a new current weather data instance and returns it.
func NewCurrentWeatherData(lat float64, lon float64, apiKey string) (*CurrentWeatherData, error) {
	w := CurrentWeatherData{}
	w.Coordinates.Lat = lat
	w.Coordinates.Lon = lon

	err := w.UpdateCurrentWeather(apiKey)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &w, nil
}

// UpdateCurrentWeather gets the current weather and updates the struct.
func (w *CurrentWeatherData) UpdateCurrentWeather(apiKey string) error {
	// get current weather from constants.WEATHER_URL get request
	// https://api.openweathermap.org/data/2.5/weather?lat={lat}&lon={lon}&appid={API key}&units=imperial
	resp, err := http.Get(constants.WEATHER_URL +
		"?lat=" + strconv.FormatFloat(w.Coordinates.Lat, 'f', -1, 64) +
		"&lon=" + strconv.FormatFloat(w.Coordinates.Lon, 'f', -1, 64) +
		"&appid=" + apiKey +
		"&units=imperial")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("request failed with error code: %d", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &w)
	if err != nil {
		return err
	}

	return nil
}
