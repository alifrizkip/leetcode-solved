func findWords(words []string) []string {
    rows := [26]int{}

    kb := []string{
        "qwertyuiop",
        "asdfghjkl",
        "zxcvbnm",
    }

    for i, row := range kb {
        for _, c := range row {
            rows[c-'a'] = i
        }
    }

    res := []string{}
    for _, w := range words {
        sameRow := true
        wL := strings.ToLower(w)

        var r int = rows[wL[0]-'a']
        for _, c := range wL {
            if rows[c-'a'] != r {
                sameRow = false
            }
        }

        if sameRow {
            res = append(res, w)
        }
    }

    return res
}