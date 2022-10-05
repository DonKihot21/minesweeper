package main

func (b *board) makeMove(x int, y int) bool {
	targetField := b.fields[x][y]
	gameOver := b.open(targetField)
	if gameOver {
		return true
	}
	return false
}

func (b *board) isBomb(x int, y int) bool {
	if b.fields[x][y].hasBomb {
		return true
	}
	return false
}

func (board *board) open(f field) bool {
	if f.hasBomb {
		return true
	}
	if !f.hasBomb && f.touchingFileds == 0 && !f.isOpened {
		board.fields[f.x][f.y].isOpened = true
		board.openSimilar(f.x, f.y)
	} else if !f.hasBomb {
		board.fields[f.x][f.y].isOpened = true
	}
	return false
}

func (b *board) openSimilar(x int, y int) {
	if x-1 >= 0 && y-1 >= 0 && !b.fields[x-1][y-1].hasBomb {
		b.open(b.fields[x-1][y-1])
	}
	if x+1 <= b.x && y+1 <= b.y && !b.fields[x+1][y+1].hasBomb {
		b.open(b.fields[x+1][y+1])
	}
	if x-1 >= 0 && y+1 <= b.y && !b.fields[x-1][y+1].hasBomb {
		b.open(b.fields[x-1][y+1])
	}
	if x+1 <= b.x && y-1 >= 0 && !b.fields[x+1][y-1].hasBomb {
		b.open(b.fields[x+1][y-1])
	}
	if x-1 >= 0 && !b.fields[x-1][y].hasBomb {
		b.open(b.fields[x-1][y])
	}
	if x+1 <= b.x && !b.fields[x+1][y].hasBomb {
		b.open(b.fields[x+1][y])
	}
	if y-1 >= 0 && !b.fields[x][y-1].hasBomb {
		b.open(b.fields[x][y-1])
	}
	if y+1 <= b.y && !b.fields[x][y+1].hasBomb {
		b.open(b.fields[x][y+1])
	}
}
