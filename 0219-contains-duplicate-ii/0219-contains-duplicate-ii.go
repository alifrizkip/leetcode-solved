func containsNearbyDuplicate(nums []int, k int) bool {
    hashMap := make(map[int]int)

    for idx, n := range nums {
        lastIdx, ok := hashMap[n]
        
        if ok && idx - lastIdx <= k {
            return true
        }
        
        hashMap[n] = idx
    }
    
    return false
}