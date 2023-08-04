import "fmt"
func reversePrefix(word string, ch byte) string {
    var idx int
    for i := 0; i < len(word); i++ {
        if word[i] == ch {
            idx = i
            break
        }
    }
    if idx == 0 {
        return word
    }

    var left, right = word[:idx+1], word[idx+1:]
    var leftByte = []byte(left)
    for i := 0; i < len(left); i++ {
        leftByte[i] = left[len(left)-1-i]
    }

    return string(leftByte) + right
}