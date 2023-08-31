/*
** BFS DP **
init queue
init MAX = m * n

visit each cell
if cell = 0
    add to queue
if cell = 1
    set value to MAX

dequeue cell
get neighbors
loop neighbors:
    if neighbor > cell value
        set neighbor value to cell+1
        add neighbor to queue
*/
func updateMatrix(mat [][]int) [][]int {
    rowL, colL := len(mat), len(mat[0])
    MAX_VALUE := rowL * colL
    q := make([][]int, 0)

    for r, rows := range mat {
        for c, _ := range rows {
            if mat[r][c] == 0 {
                q = append(q, []int{r, c})
            } else {
                mat[r][c] = MAX_VALUE
            }
        }
    }

    ngbrs := [][]int{
        []int{-1, 0},
        []int{0, 1},
        []int{1, 0},
        []int{0, -1},
    }
    for len(q) > 0 {
        c := q[0]
        q = q[1:]

        for _, n := range ngbrs {
            nR, nC := c[0]+n[0], c[1]+n[1]
            if (0 <= nR && nR < rowL) &&
                (0 <= nC && nC < colL) &&
                mat[nR][nC] > mat[c[0]][c[1]] {
                    mat[nR][nC] = mat[c[0]][c[1]] + 1
                    q = append(q, []int{nR, nC})
            }
        }
    }

    return mat
}

/*
** BFS Brute Force **
visit each tile
check tile:
 if 0 skip
if 1
get neighbors (up, right, down, left)
visit each neighbors
if neighbors 0, keep tile=1 & continue on next tile
*/
func updateMatrix_TimeLimitExceed_1x10000(mat [][]int) [][]int {
    rowL, colL := len(mat), len(mat[0])
    visited := map[[2]int]bool{}

    var getNeighbors func(int, int) [][2]int
    getNeighbors = func(row, col int) [][2]int {
        res := [][2]int{}

        up := [2]int{row-1, col}
        right := [2]int{row, col+1}
        down := [2]int{row+1, col}
        left := [2]int{row, col-1}

        if up[0] >= 0 {
            res = append(res, up)
        }

        if right[1] < colL {
            res = append(res, right)
        }

        if down[0] < rowL {
            res = append(res, down)
        }

        if left[1] >= 0 {
            res = append(res, left)
        }

        return res
    }

    var distance func([2]int) int
    distance = func(cell [2]int) int {
        var res int
        // visited := map[[2]int]bool{}

        willVisit := [][3]int{
            [3]int{cell[0], cell[1], 0},
        }

        var foundZero bool
        for !foundZero && len(willVisit) > 0 {
            // set visited
            visited[[2]int{willVisit[0][0], willVisit[0][1]}] = true

            row, col, dist := willVisit[0][0], willVisit[0][1], willVisit[0][2]
            // deque willVisit
            willVisit = willVisit[1:]

            c := mat[row][col]
            if c == 0 {
                res = dist
                foundZero = true
                break
            }

            neighbors := getNeighbors(row, col)
            for _, n := range neighbors {
                if _, ok := visited[[2]int{n[0],n[1]}]; !ok {
                    willVisit = append(
                        willVisit,
                        [3]int{n[0], n[1], dist+1},
                    )
                }
            }
        }

        return res
    }

    for rowIdx, row := range mat {
        for colIdx, _ := range row {
            cell := mat[rowIdx][colIdx]
            if cell == 0 {
                continue
            }

            mat[rowIdx][colIdx] = distance([2]int{rowIdx, colIdx})
        }
    }

    return mat
}