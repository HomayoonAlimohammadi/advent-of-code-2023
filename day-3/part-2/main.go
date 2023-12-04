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

//go:embed test.txt
var testDataB []byte

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
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if isGear(engine[j][i]) {
				adjNums := findAdjacentNumbers(i, j, engine)
				// fmt.Printf("found gear at [%d, %d], adjacent numbers: %+v\n", i, j, adjNums)
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
	x = i - 1
	if y >= 0 && x >= 0 {
		numsS = append(numsS, extractRow(i, x, y, engine)...)
	} else if y >= 0 { // left side
		x = i
		numsS = append(numsS, extractRow(i, x, y, engine)...)
	}

	// same row
	y = j

	// |__ left
	x = i - 1
	if x >= 0 && isNumber(engine[y][x]) {
		numS := string(engine[y][x])
		for i2 := x - 1; i2 >= 0; i2-- {
			if isNumber(engine[y][i2]) {
				numS = string(engine[y][i2]) + numS
			} else {
				break
			}
		}
		numsS = append(numsS, numS)
	}

	// |__ right
	x = i + 1
	if x < len(engine[0]) && isNumber(engine[y][x]) {
		numS := string(engine[y][x])
		for i2 := x + 1; i2 < len(engine[0]); i2++ {
			if isNumber(engine[y][i2]) {
				numS = numS + string(engine[y][i2])
			} else {
				break
			}
		}
		numsS = append(numsS, numS)
	}

	// bottom row
	y = j + 1
	x = i - 1
	if y < len(engine) && x >= 0 {
		numsS = append(numsS, extractRow(i, x, y, engine)...)
	} else if y < len(engine) { // left side
		x = i
		numsS = append(numsS, extractRow(i, x, y, engine)...)
	}

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

func extractRow(orgX, x, y int, engine []string) []string {
	numsS := []string{}

	numS := ""
	end := x // end is where a number is finished and we see a period
	if isNumber(engine[y][x]) {
		numS = string(engine[y][x])
		for i2 := x - 1; i2 >= 0; i2-- {
			if isNumber(engine[y][i2]) {
				numS = string(engine[y][i2]) + numS
			} else {
				break
			}
		}

		for i2 := x + 1; i2 < len(engine[0]); i2++ {
			end = i2
			if isNumber(engine[y][i2]) {
				numS = numS + string(engine[y][i2])
			} else {
				break
			}
		}
	}
	if numS != "" {
		numsS = append(numsS, numS)
	}

	for end < len(engine[0]) && end <= orgX+1 {
		numS, end = getNext(end, y, engine)
		if numS != "" {
			numsS = append(numsS, numS)
		}
	}

	return numsS
}

func getNext(start, y int, engine []string) (string, int) {
	numS := ""
	end := start
	for x := start; x < len(engine[0]); x++ {
		end++
		if isNumber(engine[y][x]) {
			numS = numS + string(engine[y][x])
		} else {
			break
		}
	}

	return numS, end
}
