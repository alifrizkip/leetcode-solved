func truncateSentence(s string, k int) string {
    sArr := strings.Split(s, " ")

    return strings.Join(sArr[:k], " ")
}