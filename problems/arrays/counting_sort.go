package problems

//For an array of integers, Each time a value occurs in the original array,
//you increment the counter at that index in the new array
//i	arr[i]	result
//0	1	[0, 1, 0, 0]
//1	1	[0, 2, 0, 0]
//2	3	[0, 2, 0, 1]
//3	2	[0, 2, 1, 1]
//4	1	[0, 3, 1, 1]

func CountingSort(arr []int32) []int32 {

	//Make new int slice with length of 100
	// make new empty map int int32
	sortA := make([]int32, 100)
	m := make(map[int]int32)

	//For loop over array, entering every value into the map
	for i := 0; i < len(arr); i++ {
		m[i] = arr[i]
	}

	//Loop over the array, if the value is in the map, increment the new array by 1, at that index
	// EG 63 in the original array increments sortA[63] by 1
	for i, v := range arr {
		if v == m[i] {
			sortA[v]++
		}
	}

	return sortA

}
