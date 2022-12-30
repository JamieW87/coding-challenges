package main

import problems "leet-code-playingz/problems/arrays"

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
	//intSlice := []int32{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//res := problems.ContainerWithMostWater(intSlice)
	//fmt.Println(res)

	//--MatchingArrays
	//resArr := problems.MatchingArrays(intSlice, reIntSlice)
	//fmt.Println(resArr)

	//--SortArray
	//sortedSlice := problems.SortArray(intSlice)
	//fmt.Println(sortedSlice)

	//--DiagonalDifference
	//sl := []int32{2, 3, 4}
	//sl2 := []int32{5, 6, 7}
	//sl3 := []int32{1, 2, 5}
	//slice := [][]int32{sl, sl2, sl3}

	//sum := problems.DiagonalDifference(slice)
	//fmt.Println("Sum: ", sum)

	//length := problems.LengthOfLastWord("hello worldddd")
	//fmt.Println(length)

	//---Merge Sorted arrays---
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{2, 5, 6, 0, 0, 0}

	problems.MergeSortedArray(nums1, 4, nums2, 3)

}
