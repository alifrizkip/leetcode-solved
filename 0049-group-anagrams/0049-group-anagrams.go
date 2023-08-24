func groupAnagrams(strs []string) [][]string {
    m := map[[26]int][]string{}

    for _, word := range strs {
        k := [26]int{}
        for _, c := range word {
            k[c - 'a'] += 1
        }
        m[k] = append(m[k], word)
    }

    res := [][]string{}
    for _, v := range m {
        res = append(res, v)
    }

    return res
}

func groupAnagramsFail(strs []string) [][]string {
    numMaps := map[int][]string{}

    for _, word := range strs {
        num := 1
        for _, c := range word {
            num *= int(c)
        }

        if _, ok := numMaps[num]; ok {
            numMaps[num] = append(numMaps[num], word)
        } else {
            numMaps[num] = []string{word}
        }
    }

    res := [][]string{}
    for _, v := range numMaps {
        res = append(res, v)
    }

    return res
}