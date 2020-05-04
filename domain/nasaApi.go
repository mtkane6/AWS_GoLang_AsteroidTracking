package domain

import (
	c "Asteroids/config"
	d "Asteroids/datastructs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CallNasa makes get request to NASA Api and fills map structure with returned JSON data
func CallNasa(date string, rawAsteroids *map[string]interface{}) {
	URL := c.GetAllAsteroidURL(date)

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(b), &rawAsteroids)
}

// FillAsteroids adds data from rawAsteroid data into Asteroid structs
func FillAsteroids(rawAsteroids map[string]interface{}, asteroidCollection []d.Asteroid, date string) []d.Asteroid {
	for _, v := range rawAsteroids["near_earth_objects"].(map[string]interface{})[date].([]interface{}) {
		var curr d.Asteroid
		curr.ID = v.(map[string]interface{})["id"].(string)
		curr.Name = v.(map[string]interface{})["name"].(string)
		curr.IsPotentiallyHazardous = v.(map[string]interface{})["is_potentially_hazardous_asteroid"].(bool)
		curr.NasaJplURL = v.(map[string]interface{})["nasa_jpl_url"].(string)
		curr.EstimatedDiameterMetersMin = v.(map[string]interface{})["estimated_diameter"].(map[string]interface{})["meters"].(map[string]interface{})["estimated_diameter_min"].(float64)
		curr.EstimatedDiameterMetersMax = v.(map[string]interface{})["estimated_diameter"].(map[string]interface{})["meters"].(map[string]interface{})["estimated_diameter_max"].(float64)
		curr.CloseApproachDate = v.(map[string]interface{})["close_approach_data"].([]interface{})[0].(map[string]interface{})["close_approach_date"].(string)
		curr.MissDistanceKilometers = v.(map[string]interface{})["close_approach_data"].([]interface{})[0].(map[string]interface{})["miss_distance"].(map[string]interface{})["kilometers"].(string)
		curr.RelativeVelocityKilometersPerHour = v.(map[string]interface{})["close_approach_data"].([]interface{})[0].(map[string]interface{})["relative_velocity"].(map[string]interface{})["kilometers_per_hour"].(string)

		asteroidCollection = append(asteroidCollection, curr)
	}
	return asteroidCollection
}
