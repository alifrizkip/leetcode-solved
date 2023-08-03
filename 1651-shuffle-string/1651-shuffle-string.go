func restoreString(s string, indices []int) string {
    sRune := make([]byte, len(s))
    
    for idx, val := range indices {
        sRune[val] = s[idx]
    }
    
    return string(sRune)
}