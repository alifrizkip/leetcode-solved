func countGoodSubstrings(s string) int {
    var res int
    var subs = []string{}

    for i := 0; i < len(s) - 2; i++ {
        sub := []byte{s[i], s[i+1], s[i+2]}
        subs = append(subs, string(sub))
    }

    for _, val := range subs {
        ht := map[byte]int{}
        a, b, c := val[0], val[1], val[2]
        if _, ok := ht[a]; ok {
            continue
        } else {
            ht[a] += 1
        }
        if _, ok := ht[b]; ok {
            continue
        } else {
            ht[b] += 1
        }
        if _, ok := ht[c]; ok {
            continue
        } else {
            ht[c] += 1
        }

        res++
    }

    fmt.Println(subs)

    return res
}