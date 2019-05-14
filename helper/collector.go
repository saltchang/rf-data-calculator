package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

var (
	accessToken = "ivUzHzaPKbGeVpnPg8zMxTAn6lVXrzZgBf3HqV0xpUS6yrKB2Ff4zD5s2nLfoaRq"
	bearingLib  = [8]float64{0, 45, 90, 135, 180, 225, 270, 315}
	const3to15  = [3]float64{3, 15, 61}
	const10to50 = [3]float64{10, 50, 41}
)

// Path struct
type Path struct {
	Start struct {
		Lat float64
		Lng float64
	}
	End struct {
		Lat float64
		Lng float64
	}
}

// Point struct
type Point struct {
	Elevation  float64   `json:"elevation"`
	Location   *Location `json:"location"`
	Resolution float64   `json:"resolution"`
}

// Location struct
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// GetHeightData func
func GetHeightData(RFLocation [2]float64) ([8][]float64, [8][]float64) {
	RFLat := RFLocation[0]
	RFLng := RFLocation[1]
	var path3to15 [8]Path
	var path10to50 [8]Path

	for index, bearing := range bearingLib {
		path3to15[index].Start.Lat, path3to15[index].Start.Lng = getLocation(RFLat, RFLng, bearing, 3)
		path3to15[index].End.Lat, path3to15[index].End.Lng = getLocation(RFLat, RFLng, bearing, 15)
		path10to50[index].Start.Lat, path10to50[index].Start.Lng = getLocation(RFLat, RFLng, bearing, 10)
		path10to50[index].End.Lat, path10to50[index].End.Lng = getLocation(RFLat, RFLng, bearing, 50)
	}

	var data3to15 [8][]float64
	var data10to50 [8][]float64

	fmt.Println()

	for i := 0; i < 8; i++ {
		// 進度
		fmt.Printf("\r(%d / 8)抓取高度資料中...", i+1)

		point3to15start := [2]float64{path3to15[i].Start.Lat, path3to15[i].Start.Lng}
		point3to15end := [2]float64{path3to15[i].End.Lat, path3to15[i].End.Lng}
		data3to15[i] = getElevation(point3to15start, point3to15end, 61)

		point10to50start := [2]float64{path10to50[i].Start.Lat, path10to50[i].Start.Lng}
		point10to50end := [2]float64{path10to50[i].End.Lat, path10to50[i].End.Lng}
		data10to50[i] = getElevation(point10to50start, point10to50end, 41)
	}
	fmt.Printf("\r\n")
	fmt.Printf("\r資料抓取完成！\n\n")

	// for _, data := range data3to15 {
	// 	fmt.Println(data)
	// }
	// for _, data := range data10to50 {
	// 	fmt.Println(data)
	// }

	return data3to15, data10to50
}

// Get elevations by api from JawgMaps
// https://www.jawg.io/docs/apidocs/elevation/#examples
func getElevation(pStart, pEnd [2]float64, samples int) []float64 {

	s := "%7C"

	url := fmt.Sprintf("https://api.jawg.io/elevations?path=%f,%f%s%f,%f&samples=%d&access-token=%s", pStart[0], pStart[1], s, pEnd[0], pEnd[1], samples, accessToken)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// sitemap, err := ioutil.ReadAll(res.Body)
	var pointList []Point
	err = json.NewDecoder(res.Body).Decode(&pointList)
	if err != nil {
		log.Fatal(err)
	}

	var result []float64
	for _, point := range pointList {
		point.Elevation = math.Round(point.Elevation)
		result = append(result, point.Elevation)
	}

	return result

	// fmt.Printf("%v", pointList)
}

// Get the target location by given an initial position, distance and the bearing angle.
// It returns two float64 type latitude and longitude as output.
// 	Input arguments:
// 		- latIn: Latitude, in decimal
// 		- lngIn: Longitude, in decimal
// 		- theta: Bearing, in degree
// 		- distance: Distance, in km
func getLocation(latIn, lngIn, theta, distance float64) (float64, float64) {

	var (
		R = 6371.0 * 1000.0 // m

		lngOut float64
		latOut float64
	)

	distance = distance * 1000.0

	theta = theta * (math.Pi / 180.0)

	latIn = latIn * (math.Pi / 180.0)
	lngIn = lngIn * (math.Pi / 180.0)

	latOut = math.Asin(math.Sin(latIn)*math.Cos(distance/R) + math.Cos(latIn)*math.Cos(theta)*math.Sin(distance/R))

	lngOut = lngIn + math.Atan((math.Cos(latIn)*math.Sin(theta)*math.Sin(distance/R))/(math.Cos(distance/R)-math.Sin(latIn)*math.Sin(latOut)))

	latOut = latOut * (180.0 / math.Pi) * 1.000002
	lngOut = lngOut * (180.0 / math.Pi) * 1.000002

	// fmt.Printf("Output location: [%.7f, %.7f]\n", latOut, lngOut)

	return latOut, lngOut

}
