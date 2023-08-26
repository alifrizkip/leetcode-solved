func merge(intervals [][]int) [][]int {
    // sort intervals
    sort.Slice(intervals, func(a, b int) bool {
        return (intervals[a][0] < intervals[b][0]) ||
            ( (intervals[a][0] == intervals[b][0]) &&
                (intervals[a][1] < intervals[b][1]) )
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

func isOverlap(a, b []int) bool {
    if a[1] < b[1] {
        return a[1] >= b[0]
    }

    return b[1] >= a[0]
}

func mergeOne(a, b []int) []int {
    start, end := a[0], b[1]

    if b[0] < start {
        start = b[0]
    }

    if a[1] > end {
        end = a[1]
    }

    return []int{start, end}
}