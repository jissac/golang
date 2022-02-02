package main

import (
	"os"
	"fmt"
	"flag"
	"encoding/csv"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	file, err := os.Open(*csvFilename) // Flag package returns pointer to the string, hence the *, & takes a variable/ addressable entity and returns a pointer, whereas * takes a pointer and returns the thing that it points to

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	// fmt.Println(lines)
	problems := parseLines(lines)
	// fmt.Printf("%v\n",problems)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)


	correct := 0
	for i,p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		// Channels: https://gobyexample.com/channels
		answerCh := make(chan string)
		// anonymous function, go routine
		// closure: similar to anon fn, except uses data that was defined outside of it
		go func() {	
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			gameOverMsg(correct,len(problems))
			return
		case answer := <- answerCh:
			if answer == p.a {
				// fmt.Println("Correct!")
				correct++
			}
		}
	}

	gameOverMsg(correct,len(problems))


}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines)) //The make function allocates a zeroed array and returns a slice that refers to that array
	
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func gameOverMsg(correct int, problemLength int) {
	fmt.Printf("\nYou scored %d out of %d.\n", correct, problemLength)
	percent := (float64(correct) / float64(problemLength)) *100
	fmt.Printf("Percentage: %.2f%%\n",percent)
}