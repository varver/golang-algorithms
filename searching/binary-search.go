package main

import "fmt"

func binarySearch(arrayzor []int, toFind int) int {

	var low, high int
	low = arrayzor[0]
	high = arrayzor[len(arrayzor)-1]

	if toFind > high || toFind < low {
		return -1
	}

	for low <= high {
		mid := low + (high-low)/2
		switch {
		case toFind < arrayzor[mid]:
			high = mid - 1

		case toFind > arrayzor[mid]:
			low = mid + 1

		case toFind == arrayzor[mid]:
			return mid
		}
	}
	return -1
}

func main() {

	fmt.Println("Binary search:")
	arrayzor := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	index := binarySearch(arrayzor, 10)
	if index == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Println("Index: ", index)
		fmt.Println("arrayzor[", index, "] = ", arrayzor[index])
	}
}
