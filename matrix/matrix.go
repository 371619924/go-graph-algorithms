package matrix

// If needed, you may define helper functions here.

// AreAdjacent returns true iff a and b are adjacent in lst.
func AreAdjacent(a, b int, lst []int) bool {
	for i := 0; i < len(lst)-1; i++ {
		if (lst[i] == a && lst[i+1] == b) || (lst[i] == b && lst[i+1] == a) {
			return true
		}
	}
	return false
}

// Transpose returns the transpose of the 2D matrix mat.
func Transpose(mat [][]int) [][]int {
	if mat == nil {
		return nil
	}
	if len(mat) == 0 {
		return [][]int{}
	}
	if len(mat[0]) == 0 {
		return [][]int{}
	}

	rows := len(mat)
	cols := len(mat[0])

	result := make([][]int, cols)
	for j := 0; j < cols; j++ {
		result[j] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = mat[i][j]
		}
	}

	return result
}
// AreNeighbors returns true iff a and b are neighbors in the 2D matrix mat.
func AreNeighbors(mat [][]int, a, b int) bool {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
		
			if j+1 < len(mat[i]) {
				x := mat[i][j]
				y := mat[i][j+1]
				if (x == a && y == b) || (x == b && y == a) {
					return true
				}
			}
			if i+1 < len(mat) && j < len(mat[i+1]) {
				x := mat[i][j]
				y := mat[i+1][j]
				if (x == a && y == b) || (x == b && y == a) {
					return true
				}
			}
		}
	}
	return false
}


