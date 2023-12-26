package utils

import (
	"log"

	"github.com/amren1254/gmtg/model"
)

func calculateCredit(statement string, resourceType string) float64 {
	var iR IUtils
	roman := NewRoman()
	iR = roman
	iR.GetRomanString(statement)
	romanInt, err := RomanToInt(roman.romanString)
	if err != nil {
		log.Printf("Roman Conversion Error : %v", err)
	}
	return float64(romanInt) * model.MapCreditsInfo[resourceType]
}
