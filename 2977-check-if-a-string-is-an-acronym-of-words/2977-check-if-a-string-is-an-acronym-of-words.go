func isAcronym(words []string, s string) bool {
    var check string
    for _, v := range words {
        check += string(v[0])
    }

    return check == s
}