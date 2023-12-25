package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	RomanI = "I"
	RomanV = "V"
	RomanX = "X"
	RomanL = "L"
	RomanC = "C"
	RomanD = "D"
	RomanM = "M"
)

var mapRomantoInteger = map[string]int{
	RomanI: 1,
	RomanV: 5,
	RomanX: 10,
	RomanL: 50,
	RomanC: 100,
	RomanD: 500,
	RomanM: 1000,
}

var mapIntergalacticToRoman = map[string]string{
	"glob": "I",
	"prok": "V",
	"pish": "X",
	"tegj": "L",
}

var mapCreditsInfo = map[string]float64{
	"Silver": 17,
	"Gold":   14450,
	"Iron":   195.5,
	"None":   1,
}
var InvalidQuestionError = "I have no idea what you are talking about"

type StatementType struct{}

func NewStatement() *StatementType {
	return &StatementType{}
}

//	func openAndReadFile(fileName string) [][]string {
//		file, err := os.Open(fileName)
//		if err != nil {
//			fmt.Printf("Failed to open file: %s", fileName)
//		}
//		defer file.Close()
//		lines, err := readFile(file)
//		if err != nil {
//			fmt.Printf("Failed to read file: %s", fileName)
//		}
//		return lines
//	}
//
//	func readFile(reader io.Reader) ([][]string, error) {
//		var lines [][]string
//		scanner := bufio.NewScanner(reader)
//		//scanner.Split(bufio.ScanWords)
//		for scanner.Scan() {
//			var line []string
//			line = append(line, scanner.Text())
//			lines = append(lines, line)
//		}
//		return lines, nil
//	}
//
//	func storeAndMapInputPrices(input [][]string) {
//		for i := 0; i < len(input); i++ {
//			if len(input[i][0]) == 0 {
//				break
//			} else {
//				words := strings.Split(input[i][0], " ")
//				word := words[0]
//				roman := words[2]
//				mapInputStringToInteger[word] = roman
//			}
//		}
//	}
type IRoman interface {
	GetRomanString(string)
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
		if val, exists := mapIntergalacticToRoman[word]; exists {
			r.romanString += val
		}
	}
}

func romantoint(s string) (int, error) {
	var numberAsInt, romanSequence int = 0, 1
	if len(s) == 1 {
		return mapRomantoInteger[s], nil
	}
	for i := 0; i < len(s)-1; i++ {

		currentNumeral := mapRomantoInteger[string(s[i])]
		nextNumeral := mapRomantoInteger[string(s[i+1])]

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

func calculateCredit(statement string, resourceType string) float64 {
	var iR IRoman
	roman := NewRoman()
	iR = roman
	iR.GetRomanString(statement)
	romanInt, err := romantoint(roman.romanString)
	if err != nil {
		log.Printf("Roman Conversion Error : %v", err)
	}
	return float64(romanInt) * mapCreditsInfo[resourceType]
}

func (s *StatementType) Statement(statement string) {
	if strings.Contains(statement, "how many Credits") && (strings.Contains(statement, "Gold") || strings.Contains(statement, "Silver") || strings.Contains(statement, "Iron")) {
		if strings.Contains(statement, "Gold") {
			fmt.Printf("%s is %v Credits \n", statement, calculateCredit(statement, "Gold"))
		} else if strings.Contains(statement, "Silver") {
			fmt.Printf("%s is %v Credits \n", statement, calculateCredit(statement, "Silver"))
		} else if strings.Contains(statement, "Iron") {
			fmt.Printf("%s is %v Credits \n", statement, calculateCredit(statement, "Iron"))
		}
	} else if strings.Contains(statement, "how much") {
		fmt.Printf("%s is %v \n", statement, calculateCredit(statement, "None"))
	} else if strings.Contains(statement, "has more") || strings.Contains(statement, "larger than") {
		var splitString []string
		if strings.Contains(statement, "has more") {
			splitString = strings.Split(statement, "has more")
		}
		if strings.Contains(statement, "larger than") {
			splitString = strings.Split(statement, "larger than")
		}
		var leftString, rightString string
		var leftCredit, rightCredit float64
		if len(splitString) >= 2 {
			leftString = splitString[0]
			rightString = splitString[1]
		}
		if strings.Contains(leftString, "Gold") {
			leftCredit = calculateCredit(leftString, "Gold")
		} else if strings.Contains(leftString, "Silver") {
			leftCredit = calculateCredit(leftString, "Silver")
		} else if strings.Contains(leftString, "Iron") {
			leftCredit = calculateCredit(leftString, "Iron")
		}

		if strings.Contains(rightString, "Gold") {
			rightCredit = calculateCredit(rightString, "Gold")
		} else if strings.Contains(rightString, "Silver") {
			rightCredit = calculateCredit(rightString, "Silver")
		} else if strings.Contains(rightString, "Iron") {
			rightCredit = calculateCredit(rightString, "Iron")
		}

		if leftCredit > rightCredit {
			fmt.Printf("%s has more %s", leftString, rightString)
		} else {
			fmt.Printf("%s has less %s", leftString, rightString)
		}
	} else if strings.Contains(statement, "has less") || strings.Contains(statement, "smaller than") {
		var splitString []string
		if strings.Contains(statement, "has less") {
			splitString = strings.Split(statement, "has less")
		}
		if strings.Contains(statement, "smaller than") {
			splitString = strings.Split(statement, "smaller  than")
		}
		var leftString, rightString string
		var leftCredit, rightCredit float64
		if len(splitString) >= 2 {
			leftString = splitString[0]
			rightString = splitString[1]
		}
		if strings.Contains(leftString, "Gold") {
			leftCredit = calculateCredit(leftString, "Gold")
		} else if strings.Contains(leftString, "Silver") {
			leftCredit = calculateCredit(leftString, "Silver")
		} else if strings.Contains(leftString, "Iron") {
			leftCredit = calculateCredit(leftString, "Iron")
		}

		if strings.Contains(rightString, "Gold") {
			rightCredit = calculateCredit(rightString, "Gold")
		} else if strings.Contains(rightString, "Silver") {
			rightCredit = calculateCredit(rightString, "Silver")
		} else if strings.Contains(rightString, "Iron") {
			rightCredit = calculateCredit(rightString, "Iron")
		}

		if leftCredit > rightCredit {
			fmt.Printf("%s has more %s", leftString, rightString)
		} else {
			fmt.Printf("%s has less %s", leftString, rightString)
		}
	} else {
		fmt.Printf("Invalid Input")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputText, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Unable to read input statement %v", err)
		return
	}

	statementType := NewStatement()
	statementType.Statement(inputText)

}
