func buildArray(nums []int) []int {
    ans := []int{}
    
    for _, num := range nums {
        ans = append(ans, nums[num])
    }
    
    return ans
}