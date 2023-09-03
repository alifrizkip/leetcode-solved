func tictactoe(moves [][]int) string {
    // build grid
    grid := [][]byte{
        {'.', '.', '.'},
        {'.', '.', '.'},
        {'.', '.', '.'},
    }
    var turn byte = 'A'
    for _, t := range moves {
        grid[t[0]][t[1]] = turn

        // change turn
        if turn == 'A' {
            turn = 'B'
        } else {
            turn = 'A'
        }
    }

    // check winner horizontal & vertical
    for i := 0; i < 3; i++ {
        if p := grid[i][i]; p != '.' {
            if grid[i][0] == grid[i][1] && grid[i][0] == grid[i][2] ||
              grid[0][i] == grid[1][i] && grid[0][i] == grid[2][i] {
                return string(p)
            }
        }
    }

    // check diagonal
    if p := grid[1][1]; p != '.' {
        if p == grid[0][0] && p == grid[2][2] ||
          p == grid[0][2] && p == grid[2][0] {
            return string(p)
        }
    }

    if len(moves) != 9 {
        return "Pending"
    }

    return "Draw"
}