package table

import (
	"testing"
)

func TestListAllTables(t *testing.T) {
	//tableList := []string{"places","places-"}
	FactTableCheck := map[string]int{

		"places":              1,
		"places-edge":         1,
		"place-categories":    1,
		"restaurants-us":      1,
		"restaurants-us-edge": 1,
		"restaurants-gb":      1,
		"hotels-us":           1,
		"world-geographies":   1,
		"crosswalk":           1,
		"products-cpg":        1,
		"products-crosswalk":  1,
	}

	tablesToCheck := ListAllTables()

	for _, tables := range tablesToCheck {
		_, ok := FactTableCheck[tables]
		errMessage := "Checking if the list of all tables is correct, table = " + tables
		assertEqual(ok, true, t, errMessage)
	}

}

var jsonTestErrorMessage = "Error while testing toJson Method"

func TestCreateTablePlaces(t *testing.T) {
	tableName := "places"
	tab, err := NewTable(tableName)

	assertNoError(err, tableName, t)
	testToJsonTablePlaces(tab, t)
}

func testToJsonTablePlaces(tab FactTable, t *testing.T) {
	assertEqual("/t/places", tab.ToJson(), t, jsonTestErrorMessage)
}

func TestCreateTablePlacesEdge(t *testing.T) {
	tableName := "places-edge"
	tab, err := NewTable(tableName)
	assertNoError(err, tableName, t)
	testToJsonTablePlacesEdge(tab, t)
}

func testToJsonTablePlacesEdge(tab FactTable, t *testing.T) {
	assertEqual("/t/places-edge", tab.ToJson(), t, jsonTestErrorMessage)
}

func TestCreateTablePlacesCategories(t *testing.T) {
	tableName := "place-categories"
	tab, err := NewTable(tableName)
	assertNoError(err, tableName, t)
	testToJsonTablePlacesCategories(tab, t)
}

func testToJsonTablePlacesCategories(tab FactTable, t *testing.T) {
	assertEqual("/t/place-categories", tab.ToJson(), t, jsonTestErrorMessage)
}

func TestCreateTableRestaurantsUS(t *testing.T) {
	tableName := "restaurants-us"
	tab, err := NewTable(tableName)
	assertNoError(err, tableName, t)
	testToJsonTableRestaurantsUS(tab, t)
}

func testToJsonTableRestaurantsUS(tab FactTable, t *testing.T) {
	assertEqual("/t/restaurants-us", tab.ToJson(), t, jsonTestErrorMessage)
}

func TestCreateTableRestaurantsUSEdge(t *testing.T) {
	tableName := "restaurants-us-edge"
	tab, err := NewTable(tableName)
	assertNoError(err, tableName, t)
	testToJsonTableRestaurantsUSEdge(tab, t)
}

func testToJsonTableRestaurantsUSEdge(tab FactTable, t *testing.T) {
	assertEqual("/t/restaurants-us-edge", tab.ToJson(), t, jsonTestErrorMessage)
}

func TestCreateTableRestaurantsUSGb(t *testing.T) {
	tableName := "restaurants-gb"
	tab, err := NewTable(tableName)
	assertNoError(err, tableName, t)
	testToJsonTableRestaurantsUSGb(tab, t)
}

func testToJsonTableRestaurantsUSGb(tab FactTable, t *testing.T) {
	assertEqual("/t/restaurants-gb", tab.ToJson(), t, jsonTestErrorMessage)
}

func TestCreateTableHotelsUs(t *testing.T) {
	tableName := "hotels-us"
	tab, err := NewTable(tableName)
	assertNoError(err, tableName, t)
	testToJsonTableHotelsUs(tab, t)
}

func testToJsonTableHotelsUs(tab FactTable, t *testing.T) {
	assertEqual("/t/hotels-us", tab.ToJson(), t, jsonTestErrorMessage)
}

func TestCreateTableWorldGeographies(t *testing.T) {
	tableName := "world-geographies"
	tab, err := NewTable(tableName)
	assertNoError(err, tableName, t)
	testToJsonTableWorldGeographies(tab, t)
}

func testToJsonTableWorldGeographies(tab FactTable, t *testing.T) {
	assertEqual("/t/world-geographies", tab.ToJson(), t, jsonTestErrorMessage)
}

func TestCreateTableCrosswalk(t *testing.T) {
	tableName := "crosswalk"
	tab, err := NewTable(tableName)
	assertNoError(err, tableName, t)
	testToJsonTableCrosswalk(tab, t)
}

func testToJsonTableCrosswalk(tab FactTable, t *testing.T) {
	assertEqual("/t/crosswalk", tab.ToJson(), t, jsonTestErrorMessage)
}

func TestCreateTableProductsCpg(t *testing.T) {
	tableName := "products-cpg"
	tab, err := NewTable(tableName)
	assertNoError(err, tableName, t)
	testToJsonTableProductsCpg(tab, t)
}

func testToJsonTableProductsCpg(tab FactTable, t *testing.T) {
	assertEqual("/t/products-cpg", tab.ToJson(), t, jsonTestErrorMessage)
}

func TestCreateTableProductsCrosswalk(t *testing.T) {
	tableName := "products-crosswalk"
	tab, err := NewTable(tableName)
	assertNoError(err, tableName, t)
	testToJsonTableProductsCrosswalk(tab, t)
}

func testToJsonTableProductsCrosswalk(tab FactTable, t *testing.T) {
	assertEqual("/t/products-crosswalk", tab.ToJson(), t, jsonTestErrorMessage)
}

func assertNoError(err error, tableName string, t *testing.T) {
	if err != nil {
		t.Errorf("Could not create table %s", tableName)
	}
}

func assertEqual(expected, got interface{}, t *testing.T, errorMessage string) {
	if expected != got {
		t.Errorf("%s => expected %s, got %s", errorMessage, expected, got)
	}
}
