package main

import "fmt"

func main() {
	arr := []int{3, 1, 2, 4, 6, 5, 7}
	SelectionSort(arr)
	fmt.Printf("%v\n", arr)
}

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
