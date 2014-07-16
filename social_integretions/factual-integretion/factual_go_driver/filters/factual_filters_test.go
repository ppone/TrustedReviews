package filters

import (
	"testing"
)

type testStruct struct {
	x int
	y int
}

var testFailStruct testStruct = testStruct{}

func TestTruncateLastZeroInFloatString(t *testing.T) {
	assertEqual("42", truncateLastZeroInFloatString("42"), t, "truncating last zero did not come up with the expected results")
	assertEqual("23.4", truncateLastZeroInFloatString("23.40"), t, "truncating last zero did not come up with the expected results")
	assertEqual("23.45", truncateLastZeroInFloatString("23.45"), t, "truncating last zero did not come up with the expected results")
	assertEqual("23.4", truncateLastZeroInFloatString("23.40"), t, "truncating last zero did not come up with the expected results")
	assertEqual("hel!!1212799", truncateLastZeroInFloatString("hel!!1212799"), t, "truncating last zero did not come up with the expected results")
	assertEqual("hel!!1212.799", truncateLastZeroInFloatString("hel!!1212.799"), t, "truncating last zero did not come up with the expected results")

}

func TestInferfaceSlice(t *testing.T) {
	//fmt.Println(NewInterfaceSlice(4, 2, "hello", true))
}

func TestCheckTypes(t *testing.T) {

	jsonType, err := checkTypes(23.4, numericT)
	assertNoError(err, t)
	if err == nil {
		assertEqual("float64", jsonType, t, "jsonType is incorrect")
	}
	jsonType, err = checkTypes("23.4", stringT)
	assertNoError(err, t)
	if err == nil {
		assertEqual("string", jsonType, t, "jsonType is incorrect")
	}
	jsonType, err = checkTypes(23, stringAndNumericT)
	assertNoError(err, t)
	if err == nil {
		assertEqual("int", jsonType, t, "jsonType is incorrect")
	}
	jsonType, err = checkTypes(23, numericT)
	assertNoError(err, t)
	if err == nil {
		assertEqual("int", jsonType, t, "jsonType is incorrect")
	}
	jsonType, err = checkTypes("23.4", numericT)
	assertError(err, t)
	jsonType, err = checkTypes(23.4, stringT)
	assertError(err, t)

	jsonType, err = checkTypes(true, stringAndNumericT)
	assertError(err, t)

	jsonType, err = checkTypes(testFailStruct, stringAndNumericT)
	assertError(err, t)
}

func TestCheckTypesInArray(t *testing.T) {
	//a := []float64{4.5, 45.2}
	//z, _ := covertSliceToInterfaceSlice(a)
	s := NewInterfaceSlice(4.5, 45.44)
	jsonType, err := checkTypesInArray(s, numericT)
	assertNoError(err, t)
	if err == nil {
		assertEqual("float64", jsonType, t, "jsonType is incorrect")
	}

	s = NewInterfaceSlice(4, 45)
	jsonType, err = checkTypesInArray(s, numericT)
	assertNoError(err, t)
	if err == nil {
		assertEqual("int", jsonType, t, "jsonType is incorrect")
	}

	s = NewInterfaceSlice("hello", "goodbye", "seeya")
	jsonType, err = checkTypesInArray(s, stringT)
	assertNoError(err, t)
	if err == nil {
		assertEqual("string", jsonType, t, "jsonType is incorrect")
	}

	s = NewInterfaceSlice("hello", "goodbye", "seeya")
	jsonType, err = checkTypesInArray(s, numericT)
	assertError(err, t)

	s = NewInterfaceSlice("hello", "goodbye", "seeya")
	jsonType, err = checkTypesInArray(s, stringAndNumericT)
	assertNoError(err, t)

	s = NewInterfaceSlice(4.5, 7.5)
	jsonType, err = checkTypesInArray(s, stringAndNumericT)
	assertNoError(err, t)

	s = NewInterfaceSlice(4, 7)
	jsonType, err = checkTypesInArray(s, stringAndNumericT)
	assertNoError(err, t)

	s = NewInterfaceSlice("23.3", "44.3")
	jsonType, err = checkTypesInArray(s, numericT)
	assertError(err, t)

	s = NewInterfaceSlice(4.5, 7.5)
	jsonType, err = checkTypesInArray(s, stringT)
	assertError(err, t)

	s = NewInterfaceSlice(4.5, "hello")
	jsonType, err = checkTypes(true, stringAndNumericT)
	assertError(err, t)

	s = NewInterfaceSlice(3, "2.5")
	jsonType, err = checkTypes(true, stringAndNumericT)
	assertError(err, t)

	s = NewInterfaceSlice(true, testFailStruct)
	jsonType, err = checkTypes(true, stringAndNumericT)
	assertError(err, t)

}

func TestFloatToString(t *testing.T) {

	assertEqual("3424.5", floatToString(3424.50), t, "Error in converting float to string")
	assertEqual("24.5", floatToString(24.50), t, "Error in converting float to string")
	assertEqual("99424.5", floatToString(99424.50), t, "Error in converting float to string")
	assertEqual("99424.78", floatToString(99424.777), t, "Error in converting float to string")
	assertEqual("73231.51", floatToString(73231.5123), t, "Error in converting float to string")
	assertEqual("100000.0", floatToString(99999.996), t, "Error in converting float to string")
	assertEqual("101.78", floatToString(101.775), t, "Error in converting float to string")
	assertEqual("6779.69", floatToString(6779.685), t, "Error in converting float to string")
	assertNotEqual("100000.0", floatToString(99999.99), t, "Error in converting float to string, values are not suppose to match")
}

func TestIntToString(t *testing.T) {

	assertEqual("23", intToString(23), t, "Error in converting int to string")
	assertEqual("234297345", intToString(234297345), t, "Error in converting int to string")
	assertEqual("723234234", intToString(723234234), t, "Error in converting int to string")
	assertEqual("1011010101", intToString(1011010101), t, "Error in converting int to string")

	assertNotEqual("23429734", intToString(2342982), t, "Error in converting int to string, values are not suppose to match")
	assertNotEqual("23429734", intToString(2342982), t, "Error in converting int to string, values are not suppose to match")
}

func TestBlank(t *testing.T) {
	assertEqual("filters={\"city\":{\"$blank\":true}}", Blank("city", true), t, "Error in testing Blank")
	assertEqual("filters={\"city\":{\"$blank\":false}}", Blank("city", false), t, "Error in testing Blank")

}

func TestBeginsWith(t *testing.T) {
	beginsWithOutput, err := BeginsWith("city", "low")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"city\":{\"$bw\":\"low\"}}", beginsWithOutput, t, "Error in testing BeginsWith")
	}

}
func TestBeginsWithAny(t *testing.T) {
	beginsWithOutput, err := BeginsWithAny("state", "mass", "cali", "flor")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"state\":{\"$bwin\":[\"mass\",\"cali\",\"flor\"]}}", beginsWithOutput, t, "Error in testing BeginsWith")
	}

}

func TestEqualsTo(t *testing.T) {
	beginsWithOutput, err := EqualTo("region", "CA")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"region\":{\"$eq\":\"CA\"}}", beginsWithOutput, t, "Error in testing EqualsTo")
	}

	beginsWithOutput, err = EqualTo("category_id", 232)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"category_id\":{\"$eq\":232}}", beginsWithOutput, t, "Error in testing EqualsTo")
	}

	beginsWithOutput, err = EqualTo("price", 101.77)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"price\":{\"$eq\":101.77}}", beginsWithOutput, t, "Error in testing EqualsTo")
	}

	beginsWithOutput, err = EqualTo("price", true)
	assertError(err, t)
	beginsWithOutput, err = EqualTo("price", testFailStruct)
	assertError(err, t)

}

func TestExcludes(t *testing.T) {
	beginsWithOutput, err := Excludes("region", "CA")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"region\":{\"$excludes\":\"CA\"}}", beginsWithOutput, t, "Error in testing Excludes")
	}

	beginsWithOutput, err = Excludes("category_id", 232)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"category_id\":{\"$excludes\":232}}", beginsWithOutput, t, "Error in testing Excludes")
	}

	beginsWithOutput, err = Excludes("price", 101.77)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"price\":{\"$excludes\":101.77}}", beginsWithOutput, t, "Error in testing Excludes")
	}

	beginsWithOutput, err = Excludes("price", true)
	assertError(err, t)
	beginsWithOutput, err = Excludes("price", testFailStruct)
	assertError(err, t)

}

func TestExcludesAny(t *testing.T) {
	beginsWithOutput, err := ExcludesAny("region", "CA", "VA", "MA")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"region\":{\"$excludes_any\":[\"CA\",\"VA\",\"MA\"]}}", beginsWithOutput, t, "Error in testing ExcludesAny")
	}

	beginsWithOutput, err = ExcludesAny("category_id", 232, 342, 232, 20284)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"category_id\":{\"$excludes_any\":[232,342,232,20284]}}", beginsWithOutput, t, "Error in testing ExcludesAny")
	}

	beginsWithOutput, err = ExcludesAny("price", 101.77, 77.67, 99.009, 67.644)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"price\":{\"$excludes_any\":[101.77,77.67,99.01,67.64]}}", beginsWithOutput, t, "Error in testing ExcludesAny")
	}

	beginsWithOutput, err = ExcludesAny("region", true)
	assertError(err, t)
	beginsWithOutput, err = ExcludesAny("region", testFailStruct)
	assertError(err, t)

}

func TestGreaterThan(t *testing.T) {
	beginsWithOutput, err := GreaterThan("rating", 7.5)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$gt\":7.5}}", beginsWithOutput, t, "Error in testing GreaterThan")
	}

	beginsWithOutput, err = GreaterThan("rating", 9)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$gt\":9}}", beginsWithOutput, t, "Error in testing GreaterThan")
	}

	beginsWithOutput, err = GreaterThan("rating", 6.64)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$gt\":6.64}}", beginsWithOutput, t, "Error in testing GreaterThan")
	}

	beginsWithOutput, err = GreaterThan("region", "6.64")
	assertError(err, t)
	beginsWithOutput, err = GreaterThan("region", testFailStruct)
	assertError(err, t)

}

func TestGreaterThanEqual(t *testing.T) {
	beginsWithOutput, err := GreaterThanEqual("rating", 7.5)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$gte\":7.5}}", beginsWithOutput, t, "Error in testing GreaterThanEqual")
	}

	beginsWithOutput, err = GreaterThanEqual("rating", 9)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$gte\":9}}", beginsWithOutput, t, "Error in testing GreaterThanEqual")
	}

	beginsWithOutput, err = GreaterThanEqual("rating", 6.64)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$gte\":6.64}}", beginsWithOutput, t, "Error in testing GreaterThanEqual")
	}

	beginsWithOutput, err = GreaterThanEqual("region", "7.5")
	assertError(err, t)
	beginsWithOutput, err = GreaterThanEqual("region", testFailStruct)
	assertError(err, t)

}

func TestEqualsAnyOf(t *testing.T) {
	beginsWithOutput, err := EqualsAnyOf("region", "MA", "VT", "NH", "RI", "CT")

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"region\":{\"$in\":[\"MA\",\"VT\",\"NH\",\"RI\",\"CT\"]}}", beginsWithOutput, t, "Error in testing EqualsAnyOf")
	}

	beginsWithOutput, err = EqualsAnyOf("rating", 9.21, 4.25, 6.50, 7.29, 9.99)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$in\":[9.21,4.25,6.5,7.29,9.99]}}", beginsWithOutput, t, "Error in testing EqualsAnyOf")
	}

	beginsWithOutput, err = EqualsAnyOf("rating", 4, 3, 100, 34, 35, 23, 66, 100)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$in\":[4,3,100,34,35,23,66,100]}}", beginsWithOutput, t, "Error in testing EqualsAnyOf")
	}

	beginsWithOutput, err = EqualsAnyOf("region", "7.5", 7.8)
	assertError(err, t)
	beginsWithOutput, err = EqualsAnyOf("region", 5, 7.8)
	assertError(err, t)
	beginsWithOutput, err = EqualsAnyOf("region", 5, "7.8")
	assertError(err, t)
	beginsWithOutput, err = EqualsAnyOf("region", 5, 6, false)
	assertError(err, t)
	beginsWithOutput, err = EqualsAnyOf("region", testFailStruct)
	assertError(err, t)
	beginsWithOutput, err = EqualsAnyOf("region", 1, 4, testFailStruct)
	assertError(err, t)

}

func TestIncludes(t *testing.T) {
	beginsWithOutput, err := Includes("region", "MA")

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"region\":{\"$includes\":\"MA\"}}", beginsWithOutput, t, "Error in testing Includes")
	}

	beginsWithOutput, err = Includes("rating", 9.21)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$includes\":9.21}}", beginsWithOutput, t, "Error in testing Includes")
	}

	beginsWithOutput, err = Includes("rating", 7)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$includes\":7}}", beginsWithOutput, t, "Error in testing Includes")
	}

	beginsWithOutput, err = Includes("rating", testFailStruct)
	assertError(err, t)
	beginsWithOutput, err = Includes("rating", true)
	assertError(err, t)

}

func TestIncludesAny(t *testing.T) {
	beginsWithOutput, err := IncludesAny("region", "MA", "VT", "NH", "RI", "CT")

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"region\":{\"$includes_any\":[\"MA\",\"VT\",\"NH\",\"RI\",\"CT\"]}}", beginsWithOutput, t, "Error in testing IncludesAny")
	}

	beginsWithOutput, err = IncludesAny("rating", 9.21, 4.25, 6.50, 7.29, 9.99)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$includes_any\":[9.21,4.25,6.5,7.29,9.99]}}", beginsWithOutput, t, "Error in testing IncludesAny")
	}

	beginsWithOutput, err = IncludesAny("category_id", 10, 100)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"category_id\":{\"$includes_any\":[10,100]}}", beginsWithOutput, t, "Error in testing IncludesAny")
	}

	beginsWithOutput, err = IncludesAny("region", "7.5", 7.8)
	assertError(err, t)
	beginsWithOutput, err = IncludesAny("region", 5, 7.8)
	assertError(err, t)
	beginsWithOutput, err = IncludesAny("region", 5, "7.8")
	assertError(err, t)
	beginsWithOutput, err = IncludesAny("region", 5, 6, false)
	assertError(err, t)
	beginsWithOutput, err = IncludesAny("region", testFailStruct)
	assertError(err, t)
	beginsWithOutput, err = IncludesAny("region", 1, 4, testFailStruct)
	assertError(err, t)

}

func TestLessThan(t *testing.T) {
	beginsWithOutput, err := LessThan("rating", 7.5)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$lt\":7.5}}", beginsWithOutput, t, "Error in testing LessThan")
	}

	beginsWithOutput, err = LessThan("rating", 9)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$lt\":9}}", beginsWithOutput, t, "Error in testing LessThan")
	}

	beginsWithOutput, err = LessThan("rating", 6.64)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$lt\":6.64}}", beginsWithOutput, t, "Error in testing LessThan")
	}

	beginsWithOutput, err = LessThan("region", "6.64")
	assertError(err, t)
	beginsWithOutput, err = LessThan("region", testFailStruct)
	assertError(err, t)

}

func TestLessThanEqual(t *testing.T) {
	beginsWithOutput, err := LessThanEqual("rating", 7.5)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$lte\":7.5}}", beginsWithOutput, t, "Error in testing LessThanEqual")
	}

	beginsWithOutput, err = LessThanEqual("rating", 9)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$lte\":9}}", beginsWithOutput, t, "Error in testing LessThanEqual")
	}

	beginsWithOutput, err = LessThanEqual("rating", 6.64)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$lte\":6.64}}", beginsWithOutput, t, "Error in testing LessThanEqual")
	}

	beginsWithOutput, err = LessThanEqual("region", "6.64")
	assertError(err, t)
	beginsWithOutput, err = LessThanEqual("region", testFailStruct)
	assertError(err, t)

}

func TestNotBeginsWith(t *testing.T) {
	beginsWithOutput, err := NotBeginWith("city", "low")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"city\":{\"$nbw\":\"low\"}}", beginsWithOutput, t, "Error in testing NotBeginsWith")
	}

}
func TestNotBeginsWithAny(t *testing.T) {
	beginsWithOutput, err := NotBeginWithAny("state", "mass", "cali", "flor")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"state\":{\"$nbwin\":[\"mass\",\"cali\",\"flor\"]}}", beginsWithOutput, t, "Error in testing NotBeginsWithAny")
	}

}

func TestNotEqualsTo(t *testing.T) {
	beginsWithOutput, err := NotEqualTo("region", "CA")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"region\":{\"$neq\":\"CA\"}}", beginsWithOutput, t, "Error in testing NotEqualsTo")
	}

	beginsWithOutput, err = NotEqualTo("category_id", 232)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"category_id\":{\"$neq\":232}}", beginsWithOutput, t, "Error in testing NotEqualsTo")
	}

	beginsWithOutput, err = NotEqualTo("price", 101.77)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"price\":{\"$neq\":101.77}}", beginsWithOutput, t, "Error in testing NotEqualsTo")
	}

	beginsWithOutput, err = NotEqualTo("price", true)
	assertError(err, t)
	beginsWithOutput, err = NotEqualTo("price", testFailStruct)
	assertError(err, t)

}

func TestNotEqualAnyOf(t *testing.T) {
	beginsWithOutput, err := NotEqualAnyOf("locality", "Los Angeles", "Santa Monica")

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"locality\":{\"$nin\":[\"Los Angeles\",\"Santa Monica\"]}}", beginsWithOutput, t, "Error in testing NotEqualAnyOf")
	}

	beginsWithOutput, err = NotEqualAnyOf("rating", 9.21, 4.25, 6.50, 7.29, 9.99)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"rating\":{\"$nin\":[9.21,4.25,6.5,7.29,9.99]}}", beginsWithOutput, t, "Error in testing NotEqualAnyOf")
	}

	beginsWithOutput, err = NotEqualAnyOf("category_id", 10, 100)

	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"category_id\":{\"$nin\":[10,100]}}", beginsWithOutput, t, "Error in testing NotEqualAnyOf")
	}

	beginsWithOutput, err = NotEqualAnyOf("region", "7.5", 7.8)
	assertError(err, t)
	beginsWithOutput, err = NotEqualAnyOf("region", 5, 7.8)
	assertError(err, t)
	beginsWithOutput, err = NotEqualAnyOf("region", 5, "7.8")
	assertError(err, t)
	beginsWithOutput, err = NotEqualAnyOf("region", 5, 6, false)
	assertError(err, t)
	beginsWithOutput, err = NotEqualAnyOf("region", testFailStruct)
	assertError(err, t)
	beginsWithOutput, err = NotEqualAnyOf("region", 1, 4, testFailStruct)
	assertError(err, t)

}

func TestSearch(t *testing.T) {
	beginsWithOutput, err := Search("name", "Charles")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("filters={\"name\":{\"$search\":\"Charles\"}}", beginsWithOutput, t, "Error in testing Search")
	}

}

func assertNoError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error was thrown with message: " + err.Error())
	}
}

func assertError(err error, t *testing.T) {
	if err == nil {
		t.Errorf("Error was not triggered")
	}
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
