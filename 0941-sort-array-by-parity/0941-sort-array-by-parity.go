func sortArrayByParity(nums []int) []int {
    var nextEvenIdx int
    for i, n := range nums {
        if n % 2 == 0 {
            nums[i], nums[nextEvenIdx] = nums[nextEvenIdx], nums[i]
            nextEvenIdx++
        }
    }

    return nums
}