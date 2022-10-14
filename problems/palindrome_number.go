package problems

import "fmt"

//EASY
//Given an integer x, return true if x is palindrome integer.
//An integer is a palindrome when it reads the same backward as forward.
//For example, 121 is a palindrome while 123 is not.

func IsPalindrome(x int) bool {

	orig := x
	result := 0
	for x > 0 {
		remainder := x % 10
		result = (result * 10) + remainder
		x /= 10
	}

	fmt.Println(result)
	fmt.Println(orig)

	if result == orig {
		return true
	}

	return false

}
