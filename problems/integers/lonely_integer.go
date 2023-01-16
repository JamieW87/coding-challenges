package problems

//Find the integer in the array that only occurs once.
// Assume that there will always be ONE integer that occurs once

func LonelyInteger(a []int32) int32 {

	m := make(map[int32]int32)

	for _, v := range a {
		m[v]++
	}

	for _, k := range a {
		if m[k] == 1 {
			return k
		}
	}
	return 0
}
