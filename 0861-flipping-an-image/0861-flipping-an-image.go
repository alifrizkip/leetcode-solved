func flipAndInvertImage(image [][]int) [][]int {
    for idx, row := range image {
        tempRow := make([]int, 0)

        for _, pixel := range row {
            newP := 0
            if pixel == 0 {
                newP = 1
            }
            
            tempRow = append([]int{newP}, tempRow...)
        }
        
        image[idx] = tempRow
    }
    
    return image
}