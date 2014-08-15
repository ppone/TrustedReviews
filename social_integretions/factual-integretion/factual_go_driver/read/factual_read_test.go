package read

import (
	"testing"
)

func TestCommaStringFromStringArray(t *testing.T) {
	s, err := commaStringFromStringArray([]string{"hello", "goodbye", "soso"})
	if err != nil {
		t.Error(err)
	}

	assertEqual("hello,goodbye,soso", s, t, "comma string from array fails")

	s, err = commaStringFromStringArray([]string{"hello,", "goodbye", "soso"})

	assertNotEqual(nil, err, t, "converstion allowed invalid arguments => ")

	s, err = commaStringFromStringArray([]string{"hello", "go,odbye", "soso"})

	assertNotEqual(nil, err, t, "converstion allowed invalid arguments => ")

	s, err = commaStringFromStringArray([]string{"hello", "goodbye", "s,oso"})

	assertNotEqual(nil, err, t, "converstion allowed invalid arguments => ")
}

func TestAddKey(t *testing.T) {
	r, err := NewRead("restaurants-us")

	if err != nil {
		t.Error(err)
	}

	r, err = r.AddKey("ASFW#Q1212987SFSDFLJLKJVXC")

	assertEqual("ASFW#Q1212987SFSDFLJLKJVXC", r.key, t, "Key not stored properly in Read struct ")

	r, err = r.AddKey("")

	assertNotEqual(nil, err, t, "error was not thrown")

}

func TestAddUser(t *testing.T) {
	r, err := NewRead("restaurants-us")

	if err != nil {
		t.Error(err)
	}

	r, err = r.AddUser("John Doe")

	assertEqual("John Doe", r.user, t, "Key not stored properly in Read struct ")

	r, err = r.AddUser("")

	assertNotEqual(nil, err, t, "error was not thrown")

}

func TestAddGeo(t *testing.T) {

	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r, err = r.AddGeoPoint(12.4234232, 232.000000000)

	if err != nil {
		t.Error(err)
	}

	assertEqual(r.geo, "geo={\"$point\":[12.4234232,232]}", t, "Geo not properly stored in read struct")

	r, err = r.AddGeoPoint(12.4234232, 232.000000001)

	if err != nil {
		t.Error(err)
	}

	assertEqual(r.geo, "geo={\"$point\":[12.4234232,232.000000001]}", t, "Geo not properly stored in read struct")

	r, err = r.AddGeoCircle(1823.123211244, 93.123211244, 25)
	if err != nil {
		t.Error(err)
	}

	assertEqual(r.geo, "geo={\"$circle\":{\"$center\":[1823.123211244,93.123211244],\"$meters\":25}}", t, "Geo not properly stored in read struct")

	r, err = r.AddGeoCircle(1823.000000001, 93.000000000, 125)
	if err != nil {
		t.Error(err)
	}

	assertEqual(r.geo, "geo={\"$circle\":{\"$center\":[1823.000000001,93],\"$meters\":125}}", t, "Geo not properly stored in read struct")

	r, err = r.AddGeoRectangle(2342.12312312, 16.323792837921, 89.2131231234, 901.232312318)
	if err != nil {
		t.Error(err)
	}

	assertEqual(r.geo, "geo={\"$rect\":[[2342.12312312,16.323792837921],[89.2131231234,901.232312318]]}", t, "Geo not properly stored in read struct")

	r, err = r.AddGeoRectangle(2342.12312312, 16.323792837921, 89.2131231234, 901.232312318)
	if err != nil {
		t.Error(err)
	}

	assertEqual(r.geo, "geo={\"$rect\":[[2342.12312312,16.323792837921],[89.2131231234,901.232312318]]}", t, "Geo not properly stored in read struct")

	r, err = r.AddGeoRectangle(2342.0000000001, 16.0000000000, 89.2131231234, 901.232312318)
	if err != nil {
		t.Error(err)
	}

	assertEqual(r.geo, "geo={\"$rect\":[[2342.0000000001,16],[89.2131231234,901.232312318]]}", t, "Geo not properly stored in read struct")
}

func TestAddLimit(t *testing.T) {
	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r = r.AddLimit(23123)

	assertEqual(r.limit, "limit=23123", t, "Limit not properly stored in read struct")

	r = r.AddLimit(231)

	assertNotEqual(r.limit, "limit=23123", t, "Limit not properly stored in read struct")

}

func TestAddThreshold(t *testing.T) {
	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r, err = r.AddThreshold("confident")

	if err != nil {
		t.Error(err)
	}

	assertEqual(r.threshold, "threshold=confident", t, "threshold not properly stored in read struct")

	r, err = r.AddThreshold("default")

	if err != nil {
		t.Error(err)
	}

	assertEqual(r.threshold, "threshold=default", t, "threshold not properly stored in read struct")

	r, err = r.AddThreshold("comprehensive")

	if err != nil {
		t.Error(err)
	}

	assertEqual(r.threshold, "threshold=comprehensive", t, "threshold not properly stored in read struct")

	r, err = r.AddThreshold("throwanerrorplease")

	assertNotEqual(nil, err, t, "threshold is accepting invalid values")

}

func TestAddOffset(t *testing.T) {
	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r = r.AddOffset(100)

	assertEqual(r.offset, "offset=100", t, "Limit not properly stored in read struct")

	r = r.AddLimit(300)

	assertNotEqual(r.offset, "offset=500", t, "Limit not properly stored in read struct")

}

func TestIncludeCount(t *testing.T) {

	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r = r.AddIncludeCount(true)

	assertEqual(r.include_count, "include_count=true", t, "Include Count not properly stored in read struct")

	r = r.AddIncludeCount(false)

	assertEqual(r.include_count, "include_count=false", t, "Include Count not properly stored in read struct")

}

func TestSort(t *testing.T) {

	s1 := NewSort("country", "asc")
	s2 := NewSort("age", "desc")
	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r, err = r.AddSort(s1, s2)

	if err != nil {
		t.Error(err)
	}
	assertEqual("sort=country:asc,age:desc", r.sort, t, "sort not stored properly")

	s3 := NewSort("$distance", "")
	s4 := NewSort("$relevance", "")

	r.query = "dummyq"
	r.geo = "dummygeo"

	r, err = r.AddSort(s3, s4)

	if err != nil {
		t.Error(err)
	}

	assertEqual("sort=$distance,$relevance", r.sort, t, "sort not stored properly")

	s5 := NewSort("distance", "40")
	s6 := NewSort("placerank", "50")

	r.query = "dummyq"
	r.geo = "dummygeo"

	r, err = r.AddSort(s5, s6)

	if err != nil {
		t.Error(err)
		return
	}

	assertEqual("sort={\"distance\":40,\"placerank\":50}", r.sort, t, "sort not stored properly")

	s7 := NewSort("placerank", "desc")

	r, err = r.AddSort(s7)

	if err != nil {
		t.Error(err)
		return
	}

	assertEqual("sort=placerank:desc", r.sort, t, "sort not stored properly")

	s1 = NewSort("country", "asc")
	s2 = NewSort("age", "a")
	r, err = NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r, err = r.AddSort(s1, s2)

	assertNotEqual(err, nil, t, "age:a should have thrown an error, alas none were thrown")

	s1 = NewSort("$distance", "asc")
	s2 = NewSort("$relevance", "desc")
	r, err = NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r.query = ""
	r.geo = "test"

	r, err = r.AddSort(s1, s2)

	assertNotEqual(err, nil, t, "$relevance:desc should have thrown an error, alas none were thrown")

	s1 = NewSort("distance", "asc")
	s2 = NewSort("$relevance", "desc")
	r, err = NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r.query = "sdf"
	r.geo = "test"

	r, err = r.AddSort(s1, s2)

	assertNotEqual(err, nil, t, "distance:asc should have thrown an error, alas none were thrown")

	s1 = NewSort("distance", "50")
	s2 = NewSort("placerank", "asc")
	r, err = NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r.geo = "test"

	r, err = r.AddSort(s1, s2)

	assertNotEqual(err, nil, t, "placerank:asc should have thrown an error, alas none were thrown")

}

func TestRead(t *testing.T) {

	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}
	assertEqual(r.fact_table.ToJson(), "/t/restaurants-us", t, "Factual Table not properly stored in read struct")

	r, err = NewRead("prone to error")
	if err == nil {
		t.Error("Read struct should have thrown an error")
	}

}

func TestSelect(t *testing.T) {
	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r, err = r.AddSelect("name", "address")

	if err != nil {
		t.Error(err)
	}

	assertEqual(r.selectq, "select=name,address", t, "Query not properly stored in read struct")

	r, err = r.AddSelect("city")

	if err != nil {
		t.Error(err)
	}

	assertEqual(r.selectq, "select=city", t, "Query not properly stored in read struct")

	r, err = r.AddSelect("name,address")

	assertNotEqual(nil, err, t, "converstion allowed invalid arguments => ")

}

func TestQuery(t *testing.T) {
	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r, err = r.AddQuery("Coffee Santa Monica")

	if err != nil {
		t.Error(err)
	}

	assertEqual(r.query, "q=Coffee Santa Monica", t, "Query not properly stored in read struct")

	r, err = r.AddQuery("Coffee", "Tea")

	if err != nil {
		t.Error(err)
	}

	assertEqual(r.query, "q=Coffee,Tea", t, "Query not properly stored in read struct")

	r, err = r.AddQuery("Coffee,Tea")

	assertNotEqual(nil, err, t, "converstion allowed invalid arguments => ")

}

func TestFilter(t *testing.T) {
	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r = r.AddFilterBlank("tel", true)
	assertEqual(r.filter, "filters={\"tel\":{\"$blank\":true}}", t, "Filter not properly stored in read struct")

	r = r.AddFilterBlank("tel", false)
	assertEqual(r.filter, "filters={\"tel\":{\"$blank\":false}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterBeginsWith("name", "b")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"name\":{\"$bw\":\"b\"}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterBeginsWithAny("name", "lt", "sg", "cpt")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"name\":{\"$bwin\":[\"lt\",\"sg\",\"cpt\"]}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterEqualTo("region", "CA")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"region\":{\"$eq\":\"CA\"}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterExcludes("category_ids", 9)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"category_ids\":{\"$excludes\":9}}", t, "Fitler not properly stored in read struct")

	r, err = r.AddFilterExcludes("category_ids", 9)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"category_ids\":{\"$excludes\":9}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterExcludesAny("category_ids", 318, 321)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"category_ids\":{\"$excludes_any\":[318,321]}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterGreaterThan("rating", 7.5)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"rating\":{\"$gt\":7.5}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterGreaterThanEqual("rating", 7.5)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"rating\":{\"$gte\":7.5}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterEqualsAnyOf("region", "MA", "VT", "NH", "RI", "CT")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"region\":{\"$in\":[\"MA\",\"VT\",\"NH\",\"RI\",\"CT\"]}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterIncludes("category_ids", 10)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"category_ids\":{\"$includes\":10}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterIncludesAny("category_ids", 10, 100)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"category_ids\":{\"$includes_any\":[10,100]}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterLessThan("age", 50)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"age\":{\"$lt\":50}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterLessThanEqual("age", 7.5)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"age\":{\"$lte\":7.5}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterNotBeginWith("name", "Mr.")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"name\":{\"$nbw\":\"Mr.\"}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterNotBeginWithAny("class", "beginner", "intermediate")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"class\":{\"$nbwin\":[\"beginner\",\"intermediate\"]}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterNotEqualTo("region", "CA")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"region\":{\"$neq\":\"CA\"}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterNotEqualAnyOf("locality", "Los Angeles", "Santa Monica")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"locality\":{\"$nin\":[\"Los Angeles\",\"Santa Monica\"]}}", t, "Filter not properly stored in read struct")

	r, err = r.AddFilterSearch("name", "Charles")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"name\":{\"$search\":\"Charles\"}}", t, "Filter not properly stored in read struct")

}

func assertEqual(expected, got interface{}, t *testing.T, errorMessage string) {
	if expected != got {
		t.Errorf("%s => expected %s, got %s", errorMessage, expected, got)
	}
}

func assertNotEqual(expected, got interface{}, t *testing.T, errorMessage string) {
	if expected == got {
		t.Errorf("%s => expected %s, got %s", errorMessage, expected, got)
	}
}
