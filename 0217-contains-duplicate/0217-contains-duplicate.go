func containsDuplicate(nums []int) bool {
    set := map[int]int{}

    for _, n := range nums {
        if _, ok := set[n]; ok {
            return true
        }

        set[n]++
    }

    return false
}