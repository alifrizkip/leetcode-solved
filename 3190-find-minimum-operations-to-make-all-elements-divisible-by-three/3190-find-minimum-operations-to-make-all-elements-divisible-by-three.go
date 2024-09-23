import (
"fmt"
)

func minimumOperations(nums []int) int {
    var res int
    
    for _, n := range nums {
        if (n-1) % 3 == 0 || (n+1) % 3 == 0 {
            res++
        }
    }
    
    return res
}