func average(salary []int) float64 {
    var sum, min, max float64 = 0, 1000000, 1000

    for _, n := range salary {
        nF := float64(n)

        sum += nF
        if nF < min {
            min = nF
        }
        if nF > max {
            max = nF
        }
    }

    return (sum - min - max) / float64(len(salary)-2)
}