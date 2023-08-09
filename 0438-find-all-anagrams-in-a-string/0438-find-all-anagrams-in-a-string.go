func findAnagrams(s string, p string) []int {
    res := []int{}
    var sL, pL int = len(s), len(p)

    pFreq := [26]int{}
    for _, c := range p {
        pFreq[c - 'a']++
    }

    freq := [26]int{}
    var start int
    for end := 0; end < sL; end++ {
        c := rune(s[end])
        freq[c - 'a']++

        if end >= pL {
            freq[s[start] - 'a']--
            start++
        }

        if freq == pFreq {
            res = append(res, start)
        }
    }

    return res
}