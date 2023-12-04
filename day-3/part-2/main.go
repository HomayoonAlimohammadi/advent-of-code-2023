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
	for j := 0; j < m; j++ {
		for i := 0; i < m; i++ {
			if isGear(engine[j][i]) {
				adjNums := findAdjacentNumbers(i, j, engine)
				if len(adjNums) == 2 {
					gearRatio := adjNums[0] * adjNums[1]
					sum = sum + gearRatio
				}
			}
		}
	}

	fmt.Println(sum)
}

func isNumber(b byte) bool {
	return 48 <= b && b <= 57
}

func isGear(b byte) bool {
	return string(b) == "*"
}

func findAdjacentNumbers(i, j int, engine []string) []int {
	numsS := []string{}

	var x, y int

	// top row
	y = j - 1

	// same row
	y = j

	// |__ left
	x = i - 1

	// |__ right
	x = i + 1

	// bottom row
	y = j + 1

	fmt.Println(x, y)

	nums := make([]int, len(numsS))
	for i := range numsS {
		num, err := strconv.Atoi(numsS[i])
		if err != nil {
			logrus.WithError(err).Errorf("failed to cast `%s` to int\n", numsS[i])
		}

		nums[i] = num
	}

	return nums
}
