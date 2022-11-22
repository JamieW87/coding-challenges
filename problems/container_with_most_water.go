package problems

func ContainerWithMostWater(h []int) int {

	// Maximum Volume
	m := 0
	// Left Index
	l := 0
	// Right Index
	r := len(h) - 1

	//Range until left is equal to right
	for l != r {
		// Volume is height of lowest wall, times the distance between walls
		v := min(h[l], h[r]) * (r - l)
		// If volume is bigger than current maximum, save new maximum
		if v > m {
			m = v
		}

		// Keep highest wall and move the other end towards it
		if h[l] > h[r] {
			r--
		} else {
			l++
		}
	}
	return m
}

//Returns smallest value
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
