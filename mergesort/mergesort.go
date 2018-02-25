package main

import (
	"fmt"
)

var (
	unsortedIntArr []int
)

func main() {
	unsortedIntArr = []int{6, 5, 4, 3, 2, 1}

	sorted := MergeSort(unsortedIntArr)
	fmt.Println(sorted)
}

// MergeSort ...
func MergeSort(arr []int) []int {
	arrLength := len(arr)
	if arrLength == 1 {
		return arr
	}

	leftArr := MergeSort(arr[:arrLength/2])
	rightArr := MergeSort(arr[arrLength/2:])

	return merge(leftArr, rightArr)
}

func merge(leftArr, rightArr []int) []int {
	retArr := []int{}

	for len(leftArr) > 0 && len(rightArr) > 0 {
		if leftArr[0] < rightArr[0] {
			retArr = append(retArr, leftArr[0])
			leftArr = rightArr[1:]
		} else {
			retArr = append(retArr, rightArr[0])
			rightArr = rightArr[1:]
		}
	}

	for len(leftArr) > 0 {
		retArr = append(retArr, leftArr[0])
		leftArr = leftArr[1:]
	}

	for len(rightArr) > 0 {
		retArr = append(retArr, rightArr[0])
		rightArr = rightArr[1:]
	}

	return retArr
}
