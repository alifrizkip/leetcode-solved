/*
2 0 2 1 1 0
0 0 2 1 1 2
0 0 1 1 2 2
0 0 1 1 2 2
        m
      l
      r
*/
func sortColors(nums []int)  {
    l, m, r := 0, 0, len(nums)-1
    
    for m <= r {
      if nums[m] == 0 {
        nums[m], nums[l] = nums[l], nums[m]
        l++
        m++
      } else if nums[m] == 2 {
        nums[m], nums[r] = nums[r], nums[m]
        r--
      } else {
        m++
      }
    }
}