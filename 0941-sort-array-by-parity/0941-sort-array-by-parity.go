func sortArrayByParity(nums []int) []int {
    var ans = make([]int,0)

    for l := 0; l <= len(nums) / 2; l++ {
        front := nums[l]
        if front % 2 == 0 {
            ans = append([]int{front}, ans...)
        } else {
            ans = append(ans, front)
        }

        // right not reach center
        if len(nums)-1-l > len(nums) / 2 {
            rear := nums[len(nums)-1-l]
            if rear % 2 == 0 {
                ans = append([]int{rear}, ans...)
            } else {
                ans = append(ans, rear)
            }
        }
    }

    return ans
}