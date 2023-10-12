func canBeTypedWords(text string, brokenLetters string) int {
    words := strings.Split(text, " ")

    brokenSets := map[rune]bool{}
    for _, c := range brokenLetters {
        brokenSets[c] = true
    }

    var res int
    Words:
        for _, w := range words {
            res++
            for _, c := range w {
                if _, ok := brokenSets[c]; ok {
                    res--
                    continue Words
                }
            }
        }

    return res
}