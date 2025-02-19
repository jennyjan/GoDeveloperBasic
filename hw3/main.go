package main

import (
	"fmt"
	"strings"
)

func main() {
	var lineCount int
	var columnCount int

	fmt.Println("Введите размер шахматной доски (кол-во строк кол-во столбцов) через пробел:")
	_, err := fmt.Scanf("%d %d", &lineCount, &columnCount)
	if err != nil {
		fmt.Print(err)
		return
	}

	chessBoard(lineCount, columnCount)
}

func chessBoard(lineCount int, columnCount int) {
	var result strings.Builder
	for line := 0; line < lineCount; line++ {
		for column := 0; column < columnCount; column++ {
			if line%2 == column%2 {
				result.WriteString("#")
			} else {
				result.WriteString(" ")
			}
		}
		result.WriteString("\n")
	}

	fmt.Print(result.String())
}
