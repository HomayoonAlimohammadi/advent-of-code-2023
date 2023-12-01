package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var numStrs = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

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
		first, err := findFirstDigit(line)
		if err != nil {
			logrus.Error(err)
		}

		last, err := findLastDigit(line)
		if err != nil {
			logrus.Error(err)
		}

		cal, _ := strconv.Atoi(first + last)
		dataCh <- cal
	}

	close(dataCh)
}

func findFirstDigit(line string) (string, error) {
	for i := 0; i < len(line); i++ {
		if isDigit(line[i]) {
			return string(line[i]), nil
		} else {
			for numStr, num := range numStrs {
				s := line[i:min(i+len(numStr), len(line))]
				if s == numStr {
					return num, nil
				}
			}
		}
	}

	return "", fmt.Errorf("failed to find number in line `%s`", line)
}

func findLastDigit(line string) (string, error) {
	for i := len(line) - 1; i >= 0; i-- {
		if isDigit(line[i]) {
			return string(line[i]), nil
		} else {
			for numStr, num := range numStrs {
				s := line[max(0, i-len(numStr)+1) : i+1]
				if s == numStr {
					return num, nil
				}
			}
		}
	}

	return "", fmt.Errorf("failed to find number in line `%s`", line)
}

func sumCalibrations(dataCh <-chan int) (sum int) {
	for cal := range dataCh {
		sum = sum + cal
	}

	return sum
}

func isDigit(r byte) bool {
	return 48 <= r && r <= 57
}
