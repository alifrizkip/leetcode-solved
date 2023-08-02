func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    var res int
    
    for _, val := range hours {
        if val >= target {
            res++
        }
    }
    
    return res
}