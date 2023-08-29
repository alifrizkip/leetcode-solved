/*
{ v d f }
dvdf
   r
 l
3
*/
func lengthOfLongestSubstring(s string) int {
    set := map[byte]bool{}
    var res, l int

    for r := 0; r < len(s); r++ {
        _, ok := set[s[r]]
        for ok {
            delete(set, s[l])
            l++

            _, ok = set[s[r]]
        }

        set[s[r]] = true
        if r-l + 1 > res {
            res = r-l + 1
        }
    }

    return res
}

// without window sliding
func lengthOfLongestSubstringNoWindowSliding(s string) int {
    dict := [256]int{}
    var res, start int

    for i := 0; i < len(s); i++ {
        if start < dict[s[i]] {
            start = dict[s[i]]
        }

        if i - start+1 > res {
            res = i - start+1
        }

        dict[s[i]] = i + 1
    }

    return res
}