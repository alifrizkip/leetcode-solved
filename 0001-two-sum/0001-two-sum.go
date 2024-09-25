func twoSum(nums []int, target int) []int {
    hash := map[int]int{}
    res := []int{}
    
    for i, n := range nums {
        if val, ok := hash[target - n]; ok {
            res = []int{i, val}
        } else {
            hash[n] = i
        }
    }
    
    return res
}