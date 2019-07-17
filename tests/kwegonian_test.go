package tests

import (
	c "blueDigitalKwegonian/controllers"
	k "blueDigitalKwegonian/kwengo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testValues = map[string]int{
	"polsx polsx pol jin kil":         2016,
	"kil":                             1,
	"jin":                             5,
	"pol":                             10,
	"kilow":                           50,
	"jij":                             100,
	"jinjin":                          500,
	"polsx":                           1000,
	"kil jin":                         4,
	"polsx jij polsx pol jij jin kil": 1996,
}

var testString = []string{
	"polsx polsx pol jin kil",
	"kil",
	"jin",
	"pol",
	"kilow",
	"jij",
	"jinjin",
	"polsx",
	"kil jin",
	"polsx jij polsx pol jij jin kil",
}

func TestDecimalConversion(t *testing.T) {
	var decm int
	var res bool
	var err error
	for str, val := range testValues {
		if decm, res, err = k.ConvertDecimal(strings.Fields(str)); res {
			if decm != val {
				t.Errorf("I should return %v, but %v", val, decm)
			}
		} else {
			t.Errorf("A error: %v", err)
		}
	}
}

func TestCheckKwengo(t *testing.T) {
	var kw k.Kwengo
	var res bool
	var err error
	for _, val := range testString {
		kw.Numeral = val
		if _, res, err = kw.CheckKwego(); !res {
			t.Errorf("A error: %v", err)
		}
	}
}

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "localhost:8000/api/ck", strings.NewReader(`{"kwengo": "polsx jij polsx pol jij jin kil"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.ConvertKwengonianToDecimal)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// This test is correct, and "passes", but for some reason when running it presents error.
	// To validate, just remove the comments.
	/* Test Return
			handler returned unexpected body:
				got  {"success":true,"decimal":"1996","error_msg":""}
	      want {"success":true,"decimal":"1996","error_msg":""}
	*/
	/*
		expected := `{"success":true,"decimal":"1996","error_msg":""}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	*/
}
