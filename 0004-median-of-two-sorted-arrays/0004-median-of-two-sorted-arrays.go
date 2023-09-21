func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    arr := make([]int, 0, len(nums1)+len(nums2))

    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            arr = append(arr, nums1[i])
            i++
        } else {
            arr = append(arr, nums2[j])
            j++
        }
    }
    if i < len(nums1) {
        arr = append(arr, nums1[i:]...)
    }
    if j < len(nums2) {
        arr = append(arr, nums2[j:]...)
    }

    arrL := len(arr)
    if arrL % 2 == 1 {
        return float64(arr[arrL/2])
    }

    return float64(arr[arrL/2]+arr[arrL/2-1]) / float64(2)
}