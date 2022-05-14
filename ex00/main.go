package main

import "piscine"

func main() {
	tetrimminoLength := 4
	width := 4
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
	piscine.FillSquare(tetrimino, square, tetrimminoLength, width, 0, 0)
	piscine.PrintSquare(square)
}
