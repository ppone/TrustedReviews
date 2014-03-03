package geocode

import (
	"encoding/json"
)

type geoPoint struct {
	Point [2]float64 `json:"$point"`
}

type geoCircleCenter struct {
	Center [2]float64 `json:"$center"`
	Radius int16      `json:"$meters"`
}

type geoCircle struct {
	Circle geoCircleCenter `json:"$circle"`
}

type geoRectangle struct {
	Rect [2][2]float64 `json:"$rect"`
}

func NewGeoPoint(longitude, latitude float64) geopoint {
	location := [2]float64{longitude, latitude}
	return geopoint{location}
}

func NewGeoCircle(longitude, latitude float64, radius int16) geocircle {
	location := [2]float64{longitude, latitude}
	geocirclec := geocirclecenter{location, radius}
	return geocircle{geocirclec}
}

func NewGeoRectangle(topRightLongitude, topRightLatitude, leftBottomLongitude, leftBottomLatitude float64) georectangle {
	rectangle := [2][2]float64{{topRightLongitude, topRightLatitude}, {leftBottomLongitude, leftBottomLatitude}}
	return georectangle{rectangle}
}

func (GeoPoint geoPoint) ToJson() (string, error) {
	jsonecoded, err := json.Marshal(GeoPoint)

	if err != nil {
		return "", err
	}

	return "geo=" + string(jsonecoded), nil
}

func (GeoCircle geoCircle) ToJson() (string, error) {
	jsonecoded, err := json.Marshal(GeoCircle)

	if err != nil {
		return "", err
	}

	return "geo=" + string(jsonecoded), nil
}

func (GeoRect geoRectangle) ToJson() (string, error) {
	jsonecoded, err := json.Marshal(GeoRect)

	if err != nil {
		return "", err
	}

	return "geo=" + string(jsonecoded), nil
}