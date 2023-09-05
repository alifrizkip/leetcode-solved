/*
BFS approach
*/
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    rL, cL := len(image), len(image[0])
    var n int = image[sr][sc]

    directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

    q := [][2]int{{sr, sc}}
    for len(q) > 0 {
        cell := q[0]
        q = q[1:]

        image[cell[0]][cell[1]] = color

        for _, dir := range directions {
            r, c := cell[0]+dir[0], cell[1]+dir[1]
            if r >= 0 && r < rL &&
              c >= 0 && c < cL {
                nDir := image[r][c]
                if nDir == n && nDir != color {
                    q = append(q, [2]int{r, c})
                }
            }
        }
    }

    return image
}