package geocode

import (
	"fmt"
	"testing"
)

func TestGeoPoint(t *testing.T) {
	var x float64
	x = 12344.342112383833

	k := New_geopoint(34322, 234)
	js, _ := k.ToJson()
	assertEqual(js, "geo={\"$point\":[34322,234]}", t)
	k.Point = [2]float64{12, 12344}
	js, _ = k.ToJson()
	assertEqual(js, "geo={\"$point\":[12,12344]}", t)
	k.Point = [2]float64{1234.342223222, 12344.342112383833}
	fmt.Println(x)
	js, _ = k.ToJson()
	assertEqual(js, "geo={\"$point\":[1234.342223222,12344.342112383833]}", t)

}

func TestGeoCircle(t *testing.T) {

}

func assertEqual(expected, got interface{}, t *testing.T) {
	if expected != got {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
