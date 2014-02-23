package geocode

import (
	"encoding/json"
)

type geopoint struct {
	Point [2]float64 `json:"$point"`
}

type geocirclecenter struct {
	Center [2]float64 `json:"$center"`
	Radius int16      `json:"$meters"`
}

type geocircle struct {
	Circle geocirclecenter `json:"$circle"`
}

type georectangle struct {
	Rect [2][2]float64 `json:"$rect"`
}

func New_geopoint(longitude, latitude float64) *geopoint {
	location := [2]float64{longitude, latitude}
	return &geopoint{location}
}

func New_geocircle(longitude, latitude float64, radius int16) *geocircle {
	location := [2]float64{longitude, latitude}
	geocirclec := geocirclecenter{location, radius}
	return &geocircle{geocirclec}
}

func New_georectangle(toprightlongitude, toprightlatitude, leftbottomlongitude, leftbottomlatitude float64) *georectangle {
	rectangle := [2][2]float64{{toprightlongitude, toprightlatitude}, {leftbottomlongitude, leftbottomlatitude}}
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
