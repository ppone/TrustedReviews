package main

import (
	"encoding/json"
	"fmt"
)

type geopoint struct {
	Point [2]float32 `json:"$point"`
}

type geocirclecenter struct {
	Center [2]float32 `json:"$center"`
	Radius int16      `json:"$meters"`
}

type geocircle struct {
	Circle geocirclecenter `json:"$circle"`
}

type georectangle struct {
	Rect [2][2]float32 `json:"$rect"`
}

func New_geopoint(longitude, latitude float32) *geopoint {
	location := [2]float32{longitude, latitude}
	return &geopoint{location}
}

func New_geocircle(longitude, latitude float32, radius int16) *geocircle {
	location := [2]float32{longitude, latitude}
	geocirclec := geocirclecenter{location, radius}
	return &geocircle{geocirclec}
}

func New_georectangle(toprightlongitude, toprightlatitude, leftbottomlongitude, leftbottomlatitude float32) *georectangle {
	rectangle := [2][2]float32{{toprightlongitude, toprightlatitude}, {leftbottomlongitude, leftbottomlatitude}}
	return &georectangle{rectangle}
}

func (GeoPoint geopoint) ToJson() (string, error) {
	jsonecoded, err := json.Marshal(GeoPoint)

	if err != nil {
		return "", err
	}

	return "geo=" + string(jsonecoded), nil
}

func (GeoCircle geocircle) ToJson() (string, error) {
	jsonecoded, err := json.Marshal(GeoCircle)

	if err != nil {
		return "", err
	}

	return "geo=" + string(jsonecoded), nil
}

func (GeoRect georectangle) ToJson() (string, error) {
	jsonecoded, err := json.Marshal(GeoRect)

	if err != nil {
		return "", err
	}

	return "geo=" + string(jsonecoded), nil
}

func main() {

	point := New_geopoint(343.22, 232.23)
	pointjson, err := point.ToJson()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pointjson)

}
