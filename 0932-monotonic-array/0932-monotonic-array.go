func isMonotonic(nums []int) bool {
    var dir string
    for i := 1; i < len(nums); i++ {
        d := nums[i] - nums[i-1]

        if dir == "" && d != 0 {
            if d > 0 {
                dir = "right"
            } else {
                dir = "left"
            }

            continue
        }

        if dir == "right" && d < 0 {
            return false
        }

        if dir == "left" && d > 0 {
            return false
        }
    }

    return true
}