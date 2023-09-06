/*
{
    3: 0
    2: 1
}
*/
func twoSum(nums []int, target int) []int {
    hm := map[int]int{}
    res := []int{}

    for i, num := range nums {
        val, ok := hm[target-num]
        if ok {
            res = []int{val, i}
        } else {
            hm[num] = i
        }
    }

    return res
}