package config

import (
	"io/ioutil"
)

// GetBaseURL returns the beginning of the API call to DarkSky
func GetBaseURL() string {
	BaseURL := "https://api.nasa.gov/neo/rest/v1/feed?"
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
