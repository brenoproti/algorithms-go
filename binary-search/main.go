package main

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	result := BinarySearch(arr, 6)

	if result >= 0 {
		println("Element found at index", result)
	} else {
		println("Element not found")
	}
}

func BinarySearch(arr []int, target int) int {
	left, rigth := 0, len(arr)-1

	for left <= rigth {
		mid := (left + rigth) / 2
		option := arr[mid]

		if option == target {
			return mid
		}

		if option > target {
			rigth = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
