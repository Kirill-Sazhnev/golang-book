package main

import (
	"fmt"
	"strings"
)

func drawBoard(board [8][8]rune) {
	for i := range board {

		for j := range board[i] {

			if board[i][j] == 0 {
				fmt.Print("_ ")
			} else {
				fmt.Printf("%c ", board[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	var board [8][8]rune
	black := "rkbkqbkr"
	white := strings.ToUpper(black)

	for i := range black {
		board[0][i] = rune(black[i])
		board[7][i] = rune(white[i])
		board[1][i] = 'p'
		board[6][i] = 'P'
	}

	board[1][1] = 0
	board[3][1] = 'p'

	drawBoard(board)
}

/*
Фигуры — это король (king ♔, ♚), ферзь (queen ♕, ♛), слон (bishop ♗, ♝),
конь (knight ♘, ♞), ладья (rook ♖, ♜), а пешки (pawns ♙, ♟) — это пешки

Допишите Листинг 8 для отображения всех шахматных фигур на их стартовых
позициях, используя символы kqrbnp для черных фигур в верхней части доски,
а также символы в верхнем регистре KQRBNP для белых фигур в нижней части
доски;

Напишите функцию для отображения доски;
Вместо строк, используйте [8][8]rune для доски. Помните, что литералы rune
должны быть окружены одинарными кавычками и могут выводиться на экран через
специальный символ %c.
*/
