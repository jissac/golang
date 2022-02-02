package main

import (
	"os"
	"fmt"
	"flag"
	"encoding/csv"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
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

	correct := 0
	for i,p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		fmt.Printf("%s\n",answer)
		if answer == p.a {
			// fmt.Println("Correct!")
			correct++
		} else {
			// fmt.Println("Incorrect!")
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
	percent := (float64(correct) / float64(len(problems))) *100
	fmt.Printf("Percentage: %.2f%%\n",percent)
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