func sortedSquares(nums []int) []int {
    var res = make([]int, 0)
    var l, r =  0, len(nums)-1
    
    for l <= r {
        vL, vR := nums[l]*nums[l], nums[r]*nums[r]
        
        if vL > vR {
            res = append([]int{vL}, res...)
            l++
        } else {
            res = append([]int{vR}, res...)
            r--
        }
    }
    
    return res
}