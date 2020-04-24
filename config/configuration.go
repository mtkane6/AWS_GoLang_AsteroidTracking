package config

import (
	"io/ioutil"
)

// GetAllAsteroidURL returns the URL for the All Asteroids call
func GetAllAsteroidURL(todaysDate string) string {
	BaseURL := "https://api.nasa.gov/neo/rest/v1/feed?start_date=" + todaysDate + "&end_date=" + todaysDate + "&api_key=" + GetAPIkey()
	return BaseURL
}

// GetAPIkey returns the unique API key to authenticate calls to DarkSky
func GetAPIkey() string {
	APIKey, err := ioutil.ReadFile("../APIKEYS/Asteroids.txt")
	if err != nil {
		panic(err)
	}
	return string(APIKey)
}
