func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    arr := append(nums1, nums2...)
    sort.Ints(arr)

    arrL := len(arr)
    if arrL % 2 == 1 {
        return float64(arr[arrL/2])
    }

    return float64(arr[arrL/2]+arr[arrL/2-1]) / float64(2)
}