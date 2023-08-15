func minAddToMakeValid(s string) int {
    st := []string{}
    for _, v := range s {
        c := string(v)

        // (
        // push to stack
        if c == "(" {
            st = append(st, c)
            continue
        }

        // )
        // if 'stack.last==(' then 'stack.pop'
        // else push ) to stack
        if c == ")" {
            if len(st) > 0 && st[len(st)-1] == "(" {
                st = st[:len(st)-1]
            } else {
                st = append(st, c)
            }
        }
    }

    return len(st)
}