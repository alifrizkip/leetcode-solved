func searchInsert(nums []int, target int) int {
    l, r := 0, len(nums) - 1

    for l < r {
        mid := (r - l) / 2
        midVal := nums[l+mid]

        if target > midVal {
            l += mid+1
        } else {
            r -= mid+1
        }
    }

    if target > nums[l] {
        return l+1
    }

    return l
}