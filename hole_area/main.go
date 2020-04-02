package main

import "fmt"

func dfs(board [][]byte, i int, j int, r int, c int, mark int) {
	if i < 0 || j < 0 || i >= r || j >= c {
		return
	}

	if board[i][j] != 'O' {
		return
	}
	board[i][j] = byte('0' + mark)
	dfs(board, i - 1, j, r, c, mark)
	dfs(board, i, j - 1, r, c, mark)
	dfs(board, i + 1, j, r, c, mark)
	dfs(board, i, j + 1, r, c, mark)
}

func solve(board [][]byte)  {
	var row int
	var column int

	row = len(board)
	if row == 0 {
		return
	}

	column = len(board[0])
	if column == 0 {
		return
	}

	dict := map[byte]bool{}
	var counter int = 1
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if board[i][j] == 'O' {
				dfs(board, i, j, row, column, counter)
				counter++
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if board[i][j] != 'X' {
				if i == 0 || j == 0 || i == row - 1 || j == column - 1 {
					dict[board[i][j]] = true
				}
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			ok, _ := dict[board[i][j]]
			if ok {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	areas := [][]byte{
		[]byte{'O', 'X', 'O', 'O', 'X', 'X'},
		[]byte{'O', 'X', 'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'O', 'X', 'O', 'O'},
		[]byte{'X', 'O', 'X', 'X', 'X', 'X'},
		[]byte{'O', 'O', 'X', 'O', 'X', 'X'},
		[]byte{'X', 'X', 'O', 'O', 'O', 'O'},
	}
	solve(areas)

	fmt.Println(areas)
}
