package filters

import (
	"fmt"
	"testing"
)

func TestAllAtributesAreCorrect(t *testing.T) {

	r := returnArrayString([]string{"ab", "bc"})
	fmt.Println(BeginsWithAny("pw", "gf", "rt"))
	fmt.Println(EqualTo("jiffy", 232332432.2342335))
	assertEqual("[\"ab\",\"bc\"]", r, t, "Did not make string array correctly")

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
