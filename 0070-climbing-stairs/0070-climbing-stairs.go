func climbStairs(n int) int {
    var res, last, secondLast int
    
    for i := 1; i <= n; i++ {
        if i == 1 {
            res = 1
        } else if i == 2 {
            res = 2
        } else {
            res = secondLast + last
        }

        secondLast = last
        last = res
    }

    return res
}