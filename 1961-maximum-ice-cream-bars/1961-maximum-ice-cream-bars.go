/*
7
0,1,2,3,4
0,2,1,1,1

[4,7,6,4,4,2,2,4,8,8]
2,4,6,7,8
2,4,1,1,2
4,16,6,7,16
  
*/
func maxIceCream(costs []int, coins int) int {
    // build counter
    counter := [100001]int{}
    for _, n := range costs {
        counter[n]++
    }

    var res int
    for i := 1; i < len(counter); i++ {
        c := counter[i]

        // iteratively deduct coins & adding res
        // according to counter idx(value) & value(occurence)
        for coins != 0 && c != 0 && i <= coins {
            res++
            c--
            coins -= i
        }
    }

    return res
}