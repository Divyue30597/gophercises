package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// What is a flag? - Command-line flags are a common way to specify options
	// for command-line programs. For example, in wc -l the -l is a command-line
	// flag

	// Creating a flag fo the user to know the type we use and the name of the
	// flag. Here flagName is "csv" and the default value is "problems.csv" while
	// the third argument provides us the message or helper text message to know
	// what has to be done.
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	// csvFileName will be a pointer to the string
	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("There was an error opening your file: %s\n", *csvFileName))
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	correct := 0

	// Displaying the problems to the user with the parseLines function we create
	// an array of problem and we loop over that here.
	for i, prob := range problems {
		fmt.Printf("Problem no #%d: %s = \n", i+1, prob.question)
		var ans string
		// We are doing this because it needs a pointer value to work with so that
		// whenever it sets the value we can access it with our variable?????
		fmt.Scanf("%s\n", &ans)
		if ans == prob.answer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	// we are creating an array of problem in structure [{Q1, A1} {Q2, A2} ...]
	// -> {Q, A} -> problem
	anArrayOfProblem := make([]problem, len(lines))

	for i, line := range lines {
		// storing the problem
		anArrayOfProblem[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return anArrayOfProblem
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
