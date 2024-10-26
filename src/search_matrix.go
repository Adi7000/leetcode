package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	targetRow := findTargetRowInMatrix(&matrix, target)
	return isTargetInRow(&(matrix[targetRow]), target)
}

func findTargetRowInMatrix(matrix *[][]int, target int) int {
	low := 0
	high := len(*matrix)
	for {
		numCandidates := high - low
		if numCandidates == 1 {
			return low
		}
		candidateRow := numCandidates/2 + low
		if target < (*matrix)[candidateRow][0] {
			high = candidateRow
		} else {
			low = candidateRow
		}
	}
}

func isTargetInRow(row *[]int, target int) bool {
	low := 0
	high := len(*row)
	for {
		numCandidates := high - low
		if numCandidates == 1 {
			if (*row)[low] == target {
				return true
			} else {
				return false
			}
		}
		candidateCol := numCandidates/2 + low
		if target < (*row)[candidateCol] {
			high = candidateCol
		} else {
			low = candidateCol
		}
	}
}
