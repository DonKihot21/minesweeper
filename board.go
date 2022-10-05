package main

import (
	"fmt"
	"math/rand"
)

type board struct {
	x      int
	y      int
	mines  int
	fields [][]field
}

func (b *board) draw() {
	for i := 0; i < b.x; i++ {
		for j := 0; j < b.y; j++ {
			if b.fields[i][j].isOpened {
				fmt.Print(b.fields[i][j].touchingFileds)
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func newBoard(x int, y int, nMines int) *board {
	board := board{
		x:     x,
		y:     y,
		mines: nMines}

	board.fields = make([][]field, x*y)

	for i := range board.fields {
		board.fields[i] = make([]field, y)
	}

	for i := 0; i < board.x; i++ {
		for j := 0; j < board.y; j++ {
			board.fields[i][j] = field{i, j, "*", false, 0, false}
		}
	}
	board.initiliazeBombs(25)
	board.calculate()
	return &board
}

func (b *board) initiliazeBombs(nMines int) {

	for i := 0; i < nMines; i++ {
		x := rand.Intn(b.x)
		y := rand.Intn(b.y)
		b.fields[x][y].value = "&"
		b.fields[x][y].hasBomb = true
	}
}
func (b *board) calculate() {
	for i := 0; i < b.x; i++ {
		for j := 0; j < b.y; j++ {
			if !b.fields[i][j].hasBomb {
				if i-1 >= 0 && j-1 >= 0 && b.fields[i-1][j-1].hasBomb {
					b.fields[i][j].touchingFileds += 1
				}
				if (i+1) <= b.x-1 && j-1 >= 0 && b.fields[i+1][j-1].hasBomb {
					b.fields[i][j].touchingFileds += 1
				}
				if i+1 < b.x-1 && j+1 <= b.y-1 && b.fields[i+1][j+1].hasBomb {
					b.fields[i][j].touchingFileds += 1
				}
				if i-1 >= 0 && j+1 <= b.y-1 && b.fields[i-1][j+1].hasBomb {
					b.fields[i][j].touchingFileds += 1
				}
				if j-1 >= 0 && b.fields[i][j-1].hasBomb {
					b.fields[i][j].touchingFileds += 1
				}
				if j+1 <= b.y-1 && b.fields[i][j+1].hasBomb {
					b.fields[i][j].touchingFileds += 1
				}
				if i-1 >= 0 && b.fields[i-1][j].hasBomb {
					b.fields[i][j].touchingFileds += 1
				}
				if i+1 <= b.x-1 && b.fields[i+1][j].hasBomb {
					b.fields[i][j].touchingFileds += 1
				}
			}
		}
	}
}
