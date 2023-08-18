func convertTime(current string, correct string) int {
    diffInMin := minutesDiff(current, correct)

    var res int
    for diffInMin != 0 {
        if diffInMin - 60 >= 0 {
            diffInMin -= 60
        } else if diffInMin - 15 >= 0 {
            diffInMin -= 15
        } else if diffInMin - 5 >= 0 {
            diffInMin -= 5
        } else {
            diffInMin -= 1
        }
        res++
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

    return (endH - startH) * 60 + (endM - startM)
}