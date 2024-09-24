func findMaxK(nums []int) int {
    sort.Ints(nums)
    
    var res int = -1
    set := make(map[int]int)
    
    for _, n := range nums {
        if n < 0 {
            set[n] = n
        }
        
        if n > 0 {
            if val, ok := set[n * -1]; ok && val < 0 {
                set[n * -1] = 0
                
                if n > res {
                    res = n
                }
            }
        }
    }
    
    return res
}