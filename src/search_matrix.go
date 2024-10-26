package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])
	if col == 0 {
		return false
	}

	//Search target row
	low := 0
	high := row
	var targetRow int
	for {
		numCandidates := high - low
		if numCandidates == 1 {
			targetRow = low
			break
		}
		candidateRow := numCandidates/2 + low
		if target < matrix[candidateRow][0] {
			high = candidateRow
		} else {
			low = candidateRow
		}
	}

	//Search target column
	low = 0
	high = col
	for {
		numCandidates := high - low
		if numCandidates == 1 {
			if matrix[targetRow][low] == target {
				return true
			} else {
				return false
			}
		}
		candidateCol := numCandidates/2 + low
		if target < matrix[targetRow][candidateCol] {
			high = candidateCol
		} else {
			low = candidateCol
		}
	}
}
