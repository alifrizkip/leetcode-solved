func countConsistentStrings(allowed string, words []string) int {
    var res int

    set := map[rune]bool{}
    for _, c := range allowed{
        set[c] = true
    }

WLoop:
    for _, w := range words {
        for _, c := range w {
            if _, ok := set[c]; !ok {
                continue WLoop
            }
        }

        res++
    }

    return res
}