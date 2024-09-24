func sortedSquares(nums []int) []int {
    var res = make([]int, 0)
    var l, r =  0, len(nums)-1
    
    for l <= r {        
        if nums[l]*nums[l] > nums[r]*nums[r] {
            res = append(res, nums[l]*nums[l])
            l++
        } else {
            res = append(res, nums[r]*nums[r])
            r--
        }
    }
    
    slices.Reverse(res)
    
    return res
}