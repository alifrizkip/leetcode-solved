func sumOfUnique(nums []int) int {
    set := [101]int{}
    for _, n := range nums {
        set[n]++
    }

    var res int
    for i := 1; i < len(set); i++ {
        if set[i] == 1 {
            res += i
        }
    }

    return res
}