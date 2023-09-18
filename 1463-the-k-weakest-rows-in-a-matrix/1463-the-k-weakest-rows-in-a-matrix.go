func kWeakestRows(mat [][]int, k int) []int {
    idxMap := map[int][]int{}
    soldiers := []int{}

    for idx, rows := range mat {
        n := 0 // soldiers num

        for _, s := range rows {
            n += s
        }

        idxMap[n] = append(idxMap[n], idx)
        soldiers = append(soldiers, n)
    }
    sort.Ints(soldiers)

    res := []int{}
    for i := 0; i < k; i++ {
        res = append(res, idxMap[soldiers[i]][0])
        idxMap[soldiers[i]] = idxMap[soldiers[i]][1:]
    }

    return res
}