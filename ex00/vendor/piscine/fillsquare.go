package piscine

func FillSquare(tetrimino [][][]rune, square [][]rune, tetriminoIndex int, tetriminoLength int, squareIndex int, blankNumber int) bool {
	squareIndexY := squareIndex / tetriminoLength
	squareIndexX := squareIndex % tetriminoLength
	squareIndexX = squareIndexY
	squareIndexY = squareIndexX
	square[0][0] = '#'
	return true
}
