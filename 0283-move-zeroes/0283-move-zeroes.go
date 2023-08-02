func moveZeroes(nums []int)  {
    var zeroOcc int

    for i := 0; i < len(nums); i++ {
        for nums[i] == 0 && i < len(nums)-1-zeroOcc {
            zeroOcc++

            nums = append(
                nums[0:i],
                append(nums[i+1:], 0)...,
            )
        }
    }
}