func shortestToChar(s string, c byte) []int {
    res := make([]int, 0)
    cPos := make([]int, 0)
    
    for i, char := range s {
        if byte(char) == c {
            cPos = append(cPos, i)
        }
    }
    
    var cI = 0
    for i := 0; i < len(s); i++ {
        far := calcFar(cPos[cI], i)
        nextFar := far
        
        if cI < len(cPos)-1 {
            nextFar = calcFar(cPos[cI+1], i)
        }
        
        if far < nextFar {
            res = append(res, far)
        } else {
            if cI < len(cPos) - 1 {
                cI++
            }
            
            res = append(res, nextFar)
        }
    }

    return res
}

func calcFar(a, b int) int {
    if a > b {
        return a - b
    }
    
    return b-a
}