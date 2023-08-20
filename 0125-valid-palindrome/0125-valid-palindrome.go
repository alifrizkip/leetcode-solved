import (
    "strings"
)

func isPalindrome(s string) bool {
    // strip
    sRune := []rune{}
    for _, char := range s {
        if ('a' <= char && char <= 'z') ||
        ('A' <= char && char <= 'Z') ||
        ('0' <= char && char <= '9') {
            sRune = append(sRune, char)
        }
    }

    // check palindrome
    newS := strings.ToLower(string(sRune))
    l, r := 0, len(newS) - 1
    for l < r {
        if newS[l] != newS[r] {
            return false
        }

        l++
        r--
    }
    
    return true
}