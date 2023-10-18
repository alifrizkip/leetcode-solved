import(
    "strings"
)

func reverseOnlyLetters(s string) string {
    sArr := strings.Split(s, "")

    l, r := 0, len(sArr)-1
    for l < r {
        if !isLetter([]byte(sArr[l])[0]) {
            l++
            continue
        }

        if !isLetter([]byte(sArr[r])[0]) {
            r--
            continue
        }

        sArr[l], sArr[r] = sArr[r], sArr[l]
        l++
        r--
    }

    return strings.Join(sArr, "")
}

func isLetter(c byte) bool {
    return (c-'a' >= 0 && c-'a' < 26) || (c-'A' >= 0 && c-'A' < 26)
}