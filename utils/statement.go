package utils

import (
	"fmt"
	"strings"
)

// type StatementType struct{}

// func NewStatement() *StatementType {
// 	return &StatementType{}
// }

func (s *Roman) Statement(statement string) {
	if strings.Contains(statement, "how many Credits") && (strings.Contains(statement, "Gold") || strings.Contains(statement, "Silver") || strings.Contains(statement, "Iron")) {
		if strings.Contains(statement, "Gold") {
			fmt.Printf("%s is %v Credits \n", statement, calculateCredit(statement, "Gold"))
		} else if strings.Contains(statement, "Silver") {
			fmt.Printf("%s is %v Credits \n", statement, calculateCredit(statement, "Silver"))
		} else if strings.Contains(statement, "Iron") {
			fmt.Printf("%s is %v Credits \n", statement, calculateCredit(statement, "Iron"))
		}
	} else if strings.Contains(statement, "how much") {
		fmt.Printf("%sis %v \n", statement[12:len(statement)-2], calculateCredit(statement, "None"))
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
