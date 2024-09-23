func intersection(nums1 []int, nums2 []int) []int {
    res := make([]int, 0)
    set := make(map[int]bool)
    
    for _, n := range nums1 {
        set[n] = false
    }
    
    for _, n := range nums2 {
        if val, ok := set[n]; ok && !val {
            set[n] = true
            res = append(res, n)
        }
    }
    
    return res
}