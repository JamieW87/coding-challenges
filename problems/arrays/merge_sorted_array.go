package problems

import (
	"fmt"
	"sort"
)

//You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
//Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
//
//Example 1:
//
//Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//Output: [1,2,2,3,5,6]
//Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
//The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {

	//Appends nums2 to nums1
	nums1 = append(nums1[:m], nums2...)

	//Sets length of nums 1 to length of new array
	c := m + n
	nums1 = nums1[:c]

	//Loop over to remove zeros
	//Loops 2 times to remove last zero... Need to find out why it does this?!
	for x := 0; x < 2; x++ {
		for i := 0; i < len(nums1); i++ {
			if nums1[i] == 0 {
				nums1 = removeZeroIndex(nums1, i)
			}
		}
	}

	//Sort in ascending order.
	sort.Slice(nums1, func(i int, j int) bool {
		return nums1[i] < nums1[j]
	})

	fmt.Println(nums1)

}

func removeZeroIndex(s []int, index int) []int {

	return append(s[:index], s[index+1:]...)
}
