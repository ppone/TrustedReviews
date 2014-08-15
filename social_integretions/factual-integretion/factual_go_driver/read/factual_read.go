package read

import (
	"../filters"
	"../geocode"
	"../table"
	"errors"
	"strconv"
	"strings"
)

type read struct {
	query         string
	filter        string
	limit         string
	fact_table    string
	geo           string
	include_count string
	sort          string
	selectq       string
	offset        string
	threshold     string
	key           string
	user          string
}

type sortData struct {
	column            string
	sortOrderOrNumber string
}

func commaStringFromStringArray(stringArray []string) (string, error) {

	commaString := ""

	for _, s := range stringArray {
		if strings.Contains(s, ",") {
			return "", errors.New("string arguments cannot contain comma ','")
		}
		commaString += s + ","
	}

	return strings.TrimRight(commaString, ","), nil

}

func NewRead(tableName string) (*read, error) {
	tab, err := table.NewTable(tableName)

	if err != nil {
		return nil, errors.New("Could not create new Read due to error creating a new table => ")
	}

	read := new(read)
	read.fact_table = tab.String()

	return read, nil

}

func NewSort(column, sortOrderOrNumber string) sortData {
	return sortData{column, sortOrderOrNumber}
}

/*
type read struct {
	query         string
	filter        string
	limit         string
	fact_table    table.FactTable
	geo           string
	include_count string
	sort          string
	selectq       string
	offset        string
	threshold     string
	key           string
	user          string
}*/

func (R *read) String() string {
	s := ""

	if R.fact_table != "" {
		s += R.fact_table
		s += "?"

	}

	if R.query != "" {
		s += R.query + "&"
	}

	if R.filter != "" {
		s += R.filter + "&"
	}

	if R.limit != "" {
		s += R.limit + "&"
	}

	if R.geo != "" {
		s += R.geo + "&"
	}

	if R.include_count != "" {
		s += R.include_count + "&"
	}

	if R.sort != "" {
		s += R.sort + "&"
	}

	if R.selectq != "" {
		s += R.selectq + "&"
	}

	if R.offset != "" {
		s += R.offset + "&"
	}

	if R.threshold != "" {
		s += R.threshold + "&"
	}

	if R.key != "" {
		s += R.key + "&"
	}

	if R.user != "" {
		s += R.user + "&"
	}

	return strings.TrimRight(s, "&")
}

func (R *read) AddKey(key string) (*read, error) {

	if key == "" {
		return nil, errors.New("key argument is empty")
	}

	R.key = "KEY=" + key

	return R, nil

}

/* Future might need validate the column we sort on instead waiting for API to return error */

func (R *read) AddSort(sorts ...sortData) (*read, error) {

	s := ""

	for _, sort := range sorts {

		column := sort.column
		sortOrderOrNumber := sort.sortOrderOrNumber

		if column == "$distance" && R.geo == "" {
			return nil, errors.New("$distance cannot be used as geo is empty")
		}

		if column == "$relevance" && R.query == "" {
			return nil, errors.New("$relevance cannot be used as query is empty")
		}

		if sort.sortOrderOrNumber == "" {
			s += column + ","
			continue
		}

		if column == "distance" {
			_, err := strconv.Atoi(sortOrderOrNumber)
			if err != nil {
				return nil, errors.New("must use number with distance")
			}

			s += "\"" + column + "\":" + sortOrderOrNumber + ","
			continue
		}

		if column == "placerank" {
			_, err := strconv.Atoi(sortOrderOrNumber)
			if err != nil && sortOrderOrNumber != "desc" {
				return nil, errors.New("must use number with placerank or specify order to be 'desc'")
			}

			if err == nil {
				s += "\"" + column + "\":" + sortOrderOrNumber + ","
			} else {
				s += column + ":" + sortOrderOrNumber + ","
			}

			continue

		}

		if sortOrderOrNumber != "asc" && sortOrderOrNumber != "desc" {
			return nil, errors.New("valid sorting values are 'asc' or 'desc' for nonvirtual columns")
		}

		s += column + ":" + sortOrderOrNumber + ","

	}

	s = strings.TrimRight(s, ",")

	if strings.Contains(s, "\"distance\":") || strings.Contains(s, "\"placerank\":") {
		s = "sort={" + s + "}"
	} else {
		s = "sort=" + s
	}

	R.sort = s

	return R, nil

}

func (R *read) AddUser(user string) (*read, error) {

	if user == "" {
		return nil, errors.New("user argument is empty")
	}

	R.user = "user=" + user

	return R, nil

}

func (R *read) AddQuery(queries ...string) (*read, error) {
	query, err := commaStringFromStringArray(queries)
	if err != nil {
		return nil, errors.New("Queries has some invalid arguments =>" + err.Error())
	}
	R.query = "q=" + query
	return R, nil
}

func (R *read) AddSelect(selects ...string) (*read, error) {

	s, err := commaStringFromStringArray(selects)
	if err != nil {
		return nil, errors.New("Select has some invalid arguments =>" + err.Error())
	}
	R.selectq = "select=" + s
	return R, nil
}

func (R *read) AddIncludeCount(include_count bool) *read {
	ic := ""

	if include_count {
		ic = "include_count=true"
	} else {
		ic = "include_count=false"
	}

	R.include_count = ic
	return R
}

func (R *read) AddLimit(limit int) *read {
	R.limit = "limit=" + strconv.Itoa(limit)
	return R
}

func (R *read) AddOffset(offset int) *read {
	R.offset = "offset=" + strconv.Itoa(offset)
	return R
}

func (R *read) AddThreshold(threshold string) (*read, error) {
	t := ""

	switch threshold {
	case "confident", "default", "comprehensive":
		t = "threshold=" + threshold
	default:
		return nil, errors.New("not valid threshold value, only 'confident','default', and 'comprehensive' allowed")

	}

	R.threshold = t

	return R, nil

}

func (R *read) AddGeoPoint(longitude, latitude float64) (*read, error) {
	g := geocode.NewGeoPoint(longitude, latitude)

	s, err := g.ToJsonFromGeo()

	if err != nil {
		return nil, err
	}

	R.geo = s

	return R, nil
}

func (R *read) AddGeoCircle(longitude, latitude float64, radius int16) (*read, error) {
	g := geocode.NewGeoCircle(longitude, latitude, radius)

	s, err := g.ToJsonFromGeo()

	if err != nil {
		return nil, err
	}

	R.geo = s

	return R, nil
}

func (R *read) AddGeoRectangle(topRightLongitude, topRightLatitude, leftBottomLongitude, leftBottomLatitude float64) (*read, error) {
	g := geocode.NewGeoRectangle(topRightLongitude, topRightLatitude, leftBottomLongitude, leftBottomLatitude)

	s, err := g.ToJsonFromGeo()

	if err != nil {
		return nil, err
	}

	R.geo = s

	return R, nil
}

func (R *read) AddFilterBlank(keyword string, b bool) *read {
	f := filters.Blank(keyword, b)
	R.filter = f
	return R
}
func (R *read) AddFilterBeginsWith(keyword string, value string) (*read, error) {
	f, err := filters.BeginsWith(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterBeginsWithAny(keyword string, values ...string) (*read, error) {
	f, err := filters.BeginsWithAny(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterEqualTo(keyword string, value string) (*read, error) {
	f, err := filters.EqualTo(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterExcludes(keyword string, value interface{}) (*read, error) {
	f, err := filters.Excludes(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterExcludesAny(keyword string, values ...interface{}) (*read, error) {
	f, err := filters.ExcludesAny(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterGreaterThan(keyword string, value interface{}) (*read, error) {
	f, err := filters.GreaterThan(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterGreaterThanEqual(keyword string, value interface{}) (*read, error) {
	f, err := filters.GreaterThanEqual(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterEqualsAnyOf(keyword string, values ...interface{}) (*read, error) {
	f, err := filters.EqualsAnyOf(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterIncludes(keyword string, value interface{}) (*read, error) {
	f, err := filters.Includes(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterIncludesAny(keyword string, values ...interface{}) (*read, error) {
	f, err := filters.IncludesAny(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterLessThan(keyword string, value interface{}) (*read, error) {
	f, err := filters.LessThan(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterLessThanEqual(keyword string, value interface{}) (*read, error) {
	f, err := filters.LessThanEqual(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterNotBeginWith(keyword string, value string) (*read, error) {
	f, err := filters.NotBeginWith(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterNotBeginWithAny(keyword string, values ...string) (*read, error) {
	f, err := filters.NotBeginWithAny(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterNotEqualTo(keyword string, value interface{}) (*read, error) {
	f, err := filters.NotEqualTo(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterNotEqualAnyOf(keyword string, values ...interface{}) (*read, error) {
	f, err := filters.NotEqualAnyOf(keyword, values...)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}

func (R *read) AddFilterSearch(keyword string, value string) (*read, error) {
	f, err := filters.Search(keyword, value)
	if err != nil {
		return nil, err
	}
	R.filter = f
	return R, nil
}
