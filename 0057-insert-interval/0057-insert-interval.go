func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {
        return [][]int{newInterval}
    }

    res := [][]int{}
    anyOverlap := false
    for i := 0; i < len(intervals); i++ {
        if isOverlap(intervals[i], newInterval) {
            newItem := mergeOne(intervals[i], newInterval)
            anyOverlap = true

            if len(res) > 0 && isOverlap(res[len(res)-1], newItem) {
                replaceItem := mergeOne(res[len(res)-1], newItem)
                res[len(res)-1] = replaceItem
            } else {
                res = append(res, newItem)
            }
        } else {
            res = append(res, intervals[i])
        }
    }
    if !anyOverlap {
        res = append(res, newInterval)
        sort.Slice(res, func(a, b int) bool {
            return res[a][0] < res[b][0]
        })
    }

    return res
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
    start, end := a[0], b[1]

    if b[0] < start {
        start = b[0]
    }

    if a[1] > end {
        end = a[1]
    }

    return []int{start, end}
}