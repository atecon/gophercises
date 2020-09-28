package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	timeLimit := flag.Int("limit", 30, "time limit for answering the quiz in sec.")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file.")
	}
	// fmt.Println(lines)  // 2-d slice array of strings
	problems := parseLines(lines)
	// fmt.Println(problems)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) // immediately starte timer
	**** HIER BEI MINUTE 8 2. Video weitermachen
	<-timer.C                                                       // wait for a message from channel

	nCorrect := 0
	for i, p := range problems {
		// Print every single problem
		fmt.Printf("Problem #%d : %s = \n", i+1, p.question)
		var answer string // declare empty string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			fmt.Println("Correct")
			nCorrect++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", nCorrect, len(problems))
}

func parseLines(lines [][]string) []problem {
	/* Input: 2-d string slice
	Return: single slice holding both question + answer in a struct like form. */

	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct {
	answer   string
	question string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
