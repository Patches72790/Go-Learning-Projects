package main

import (
	"fmt"
)

func rank_list(l []int) map[int]float32 {
	//ranks_map := make(map[int]float32, len(l))
	//sort.Ints(l)

	return nil
}

func map_list_type(list []int, name string) map[int]string {
	list_map := make(map[int]string, len(list))

	for _, el := range list {
		list_map[el] = name
	}

	return list_map
}

func wilcoxon_sum_rank_test(l1, l2 []int) float64 {
	// determine n1 and n2 : n1 <= n2
	var n1, n2 map[int]string
	if len(l1) <= len(l2) {
		n1 = map_list_type(l1, "n1")
		n2 = map_list_type(l2, "n2")
	} else {
		n1 = map_list_type(l2, "n1")
		n2 = map_list_type(l1, "n2")
	}

	return 0.0
}

func main() {
	l1 := []int{60, 68, 53, 26, 50, 72, 60, 54, 71, 46, 71}
	l2 := []int{29, 44, 33, 22, 44, 41, 22, 36, 35, 34}

	result := wilcoxon_sum_rank_test(l1, l2)

	fmt.Printf("Result of test: %f\n", result)
}
