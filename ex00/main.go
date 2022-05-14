package main

import "piscine"

func main() {
	tetrimminoLength := 4
	tetrimino := make([][][]rune, tetrimminoLength)
	square := [][]rune{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'}}
	tetrimino[0] = [][]rune{
		{'#', '#', '#', '#'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'}}
	tetrimino[1] = [][]rune{
		{'#', '#', '#', '#'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'}}
	tetrimino[2] = [][]rune{
		{'#', '#', '#', '#'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'}}
	tetrimino[3] = [][]rune{
		{'#', '#', '#', '#'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'}}
	tetriminoIndex := 0
	blankNumber := 0
	squareIndex := 0
	success := piscine.FillSquare(tetrimino, square, tetriminoIndex, tetrimminoLength, squareIndex, blankNumber)
	for success == false {
		// 正方形を拡張
		//expand(square)
		// 空白文字を増やす
		//increase(blankNumber)
		success = piscine.FillSquare(tetrimino, square, tetriminoIndex, tetrimminoLength, squareIndex, blankNumber)
	}
	piscine.PrintSquare(square)
}
