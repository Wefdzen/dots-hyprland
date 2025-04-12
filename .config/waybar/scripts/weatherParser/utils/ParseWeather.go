package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Wefdzen/WayWeather/weather"
)

// parseWeatherFormat3: parsing string like that Tokyo, Japan: 🌩 +15°C.
func parseWeatherFormat3(input string) (weather.Weather, error) {
	// Split the input string by ":"
	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		return weather.Weather{}, fmt.Errorf("invalid input format")
	}

	// Get the city and country
	location := strings.TrimSpace(parts[0])
	cityCountry := strings.Split(location, ",")
	if len(cityCountry) != 2 {
		return weather.Weather{}, fmt.Errorf("invalid location format")
	}

	city := strings.TrimSpace(cityCountry[0])
	country := strings.TrimSpace(cityCountry[1])

	// Get the weather info
	weatherInfo := strings.TrimSpace(parts[1])
	weatherParts := strings.Fields(weatherInfo)

	// Extract emoji and temperature
	emoji := weatherParts[0]
	tempStr := strings.TrimPrefix(weatherParts[1], "+")
	tempStr = strings.TrimSuffix(tempStr, "°C")

	// Convert temperature to int16
	temp, err := strconv.ParseInt(tempStr, 10, 16)
	if err != nil {
		return weather.Weather{}, err
	}

	return weather.Weather{
		City:           city,
		Country:        country,
		Celsius:        int16(temp),
		EmojiOfWeather: emoji,
	}, nil
}
