func lemonadeChange(bills []int) bool {
    wallet := map[int]int{}
    for _, v := range bills {
        wallet[v]++
        if v == 5 {
            continue
        }

        diff := v - 5
        ten, _ := wallet[10]
        five, _ := wallet[5]

        if diff == 15 {
            if ten >= 1 && five >= 1 {
                wallet[10]--
                wallet[5]--
            } else if five >= 3 {
                wallet[5] -= 3
            } else {
                return false
            }
        } else {
            if five >= 1 {
                wallet[5]--
            } else {
                return false
            }
        }
    }

    return true
}