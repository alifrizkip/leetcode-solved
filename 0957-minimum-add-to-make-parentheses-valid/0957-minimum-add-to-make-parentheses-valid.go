/*
((())

*/
func minAddToMakeValid(s string) int {
    st := []string{}
    for _, v := range s {
        c := string(v)

        if c == "(" {
            st = append(st, c)
        } else {
            if len(st) > 0 && st[len(st)-1] == "(" {
                st = st[:len(st)-1]
            } else {
                st = append(st, c)
            }
        }
    }

    return len(st)
}