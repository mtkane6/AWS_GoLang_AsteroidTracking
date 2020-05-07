package domain

import (
	c "Asteroids/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CallNasaAllNeos makes get request to NASA Api and fills map structure with returned JSON data
func CallNasaAllNeos(date string, rawAsteroids *map[string]interface{}) {
	URL := c.GetAllAsteroidURL(date)

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(b), &rawAsteroids)
}
