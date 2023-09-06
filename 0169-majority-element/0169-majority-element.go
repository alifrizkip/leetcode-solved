/*
sort
majority element always in the middle
*/
func majorityElementV1(nums []int) int {
    sort.Ints(nums)

    return nums[len(nums)/2]
}

/*
use hash table as element counter
*/
func majorityElementV2(nums []int) int {
    m := map[int]int{}
    for _, n := range nums {
        m[n]++
    }

    res, mid := 0, len(nums)/2
    for k, v := range m {
        if v > mid {
            res = k
            break
        }
    }

    return res
}

// moore voting algorithm
func majorityElement(nums []int) int {
    count, candidate := 0, nums[0]
    for _, n := range nums{
        if count == 0 {
            candidate = n
        }

        if candidate == n {
            count++
        } else {
            count--
        }
    }
    return candidate
}