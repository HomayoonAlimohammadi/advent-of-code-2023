package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

//go:embed input.txt
var inputDataB []byte

type number struct {
	value    string
	row      int
	startCol int
	endCol   int
}

func main() {
	input := string(inputDataB)
	splitted := strings.Split(input, "\n")
	n := len(splitted)
	m := len(splitted[0])

	engine := make([]string, n)
	for j := 0; j < n; j++ {
		engine[j] = splitted[j]
	}

	sum := 0
	var num *number
	for j := 0; j < m; j++ {
		num = nil // reset number on each line start
		for i := 0; i < m; i++ {
			if char := engine[j][i]; isNumber(char) {
				if num == nil {
					num = &number{row: j, startCol: i}
				}

				num.value = num.value + string(char)
			} else {
				if num == nil {
					continue // no number, just iterating symbols and periods
				}

				num.endCol = i - 1

				fmt.Printf("end of number: %+v\n", *num)

				if isAdjacentToSymbol(num, engine) {
					fmt.Println("is adjacent:", num.value)

					val, err := strconv.Atoi(num.value)
					if err != nil {
						logrus.WithError(err).Errorf("failed to cast `%s` to int\n", num.value)
					}

					sum = sum + val
				}

				num = nil // reset number
			}
		}

		// end of line
		if num == nil {
			continue // no number, just iterating symbols and periods
		}

		num.endCol = m - 1

		fmt.Printf("end of number: %+v\n", *num)

		if isAdjacentToSymbol(num, engine) {
			fmt.Println("is adjacent:", num.value)

			val, err := strconv.Atoi(num.value)
			if err != nil {
				logrus.WithError(err).Errorf("failed to cast `%s` to int\n", num.value)
			}

			sum = sum + val
		}

		num = nil // reset number
	}

	fmt.Println(sum)
}

func isNumber(b byte) bool {
	return 48 <= b && b <= 57
}

func isAdjacentToSymbol(num *number, engine []string) bool {
	var x, y int
	n := len(engine)    // num rows
	m := len(engine[0]) // num cols

	// top left
	x = num.startCol - 1
	y = num.row - 1
	if x >= 0 && y >= 0 && isSymbol(engine[y][x]) {
		return true
	}

	// top right
	x = num.endCol + 1
	y = num.row - 1
	if x <= m-1 && y >= 0 && isSymbol(engine[y][x]) {
		return true
	}

	// top
	y = num.row - 1
	if y >= 0 {
		for i := num.startCol; i <= num.endCol; i++ {
			if isSymbol(engine[y][i]) {
				return true
			}
		}
	}

	// bottom left
	x = num.startCol - 1
	y = num.row + 1
	if x >= 0 && y <= n-1 && isSymbol(engine[y][x]) {
		return true
	}

	// bottom right
	x = num.endCol + 1
	y = num.row + 1
	if x <= m-1 && y <= n-1 && isSymbol(engine[y][x]) {
		return true
	}

	// bottom
	y = num.row + 1
	if y <= n-1 {
		for i := num.startCol; i <= num.endCol; i++ {
			if isSymbol(engine[y][i]) {
				return true
			}
		}
	}

	// left
	x = num.startCol - 1
	y = num.row
	if x >= 0 && isSymbol(engine[y][x]) {
		return true
	}

	// right
	x = num.endCol + 1
	y = num.row
	if x <= m-1 && isSymbol(engine[y][x]) {
		return true
	}

	return false
}

func isSymbol(b byte) bool {
	return b != 46 && (b < 48 || b > 57) // not dot and not a number
}
