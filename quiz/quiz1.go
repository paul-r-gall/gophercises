package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"bufio"
	"regexp"
	"strings"
)

func parseQuest(q string) string {
	re := regexp.MustCompile("([0-9, ]+)([+])([0-9, ]+)")
	return strings.Replace(re.FindString(q), " ","",-1)
}

func parseAns(a string) string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindString(a)
}

func main() {
	csvFile, _ := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	correct := 0
	total := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		total += 1
		fmt.Println(parseQuest(line[0]))
		var input string
		fmt.Scanln(&input)
		if input == parseAns(line[1]) {
			fmt.Println("Correct!")
			correct+=1
		} else {
			fmt.Println("Incorrect.")
		}
	}

	fmt.Printf("You got %d out of %d correct", correct, total)

}
