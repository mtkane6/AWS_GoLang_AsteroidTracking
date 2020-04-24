package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	c "../config"
	// d "../datastructs"
)

// SeeAllAsteroids returns data for all asteroids
func SeeAllAsteroids() {
	// var rawAsteroids d.RawAsteroids
	var rawAsteroids map[string]interface{}
	// date := time.Now()
	// newDate := date.Format("2006-01-02")
	// fmt.Println(newDate)

	URL := c.GetAllAsteroidURL(time.Now().Format("2006-01-02"))
	fmt.Println("URL:", URL)
	// URL += "start_date=2020-04-22&end_date=2020-04-22&api_key=" + c.GetAPIkey()

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Resp:", resp)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(b), &rawAsteroids)
	// fmt.Println("rawAsteroidsxxx:", rawAsteroids["near_earth_objects"].(map[string]interface{})["2020-04-23"])

	for _, v := range rawAsteroids["near_earth_objects"].(map[string]interface{})["2020-04-23"].([]interface{}) {
		// create new Asteroid object
		// add all the data from Raw into Asteroid
		fmt.Println("ID:", v.(map[string]interface{})["id"])
	}
}
