/*
[4,4,4,11,4,4,11,4,4,4,4,4]
 ^                       ^
*/
func maxDistance(colors []int) int {
    var ltr, rtl int
    var l, r int = 1, len(colors) - 2
    for l < len(colors) || r >= 0 {
        // check l to r
        if colors[0] != colors[l] {
            ltr = l - 0
        }
        l++


        // check r to l
        if colors[len(colors)-1] != colors[r] {
            rtl = (len(colors) - 1) - r
        }
        r--
    }

    if ltr > rtl {
        return ltr
    }

    return rtl
}