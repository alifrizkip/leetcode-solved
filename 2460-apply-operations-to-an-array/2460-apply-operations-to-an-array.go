func applyOperations(nums []int) []int {
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] == nums[i+1] {
            nums[i] *= 2
            nums[i+1] = 0
        }
    }
    
    res := make([]int, len(nums))
    var p int
    for _, n := range nums {
        if n != 0 {
            res[p] = n
            p++
        }
    }
    
    return res
}