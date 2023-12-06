package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

//go:embed test.txt
var testB []byte

//go:embed input.txt
var inputB []byte

func main() {
	input := string(inputB)

	lines := strings.Split(input, "\n")
	timesS := strings.Split(lines[0], ":")
	distsS := strings.Split(lines[1], ":")

	var times, dists []int
	for _, tS := range strings.Split(timesS[1], " ") {
		if len(tS) == 0 {
			continue
		}

		t, err := strconv.Atoi(strings.TrimSpace(tS))
		if err != nil {
			logrus.WithError(err).Error("failed to convert time string to int")
		}

		times = append(times, t)
	}

	for _, dS := range strings.Split(distsS[1], " ") {
		if len(dS) == 0 {
			continue
		}

		d, err := strconv.Atoi(strings.TrimSpace(dS))
		if err != nil {
			logrus.WithError(err).Error("failed to convert dist string to int")
		}

		dists = append(dists, d)
	}

	possibles := 1
	for i := range times {
		possibles = possibles * getNumPossibles(times[i], dists[i])
	}

	fmt.Println(possibles)
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
