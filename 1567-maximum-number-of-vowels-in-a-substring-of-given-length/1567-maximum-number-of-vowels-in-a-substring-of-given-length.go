func maxVowels(s string, k int) int {
    var res int
    var start, windowVow int
    for end := 0; end < len(s); end++ {
        if isVowel(string(s[end])) {
            windowVow++
        }

        if end >= k-1 {
            fmt.Println(windowVow)
            if windowVow > res {
                res = windowVow
            }

            // 0 < windowVow <= k
            if (windowVow > 0 && windowVow <= k) &&
              isVowel(string(s[start])) {
                windowVow--
            }
            start++
        }
    }

    return res
}

func isVowel(char string) bool {
    return char == "a" ||
        char == "e" ||
        char == "i" ||
        char == "o" ||
        char == "u"
}