func spiralOrder(matrix [][]int) []int {
    const (
        RIGHT = iota
        DOWN
        LEFT
        UP
    )

    res := []int{}

    rowL, colL := len(matrix), len(matrix[0])
    row, col := 0, 0
    i, size := 0, rowL * colL

    dir := RIGHT
    for i < size {
        res = append(res, matrix[row][col])
        i++

        switch dir {
        case RIGHT:
            if row+col == colL-1 {
                dir = DOWN
                row++
            } else {
                col++
            }
        case DOWN:  
            if rowL-row == colL-col {
                dir = LEFT
                col--
            } else {
                row++
            }
        case LEFT:
            if row+col == rowL-1 {
                dir = UP
                row--
            } else {
                col--
            }
        case UP:
            if row-col == 1 {
                dir = RIGHT
                col++
            } else {
                row--
            }
        }
    }

    return res
}