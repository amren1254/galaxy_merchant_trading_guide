package main

import (
	"bufio"
	"log"
	"os"

	"github.com/amren1254/gmtg/utils"
)

var InvalidQuestionError = "I have no idea what you are talking about"

func main() {
	log.Printf("Press ctrl-c to exit")
	for {
		reader := bufio.NewReader(os.Stdin)
		inputText, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Unable to read input statement %v", err)
			return
		}

		// statementType := utils.NewStatement()
		utils.NewRoman().Statement(inputText)
	}
}
