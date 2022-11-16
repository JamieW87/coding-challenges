package problems

func ReverseArray(a []int32) []int32 {
	//If incoming array is nil, return a nil array
	if a == nil {
		return nil
	}

	var rev []int32

	//Iterate through a backwards, append them in reverse order
	for i := len(a) - 1; i >= 0; i-- {
		rev = append(rev, a[i])
	}

	return rev
}
