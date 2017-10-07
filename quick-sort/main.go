package main

import (
	"fmt"
)

var _ = fmt.Errorf

func main() {
	var a = []int{2, 4, 6, 21, 87, 43, 9, 31}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func QuickSort(a []int, startIndex, endIndex int) {
	if startIndex < endIndex {
		q := Partition(a, startIndex, endIndex)
		QuickSort(a, startIndex, q-1)
		QuickSort(a, q+1, endIndex)
	}
}

func Partition(a []int, startIndex, endIndex int) int {
	pivot := a[endIndex]
	i := startIndex - 1

	for j := startIndex; j <= endIndex-1; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[endIndex] = a[endIndex], a[i+1]
	return i + 1
}
