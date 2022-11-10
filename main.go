package main

import (
	"fmt"
	"leet-code-playingz/problems"
)

func main() {

	//--TwoSum--
	//intSlice := []int{3, 2, 4}
	//result := problems.TwoSum(intSlice, 6)
	//fmt.Println("Indices of numbers that reach target:", result)

	//--Palindrome--
	//isPal := problems.IsPalindrome(121)
	//fmt.Println(isPal)

	//--RemoveElement--
	//reIntSlice := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//res := problems.RemoveElement(reIntSlice, 2)

	//--RomanToInteger--
	//res := problems.RomanToInt("VI")
	//fmt.Println(res)

	//--FirstUniqueCharInString
	//charIndex := problems.FirstUniqChar("loveleetcode")
	//fmt.Println(charIndex)

	//--ContainerWithMostWater
	intSlice := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := problems.ContainerWithMostWater(intSlice)
	fmt.Println(res)

}
