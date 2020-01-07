//
// A solution to https://github.com/gophercises/quiz
//

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	csvFP, err := os.Open(*csvFilename)
	if err != nil {
		panic(err)
	}
	defer csvFP.Close()

	problems, err := csv.NewReader(csvFP).ReadAll()
	if err != nil {
		panic(err)
	}
	score := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	go func() {
		<-timer.C
		fmt.Printf("\nYou scored %d out of %d", score, len(problems))
		os.Exit(0)
	}()
	for index, problem := range problems {
		ans := ""
		fmt.Printf("Problem #%d: %s = ", index, problem[0])
		fmt.Scanln(&ans)
		if problem[1] == ans {
			score++
		}
	}
}
