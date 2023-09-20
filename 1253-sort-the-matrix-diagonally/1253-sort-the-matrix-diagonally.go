func diagonalSort(mat [][]int) [][]int {
    rowL, colL := len(mat), len(mat[0])
    
    // col
    for i := 0; i < colL; i++ {
        arr := []int{}

        // visit diagonally & add to arr
        r, c := 0, i
        for r < rowL && c < colL {
            arr = append(arr, mat[r][c])

            r++
            c++
        }
        sort.Ints(arr)

        // replace sorted arr to mat diagonally
        idx, r, c := 0, 0, i
        for r < rowL && c < colL {
            mat[r][c] = arr[idx]

            idx++
            r++
            c++
        }
    }

    // row
    for i := 1; i < rowL; i++ {
        arr := []int{}

        // visit diagonally & add to arr
        r, c := i, 0
        for r < rowL && c < colL {
            arr = append(arr, mat[r][c])

            r++
            c++
        }
        sort.Ints(arr)

        // replace sorted arr to mat diagonally
        idx, r, c := 0, i, 0
        for r < rowL && c < colL {
            mat[r][c] = arr[idx]

            idx++
            r++
            c++
        }
    }

    return mat
}