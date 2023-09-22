/*
bikin 2 pointer untuk s, t
loop t -> char
  - if check char with s[sIdx]
    - true -> sIdx += 1
    - false skip 

return sIdx == len(s) - 1
*/
func isSubsequence(s string, t string) bool {
    var sIdx int

    if s == "" {
        return true
    }

    for _, char := range t {
        if string(char) == string(s[sIdx]) {
            sIdx += 1
        }

        if sIdx == len(s) {
            break
        }
    }

    return sIdx == len(s)
}