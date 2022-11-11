package problems

//Given two integer arrays. Return an integer array containing the values that appear in both arrays.

func MatchingArrays(arrA []int, arrB []int) []int {

	//Loop through first array and add to map,
	//then loop and lookup in map, if its in there, add to array
	var newArr []int
	m := make(map[int]int)

	for i := 0; i < len(arrA); i++ {
		//Add values for arrA to map
		m[i] = arrA[i]
	}

	for x := 0; x < len(arrB); x++ {

		for _, val := range m {
			if val == arrB[x] {
				newArr = append(newArr, arrB[x])
			}
		}
	}

	return newArr
}
