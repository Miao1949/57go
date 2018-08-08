package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
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
	displayWeather(location, weather)
}

func displayWeather(location string, weatherData WeatherData) {
	fmt.Printf("%s weather:\n", location)
	fmt.Printf("%.0f degrees Celisius\n", kelvinToCelsius(weatherData.Main.Temp))
	fmt.Printf("%.0f degrees Fahrenheit\n", kelvinToFahrenheit(weatherData.Main.Temp))
	fmt.Printf("Wind direction: %s\n", degreesToWindDirectionInWords(weatherData.Wind.Deg))
	fmt.Printf("Sunrise: %s\n",  epochToPrintableTime(weatherData.Sys.Sunrise))
	fmt.Printf("Sunset: %s\n" , epochToPrintableTime(weatherData.Sys.Sunset))
	fmt.Printf("Humidity: %d\n" , weatherData.Main.Humidity)
	fmt.Print("Description: ")
	fmt.Println(weatherData.Weather[0].Description)
	fmt.Printf("Categorization: %s\n", categorizeWeather(weatherData))
}

func epochToPrintableTime(epoch int) (epochAsString string) {
	epochAsTime := time.Unix(int64(epoch), 0)
	// Mon Jan 2 15:04:05 MST 2006
	layout := "2006-01-02 15:04:05"
	epochAsString = epochAsTime.Format(layout)
	return
}

func categorizeWeather(weatherData WeatherData) (categorization string) {
	description := weatherData.Weather[0].Description
	if strings.Contains(description, "rain") {
		categorization = "Rainy"
	} else if  strings.Contains(description, "cloud") {
		categorization = "Cloudy"
	} else if  strings.Contains(description, "clear sky") {
		categorization = "Nice"
	}
	categorization += ", "

	temperature := kelvinToCelsius(weatherData.Main.Temp)
	if temperature < 0 {
		categorization += "cold"
	} else if temperature < 20 {
		categorization += "normal temp"
	} else {
		categorization += "warm"
	}

	return
}

func degreesToWindDirectionInWords(degrees float64) (windDirection string) {
	degreesAsFloat := float64(int(degrees) % 360)
	if degreesAsFloat > 337.5 || degreesAsFloat <= 22.5 {
		windDirection = "East"
	} else if degreesAsFloat > 22.5 && degreesAsFloat <= 67.5 {
		windDirection = "NorthEast"
	} else if degreesAsFloat > 67.5 && degreesAsFloat <= 112.5 {
		windDirection = "North"
	} else if degreesAsFloat > 112.5 && degreesAsFloat <= 157.5 {
		windDirection = "NorthWest"
	} else if degreesAsFloat > 157.5 && degreesAsFloat <= 202.5 {
		windDirection = "West"
	} else if degreesAsFloat > 202.5 && degreesAsFloat <= 247.5 {
		windDirection = "SouthWest"
	} else if degreesAsFloat > 247.5 && degreesAsFloat <= 292.5 {
		windDirection = "South"
	} else if degreesAsFloat > 292.5 && degreesAsFloat <= 337.5 {
		windDirection = "SouthEast"
	}
	return
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





