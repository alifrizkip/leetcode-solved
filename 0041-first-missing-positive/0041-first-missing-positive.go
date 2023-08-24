func firstMissingPositive(nums []int) int {
    hashMap := [100001]int{}

    for _, n := range nums {
        if 0 < n && n <= 100000 {
            hashMap[n] = 1
        }
    }

    var prev, res int
    for i := 1; i <= 100000; i++ {
        if i == 1 && hashMap[i] == 0 {
            res = 1
            break
        }

        if prev == 1 && hashMap[i] == 0 {
            res = i
            break
        }

        prev = hashMap[i]
    }

    if res == 0 {
        return 100000+1
    }

    return res
}