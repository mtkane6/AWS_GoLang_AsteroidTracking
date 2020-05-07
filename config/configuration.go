package config

import (
	"io/ioutil"
)

// GetAllAsteroidURL returns the URL for the All Asteroids call
func GetAllAsteroidURL(todaysDate string) string {
	BaseURL := "https://api.nasa.gov/neo/rest/v1/feed?start_date=" + todaysDate + "&end_date=" + todaysDate + "&api_key=" + getAPIkey()
	return BaseURL
}

// GetSingleAsteroidURL returns the URL to make the single asteroid call
func GetSingleAsteroidURL(id string) string {
	BaseURL := "https://api.nasa.gov/neo/rest/v1/neo/" + id + "?api_key=" + getAPIkey()
	return BaseURL
}

// GetAPIkey returns the unique API key to authenticate calls to DarkSky
func getAPIkey() string {
	APIKey, err := ioutil.ReadFile("../APIKEYS/Asteroids.txt")
	if err != nil {
		panic(err)
	}
	return string(APIKey)
}
