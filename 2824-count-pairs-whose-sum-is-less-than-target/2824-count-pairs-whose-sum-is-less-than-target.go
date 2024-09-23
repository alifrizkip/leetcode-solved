func countPairs(nums []int, target int) int {
    var res int
    var l, r = 0, 1
    
    for r < len(nums) {        
        if nums[l] + nums[r] < target {
            res++
        }
        r++
        
        if r == len(nums) {
            l++
            r = l+1
        }
    }
    
    return res
}