package utils

import (
	"errors"

	"github.com/amren1254/gmtg/model"
)

func RomanToInt(s string) (int, error) {
	var numberAsInt, romanSequence int = 0, 1
	if len(s) == 1 {
		return model.MapRomantoInteger[s], nil
	}
	for i := 0; i < len(s)-1; i++ {

		currentNumeral := model.MapRomantoInteger[string(s[i])]
		nextNumeral := model.MapRomantoInteger[string(s[i+1])]

		if currentNumeral >= nextNumeral {
			if (currentNumeral == 5 && nextNumeral == 5) ||
				(currentNumeral == 50 && nextNumeral == 50) ||
				(currentNumeral == 500 && nextNumeral == 500) {

				return 0, errors.New("Input is incorrect")
			}
			if currentNumeral == nextNumeral {
				romanSequence++
			}
			if romanSequence > 3 {
				return 0, errors.New("Requested number is in invalid format")
			}
			numberAsInt += currentNumeral
		} else {
			numberAsInt -= currentNumeral
		}
		if i == len(s)-2 {
			numberAsInt += nextNumeral
			return numberAsInt, nil
		}
	}
	return numberAsInt, nil
}
