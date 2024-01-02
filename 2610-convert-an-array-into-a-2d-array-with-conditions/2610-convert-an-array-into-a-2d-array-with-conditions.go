func findMatrix(nums []int) [][]int {
    nFreq := map[int]int{}
    for _, n := range nums {
        nFreq[n]++
    }

    res := [][]int{}
    for true {
        var rest int

        item := []int{}
        for k, v := range nFreq {
            rest += v

            if v > 0 {
                item = append(item, k)
                nFreq[k]--
            }
        }

        if rest == 0 {
            break
        }

        res = append(res, item)
    }

    return res
}