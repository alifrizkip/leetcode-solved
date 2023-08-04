func intersection(nums1 []int, nums2 []int) []int {
    var leftSet = make([]int, 1000)
    for _, n := range nums1 {
        leftSet[n]++
    }

    var rightSet = make([]int, 1000)
    for _, n := range nums2 {
        if leftSet[n] > 0 {
            rightSet[n]++
        }
    }

    var ans []int
    for i := 0; i < 1000; i++ {
        if rightSet[i] > 0 {
            ans = append(ans, i)
        }
    }
    return ans
}