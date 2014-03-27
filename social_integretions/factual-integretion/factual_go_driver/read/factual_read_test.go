package main

import (
	"testing"
)

func TestRead(t *testing.T) {

}

func assertEqual(expected, got interface{}, t *testing.T, errorMessage string) {
	if expected != got {
		t.Errorf("%s => expected %s, got %s", errorMessage, expected, got)
	}
}
