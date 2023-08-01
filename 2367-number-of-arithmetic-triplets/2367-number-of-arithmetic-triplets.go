func arithmeticTriplets(nums []int, diff int) int {
    var ans int

    numsHashMap := map[int]bool{}
    for _, n := range(nums) {
        numsHashMap[n] = true
    }

    for i := 0; i < len(nums); i++ {
        kNum := nums[i]
        jExist := numsHashMap[kNum - diff]
        iExist := numsHashMap[kNum - (diff * 2)]

        if jExist && iExist {
            ans++
        }
    }

    return ans
}