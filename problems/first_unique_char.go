package problems

func FirstUniqChar(s string) int {

	// Makes a map rune int
	//Ranges over the string(s) and assigns every rune (j) to the map
	m := make(map[rune]int, len(s))
	for _, j := range s {
		m[j]++
	}

	//Ranges through the index and rune of the string
	//For the first character that only appears once, return the index of that character (i)
	for i, j := range s {
		if m[j] == 1 {
			return i
		}
	}

	return -1

}
