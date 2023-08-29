/*
target = 11
res = -1
1 2 3 4 5
        r
    l
sum = 12
*/
func minSubArrayLen(target int, nums []int) int {
    res, l, sum := 100001, 0, 0

    for r := 0; r < len(nums); r++ {
        sum += nums[r]

        for sum >= target {
            if r-l + 1 < res {
                res = r-l + 1
            }

            sum -= nums[l]
            l++
        }
    }

    if res == 100001 {
        return 0
    }

    return res
}