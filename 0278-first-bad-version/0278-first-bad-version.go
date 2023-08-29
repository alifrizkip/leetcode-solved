/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

/*
1 2 3 4 5
F F F T T
      m
    l
        r
--
1 2
F T
m
l
  r
*/
func firstBadVersion(n int) int {
    const MAX int = 2147483647

    res, l, r := 1, 1, n
    for l <= r {
        mid := (r - l) / 2

        if isBadVersion(l+mid) {
            res = l+mid
            r = l+mid - 1
        } else {
            l = l+mid + 1
        }
    }

    return res
}

func isOdd(n int) bool {
    return (n % 2) == 1
}