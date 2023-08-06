/*
IV -> 4
XXI -> 21
*/
func romanToInt(s string) int {
    var res int
    syms := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    var i int
    for i < len(s) {
        curr := syms[string(s[i])]

        var next int
        if i < len(s)-1 {
            next = syms[string(s[i+1])]
        }

        if curr < next {
            res += next - curr
            i++
        } else {
            res += curr
        }
        i++
    }

    return res
}