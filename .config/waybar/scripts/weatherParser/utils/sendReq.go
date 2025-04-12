package utils

import (
	"bufio"
	"errors"
	"net/http"

	"github.com/Wefdzen/WayWeather/weather"
)

// If you frequently use a VPN, the weather information will be based on the region of your VPN.
// To get the weather for a specific city, use the following format:
// "https://wttr.in/<Your_City_Name>?format=3"
// For example, to get the weather for London, use:
// "https://wttr.in/London?format=3"
var weatherAPI = "https://wttr.in/?format=3"

// SendReq: send requst to api of "wttr.in".
func SendReq() (weather.Weather, error) {
	resp, err := http.Get(weatherAPI)
	if err != nil {
		return weather.Weather{}, err
	}
	defer resp.Body.Close()

	//Checking status code.
	if resp.StatusCode != http.StatusOK {
		return weather.Weather{}, errors.New("wttr.in return not 200.")
	}

	scanner := bufio.NewScanner(resp.Body)
	var res weather.Weather
	for i := 0; scanner.Scan() && i < 1; i++ {
		res, err = parseWeatherFormat3(scanner.Text())
		if err != nil {
			return weather.Weather{}, err
		}
	}

	if err := scanner.Err(); err != nil {
		return weather.Weather{}, err
	}

	return res, nil
}
