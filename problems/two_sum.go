package problems

//EASY
//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//You can return the answer in any order.

func TwoSum(nums []int, target int) []int {

	var slint []int
	for i := 0; i < len(nums); i++ {
		for x := i + 1; x < len(nums); x++ {
			if nums[i]+nums[x] == target {
				slint = append(slint, i)
				slint = append(slint, x)
			}
		}
	}
	return slint
}
