func convertTime(current string, correct string) int {
    mins := []int{60, 15, 5, 1}

    diffInMin := minutesDiff(current, correct)
    fmt.Println("diffInMin:", diffInMin)
    
    var res int
    for diffInMin != 0 {
        if diffInMin - mins[0] >= 0 {
            diffInMin -= mins[0]
            res++
            continue
        } else if diffInMin - mins[1] >= 0 {
            diffInMin -= mins[1]
            res++
            continue
        } else if diffInMin - mins[2] >= 0 {
            diffInMin -= mins[2]
            res++
            continue
        } else {
            diffInMin -= mins[3]
            res++
            continue
        }
    }

    return res
}

func minutesDiff(start, end string) int {
    startArr := strings.Split(start, ":")
    endArr := strings.Split(end, ":")

    startH, _ := strconv.Atoi(startArr[0])
    startM, _ := strconv.Atoi(startArr[1])

    endH, _ := strconv.Atoi(endArr[0])
    endM, _ := strconv.Atoi(endArr[1])

    diffH := endH - startH
    diffM := endM - startM

    res := diffH * 60 + diffM
    // if diffM < 0 {
    //     res -= diffM
    // } else {
    //     res += diffM
    // }

    return res
}