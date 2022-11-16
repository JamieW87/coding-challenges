package problems

import "fmt"

func FirstUniqChar(s string) int {

	// Makes a map rune int
	//Ranges over the string(s) and assigns a 1 to every rune
	m := make(map[rune]int, len(s))
	for _, j := range s {
		m[j]++
	}

	//Ranges through the index and rune of the string
	//Returns the index of the first character that has a 1 assigned to it
	for i, j := range s {
		if m[j] == 1 {
			fmt.Println(string(s[i]))
			return i
		}
	}

	return -1

}
