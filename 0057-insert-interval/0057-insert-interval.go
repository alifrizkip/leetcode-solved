func insert(intervals [][]int, newInterval []int) [][]int {
    intervals = append(intervals, newInterval)
    
    // sort intervals
    sort.Slice(intervals, func(a, b int) bool {
        return intervals[a][0] < intervals[b][0]
    })

    var left, right [][]int = intervals[0:1], intervals[1:]
    
    // merge right to left    
    for len(left) > 0 && len(right) > 0 {
        // left[last], right[first]
        leftItem, rightItem := left[len(left)-1], right[0]

        if isOverlap(leftItem, rightItem) {
            newItem := mergeOne(leftItem, rightItem)
            
            // len(left) == 2 && newItem == right[first]
            if (newItem[0] == rightItem[0] &&
                newItem[1] == rightItem[1]) && len(left) > 1 {
                left = left[:len(left)-1] // remove left[last]
            } else {
                left[len(left)-1] = newItem
                right = right[1:] // remove right[first]
            }
        } else {
            left = append(left, rightItem)
            right = right[1:]
        }
    }

    return left
}

/*
1,3 | 2,5
4,8 | 3,5
*/
func isOverlap(a, b []int) bool {
    if a[1] < b[1] {
        return a[1] >= b[0]
    }

    return b[1] >= a[0]
}

func mergeOne(a, b []int) []int {
    start, end := a[0], a[1]

    if b[0] < start {
        start = b[0]
    }

    if b[1] > end {
        end = b[1]
    }

    return []int{start, end}
}