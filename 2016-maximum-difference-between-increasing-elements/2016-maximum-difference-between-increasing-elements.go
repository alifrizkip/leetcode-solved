func maximumDifference(nums []int) int {
    i, j := nums[0], 0

    for _, n := range nums {
        if n < i {
            i = n
        }
        
        diff := n - i
        if diff > j {
            j = diff
        }
    }
    
    if j == 0 {
        return -1
    }
    
    return j
}