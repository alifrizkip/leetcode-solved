func checkIfPangram(sentence string) bool {
    var arr = [26]int{}

    for _, val := range sentence {
        arr[val - 'a'] += 1
    }

    for _, n := range arr {
        if n < 1 {
            return false
        }
    }
    
    return true
}