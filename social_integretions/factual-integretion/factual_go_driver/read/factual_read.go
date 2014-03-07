package main

import (
	"../geocode"
	"../table"
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

func (R *Read) AddQuery(query string) {
	R.query = query
}

func (R *Read) AddFilters(filters string) {
	R.filters = filters
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
