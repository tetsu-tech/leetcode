package main

import "fmt"

func main() {
	array := []int{5, 6, 40, 4, 1, 3, 2, 3}
	sorted_array := selectionSort(array)
	fmt.Printf("%d\n", sorted_array)
}

func selectionSort(array []int) []int {
	n := len(array)
	for i := range array {
		for j := i + 1; j < n; j++ {
			fmt.Println(array[i], array[j])
			if array[i] > array[j] {
				temp := array[j]
				array[j] = array[i]
				array[i] = temp
			}
		}
	}
	return array
}
