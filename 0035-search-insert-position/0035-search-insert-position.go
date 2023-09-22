func searchInsert(nums []int, target int) int {
    l, r := 0, len(nums) - 1

    for l <= r {
        m := l + (r-l)/2

        if nums[m] == target {
            return m
        } else if nums[m] < target {
            l, r = m+1, r
        } else {
            l, r = l, m-1
        }
    }

    return l
}