package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/patches72790/mapreduce/worker"

	//"log"
	"os"
)

//
// Reads from a file given by the dirname and filename
// and returns the files contents ready to be mapped
//
func read_filecontents(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file: ", filename)
	}

	return string(bytes[:])
}

func read_directory(dirname string) {
	d_entries, err := os.ReadDir(dirname)
	if err != nil {
		os.Exit(1)
	}

	var filenames []string
	for _, file := range d_entries {
		fmt.Println(file.Name())
		filenames = append(filenames, file.Name())
	}
}

func clean_filename(filename string) string {
	file_with_path_list := strings.Split(filename, string(os.PathSeparator))
	file_with_type := file_with_path_list[len(file_with_path_list)-1]

	return strings.Split(file_with_type, ".")[0]
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Error not enough args given.")
		fmt.Println("Usage: ./main <files...>")
		os.Exit(1)
	}

	filename := os.Args[1]
	file_contents := read_filecontents(filename)
	truncated_filename := clean_filename(filename)
	worker.Map(truncated_filename, file_contents)

	mapped_file_contents := read_filecontents(truncated_filename + ".out")
	worker.Reduce(mapped_file_contents)
}
