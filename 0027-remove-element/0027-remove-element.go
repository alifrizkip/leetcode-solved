func removeElement(nums []int, val int) int {
    var occ int
    
    for _, n := range nums {
        if n != val {
            nums[occ] = n
            occ++
        }
    }
    
    return occ
}