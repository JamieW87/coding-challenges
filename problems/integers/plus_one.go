package problems

//You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.
//
//Increment the large integer by one and return the resulting array of digits.
//
//
//
//Example 1:
//
//Input: digits = [1,2,3]
//Output: [1,2,4]
//Explanation: The array represents the integer 123.
//Incrementing by one gives 123 + 1 = 124.
//Thus, the result should be [1,2,4].

func PlusOne(digits []int) []int {

	//Loop over the slice backwards
	for i := len(digits) - 1; i >= 0; i-- {
		//If the digit is a 9, set it to 0
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			//If the digit is not a 9, increment it by one and break out of the loop
			digits[i]++
			break
		}
	}

	//Then if the first digit is a zero, append a 1 to the front of the slice
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
