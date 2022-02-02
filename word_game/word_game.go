
// https://www.youtube.com/watch?v=LHhsNa_Kgns
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Awesome Quiz Game!")

	// var name string = "Tim" // string is a sequence of characters, wrapped in double quotes
	// fmt.Println(name)

	// nameImplicit := "Tim Implicit" // implicitly determine variable type
	// age := 24
	// fmt.Printf("Hello %v, you are %v years old.", nameImplicit,age)

	// var userInput string

	// fmt.Scan(&userInput) // & references memory RAM address of declared variable
	// fmt.Println(userInput)

	var userName string
	fmt.Printf("Please enter your name:")
	fmt.Scan(&userName)
	fmt.Printf("Hello, %v\n", userName)

	var userAge uint
	fmt.Printf("Enter your age:")
	fmt.Scan(&userAge)

	if userAge >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("Oh no, you can't play :(")
		return // breaks out of main function
	}

	userScore := 0
	numQuestions := 2

	fmt.Println("Which is better, $DOT or $BTC?")
	var answer string
	fmt.Scan(&answer)

	if strings.ToLower(answer) == "dot" {
		fmt.Println("Good job!")
		userScore++
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Println("What is the closest interest rate staking $UST, 5 or 20 percent?")
	var staking uint
	fmt.Scan(&staking)
	if staking == 20 {
		fmt.Println("Good job!")
		userScore++
	} else {
		fmt.Println("Incorrect")
	}

	percent := (float64(userScore) / float64(numQuestions)) *100
	fmt.Printf("You score %v out of %v, or %v%%\n", userScore, numQuestions, percent)
}