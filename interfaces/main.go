package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ByteCounter int

type WordCounter int

type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *ByteCounter) String() string {
	return fmt.Sprintf("%v", *c)
}

func (c *WordCounter) Write(p []byte) (int, error) {

	var words_count int
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words_count = words_count + 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

	*c += WordCounter(words_count)
	return words_count, nil
}
func (c *WordCounter) String() string {
	return fmt.Sprintf("%v", *c)
}

func (c *LineCounter) Write(p []byte) (int, error) {
	count := scanStuff(p, bufio.ScanWords)
	*c += LineCounter(count)
	return count, nil
}

func scanStuff(p []byte, fn bufio.SplitFunc) int {
	var count int
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(fn)

	for scanner.Scan() {
		count++
	}

	return count
}
func (c *LineCounter) String() string {
	return fmt.Sprintf("%v", *c)
}

func main() {

	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c.String())

	var wc WordCounter
	wc.Write([]byte("Hello my name is Stradivarius, and I play the violin"))
	fmt.Println(wc.String())

	var lc LineCounter
	lc.Write([]byte("Hello\nmy\nname\nis\nStradivarius,\nand\nI\nplay\nthe\nviolin\n"))
	fmt.Println(lc)
}
