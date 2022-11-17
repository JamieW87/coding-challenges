package problems

import (
	"fmt"
)

//Brute forcing it. Need to convert to int64s to allow for int overflow (could be big number)
func miniMaxSum(arr []int32) {
	if arr == nil {
		fmt.Println("0 0")
	}

	var sum int64
	largest := arr[0]
	smallest := arr[0]

	for i := 0; i < len(arr); i++ {
		sum = sum + int64(arr[i])
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}

	fmt.Println(sum-int64(largest), sum-int64(smallest))

}

//Way to do it by appending them into an array then just sorting the array and taking the bottom one and top one off the sum.
//func miniMaxSum(arr []int) {
//	var _arr []int
//	var i, j int
//
//	arrayLength := int(len(arr))
//
//	for i = 0; i < arrayLength; i++ {
//		var sum int
//		for j = 0; j < arrayLength; j++ {
//			if i != j {
//				sum += arr[j]
//			}
//		}
//		_arr = append(_arr, sum)
//	}
//
//	sort.Ints(_arr)
//	fmt.Println(_arr[0], _arr[arrayLength-1])
//}
