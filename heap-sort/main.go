package main

import(
	"fmt"	 
 )

var _ = fmt.Errorf

func main(){
	var arr = []int{81,72,35,8,1,2,9,7,5}
	HeapSort(arr)
	fmt.Println(arr)
}

func MaxHeapify(arr []int, parentIndex int){
	var (
		leftSubtreeIndex = parentIndex*2 + 1
		rightSubtreeIndex = parentIndex*2 + 2
		largestIndex int
	)

	if leftSubtreeIndex <= len(arr)-1 && arr[parentIndex] < arr[leftSubtreeIndex]{
		largestIndex = leftSubtreeIndex
	} else {
		largestIndex = parentIndex
	}

	if rightSubtreeIndex <= len(arr) -1 && arr[rightSubtreeIndex] > arr[largestIndex]{
		largestIndex = rightSubtreeIndex
	}

	if largestIndex != parentIndex {
		arr[largestIndex],arr[parentIndex] = arr[parentIndex],arr[largestIndex]
		MaxHeapify(arr,largestIndex)
	}
}

func BuildMaxHeap(arr []int){
	for i := len(arr)/2 - 1; i >= 0; i--{
		MaxHeapify(arr,i)
	}
}

func HeapSort(arr []int){
	BuildMaxHeap(arr)
	for i := len(arr)-1; i >=0; i-- {
		arr[0],arr[i] = arr[i],arr[0]
		MaxHeapify(arr[:i],0)
	}
}