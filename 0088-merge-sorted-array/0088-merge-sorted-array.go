func merge(nums1 []int, m int, nums2 []int, n int)  {
    // last index nums1
    last := m + n - 1

    // merge from back
    for m > 0 && n > 0 {
        if nums1[m - 1] > nums2[n - 1] {
            nums1[last] = nums1[m - 1]
            m -= 1
        } else {
            nums1[last] = nums2[n - 1]
            n -= 1
        }
        last -= 1
    }

    // fill nums1 with leftover nums2 element
    for n > 0 {
        nums1[last] = nums2[n - 1]
        n -= 1
        last -= 1
    }
}