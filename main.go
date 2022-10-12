package main

import (
	"fmt"
	"leet-code-playingz/problems"
)

func main() {

	intSlice := []int{3, 2, 4}

	result := problems.TwoSum(intSlice, 6)
	fmt.Println("Indices of numbers that reach target:", result)

}
