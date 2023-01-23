package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func parse_csv_from_string(contents string) [][]string {
	csv_reader := csv.NewReader(strings.NewReader(contents))
	result, error := csv_reader.ReadAll()
	if error != nil {
		fmt.Fprintf(os.Stderr, "Error reading from csv: %v", error)
		os.Exit(3)
	}

	return result
}

func apply_questions(csv_lines [][]string, time_limit int) {

	var correct int
	total := len(csv_lines)
	reset := make(chan bool)

	for i, qa_pair := range csv_lines {
		reached_limit := time.After(time.Duration(time_limit) * time.Second)
		go func() {
			for {
				select {
				case <-reached_limit:
					fmt.Println("\nReached time limit... Stopping quiz game.")
					os.Exit(1)
				case <-reset:
					fmt.Println("\nResetting timer...")
					return
				}
			}
		}()

		question := qa_pair[0]
		answer := qa_pair[1]
		fmt.Printf("Problem #%v: %v = ", i, question)

		input := read_answer()
		if input == answer {
			correct += 1
		}
		reset <- true
	}

	fmt.Printf("You scored %v out of %v.\n", correct, total)
}

/*
Reads a line of text from the terminal
to be passed to the quiz game for checking
against the answer.
*/
func read_answer() string {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	cleaned_line := strings.Trim(string(line), " ")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading line: %v\n", err)
		os.Exit(1)
	}

	return cleaned_line
}

func main() {

	// parse command line input
	filename, limit := parse_flags()

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from file")
		os.Exit(1)
	}

	contents := string(file[:])

	// read csv and put into suitable data structure with question answer format
	csv_lines := parse_csv_from_string(contents)

	// apply operators and check that answer is correct or not keeping track of totals
	apply_questions(csv_lines, limit)
}

func parse_flags() (string, int) {
	filename := flag.String("csv", "problems.csv",
		"\t-csv string\n\t\ta csv file in format 'question,answer' (default 'problems.csv')\n")

	limit := flag.Int("limit", 30, "\t-limit int\n\t\tthe time limit for the quiz in seconds\n")

	flag.Parse()

	return *filename, *limit
}
