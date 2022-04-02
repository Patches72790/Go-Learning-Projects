package worker

import (
	"fmt"
	"os"
	"sort"
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
		next_pair := struct {
			string
			int
		}{key, (*words)[key]}

		sorted_key_tuple = append(sorted_key_tuple, next_pair)
	}

	return &sorted_key_tuple
}

func Map(contents string) {
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

	file, err := os.Create("frankenstein.out")
	if err != nil {
		panic("Error opening output file")
	}

	for _, pair := range *sorted_key_tuples {
		out_str := string(pair.string) + ":" + string(fmt.Sprint(pair.int)) + "\n"
		file.WriteString(out_str)
	}
}

func Reduce() {
	panic("Unimplemented reduce!")
}
