package piscine

import "ft"

func PrintSquare(square [][]rune) {
	for _, column := range square {
		for _, v := range column {
			if v == ' ' {
				ft.PrintRune(' ')
			} else {
				ft.PrintRune(v)
			}
		}
		ft.PrintRune('\n')
	}
}
