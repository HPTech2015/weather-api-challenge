package main

import (
	"log"
	"net/http"

	"weatherapi.com/weather/constants"
	"weatherapi.com/weather/handlers"
)

func urls() {
	http.HandleFunc(constants.API_V1_PATH+"/weather/loc/", handlers.CurrentLocWeather)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	urls()
}
