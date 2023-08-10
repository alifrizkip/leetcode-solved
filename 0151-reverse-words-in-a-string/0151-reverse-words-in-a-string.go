import (
    "strings"
)

func reverseWords(s string) string {
    // split by space
    // if s[i] = space AND s[i-1]=space skip
    res := []string{}
    sArr := strings.Split(s, " ")
    var prevW string
    for _, w := range sArr {
        if prevW == "" && w == "" {
            prevW = w
            continue
        } else {
            res = append([]string{w}, res...)
        }
    }

    return strings.Join(res, " ")
}