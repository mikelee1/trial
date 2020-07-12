package main

import (
	"fmt"
	"github.com/gansidui/geohash"
	"math"
	"strings"
)

var (
	datas = []Point{
		Point{Latitude: 39.89, Longitude: 116.4},
		Point{Latitude: 39.90234375, Longitude: 116.3671875},
		Point{Latitude: 39.91, Longitude: 116.4},
		Point{Latitude: 39.92, Longitude: 116.4},
		Point{Latitude: 39.93, Longitude: 116.4},
		Point{Latitude: 39.9462890625, Longitude: 116.4111328125},
		Point{Latitude: 39.95, Longitude: 116.4},
		Point{Latitude: 39.96, Longitude: 116.4},
		Point{Latitude: 39.97, Longitude: 116.4},
		Point{Latitude: 39.98, Longitude: 116.4},
		Point{Latitude: 39.99, Longitude: 116.4},
	}
)

type Point struct {
	Latitude  float64
	Longitude float64
}

func main() {
	latitude := 39.92324
	longitude := 116.3906
	precision := 5

	hash, box := geohash.Encode(latitude, longitude, precision)
	fmt.Println(box.MaxLat, box.MaxLng, box.MinLat, box.MinLng)
	step := 0.0001
	corner := []Point{
		Point{box.MinLat - box.Height() + step, box.MinLng - box.Width() + step},
		Point{box.MinLat - box.Height() + step, box.MaxLng + box.Width() - step},
		Point{box.MaxLat + box.Height() - step, box.MinLng - box.Width() + step},
		Point{box.MaxLat + box.Height() - step, box.MaxLng + box.Width() - step},
	}
	//corner := []Point{
	//	Point{box.MinLat + step, box.MinLng + step},
	//	Point{box.MinLat + step, box.MaxLng - step},
	//	Point{box.MaxLat - step, box.MinLng + step},
	//	Point{box.MaxLat - step, box.MaxLng - step},
	//}
	for _, data := range datas {
		hash, _ := geohash.Encode(data.Latitude, data.Longitude, precision)
		fmt.Printf("%s 距离：%f公里\n", hash, distance(latitude, longitude, data.Latitude, data.Longitude))
	}
	for _, data := range corner {
		hash, _ := geohash.Encode(data.Latitude, data.Longitude, precision)
		fmt.Printf("四角： %s 距离：%f公里\n", hash, distance(latitude, longitude, data.Latitude, data.Longitude))
	}

	fmt.Println(hash)
	neighbors := geohash.GetNeighbors(latitude, longitude, precision)
	for _, hash = range neighbors {
		fmt.Print(hash, " ")
	}
	fmt.Println("hashs: ", strings.Join(neighbors[1:],"|"))
}

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
