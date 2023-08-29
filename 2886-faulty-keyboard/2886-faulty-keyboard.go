func finalString(s string) string {
    var res string
    for _, c := range s {
        if c == 'i' {
            res = reverse(res)
            continue
        }

        res = res + string(c)
    }

    return res
}

func reverse(s string) string {
    var res string
    for _, c := range s {
        res = string(c) + res
    }

    return res
}