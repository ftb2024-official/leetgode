/*
Determine if `9x9` Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
	- each row must contain the digits `1-9` without repetition
	- each column must contain the digits `1-9` without repetition
	- each of the `3x3` sub-boxes of the grid must contain the digits `1-9` without repetition

Note:
	- a Sudoku board (partially filled) could be valid but is not necessarily solvable
	- only the filled cells need to be validated according to the mentioned rules
*/

package neetcode

func IsValidSudoku(board [][]byte) bool {
	rows := make(map[byte]struct{}, len(board))
	cols := make(map[byte]struct{}, len(board))
	// squares := make(map[byte]struct{}, len(board))

	for r := range 9 {
		for c := range 9 {
			if board[r][c] == '.' {
				continue
			}

			if _, exists := rows[byte(r)]; exists {
				return false
			}

			if _, exists := cols[byte(c)]; exists {
				return false
			}
		}
	}
	return true
}
