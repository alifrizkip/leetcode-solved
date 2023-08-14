func maximizeSum(nums []int, k int) int {
    var res int
    var maxNum int
    for _, n := range nums {
        if n > maxNum {
            maxNum = n
        }
    }

    for k > 0 {
        res += maxNum
        maxNum += 1

        k--
    }

    return res
}