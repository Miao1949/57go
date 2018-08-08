package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const ApiKey="8028ac32013fb0a39f9043a36198a213"
const UrlToService= "http://api.openweathermap.org/data/2.5/weather?APPID=" + ApiKey + "&q="
const KELVIN_TO_CELSIUS_DIFF=273.15

// {"coord":{"lon":18.07,"lat":59.33},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03d"}],"base":"stations","main":{"temp":297.181,"pressure":1024.01,"humidity":53,"temp_min":297.181,"temp_max":297.181,"sea_level":1026.78,"grnd_level":1024.01},"wind":{"speed":2.2,"deg":219.504},"clouds":{"all":44},"dt":1533652785,"sys":{"message":0.0026,"country":"SE","sunrise":1533610051,"sunset":1533668262},"id":2673730,"name":"Stockholm","cod":200}

type Coord struct {
	Lon float64
	Lat float64
}

type Weather struct {
	Main string
	Description string
}

type Main struct {
	Temp float64
	Pressure float64
	Humidity int
	TempMin float64 `json:"temp_min"`
	TempMax float64 `json:"temp_max"`
}
type Wind struct {
	Speed float64
	Deg float64
}

type Sys struct {
	Message float64
	Country string
	Sunrise int
	Sunset int
}
type WeatherData struct {
	coord Coord
	Weather []Weather
	Main Main
	Wind Wind
	Sys Sys
}


func main() {
	location := getNonEmptyInput("Where are you? ")
	weather, _ := readWeatherData(location)
	//fmt.Println(weather)
	displayWeather(location, weather)
}

func displayWeather(location string, weatherData WeatherData) {
	fmt.Printf("%s weather:\n", location)
	fmt.Printf("%.0f degrees Fahrenheit\n", kelvinToFahrenheit(weatherData.Main.Temp))
}

func kelvinToFahrenheit(degreesKelvin float64) (degreesFahrenheit float64) {
	degreesFahrenheit = celsiusToFahrenheit(kelvinToCelsius(degreesKelvin))
	return
}

func celsiusToFahrenheit(degreesCelsius float64) (degreesFahrenheit float64) {
	degreesFahrenheit = (degreesCelsius * 9.0 / 5.0) + 32.0
	return
}

func kelvinToCelsius(degreesKelvin float64) (degreesCelsius float64) {
	degreesCelsius = degreesKelvin - 273.15
	return
}


func getNonEmptyInput(msg string) (input string) {
	done := false
	for ; !done; {
		fmt.Print(msg)
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			done = true
		}
	}

	return input
}

func readWeatherData(location string) (weatherData WeatherData, errorToReturn error) {
	url := UrlToService + location
	fmt.Printf("Fetching data from %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read data from URL! Error: %v", err)
		return
	} else {
		// Make sure the connection is closed.
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(contents)
			fmt.Fprintf(os.Stderr,"Could not read from response! Error: %v", err)
			errorToReturn = err
			return
		} else {
			if err := json.Unmarshal(contents, &weatherData); err != nil {
				fmt.Fprintf(os.Stderr,"Could not urmarshal contents as json!")
				errorToReturn = err
				return
			}
		}
	}
	return
}





