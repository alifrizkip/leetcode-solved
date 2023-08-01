func makeSmallestPalindrome(s string) string {
    sRune := []rune(s)
    for i := 0; i < len(sRune) / 2; i++ {
        if sRune[i] < sRune[len(sRune)-i-1] {
            sRune[len(sRune)-i-1] = sRune[i]
        }

        if sRune[i] > sRune[len(sRune)-i-1] {
            sRune[i] = sRune[len(sRune)-i-1]
        }
    }

    return string(sRune)
}