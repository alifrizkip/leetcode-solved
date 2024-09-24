func shortestToChar(s string, c byte) []int {
    res := make([]int, 0)
    cPos := make([]int, 0)
    
    for i, char := range s {
        if byte(char) == c {
            cPos = append(cPos, i)
        }
    }
    
    fmt.Println(cPos)
    
    var cI = 0
    for i := 0; i < len(s); i++ {
        far := calcFar(cPos[cI], i)
        nextFar := far
        // nextFar := calcFar(cPos[cI+1], i)
        
        if cI < len(cPos)-1 {
            nextFar = calcFar(cPos[cI+1], i)
        }
        
        if abs(far) < abs(nextFar) {
            res = append(res, abs(far))
        } else {
            if cI < len(cPos) - 1 {
                cI++
            }
            
            res = append(res, abs(nextFar))
        }
    }
    
    
    return res
}

func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    
    return n
}

func calcFar(a, b int) int {
    if a > b {
        return a - b
    }
    
    return b-a
}