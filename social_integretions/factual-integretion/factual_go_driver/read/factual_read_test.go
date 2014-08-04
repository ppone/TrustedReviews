package read

import (
	"testing"
)

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

func TestQuery(t *testing.T) {
	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r = r.AddQuery("Coffee Santa Monica")

	assertEqual(r.query, "q=Coffee Santa Monica", t, "Query not properly stored in read struct")

	r = r.AddQuery("Coffee,Tea")

	assertEqual(r.query, "q=Coffee,Tea", t, "Query not properly stored in read struct")

}

func TestFilter(t *testing.T) {
	r, err := NewRead("restaurants-us")
	if err != nil {
		t.Error(err)
	}

	r = r.AddFilterBlank("tel", true)
	assertEqual(r.filter, "filters={\"tel\":{\"$blank\":true}}", t, "Query not properly stored in read struct")

	r = r.AddFilterBlank("tel", false)
	assertEqual(r.filter, "filters={\"tel\":{\"$blank\":false}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterBeginsWith("name", "b")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"name\":{\"$bw\":\"b\"}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterBeginsWithAny("name", "lt", "sg", "cpt")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"name\":{\"$bwin\":[\"lt\",\"sg\",\"cpt\"]}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterEqualTo("region", "CA")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"region\":{\"$eq\":\"CA\"}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterExcludes("category_ids", 9)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"category_ids\":{\"$excludes\":9}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterExcludes("category_ids", 9)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"category_ids\":{\"$excludes\":9}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterExcludesAny("category_ids", 318, 321)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"category_ids\":{\"$excludes_any\":[318,321]}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterGreaterThan("rating", 7.5)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"rating\":{\"$gt\":7.5}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterGreaterThanEqual("rating", 7.5)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"rating\":{\"$gte\":7.5}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterEqualsAnyOf("region", "MA", "VT", "NH", "RI", "CT")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"region\":{\"$in\":[\"MA\",\"VT\",\"NH\",\"RI\",\"CT\"]}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterIncludes("category_ids", 10)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"category_ids\":{\"$includes\":10}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterIncludesAny("category_ids", 10, 100)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"category_ids\":{\"$includes_any\":[10,100]}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterLessThan("age", 50)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"age\":{\"$lt\":50}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterLessThanEqual("age", 7.5)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"age\":{\"$lte\":7.5}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterNotBeginWith("name", "Mr.")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"name\":{\"$nbw\":\"Mr.\"}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterNotBeginWithAny("class", "beginner", "intermediate")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"class\":{\"$nbwin\":[\"beginner\",\"intermediate\"]}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterNotEqualTo("region", "CA")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"region\":{\"$neq\":\"CA\"}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterNotEqualAnyOf("locality", "Los Angeles", "Santa Monica")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"locality\":{\"$nin\":[\"Los Angeles\",\"Santa Monica\"]}}", t, "Query not properly stored in read struct")

	r, err = r.AddFilterSearch("name", "Charles")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	assertEqual(r.filter, "filters={\"name\":{\"$search\":\"Charles\"}}", t, "Query not properly stored in read struct")

}

func TestGeo(t *testing.T) {

}

func TestLimit(t *testing.T) {

}

func TestIncludeCount(t *testing.T) {

}

func TestSort(t *testing.T) {

}

func assertEqual(expected, got interface{}, t *testing.T, errorMessage string) {
	if expected != got {
		t.Errorf("%s => expected %s, got %s", errorMessage, expected, got)
	}
}
