package worker

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type key_pair struct {
	string
	int
}

func sort_keys(words *map[string]int) *[]key_pair {
	keys := make([]string, 0, len(*words))

	for k := range *words {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var sorted_key_tuple []key_pair = make([]key_pair, 0, len(keys))

	for _, key := range keys {
		next_pair := key_pair{key, (*words)[key]}
		sorted_key_tuple = append(sorted_key_tuple, next_pair)
	}

	return &sorted_key_tuple
}

func Map(filename string, contents string) {
	word_count := map[string]int{}

	lines := strings.Split(contents, "\n")

	var tokens []string
	for _, line := range lines {
		words_in_line := strings.Split(line, " ")
		tokens = append(tokens, words_in_line...)
	}

	for _, s := range tokens {
		var cleaned_string string
		for _, c := range s {
			if unicode.IsLetter(c) {
				cleaned_string += string(c)
			} else if !unicode.IsSpace(c) || !unicode.IsNumber(c) {
				continue
			}
		}

		cleaned_string = strings.ToLower(cleaned_string)
		if value, present := word_count[cleaned_string]; present {
			word_count[cleaned_string] = value + 1
		} else {
			word_count[cleaned_string] = 1
		}
	}

	sorted_key_tuples := sort_keys(&word_count)
	file := file_wrangling(filename)

	for _, pair := range *sorted_key_tuples {
		out_str := string(pair.string) + ":" + string(fmt.Sprint(pair.int)) + "\n"
		file.WriteString(out_str)
	}
}

func file_wrangling(filename string) *os.File {
	file, err := os.Create(filename + ".out")
	if err != nil {
		panic("Error opening output file" + err.Error())
	}

	return file
}

func Reduce(file_contents string) {
	lines := strings.Split(file_contents, "\n")

	key_pairs := []key_pair{}

	// build word map from file contents
	for _, line := range lines {
		line_contents := strings.Split(line, ":")
		if len(line_contents) < 2 || len(line_contents[0]) == 0 {
			continue
		}
		count, err := strconv.Atoi(line_contents[1])
		if err != nil {
			panic("Error converting string to int")
		}
		key_pairs = append(key_pairs, key_pair{line_contents[0], count})
	}

	// sort by number of occurrences
	sort.Slice(key_pairs, func(i, j int) bool { return key_pairs[i].int > key_pairs[j].int })

	file, err := os.Create("reduced.out")
	if err != nil {
		panic("Error creating reduce file")
	}

	// group by number of occurrences
	for _, pair := range key_pairs {
		out_str := pair.string + "==" + string(fmt.Sprint(pair.int)) + "\n"
		file.WriteString(out_str)
	}
}
