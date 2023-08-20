func isAnagram(s string, t string) bool {
    if len(t) < len(s) {
        s, t = t, s
    }

    dict := map[rune]int{}
    for _, c := range s {
        dict[c] += 1
    }

    for _, c := range t {
        v, ok := dict[c]
        if ok && v > 0 {
            dict[c]--
        } else {
            return false
        }
    }

    return true
}