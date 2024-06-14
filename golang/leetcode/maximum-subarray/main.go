package maximum_subarray

/**
53. Maximum Subarray
https://leetcode.com/problems/maximum-subarray/description/
https://en.wikipedia.org/wiki/Maximum_subarray_problem#Kadane's_algorithm

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
*/

func maxSubArray(nums []int) int {
	var currentSum = 0
	var finalSum = int(-10e4)
	var lenNums = len(nums)

	for i := 0; i < lenNums; i++ {
		currentSum += nums[i]

		if currentSum > finalSum {
			finalSum = currentSum
		}

		if currentSum < 0 {
			currentSum = 0
		}
	}

	return finalSum
}
