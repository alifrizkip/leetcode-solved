// DFS
func countBattleships(board [][]byte) int {
    res, rowL, colL := 0, len(board), len(board[0])

    var dfs func(int, int)
    dfs = func(row, col int) {
        if row < rowL &&
          col < colL &&
          board[row][col] == 'X' {
            board[row][col] = '.'

            dfs(row+1, col)
            dfs(row, col+1)
        }
    }

    for rowIdx, row := range board {
        for colIdx, cell := range row {
            if cell == '.' {
                continue
            }

            dfs(rowIdx, colIdx)

            res++
        }
    }

    return res
}

// BFS
func countBattleships_BFS(board [][]byte) int {
    var res int
    rowL, colL := len(board), len(board[0])

    hor := [1][2]int{{0, 1}}
    ver := [1][2]int{{1, 0}}

    var bfs func(int, int, [1][2]int)
    bfs = func(rowIdx, colIdx int, dirs [1][2]int) {
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

            bfs(rowIdx, colIdx, hor)
            bfs(rowIdx, colIdx, ver)

            res++
        }
    }

    return res
}