package kwengo

import (
	"errors"
	"strings"
)

// Kwengo is a default struct
type Kwengo struct {
	Numeral string `json:"kwengo"`
}

var tabConv = map[string]int{
	"kil":    1,
	"jin":    5,
	"pol":    10,
	"kilow":  50,
	"jij":    100,
	"jinjin": 500,
	"polsx":  1000,
}

// CheckKwego check if the numeral is correct
func (kwen *Kwengo) CheckKwego() ([]string, bool, error) {
	numList := strings.Fields(kwen.Numeral)
	for _, val := range numList {
		if _, ok := tabConv[val]; !ok {
			return nil, false, errors.New("String has a invalid number")
		}
	}
	return numList, true, nil
}

// ConvertDecimal converts the list of kwengonian strings in int
func ConvertDecimal(algRaw []string) (int, bool, error) {
	finalDecimal := 0
	algLen := len(algRaw)
	for i := 0; i < algLen; i++ {
		curr := algRaw[i]
		currValue := tabConv[curr]
		if i < algLen-1 {
			nextNum := algRaw[i+1]
			nextVal := tabConv[nextNum]
			if currValue < nextVal {
				finalDecimal += nextVal - currValue
				i++
			} else {
				finalDecimal += currValue
			}
		} else {
			finalDecimal += currValue
		}
	}
	return finalDecimal, true, nil
}
