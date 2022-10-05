package main

import "fmt"

func main() {
	board := newBoard(5, 10, 5)
	board.draw()

	for {
		var x int = 0
		var y int = 0
		fmt.Scan(&x, &y)

		isOver := board.makeMove(x, y)
		if isOver {
			fmt.Println("BOMB, GAME OVER!")
			break
		}
		board.draw()
	}
}
