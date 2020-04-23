package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	c "../config"
	d "../datastructs"
)

// MakeCall makes the GET request call
func MakeCall() {
	var rawAsteroids d.RawAsteroids
	URL := c.GetBaseURL()
	URL += "start_date=2020-04-22&end_date=2020-04-22&api_key=" + c.GetAPIkey()

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(b), &rawAsteroids)

	for _, v := range rawAsteroids.NearEarthObjects.Two200422 {
		// create new Asteroid object
		// add all the data from Raw into Asteroid
		fmt.Printf("%s, %s\n", v.ID, v.Name)
	}
}
