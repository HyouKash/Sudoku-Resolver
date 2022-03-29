package main

import (
	"os"

	"github.com/01-edu/z01"
)

var Tested = [][9][9]int{}

func main() {
	tab := [9][9]int{}
	if len(os.Args) == 10 {
		for i := 1; i < len(os.Args); i++ {
			if len(os.Args[i]) == 9 {
				for j := 0; j < len(os.Args[i]); j++ {
					if rune(os.Args[i][j]) > 47 && rune(os.Args[i][j]) < 58 {
						tab[i-1][j] = int(os.Args[i][j] - 48)
					} else if rune(os.Args[i][j]) != 46 {
						Error()
					}
				}
			} else {
				Error()
			}
		}
		Beat(tab, 0, 0)
		if len(Tested) != 1 {
			Error()
		} else {
			for t := 0; t < len(Tested[0]); t++ {
				for s := 0; s < len(Tested[0][t]); s++ {
					if !Testeur(Tested[0], t, s, Tested[0][t][s]) {
						Error()
					}
				}
			}
			for i := 0; i < len(Tested[0]); i++ {
				for j := 0; j < len(Tested[0][i]); j++ {
					z01.PrintRune(rune(Tested[0][i][j] + 48))
					if j != len(Tested[0][i])-1 {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}
		}
	} else {
		Error()
	}
}

func Beat(magrows [9][9]int, y int, x int) {
	if y > 8 {
		if len(Tested) < 1 {
			Tested = append(Tested, magrows)
		} else {
			Error()
		}
	} else {
		if magrows[y][x] == 0 {
			for i := 1; i <= 9; i++ {
				valide := Testeur(magrows, y, x, i)
				if valide {
					magrows[y][x] = i
					if x == 8 {
						Beat(magrows, y+1, 0)
					} else {
						Beat(magrows, y, x+1)
					}
				}
			}
		} else {
			if x == 8 {
				Beat(magrows, y+1, 0)
			} else {
				Beat(magrows, y, x+1)
			}
		}
	}
}

func Error() {
	str := "Error\n"
	for i := 0; i < len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}
	os.Exit(0)
}

func Testeur(magrows [9][9]int, y int, x int, i int) bool {
	for o := 0; o < 9; o++ {
		if i == magrows[o][x] && o != y {
			return false
		}
	}
	for a := 0; a < 9; a++ {
		if i == magrows[y][a] && a != x {
			return false
		}
	}
	for o := (y / 3) * 3; o < (y/3)*3+3; o++ {
		for a := (x / 3) * 3; a < (x/3)*3+3; a++ {
			if magrows[o][a] == i && (o != y || a != x) {
				return false
			}
		}
	}
	return true
}
