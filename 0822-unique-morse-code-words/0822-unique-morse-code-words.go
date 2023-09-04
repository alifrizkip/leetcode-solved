func uniqueMorseRepresentations(words []string) int {
    codes := [26]string{
        ".-","-...","-.-.","-..",".",
        "..-.","--.","....","..",".---",
        "-.-",".-..","--","-.","---",
        ".--.","--.-",".-.","...","-",
        "..-","...-",".--","-..-","-.--","--..",
    }

    collection := map[string]int{}
    for _, w := range words {
        var wCode string
        for _, c := range w {
            wCode += codes[c-'a']
        }
        collection[wCode]++
    }

    return len(collection)
}