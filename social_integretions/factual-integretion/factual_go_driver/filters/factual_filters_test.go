package filters

import (
	"reflect"
	"testing"
)

func TestAllAtributesAreCorrect(t *testing.T) {
	dummyStruct := new(placesData)

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
