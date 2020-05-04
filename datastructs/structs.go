package datastructs

// RawAsteroids holds the asteroid data
type RawAsteroids struct {
	Links struct {
		Next string `json:"next"`
		Prev string `json:"prev"`
		Self string `json:"self"`
	} `json:"links"`
	ElementCount     int `json:"element_count"`
	NearEarthObjects struct {
		Two200422 []struct {
			Links struct {
				Self string `json:"self"`
			} `json:"links"`
			ID                 string  `json:"id"`
			NeoReferenceID     string  `json:"neo_reference_id"`
			Name               string  `json:"name"`
			NasaJplURL         string  `json:"nasa_jpl_url"`
			AbsoluteMagnitudeH float64 `json:"absolute_magnitude_h"`
			EstimatedDiameter  struct {
				Kilometers struct {
					EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
					EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
				} `json:"kilometers"`
				Meters struct {
					EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
					EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
				} `json:"meters"`
				Miles struct {
					EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
					EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
				} `json:"miles"`
				Feet struct {
					EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
					EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
				} `json:"feet"`
			} `json:"estimated_diameter"`
			IsPotentiallyHazardousAsteroid bool `json:"is_potentially_hazardous_asteroid"`
			CloseApproachData              []struct {
				CloseApproachDate      string `json:"close_approach_date"`
				CloseApproachDateFull  string `json:"close_approach_date_full"`
				EpochDateCloseApproach int64  `json:"epoch_date_close_approach"`
				RelativeVelocity       struct {
					KilometersPerSecond string `json:"kilometers_per_second"`
					KilometersPerHour   string `json:"kilometers_per_hour"`
					MilesPerHour        string `json:"miles_per_hour"`
				} `json:"relative_velocity"`
				MissDistance struct {
					Astronomical string `json:"astronomical"`
					Lunar        string `json:"lunar"`
					Kilometers   string `json:"kilometers"`
					Miles        string `json:"miles"`
				} `json:"miss_distance"`
				OrbitingBody string `json:"orbiting_body"`
			} `json:"close_approach_data"`
			IsSentryObject bool `json:"is_sentry_object"`
		} `json:"2020-04-22"`
	} `json:"near_earth_objects"`
}

// Asteroid is the formatted data I want
type Asteroid struct {
	ID                                string
	Name                              string
	NasaJplURL                        string
	EstimatedDiameterMetersMin        float64
	EstimatedDiameterMetersMax        float64
	IsPotentiallyHazardous            bool
	CloseApproachDate                 string
	RelativeVelocityKilometersPerHour string
	MissDistanceKilometers            string
}

// FillAsteroids adds data from rawAsteroid data into Asteroid structs
func FillAsteroids(rawAsteroids map[string]interface{}, asteroidCollection []Asteroid, date string) []Asteroid {
	for _, v := range rawAsteroids["near_earth_objects"].(map[string]interface{})[date].([]interface{}) {
		var curr Asteroid
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
