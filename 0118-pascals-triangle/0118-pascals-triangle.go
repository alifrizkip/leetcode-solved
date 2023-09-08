/*
[
    [1],
    [1,1],
    [1,2,1],
    [1,3,3,1],
    [1,4,6,4,1]
]
*/

func generate(numRows int) [][]int {
    res := [][]int{
        []int{1},
    }

    for i := 1; i < numRows; i++ {
        items := []int{}
        for j := 0; j <= i; j++ {
            
            if j == 0 || j == i {
                items = append(items, 1)
            } else {
                up := res[i - 1][j]
                behind := res[i - 1][j - 1]
                item := up + behind
                items = append(items, item)
            }
        }

        res = append(res, items)
    }

    return res
}