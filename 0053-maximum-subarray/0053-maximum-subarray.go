/*
- track current sum
- reset current sum if negative
- compare current sum with max sum
*/
func maxSubArray(nums []int) int {
    maxSum, currSum := nums[0], nums[0]

    for i := 1; i < len(nums); i++ {
        if currSum < 0 {
            currSum = nums[i]
        } else {
            currSum += nums[i]
        }

        if currSum > maxSum {
            maxSum = currSum
        }
    }

    return maxSum
}