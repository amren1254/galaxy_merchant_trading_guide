package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

type Roman struct{}

func NewRoman() *Roman {
	return &Roman{}
}

func romantoint(s string) int {
	var numberAsInt int
	for i := 0; i < len(s)-1; i++ {
		currentNumeral := mapRomantoInteger[string(s[i])]
		nextNumeral := mapRomantoInteger[string(s[i+1])]
		if currentNumeral >= nextNumeral {
			numberAsInt += currentNumeral
		} else {
			numberAsInt -= currentNumeral
		}
		if i == len(s)-2 {
			numberAsInt += nextNumeral
			return numberAsInt
		}
	}
	return numberAsInt
}

func main() {
	data := openAndReadFile("input.txt")
	fmt.Printf("%v \n", data)
	fmt.Println(romantoint("XVIII"))
}

func openAndReadFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %s", fileName)
	}
	defer file.Close()
	lines, err := readFile(file)
	if err != nil {
		fmt.Printf("Failed to read file: %s", fileName)
	}
	return lines
}

func readFile(reader io.Reader) ([][]string, error) {
	var lines [][]string
	scanner := bufio.NewScanner(reader)
	//scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		var line []string
		line = append(line, scanner.Text())
		lines = append(lines, line)
	}
	return lines, nil
}
