package utils

import (
	"strings"

	"github.com/amren1254/gmtg/model"
)

type IUtils interface {
	GetRomanString(string)
	Statement(string)
}

type Roman struct {
	romanString string
}

func NewRoman() *Roman {
	return &Roman{}
}

func (r *Roman) GetRomanString(line string) {
	words := strings.Fields(line)

	for _, word := range words {
		if val, exists := model.MapIntergalacticToRoman[word]; exists {
			r.romanString += val
		}
	}
}
