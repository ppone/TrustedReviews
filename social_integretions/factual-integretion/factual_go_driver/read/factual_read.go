package main

import (
	"../geocode"
	"../table"
	"errors"
	"fmt"
)

type Read struct {
	query         string
	filters       string
	limit         int
	fact_table    table.FactTable
	geo           geocode.GeoShape
	include_count bool
	sort          string
	offset        int
}

func NewRead(tableName string) (*Read, error) {
	tab, err := table.NewTable(tableName)

	if err != nil {
		return nil, errors.New("Could not create new Read due to error creating a new table")
	}

	read := &Read{}
	read.fact_table = tab

	return read, nil

}

func (R *Read) AddQuery(query string) *Read {
	R.query = query
	return R
}

func (R *Read) AddFilters(filters string) *Read {
	R.filters = filters
	return R
}
func (R *Read) AddLimit(limit int) *Read {
	R.filters = filters
	return R
}

func (R *Read) AddGeoPoint(longitude, latitude float64) *Read {
	R.geo = geocode.NewGeoPoint(longitude, latitude)
	return R
}

func (R *Read) AddGeoCircle(longitude, latitude float64, radius int16) *Read {
	R.geo = geocode.NewGeoPoint(longitude, latitude)
	return R
}

func (R *Read) AddGeoRectangle(topRightLongitude, topRightLatitude, leftBottomLongitude, leftBottomLatitude float64) *Read {
	R.geo = geocode.NewGeoRectangle(topRightLongitude, topRightLatitude, leftBottomLongitude, leftBottomLatitude)
	return R
}

func (R *Read) AddGeo(geoType string) *Read {
	R.filters = filters
	return R
}

func (R *Read) ToJson() (string, error) {

}
func main() {

	r := read{}
	r.query = "hello world"
	r.limit = 34
	r.fact_table, _ = table.NewTable("places")
	r.geo = geocode.NewGeoPoint(324.1, 213.23)

	fmt.Println(r)

}
