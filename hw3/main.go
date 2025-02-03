package main

import "fmt"

func main() {
	chessBoard(8, 8)
}

func chessBoard(lineCount int, columnCount int) {
	var result string
	for line := 0; line < lineCount; line++ {
		for column := 0; column < columnCount; column++ {
			if line%2 == column%2 {
				result += "#"
			} else {
				result += " "
			}
		}
		result += "\n"
	}

	fmt.Print(result)
}
