package main

import (
	"fmt"
	"math"
)

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515 * 1.609344
	return dist
}

func main() {
	//fmt.Printf("%f Miles\n", distance(32.9697, -96.80322, 29.46786, -98.53506, "M"))
	fmt.Printf("%f Kilometers\n", distance(32.9697, -96.80322, 29.46786, -98.53506))
	//fmt.Printf("%f Nautical Miles\n", distance(32.9697, -96.80322, 29.46786, -98.53506, "N"))
}
