/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
import "math"

func guessNumber(n int) int {
    var ans, div int = n, n

    g := guess(ans)
    for g != 0 {
        // greater
        if g == -1 {
            div = int(math.Round(float64(div) / 2))
            ans = ans - div
        }

        // lower
        if g == 1 {
            div = int(math.Round(float64(div) / 2))
            ans = ans + div
        }

        g = guess(ans)
    }

    return ans
}