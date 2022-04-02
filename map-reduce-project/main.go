package main

import (
	"fmt"
	"github.com/patches72790/mapreduce/worker"
	//"log"
	"os"
)

//
// Reads from a file given by the dirname and filename
// and returns the files contents ready to be mapped
//
func read_filecontents(dirname string) []string { return nil }

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Error not enough args given.")
		fmt.Println("Usage: ./main <files...>")
		os.Exit(1)
	}

	dirname := os.Args[1]
	//bytes, err := os.ReadFile(os.Args[1])
	//if err != nil {
	//	log.Fatal("Error reading file")
	//}

	d_entries, err := os.ReadDir(dirname)
	if err != nil {
		os.Exit(1)
	}

	var filenames []string
	for _, file := range d_entries {
		fmt.Println(file.Name())
		filenames = append(filenames, file.Name())
	}
	//file_contents := string(bytes[:])
	fmt.Println(filenames)

	//worker.Map(file_contents)
	worker.Reduce()
}
