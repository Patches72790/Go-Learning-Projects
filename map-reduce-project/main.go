package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Map(contents string) {
	word_count := map[string]int{}

    tokens := strings.Split(contents, " ")

}

func Reduce() {}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Error reading file")
	}

	file_contents := string(bytes[:])

	fmt.Println("Here are the contents of the file\n:%s", file_contents)
}
