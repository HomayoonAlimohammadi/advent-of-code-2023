package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	b, err := os.ReadFile("../puzzle.txt")
	if err != nil {
		logrus.WithError(err).Fatal("failed to read input data")
	}

	lines := strings.Split(string(b), "\n")
	dataCh := make(chan int, len(lines))

	go findCalibration(dataCh, lines)

	sum := sumCalibrations(dataCh)
	fmt.Println(sum)
}

func findCalibration(dataCh chan<- int, data []string) {
	for _, line := range data {
		var first, last string
		// forward
		for _, char := range line {
			if isDigit(char) {
				first = string(char)
				break
			}
		}

		// reverse
		for i := len(line) - 1; i >= 0; i-- {
			if isDigit(rune(line[i])) {
				last = string(line[i])
				break
			}
		}

		cal, _ := strconv.Atoi(first + last)
		dataCh <- cal
	}

	close(dataCh)
}

func sumCalibrations(dataCh <-chan int) (sum int) {
	for cal := range dataCh {
		sum = sum + cal
	}

	return sum
}

func isDigit(r rune) bool {
	return 48 <= r && r <= 57
}
