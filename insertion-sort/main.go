package main

import (
	"fmt"
)

var _ = fmt.Errorf

func main(){
	var arr = []int{7,612,64,2,4,55,34,66,32,12,24,54,67,123,45}
	fmt.Println("Arr : ",arr)
	InsertionSort(arr)
	fmt.Println("Sorted Arr : ",arr)
}

func InsertionSort(arr []int){
	for j:= 1; j <= len(arr)-1; j++{
		key := arr[j]
		i := j-1
		for i >= 0 && key < arr[i] {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = key
	}
}