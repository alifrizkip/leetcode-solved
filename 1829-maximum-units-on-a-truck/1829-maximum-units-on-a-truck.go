func maximumUnits(boxTypes [][]int, truckSize int) int {
    // sort by numberOfUnitsPerBox
    sort.Slice(boxTypes, func(i, j int) bool {
        return boxTypes[i][1] > boxTypes[j][1]
    })
    fmt.Println(boxTypes)

    var res int
    for _, b := range boxTypes {
        var nB, nUB int = b[0], b[1]

        for nB > 0 && truckSize > 0 {
            res += nUB
            nB--
            truckSize--
        }
    }

    return res
}