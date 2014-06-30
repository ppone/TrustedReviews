package filters

import (
	"testing"
)

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
}

func TestCheckTypesInArray(t *testing.T) {
	a :=  []interface{4.5,45.2}
	jsonType, err := checkTypesInArray(a, numericT)
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
	jsonType, err = checkTypes("23.4", numericT)
	assertError(err, t)
	jsonType, err = checkTypes(23.4, stringT)
	assertError(err, t)

	jsonType, err = checkTypes(true, stringAndNumericT)
	assertError(err, t)
}

func TestFloatToString(t *testing.T) {

	assertEqual("3424.50", floatToString(3424.50), t, "Error in converting float to string")
	assertEqual("24.50", floatToString(24.50), t, "Error in converting float to string")
	assertEqual("99424.50", floatToString(99424.50), t, "Error in converting float to string")
	assertEqual("73231.51", floatToString(73231.5123), t, "Error in converting float to string")
	assertEqual("100000.00", floatToString(99999.996), t, "Error in converting float to string")
	assertEqual("101.78", floatToString(101.775), t, "Error in converting float to string")
	assertNotEqual("100000.00", floatToString(99999.99), t, "Error in converting float to string, values are not suppose to match")
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
	assertEqual("{\"city\":{\"$blank\":true}}", Blank("city", true), t, "Error in testing Blank")
	assertEqual("{\"city\":{\"$blank\":false}}", Blank("city", false), t, "Error in testing Blank")

}

func TestBeginsWith(t *testing.T) {
	beginsWithOutput, err := BeginsWith("city", "low")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("{\"city\":{\"$bw\":\"low\"}}", beginsWithOutput, t, "Error in testing BeginsWith")
	}

}
func TestBeginsWithAny(t *testing.T) {
	beginsWithOutput, err := BeginsWithAny("state", "mass", "cali", "flor")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("{\"state\":{\"$bwin\":[\"mass\",\"cali\",\"flor\"]}}", beginsWithOutput, t, "Error in testing BeginsWith")
	}

}

func TestEqualsTo(t *testing.T) {
	beginsWithOutput, err := EqualTo("region", "CA")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("{\"region\":{\"$eq\":\"CA\"}}", beginsWithOutput, t, "Error in testing EqualsTo")
	}

	beginsWithOutput, err = EqualTo("category_id", 232)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("{\"category_id\":{\"$eq\":232}}", beginsWithOutput, t, "Error in testing EqualsTo")
	}

	beginsWithOutput, err = EqualTo("price", 101.77)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("{\"price\":{\"$eq\":101.77}}", beginsWithOutput, t, "Error in testing EqualsTo")
	}

}

func TestExcludes(t *testing.T) {
	beginsWithOutput, err := Excludes("region", "CA")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("{\"region\":{\"$excludes\":\"CA\"}}", beginsWithOutput, t, "Error in testing EqualsTo")
	}

	beginsWithOutput, err = Excludes("category_id", 232)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("{\"category_id\":{\"$excludes\":232}}", beginsWithOutput, t, "Error in testing EqualsTo")
	}

	beginsWithOutput, err = Excludes("price", 101.77)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("{\"price\":{\"$excludes\":101.77}}", beginsWithOutput, t, "Error in testing EqualsTo")
	}

}

func TestExcludesAny(t *testing.T) {
	beginsWithOutput, err := ExcludesAny("region", "CA", "VA", "MA")
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("{\"region\":{\"$excludes_any\":[\"CA\",\"VA\",\"MA\"]}}", beginsWithOutput, t, "Error in testing EqualsTo")
	}

	beginsWithOutput, err = ExcludesAny("category_id", 232, 342, 232, 20284)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("{\"category_id\":{\"$excludes_any\":[232,342,232,20284]}}", beginsWithOutput, t, "Error in testing EqualsTo")
	}

	beginsWithOutput, err = ExcludesAny("price", 101.77, 77.67, 99.009, 67.644)
	assertEqual(nil, err, t, "function error was thrown")
	if err == nil {
		assertEqual("{\"price\":{\"$excludes_any\":[101.77,77.67,99.01,67.64]}}", beginsWithOutput, t, "Error in testing EqualsTo")
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
