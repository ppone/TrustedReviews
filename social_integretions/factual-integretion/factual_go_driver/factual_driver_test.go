package factual

import (
	"fmt"
	"strings"
	"testing"
)

func TestToken(t *testing.T) {
	fR := NewToken("Pbu7jRdBErgLW07g9c25JtGcwwt1KmpoxRTfFL3x", "vC4AgocPBhxe0GFkTsetoiuEAJEgqz6MCbAnXEoO")

	nR, err := NewRead(fR, "restaurants-us")

	if err != nil {
		fmt.Println(err)
		return
	}

	nR.AddQuery("pizza")
	nR.AddLimit(5)
	nR.AddSelect("factual_id")

	//fmt.Println("frData", nR.GetReadURL())

	s, err := nR.Get()

	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	if strings.Contains(s, "error") {
		t.Error("response returned an error")
	}

	//fmt.Println(s)

}

/*
func TestSchema(t *testing.T) {
	fR := NewToken("Pbu7jRdBErgLW07g9c25JtGcwwt1KmpoxRTfFL3x", "vC4AgocPBhxe0GFkTsetoiuEAJEgqz6MCbAnXEoO")
	s, err := fR.Get("http://api.v3.factual.com/t/restaurants-us/schema")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(s)

}*/
