func numIslands(grid [][]byte) int {
    rowL, colL := len(grid), len(grid[0])

    dirs := [][2]int{
        {0, 1},
        {1, 0},
        {0, -1},
        {-1, 0},
    }

    var res int
    for row, rows := range grid {
        for col, val := range rows {
            if val == byte('0') {
                continue
            }

            q := [][2]int{{row, col}}
            for len(q) > 0 {
                cell := q[0]
                q = q[1:]

                if grid[cell[0]][cell[1]] == byte('0'){
                    continue
                }

                grid[cell[0]][cell[1]] = byte('0')
                for _, dir := range dirs {
                    r, c := cell[0]+dir[0], cell[1]+dir[1]
                    if r >= 0 && r < rowL &&
                      c >= 0 && c < colL &&
                      grid[r][c] == byte('1') {
                        q = append(q, [2]int{r, c})
                    }
                }
            }
            res++
        }
    }

    return res
}