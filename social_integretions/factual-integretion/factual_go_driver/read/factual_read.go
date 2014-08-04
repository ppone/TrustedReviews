package read

import (
	"../filters"
	"../geocode"
	"../table"
	"errors"
)

type Read struct {
	query         string
	filter        string
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
		return nil, errors.New("Could not create new Read due to error creating a new table => ")
	}

	read := new(Read)
	read.fact_table = tab

	return read, nil

}

func (R *Read) AddQuery(query string) *Read {
	R.query = "q=" + query
	return R
}

func (R *Read) AddFilterBlank(keyword string, b bool) *Read {
	f := filters.Blank(keyword, b)
	R.filter = f
	return R
}
func (R *Read) AddFilterBeginsWith(keyword string, value string) (*Read, error) {
	f, err := filters.BeginsWith(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterBeginsWithAny(keyword string, values ...string) (*Read, error) {
	f, err := filters.BeginsWithAny(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterEqualTo(keyword string, value string) (*Read, error) {
	f, err := filters.EqualTo(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterExcludes(keyword string, value interface{}) (*Read, error) {
	f, err := filters.Excludes(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterExcludesAny(keyword string, values ...interface{}) (*Read, error) {
	f, err := filters.ExcludesAny(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterGreaterThan(keyword string, value interface{}) (*Read, error) {
	f, err := filters.GreaterThan(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterGreaterThanEqual(keyword string, value interface{}) (*Read, error) {
	f, err := filters.GreaterThanEqual(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterEqualsAnyOf(keyword string, values ...interface{}) (*Read, error) {
	f, err := filters.EqualsAnyOf(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterIncludes(keyword string, value interface{}) (*Read, error) {
	f, err := filters.Includes(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterIncludesAny(keyword string, values ...interface{}) (*Read, error) {
	f, err := filters.IncludesAny(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterLessThan(keyword string, value interface{}) (*Read, error) {
	f, err := filters.LessThan(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterLessThanEqual(keyword string, value interface{}) (*Read, error) {
	f, err := filters.LessThanEqual(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterNotBeginWith(keyword string, value string) (*Read, error) {
	f, err := filters.NotBeginWith(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterNotBeginWithAny(keyword string, values ...string) (*Read, error) {
	f, err := filters.NotBeginWithAny(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterNotEqualTo(keyword string, value interface{}) (*Read, error) {
	f, err := filters.NotEqualTo(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterNotEqualAnyOf(keyword string, values ...interface{}) (*Read, error) {
	f, err := filters.NotEqualAnyOf(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddFilterSearch(keyword string, value string) (*Read, error) {
	f, err := filters.Search(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *Read) AddLimit(limit int) *Read {
	R.limit = limit
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

func (R *Read) ToJsonFromGeo() (string, error) {
	return "", nil
}
