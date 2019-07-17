package kwengo

import (
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
		if decm, res, err = convertDecimal(strings.Fields(str)); res {
			if decm != val {
				t.Errorf("I should return %v, but %v", val, decm)
			}
		} else {
			t.Errorf("A error: %v", err)
		}
	}
}

func TestCheckKwengo(t *testing.T) {
	var res bool
	var err error
	for _, val := range testString {
		if _, res, err = checkKwego(val); !res {
			t.Errorf("A error: %v", err)
		}
	}
}
