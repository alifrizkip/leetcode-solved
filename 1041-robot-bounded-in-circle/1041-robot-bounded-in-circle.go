func isRobotBounded(instructions string) bool {
    var x, y int
    var addX, addY int = 0, 1
    
    for _, c := range instructions {
        if c == 'G' {
            x += addX
            y += addY
        }
        
        if c == 'R' {
            // UP
            if addX == 0 && addY == 1 {
                addX = 1
                addY = 0
            // RIGHT
            } else if addX == 1 && addY == 0 {
                addX = 0
                addY = -1
            // DOWN
            } else if addX == 0 && addY == -1 {
                addX = -1
                addY = 0
            // LEFT
            } else if addX == -1 && addY == 0 {
                addX = 0
                addY = 1
            }
        }
        
        if c == 'L' {
            // UP
            if addX == 0 && addY == 1 {
                addX = -1
                addY = 0
            // LEFT
            } else if addX == -1 && addY == 0 {
                addX = 0
                addY = -1
            // DOWN
            } else if addX == 0 && addY == -1 {
                addX = 1
                addY = 0
            // RIGHT
            } else if addX == 1 && addY == 0 {
                addX = 0
                addY = 1
            }
        }
    }
    
    if x == 0 && y == 0 {
        return true
    }
    
    if addX == 0 && addY == 1 {
        return false
    }
    
    return true
}