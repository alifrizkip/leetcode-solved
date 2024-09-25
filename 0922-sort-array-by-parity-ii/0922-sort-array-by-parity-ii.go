func sortArrayByParityII(nums []int) []int {
    evenIdx, oddIdx := 0, 1
    
    for evenIdx < len(nums) && oddIdx < len(nums) {
        if nums[evenIdx] % 2 == 0 {
            evenIdx += 2
        } else if nums[oddIdx] % 2 == 1 {
            oddIdx += 2
        } else {
            nums[evenIdx], nums[oddIdx] = nums[oddIdx], nums[evenIdx]
            evenIdx += 2
            oddIdx += 2
        }
    }
    
    return nums
}
