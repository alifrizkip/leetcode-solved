import "strings"

func reverseWords(s string) string {
    sArr := strings.Split(s, " ")
    
    for wIdx, w := range sArr {
        tempW := []rune{}

        for _, char := range w {
            tempW = append([]rune{char}, tempW...)            
        }
        
        sArr[wIdx] = string(tempW)
    }
    
    return strings.Join(sArr, " ")
}