func longestSubarray(nums []int) int {
    var res int
    var left, right int
    var foundZero bool
    for end := 0; end < len(nums); end++ {
        n := nums[end]

        if n == 1 {
            right++

            if left+right > res {
                res = left+right
            }
        }

        if n == 0 {
            left = right
            right = 0
            foundZero = true
        }
    }

    if !foundZero {
        return len(nums)-1
    }

    return res
}