func longestPalindrome(s string) int {
    dicts := map[string]int{}
    for _, c := range s {
        dicts[string(c)]++
    }

    var res int
    var one int
    for _, val := range dicts {
        if val % 2 == 0 {
            res += val
        } else {
            res += val - 1

            if one == 0 {
                one = 1
            }
        }
    }
    return res+one
}