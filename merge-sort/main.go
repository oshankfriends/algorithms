package main
import (
	"fmt"
	"math"
)

var _ = fmt.Errorf

func main(){
	var arr = []int{7,6,64,2,4,55,34,66,32,12,24,54,67,123,45}
	fmt.Println("Arr : ",arr)
	MergeSort(arr,0,len(arr)-1)
	fmt.Println("Sorted Arr : ",arr)
}

func MergeSort(arr []int, startIndex, endIndex int){
	if startIndex >= endIndex{
		return
	}
	midIndex := (startIndex + endIndex ) / 2
	MergeSort(arr,startIndex,midIndex)
	MergeSort(arr,midIndex+1,endIndex)
	Merge(arr,startIndex,midIndex,endIndex)
}

func Merge(arr []int, start, mid, end int){
	leftSubArr := make([]int,mid-start+2)
	rightSubArr :=make([]int,end-mid+1)
	copy(leftSubArr,arr[start:mid+1])
	copy(rightSubArr,arr[mid+1:end+1])
	leftSubArr[len(leftSubArr)-1] = int(math.MaxInt8)
	rightSubArr[len(rightSubArr)-1] = int(math.MaxInt8)
	var j,k = 0,0
	for i := start; i <= end; i++{
		if leftSubArr[j] < rightSubArr[k]{
			arr[i] = leftSubArr[j]
			j++
		} else {
			arr[i] = rightSubArr[k]
			k++
		}
	}
}