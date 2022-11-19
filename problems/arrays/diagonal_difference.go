package problems

import (
	"fmt"
	"math"
)

func DiagonalDifference(arr [][]int32) int32 {

	//Create new arrays to append the diagonal values to
	var d1 []int32
	var d2 []int32

	//Instantiate new ints to sum up the diagonal values
	var d1Sum int32
	var d2Sum int32

	//The double loop loops over the indexes, when they match (0, 0 - 1, 1 etc) it appends the value arr[x][y] to d1
	//Then adds it to d1Sum to do the calculation at the end
	for x := range arr {
		for y := range arr {
			fmt.Println("here", arr[x], arr[y], arr[x][y])
			fmt.Println(x, y)
			if x == y {
				d1 = append(d1, arr[x][y])
				d1Sum += arr[x][y]
			}

			//The opposite for the right diagonal. When the index is the length of the array minus
			//the index of y and -1 it appends to d2 and adds it to d2Sum
			if x == len(arr)-y-1 {
				d2 = append(d2, arr[x][y])
				d2Sum += arr[x][y]
			}
		}
	}
	fmt.Println("Left Diagonal Values", d1)
	fmt.Println("Right Diagonal Values", d2)

	return int32(math.Abs(float64(d1Sum) - float64(d2Sum)))
}
