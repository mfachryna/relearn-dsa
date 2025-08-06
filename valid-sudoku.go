package relearndsa

func isValidSudoku(board [][]byte) bool {
    row := make([]map[byte]bool, 9)
    col := make([]map[byte]bool, 9)
    square := make([]map[byte]bool, 9)

    for i := 0; i < 9; i++{    
        row[i] = make(map[byte]bool)
        col[i] = make(map[byte]bool)
        square[i] = make(map[byte]bool)
    }
    for i := 0; i < 9; i++{   
        for j := 0; j < 9; j++{    
            temp := board[i][j]
            if temp == '.' {
                continue
            }

            if row[i][temp]{
                return false
            }
            row[i][temp] = true

            if col[j][temp]{
                return false
            }
            col[j][temp] = true

            squareVal := (i/3)*3 +(j/3)
            if square[squareVal][temp]{
                return false
            }
            square[squareVal][temp] = true
        } 
    }
    return true
}