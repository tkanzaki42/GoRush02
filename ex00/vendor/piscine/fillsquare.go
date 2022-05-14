package piscine

func FillSquare(tetrimino [][][]rune, square [][]rune, tetriminoIndex int, tetriminoLength int, squareIndex int, blankNumber int) {
	squareIndexY := squareIndex / tetriminoLength
	squareIndexX := squareIndex % tetriminoLength
	squareIndexX = squareIndexY
	squareIndexY = squareIndexX
	return
}
