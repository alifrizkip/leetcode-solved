func longestPalindrome(s string) int {
    // rune 'z' - 'A'
    dicts := [58]int{}
    for _, c := range s {
        dicts[c - 'A']++
    }

    var res, one int
    for _, val := range dicts {
        if val % 2 == 0 {
            res += val
        } else {
            res += val - 1
            one = 1
        }
    }

    return res+one
}