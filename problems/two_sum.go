package problems

//EASY
//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//You can return the answer in any order.

func TwoSum(nums []int, target int) []int {

	var sliceInt []int

	for i := 0; i <= len(nums); i++ {
		for x := 0; x <= len(nums); x++ {
			if nums[i]+nums[x] == target {
				sliceInt = append(sliceInt, i)
				sliceInt = append(sliceInt, x)
			}
		}
	}

	return sliceInt

}
