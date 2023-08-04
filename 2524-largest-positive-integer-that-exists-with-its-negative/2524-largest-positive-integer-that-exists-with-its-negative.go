func findMaxK(nums []int) int {
    sort.Ints(nums)
    var ht = map[int]int{}
    var ans int = -1

    for i, n := range nums {
        if n < 0 {
            ht[n] = i
            continue
        }

        tmp := -n
        if _, ok := ht[tmp]; ok {
            ans = n
        }
    }

    return ans
}