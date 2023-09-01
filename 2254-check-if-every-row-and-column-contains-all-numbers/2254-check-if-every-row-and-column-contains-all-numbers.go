func checkValid(matrix [][]int) bool {
/*
{
    [0,0]: {
        1: 1,
        2: 0,
        3: 0,
    }
}
*/
    rows := make(map[int]map[int]int)
    cols := make(map[int]map[int]int)

    for rowIdx, row := range matrix {
        for colIdx, val := range row {
            // set map from nil
            if rows[rowIdx+1] == nil {
                rows[rowIdx+1] = make(map[int]int)
            }
            if cols[colIdx+1] == nil {
                cols[colIdx+1] = make(map[int]int)
            }

            // check invalid
            if rows[rowIdx+1][val] != 0 {
                return false
            }
            if cols[colIdx+1][val] != 0 {
                return false
            }

            rows[rowIdx+1][val]++
            cols[colIdx+1][val]++
        }
    }

    return true
}

func checkValid_Invalid(matrix [][]int) bool {
    var maxRowSum int
    for i := 1; i <= len(matrix); i++ {
        maxRowSum += i
    }

    for _, row := range matrix {
        var currSum int
        for _, col := range row {
            currSum += col
        }

        if currSum != maxRowSum {
            return false
        }
    }

    return true
}