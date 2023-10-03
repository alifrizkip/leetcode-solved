func sortPeople(names []string, heights []int) []string {
    hOrder := map[int]int{}
    for idx, h := range heights {
        hOrder[h] = idx
    }
    sort.Ints(heights)

    res := []string{}
    for _, h := range heights {
        idx := hOrder[h]
        res = append(
            []string{names[idx]},
            res...,
        )
    }

    return res
}