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
    for i := 0; i < len(newS)/2; i++ {
        if newS[i] != newS[len(newS)-1-i] {
            return false
        }   
    }
    
    return true
}