package main

import "fmt"

func main() {
	array := []int{2, 8, 3, 0, 24, 18}
	sorted_array := bubbleSort(array)
	fmt.Printf("%d\n", sorted_array)
}

func bubbleSort(array []int) []int {
	array_len := len(array)
	for i := range array {
		for j := array_len - 1; j > i; j-- {
			if array[j-1] > array[j] {
				temp := array[j-1]
				array[j-1] = array[j]
				array[j] = temp
			}
		}
	}
	return array
}
