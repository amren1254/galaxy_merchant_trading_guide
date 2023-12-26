package model

const (
	RomanI = "I"
	RomanV = "V"
	RomanX = "X"
	RomanL = "L"
	RomanC = "C"
	RomanD = "D"
	RomanM = "M"
)

var MapRomantoInteger = map[string]int{
	RomanI: 1,
	RomanV: 5,
	RomanX: 10,
	RomanL: 50,
	RomanC: 100,
	RomanD: 500,
	RomanM: 1000,
}

var MapCreditsInfo = map[string]float64{
	"Silver": 17,
	"Gold":   14450,
	"Iron":   195.5,
	"None":   1,
}

var MapIntergalacticToRoman = map[string]string{
	"glob": "I",
	"prok": "V",
	"pish": "X",
	"tegj": "L",
}
