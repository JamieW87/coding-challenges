package problems

import (
	"strings"
)

func LengthOfLastWord(s string) int {

	//Removes extra spaces from string
	s2 := strings.TrimSpace(strings.Replace(s, "  ", " ", -1))
	//Splits string into seperate words in string array
	subStr := strings.Split(s2, " ")
	//Gets length of last entry
	last := subStr[len(subStr)-1]

	return len(last)
}
