func decodeMessage(key string, message string) string {
    keysMap := map[rune]rune{}
    keys := [26]int{}

    incIdx := 0
    for _, c := range key {
        if c == ' ' {
            continue
        }

        if _, ok := keysMap[c - 'a']; ok {
            continue
        }

        keys[c - 'a'] = incIdx
        incIdx++
        keysMap[c - 'a'] = c - 'a'
    }
    
    var res string
    for _, c := range message {
        if c == ' ' {
            res += string(c)
            continue
        }

        res += string(keys[c-'a'] + 'a')
    }

    return res
}