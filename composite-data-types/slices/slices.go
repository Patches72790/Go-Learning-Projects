package slices

import (
	"fmt"
)

func reverse(s []int) []int {
	var s_copy = make([]int, len(s))

	copy(s_copy, s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s_copy[i], s_copy[j] = s_copy[j], s_copy[i]
	}

	return s_copy
}

func reverse_in_place(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}

func rotate_n(s []int, n int) []int {

	//	if n > len(s) {
	//		panic("cannot rotate longer than length")
	//	}

	var s_copy = make([]int, len(s))

	copy(s_copy, s)

	n = n % len(s)

	reverse_in_place(s_copy[:n])
	reverse_in_place(s_copy[n:])
	reverse_in_place(s_copy)

	return s_copy
}

func int_slice_eq(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// concept of slice type as a ptr with len and capacity
type Slice[T any] struct {
	ptr      *T
	len, cap int
}

func test_append() {

}

func TestSlices() {
	// slices and reverse and equality
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(reverse(a))
	fmt.Println(rotate_n(a, 6))
	fmt.Printf("a{%v} == rotate(a, 6){%v} ? %v\n", a,
		rotate_n(a, 6),
		int_slice_eq(a, rotate_n(a, 6)),
	)

	// append

}
