func findMaxAverage(nums []int, k int) float64 {
    var res float64 = math.Inf(-1)

    start := 0
    windowSum := 0
    for end := 0; end < len(nums); end++ {
        windowSum += nums[end]

        if end >= k-1 {
            tempAvg := float64(windowSum) / float64(k)
            if tempAvg > res {
                res = tempAvg
            }

            windowSum -= nums[start]
            start++
        }
    }

    return res
}