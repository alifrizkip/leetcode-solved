func flipAndInvertImage(image [][]int) [][]int {
    for rowIdx, row := range image {
        flippedRow := []int{}
        
        for i := 0; i < len(row); i++  {
            newP := 0
            
            if row[len(row)-1-i] == 0 {
                newP = 1
            }
            
            flippedRow = append(flippedRow, newP)
        }
        
        image[rowIdx] = flippedRow
    }
    
    return image
}