package arrays

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"time"
)

func sha_diff_bit_shift(hash1, hash2 [32]byte) int {
	count := 0

	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {

			first := hash1[i] & (1 << j)
			second := hash2[i] & (1 << j)

			if first != second {
				count++
			}
		}
	}

	return count
}

func sha_diff_loop(hash1, hash2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {

			first := hash1[i] & (1 << j)
			second := hash2[i] & (1 << j)

			if first != second {
				count++
			}
		}
	}

	return count
}

func printPath() {
	file, b := os.LookupEnv("PATH")

	if !b {
		fmt.Fprintln(os.Stderr, "Problem reading file")
		os.Exit(1)
	}
	var contents string = string(file)
	paths := strings.Split(contents, ":")

	fmt.Println("Path:")
	for _, s := range paths {
		fmt.Println(s)
	}
}

func measureFnTime(f func() int, name string) int {
	start := time.Now()

	//diff := f()

	end := time.Since(start)
	//fmt.Printf("Number of diff bits: %d\n", diff)
	//fmt.Printf("Time to complete: %d\n", end)

	val := int(end)

	return val
}

func sum(a *[1000000]int) int {
	count := 0
	for _, el := range *a {
		count += el
	}

	return count
}

func TestShaTiming() {
	measureFnTime(func() int {
		return sha_diff_bit_shift(sha256.Sum256([]byte("1")), sha256.Sum256([]byte("2")))
	}, "hash_diff")

	iterations := 1000000
	box := [1000000]int{}
	first := sha256.Sum256([]byte("abc"))
	second := sha256.Sum256([]byte("123"))

	for i := 0; i < iterations; i++ {
		box[i] = measureFnTime(func() int {
			return sha_diff_bit_shift(first, second)
		}, "hash_diff")
	}

	fmt.Printf("Avg time: %d ns\n", (sum(&box) / iterations))

}
