func findMaxAverage(nums []int, k int) float64 {
    var res float64 = math.Inf(-1)

    for i := 0; i < len(nums) - k + 1; i++ {
        var tmp float64
        for _, val := range nums[i:i+k] {
            tmp += float64(val)
        }

        var avg float64 = float64(tmp) / float64(k)
        if avg > res {
            res = avg
        }
    }

    return res
}