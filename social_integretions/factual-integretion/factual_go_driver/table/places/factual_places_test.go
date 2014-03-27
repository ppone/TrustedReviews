package places

import (
	"reflect"
	"testing"
)

func TestAllAtributesAreCorrect(t *testing.T) {
	dummyStruct := new(placesData)
	val := reflect.ValueOf(dummyStruct).Elem()

	assertEqual(val.Type().Field(0).Name, "factual_id", t, "Field name does not match ")
	assertEqual(val.Type().Field(0).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(1).Name, "name", t, "Field name does not match ")
	assertEqual(val.Type().Field(1).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(2).Name, "address", t, "Field name does not match ")
	assertEqual(val.Type().Field(2).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(3).Name, "address_extended", t, "Field name does not match ")
	assertEqual(val.Type().Field(3).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(4).Name, "locality", t, "Field name does not match ")
	assertEqual(val.Type().Field(4).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(5).Name, "region", t, "Field name does not match ")
	assertEqual(val.Type().Field(5).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(6).Name, "postcode", t, "Field name does not match ")
	assertEqual(val.Type().Field(6).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(7).Name, "country", t, "Field name does not match ")
	assertEqual(val.Type().Field(7).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(8).Name, "neighborhood", t, "Field name does not match ")
	assertEqual(val.Type().Field(8).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(9).Name, "tel", t, "Field name does not match ")
	assertEqual(val.Type().Field(9).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(10).Name, "fax", t, "Field name does not match ")
	assertEqual(val.Type().Field(10).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(11).Name, "website", t, "Field name does not match ")
	assertEqual(val.Type().Field(11).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(12).Name, "latitude", t, "Field name does not match ")
	assertEqual(val.Type().Field(12).Type.String(), "float64", t, "Field type does not match ")

	assertEqual(val.Type().Field(13).Name, "longitude", t, "Field name does not match ")
	assertEqual(val.Type().Field(13).Type.String(), "float64", t, "Field type does not match ")

	assertEqual(val.Type().Field(14).Name, "status", t, "Field name does not match ")
	assertEqual(val.Type().Field(14).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(15).Name, "hours_display", t, "Field name does not match ")
	assertEqual(val.Type().Field(15).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(16).Name, "chain_name", t, "Field name does not match ")
	assertEqual(val.Type().Field(16).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(17).Name, "email", t, "Field name does not match ")
	assertEqual(val.Type().Field(17).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(18).Name, "spost_town", t, "Field name does not match ")
	assertEqual(val.Type().Field(18).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(19).Name, "category_ids", t, "Field name does not match ")
	assertEqual(val.Type().Field(19).Type.String(), "int", t, "Field type does not match ")

	assertEqual(val.Type().Field(20).Name, "admin_region", t, "Field name does not match ")
	assertEqual(val.Type().Field(20).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(21).Name, "chain_id", t, "Field name does not match ")
	assertEqual(val.Type().Field(21).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(22).Name, "hours", t, "Field name does not match ")
	assertEqual(val.Type().Field(22).Type.String(), "string", t, "Field type does not match ")

	assertEqual(val.Type().Field(23).Name, "po_box", t, "Field name does not match ")
	assertEqual(val.Type().Field(23).Type.String(), "string", t, "Field type does not match ")

}

func assertEqual(expected, got interface{}, t *testing.T, errorMessage string) {
	if expected != got {
		t.Errorf("%s => expected %s, got %s", errorMessage, expected, got)
	}
}
