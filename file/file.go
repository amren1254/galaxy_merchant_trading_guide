package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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

// func storeAndMapInputPrices(input [][]string) {
// 	for i := 0; i < len(input); i++ {
// 		if len(input[i][0]) == 0 {
// 			break
// 		} else {
// 			words := strings.Split(input[i][0], " ")
// 			word := words[0]
// 			roman := words[2]
// 			utils.MapInputStringToInteger[word] = roman
// 		}
// 	}
// }
