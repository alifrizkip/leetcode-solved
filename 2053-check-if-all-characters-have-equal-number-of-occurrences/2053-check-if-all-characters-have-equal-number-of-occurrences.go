func areOccurrencesEqual(s string) bool {
    charOcc := [26]int{}

    for _, c := range s {
        charOcc[c-'a']++
    }


    var prev int
    for _, num := range charOcc {
        if num == 0 {
            continue
        }

        if prev == 0 {
            prev = num
        }

        if prev != num {
            return false
        }

        prev = num
    }

    return true
}