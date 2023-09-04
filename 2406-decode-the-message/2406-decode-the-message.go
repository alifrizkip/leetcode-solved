func decodeMessage(key string, message string) string {
    // can't use 0 because 'a' is 0
    keys := [26]int{
        -1, -1, -1, -1, -1,
        -1, -1, -1, -1, -1,
        -1, -1, -1, -1, -1,
        -1, -1, -1, -1, -1,
        -1, -1, -1, -1, -1,
        -1,
    }

    incIdx := 0
    for _, c := range key {
        if c == ' ' {
            continue
        }

        if keys[c-'a'] != -1 {
            continue
        }

        keys[c - 'a'] = incIdx
        incIdx++
    }
    fmt.Println(keys)
    
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