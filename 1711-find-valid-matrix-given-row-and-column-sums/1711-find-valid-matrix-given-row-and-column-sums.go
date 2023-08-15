/*
- build matrix
- iterate through every cell
- get min between row[i] & col[j]
- set min to cell
- substract row[i] & col[j] with min

      8,6,8  |       0,0,0
      | | |  |       | | |
5  -  0,0,0  | 0  -  5,0,0
7  -  0,0,0  | 0  -  3,4,0
10 -  0,0,0  | 0  -  0,2,8
*/

func restoreMatrix(rowSum []int, colSum []int) [][]int {
    rowL, colL := len(rowSum), len(colSum)
    res := make([][]int, rowL)

    for i := 0; i < rowL; i++ {
        arr := make([]int, colL)

        for j := 0; j < colL; j++ {
            // get min
            n := min(rowSum[i], colSum[j])

            // set min to cell
            arr[j] = n
            // substract row[i] & col[j] with min
            rowSum[i], colSum[j] = rowSum[i]-n, colSum[j]-n
        }
        res[i] = arr
    }

    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}