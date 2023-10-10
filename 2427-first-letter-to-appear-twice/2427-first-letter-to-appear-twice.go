func repeatedCharacter(s string) byte {
    set := [26]int{}

    var res byte
    for _, c := range s {
        idx := c - 'a'
        set[idx]++

        if set[idx] == 2 {
            res = byte(c)
            break
        }
    }

    return res
}