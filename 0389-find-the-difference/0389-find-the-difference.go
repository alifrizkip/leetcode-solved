func findTheDifference(s string, t string) byte {
    if len(s) == 0 {
        return byte(t[0])
    }

    sArr, tArr := strings.Split(s, ""), strings.Split(t, "")
    sort.Strings(sArr)
    sort.Strings(tArr)
    for i := 0; i < len(s); i++ {
        if sArr[i] != tArr[i] {
            return []byte(tArr[i])[0]
        }
    }

    return []byte(tArr[len(tArr)-1])[0]
}