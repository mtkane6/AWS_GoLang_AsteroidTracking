package main

import (
	d "Asteroids/domain"
)

func main() {
	d.SeeAllAsteroids()
}

/*

P^2 =kd^3

where k is the proportionality constant = 4pi / G*M

where G is the gravitational constant, M is mass of sun or mass of planet+sun
	(in this case, M will be mass of Earth + mass of asteroid)

*/
