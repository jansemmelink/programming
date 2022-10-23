/* Consider a black and white digitized image of a maze - white pixels represent open areas and
black spaces (X) are walls. There are two special white pixels: one is designated the entrance
and the other is the exit. The goal in this problem is to find a way of getting from the
entrance to the exit. Given a 2D array of black and white entries representing a maze with
designated entrance and exit points, find a path from the entrance to the exit, if one exists.



___________
| | | |X|E|
-----------
| |X| | | |
-----------
| | |X| |X|
-----------
|X|S| | | |
-----------

X - Wall
S - Start
E - End

*/

package main

import "fmt"

func main() {
	board := []string{
		"PPPXE",
		"PXPPP",
		"PSX X",
	}
	path := solve(
		board,
		[]Point{{2, 1}},
		Point{0, 4})

	fmt.Printf("path: %+v\n", path)
	//path: [{row:2 col:1} {row:2 col:0} {row:1 col:0} {row:0 col:0} {row:0 col:1} {row:0 col:2} {row:1 col:2} {row:1 col:3} {row:1 col:4} {row:0 col:4}]

}

type Point struct {
	row int
	col int
}

func solve(board []string, path []Point, end Point) []Point {
	nrRows := len(board)
	nrCols := len(board[0])
	last := path[len(path)-1]
	for dRow := -1; dRow <= 1; dRow += 2 {
		next := Point{last.row + dRow, last.col}
		if next.row < 0 || next.row >= nrRows {
			continue //exceeded boundary
		}
		solution := rSolve(board, path, next, end)
		if solution != nil {
			return solution
		}
	}

	for dCol := -1; dCol <= 1; dCol += 2 {
		next := Point{last.row, last.col + dCol}
		if next.col < 0 || next.col >= nrCols {
			continue //exceeded boundary
		}
		solution := rSolve(board, path, next, end)
		if solution != nil {
			return solution
		}
	}

	return nil
}

func rSolve(board []string, path []Point, next Point, end Point) []Point {
	//check for wall...
	if board[next.row][next.col] == 'X' {
		return nil
	}

	//check for back track
	for _, p := range path {
		if p == next {
			return nil
		}
	}

	if next == end {
		return append(path, next) //this is a solution
	}
	return solve(board, append(path, next), end)
}
