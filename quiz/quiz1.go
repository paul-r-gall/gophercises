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
	"time"
	"flag"
)

func parseQuest(q string) string {
	re := regexp.MustCompile("([0-9, ]+)([+])([0-9, ]+)")
	return strings.Replace(re.FindString(q), " ","",-1)
}

func parseAns(a string) string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindString(a)
}

func awaitAns(ansCha chan string) {
	var input string
	fmt.Scanln(&input)
	ansCha <- input
	return
}

func main() {
	tChannel := make(chan bool, 1)
	ansChannel := make(chan string, 1)
	filePtr := flag.String("file","problems.csv","quiz file (csv format)")
	flag.Parse()
	
	csvFile, _ := os.Open(*filePtr)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	
	correct := 0
	total := 0
	timeout := false
	fmt.Println("Hit Enter to start the quiz")
	fmt.Scanln()
	// start the timer
	go func() {
		time.Sleep(10*time.Second)
		tChannel <- true
	}()
	// run the quiz
	for !timeout {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		total += 1
		fmt.Println(parseQuest(line[0]))
		go awaitAns(ansChannel)
		select {
		case ans := <-ansChannel:
			if ans == parseAns(line[1]) {
				fmt.Println("Correct!")
				correct+=1
			} else {
				fmt.Println("Incorrect.")
			}
		case timeout = <-tChannel:
		}	
	}
	// print out unanswered questions
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(line)
		total += 1
	}

	fmt.Printf("You got %d out of %d correct", correct, total)

}
