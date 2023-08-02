func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    var ans [][]int
    var leftIdx, rightIdx int
    
    for leftIdx < len(nums1) || rightIdx < len(nums2) {
        var left, right []int
        if leftIdx < len(nums1) {
            left = nums1[leftIdx]
        }
        if rightIdx < len(nums2) {
            right = nums2[rightIdx]
        }
        
        if left != nil && right != nil {
            if left[0] == right[0] {
                ans = append(ans, []int{left[0], left[1]+right[1]})
                leftIdx++
                rightIdx++
            }
            
            if left[0] < right[0] {
                ans = append(ans, []int{left[0], left[1]})
                leftIdx++
            }
            
            if left[0] > right[0] {
                ans = append(ans, []int{right[0], right[1]})
                rightIdx++
            }
        } else if left != nil && right == nil {
            ans = append(ans, []int{left[0], left[1]})
            leftIdx++
        } else if left == nil && right != nil {
            ans = append(ans, []int{right[0], right[1]})
            rightIdx++
        }
    }
    
    return ans
}