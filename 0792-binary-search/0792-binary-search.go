/*
    Iterative
func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        m := l + (r-l)/2

        if target == nums[m] {
            return m
        } else if target > nums[m] {
            l = m+1
        } else if target < nums[m] {
            r = m-1
        }
    }

    return -1
}
*/

/*
    Recursive
*/
func search(nums []int, target int) int {
    return recursive(nums, target, 0, len(nums)-1)
}

func recursive(nums []int, target, l, r int) int {
    if l > r {
        return -1
    }

    m := l + (r-l)/2
    if target > nums[m] {
        return recursive(nums, target, m+1, r)
    }
    if target < nums[m] {
        return recursive(nums, target, l, m-1)
    }

    return m
}