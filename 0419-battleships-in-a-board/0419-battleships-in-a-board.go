func countBattleships(board [][]byte) int {
    var res int
    rowL, colL := len(board), len(board[0])

    hor := [2][2]int{{0, -1}, {0, 1}}
    ver := [2][2]int{{-1, 0}, {1, 0}}

    var visit func(int, int, [2][2]int)
    visit = func(rowIdx, colIdx int, dirs [2][2]int) {
        q := [][2]int{{rowIdx, colIdx}}

        for len(q) > 0 {
            cR, cC := q[0][0], q[0][1]
            board[cR][cC] = '.'

            q = q[1:]
            for _, d := range dirs {
                dR, dC := cR+d[0], cC+d[1]
                if 0 <= dR && dR < rowL &&
                    0 <= dC && dC < colL &&
                    board[dR][dC] == 'X' {
                    q = append(q, [2]int{dR, dC})
                }
            }
        }
    }

    for rowIdx, row := range board {
        for colIdx, cell := range row {
            if cell == '.' {
                continue
            }

            // visit
            visit(rowIdx, colIdx, hor)
            visit(rowIdx, colIdx, ver)

            res++
        }
    }

    return res
}