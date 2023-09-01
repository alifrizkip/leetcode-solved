func diagonalSum(mat [][]int) int {
    var res int
    colL, colR := 0, len(mat)-1

    for i := 0; i < len(mat); i++ {
        if mat[i][colL+i] != 0 {
            res += mat[i][colL+i]
            mat[i][colL+i] = 0
        }

        if mat[i][colR-i] != 0 {
            res += mat[i][colR-i]
            mat[i][colR-i] = 0
        }
    }

    return res
}