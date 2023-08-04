func mergeAlternately(word1 string, word2 string) string {
    ans := make([]byte, 0)
    var idx1, idx2 int

    for idx1 < len(word1) || idx2 < len(word2) {
        if idx1 < len(word1) {
            ans = append(ans, word1[idx1])
            idx1++
        } else {
            ans = append(ans, word2[idx2:]...)
            idx2 += len(word2) - 1 + idx2
        }

        if idx2 < len(word2) {
            ans = append(ans, word2[idx2])
            idx2++
        } else {
            ans = append(ans, word1[idx1:]...)
            idx1 += len(word1) - 1 + idx1
        }
    }

    return string(ans)
}