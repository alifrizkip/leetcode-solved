func canConstruct(ransomNote string, magazine string) bool {
    magz := [26]int{}
    for _, c := range magazine {
        magz[c-'a']++
    }

    for _, c := range ransomNote {
        if magz[c-'a'] < 1 {
            return false
        }

        magz[c-'a']--
    }

    return true
}