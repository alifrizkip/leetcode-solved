func firstPalindrome(words []string) string {
    for _, val := range words {
        if isPalindrome(val) {
            return val
        }
    }

    return ""
}

func isPalindrome(word string) bool {
    for i := 0; i < len(word)/2; i++ {
        if word[i] != word[len(word)-1-i] {
            return false
        }
    }

    return true
}