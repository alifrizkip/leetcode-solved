/*
Binary Search
[1,2,3,4,5,6,7,8] 2

[1,2,3,4,5,6,7,8]
 ^       ^

[1,2,3,4,5,6,7,8]
 ^   ^

[1,2,3,4,5,6,7,8]
 ^ ^
*/
func searchMatrix(matrix [][]int, target int) bool {
    /*
        L -> Left
        R -> Right
        M -> Middle
    */
    var rowIdx int
    rowLIdx, rowRIdx := 0, len(matrix) - 1

    for rowLIdx <= rowRIdx {
        rowMIdx := (rowLIdx + rowRIdx) / 2
        rowM := matrix[rowMIdx]

        if target < rowM[0] {
            rowRIdx = rowMIdx - 1
        } else if target > rowM[len(rowM) - 1] {
            rowLIdx = rowMIdx + 1
        } else {
            rowIdx = rowMIdx
            break
        }
    }

    l, r := 0, len(matrix[rowIdx]) - 1
    for l <= r {
        m := (l + r) / 2
        if target < matrix[rowIdx][m] {
            r = m - 1
        } else if target > matrix[rowIdx][m] {
            l = m + 1
        } else {
            return true
        }
    }

    return false
}