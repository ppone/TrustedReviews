package places_categories

import (
	"reflect"
	"testing"
)

func TestAllAtributesAreCorrect(t *testing.T) {
	dummyStruct := new(placesCategoriesData)
	val := reflect.ValueOf(dummyStruct).Elem()

	assertEqual(val.Type().Field(0).Name, "category_id", t, "Field name does not match ")
	assertEqual(val.Type().Field(0).Type.String(), "int", t, "Field type does not match ")

	assertEqual(val.Type().Field(1).Name, "parents", t, "Field name does not match ")
	assertEqual(val.Type().Field(1).Type.String(), "int", t, "Field type does not match ")

	assertEqual(val.Type().Field(2).Name, "en", t, "Field name does not match ")
	assertEqual(val.Type().Field(2).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(3).Name, "de", t, "Field name does not match ")
	assertEqual(val.Type().Field(3).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(4).Name, "es", t, "Field name does not match ")
	assertEqual(val.Type().Field(4).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(5).Name, "fr", t, "Field name does not match ")
	assertEqual(val.Type().Field(5).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(6).Name, "it", t, "Field name does not match ")
	assertEqual(val.Type().Field(6).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(7).Name, "jp", t, "Field name does not match ")
	assertEqual(val.Type().Field(7).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(8).Name, "kr", t, "Field name does not match ")
	assertEqual(val.Type().Field(8).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(9).Name, "zh", t, "Field name does not match ")
	assertEqual(val.Type().Field(9).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(10).Name, "zh_hant", t, "Field name does not match ")
	assertEqual(val.Type().Field(10).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(11).Name, "pt", t, "Field name does not match ")
	assertEqual(val.Type().Field(11).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(12).Name, "abstract", t, "Field name does not match ")
	assertEqual(val.Type().Field(12).Type.String(), "bool", t, "Field type does not match ")

}

func assertEqual(expected, got interface{}, t *testing.T, errorMessage string) {
	if expected != got {
		t.Errorf("%s => expected %s, got %s", errorMessage, expected, got)
	}
}
