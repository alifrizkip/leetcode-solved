func firstPalindrome(words []string) string {
    for _, word := range words {
        if isPalindromic(word) {
            return word
        }
    }
    
    return ""
}

func isPalindromic(word string) bool {
    for i := 0; i < len(word) / 2; i++ {
        if word[i] != word[len(word)-1-i] {
            return false
        }
    }
    
    return true
}