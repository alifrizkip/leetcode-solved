func sortArrayByParity(nums []int) []int {
    res := make([]int, 0)
    
    for _, n := range nums {
        if n % 2 ==0 {
            res = append([]int{n}, res...)
        } else {
            res = append(res, n)
        }
    }
    
    return res
}