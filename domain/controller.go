package domain

import (
	"time"

	a "Asteroids/aws"
	d "Asteroids/datastructs"
)

// SeeAllAsteroids returns data for all asteroids
func SeeAllAsteroids() {
	var rawAsteroids map[string]interface{}
	var asteroidCollection []d.Asteroid

	date := time.Now().Format("2006-01-02")

	// fills asteroid info into rawAsteroids map
	CallNasaAllNeos(date, &rawAsteroids)

	// fills asteroidCollection[] and orbital data
	asteroidCollection = d.FillAsteroids(rawAsteroids, asteroidCollection, date)
	d.FillOrbitalData(asteroidCollection)

	// for i, v := range asteroidCollection {
	// 	fmt.Printf("Collection %d: %+v\n", i, v.Orbital)

	dyna := a.DynamoInit()
	// fmt.Printf("%+v", &dyna)
	a.InputItem(asteroidCollection, dyna)

}
