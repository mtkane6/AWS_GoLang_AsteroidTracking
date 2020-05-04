package domain

import (
	"fmt"
	"time"

	d "Asteroids/datastructs"
)

// SeeAllAsteroids returns data for all asteroids
func SeeAllAsteroids() {
	var rawAsteroids map[string]interface{}
	var asteroidCollection []d.Asteroid

	date := time.Now().Format("2006-01-02")

	// fills asteroid info into rawAsteroids map
	CallNasa(date, &rawAsteroids)

	// fills asteroidCollection[]
	asteroidCollection = d.FillAsteroids(rawAsteroids, asteroidCollection, date)

	for i, v := range asteroidCollection {
		fmt.Printf("Collection %d: %+v\n", i, v)

	}

}
