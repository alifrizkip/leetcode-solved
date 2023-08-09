import "math"

func minimumCardPickup(cards []int) int {
    var maxLength int = int(math.Pow(10, 5)) + 1
    var res int = maxLength

    cardFreq := map[int]int{}
    for i, val := range cards {
        if firstIdx, ok := cardFreq[val]; ok {
            diff := i - firstIdx

            if diff+1 < res {
                res = diff+1
            }
            cardFreq[val] = i
        } else {
            cardFreq[val] = i
        }
    }

    if res == maxLength {
        return -1
    }

    return res
}