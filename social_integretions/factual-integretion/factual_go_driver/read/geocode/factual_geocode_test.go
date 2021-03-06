package geocode

import (
	"testing"
)

func TestGeoPoint(t *testing.T) {
	k := NewGeoPoint(34322, 234)
	js, _ := k.ToJsonFromGeo()
	assertEqual(js, "geo={\"$point\":[34322,234]}", t)

	k = NewGeoPoint(1234.342223222, 12344.342112383833)
	js, _ = k.ToJsonFromGeo()
	assertEqual(js, "geo={\"$point\":[1234.342223222,12344.342112383833]}", t)

	k = NewGeoPoint(1234.342223222, 12344.0000001)
	js, _ = k.ToJsonFromGeo()

	assertEqual(js, "geo={\"$point\":[1234.342223222,12344.0000001]}", t)

	k = NewGeoPoint(1234.342223222, 12344.00000000000000000000)
	js, _ = k.ToJsonFromGeo()

	assertEqual(js, "geo={\"$point\":[1234.342223222,12344]}", t)

	k = NewGeoPoint(1234.342223222, 12344.000000001)
	js, _ = k.ToJsonFromGeo()
	assertEqual(js, "geo={\"$point\":[1234.342223222,12344.000000001]}", t)

}

func TestGeoCircle(t *testing.T) {
	k := NewGeoCircle(34322, 234, 34)
	js, _ := k.ToJsonFromGeo()
	assertEqual(js, "geo={\"$circle\":{\"$center\":[34322,234],\"$meters\":34}}", t)

	k = NewGeoCircle(34322.7765765, 234.544, 500)
	js, _ = k.ToJsonFromGeo()
	assertEqual(js, "geo={\"$circle\":{\"$center\":[34322.7765765,234.544],\"$meters\":500}}", t)

	k = NewGeoCircle(34322.000000005, 234.544, 500)
	js, _ = k.ToJsonFromGeo()
	assertEqual(js, "geo={\"$circle\":{\"$center\":[34322.000000005,234.544],\"$meters\":500}}", t)

	k = NewGeoCircle(34322.000000005, 234.00000000000, 500)
	js, _ = k.ToJsonFromGeo()
	assertEqual(js, "geo={\"$circle\":{\"$center\":[34322.000000005,234],\"$meters\":500}}", t)

	k = NewGeoCircle(34322.000000005, 234.000000000001, 500)
	js, _ = k.ToJsonFromGeo()
	assertEqual(js, "geo={\"$circle\":{\"$center\":[34322.000000005,234.000000000001],\"$meters\":500}}", t)

}

func TestGeoRectangle(t *testing.T) {
	k := NewGeoRectangle(234, 1123, 1212, 199)
	js, _ := k.ToJsonFromGeo()
	assertEqual(js, "geo={\"$rect\":[[234,1123],[1212,199]]}", t)

	k = NewGeoRectangle(234.34, 1123.2342332423, 1212.23487987, 199.188823482234)
	js, _ = k.ToJsonFromGeo()
	assertEqual(js, "geo={\"$rect\":[[234.34,1123.2342332423],[1212.23487987,199.188823482234]]}", t)

	k = NewGeoRectangle(234.0000000000, 1123.2342332423, 1212.00000000001, 199.188823482234)
	js, _ = k.ToJsonFromGeo()
	assertEqual(js, "geo={\"$rect\":[[234,1123.2342332423],[1212.00000000001,199.188823482234]]}", t)

}

func assertEqual(expected, got interface{}, t *testing.T) {
	if expected != got {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
