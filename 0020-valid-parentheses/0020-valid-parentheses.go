func isValid(s string) bool {
    parentheses := map[byte]byte{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    stack := []byte{}

    for _, c := range s {
        matchOpen, isClose := parentheses[byte(c)]

        if !isClose {
            stack = append(stack, byte(c))
            continue
        }

        if len(stack) == 0 {
            return false
        }

        lastOpen := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if matchOpen != lastOpen {
            return false
        }
    }

    return len(stack) == 0
}