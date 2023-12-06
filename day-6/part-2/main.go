package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed test.txt
var testB []byte

//go:embed input.txt
var inputB []byte

func main() {
	input := string(inputB)

	lines := strings.Split(input, "\n")
	timeS := strings.Split(lines[0], ":")
	distS := strings.Split(lines[1], ":")

	timeStr := strings.Join(strings.Split(timeS[1], " "), "")
	distStr := strings.Join(strings.Split(distS[1], " "), "")

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		panic(err)
	}

	dist, err := strconv.Atoi(distStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(getNumPossibles(time, dist))
}

func getNumPossibles(totalT, record int) int {
	num := 0
	for wait := 0; wait <= totalT; wait++ {
		speed := wait // speed goes up 1mm/ms for each ms of waiting
		dist := speed * (totalT - wait)
		if dist > record {
			num++
		}
	}

	return num
}
