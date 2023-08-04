func removePalindromeSub(s string) int {
    for l, r := 0, len(s)-1; l < r; l++ {
        if s[l] != s[r] {
            return 2
        }
        r--
    }

    return 1
}