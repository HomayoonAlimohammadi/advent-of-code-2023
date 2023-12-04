package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDataB []byte

//go:embed test.txt
var testDataB []byte

func main() {
	input := string(inputDataB)
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.Split(lines[i], ":")[1]
	}

	cards := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		cards[i] = 1
	}

	for i, l := range lines {
		splitted := strings.Split(l, "|")
		winning := strings.Split(strings.TrimSpace(splitted[0]), " ")
		ours := strings.Split(strings.TrimSpace(splitted[1]), " ")

		oursMap := map[string]struct{}{}

		for _, o := range ours {
			o = strings.TrimSpace(o)
			if o == "" {
				continue
			}
			if _, found := oursMap[o]; !found {
				oursMap[o] = struct{}{}
			}
		}

		matches := []string{}
		for _, w := range winning {
			w = strings.TrimSpace(w)
			if w == "" {
				continue
			}
			if _, found := oursMap[w]; found {
				matches = append(matches, w)
			}
		}

		// fmt.Println("oursMap:", oursMap)
		// fmt.Println("ours:", ours)
		// fmt.Println("winning:", winning)
		// fmt.Println("matches:", matches)

		if len(matches) != 0 {
			for j := 1; j <= len(matches); j++ {
				cards[i+j] = cards[i+j] + (cards[i])
			}
		}

		// fmt.Println()
	}

	fmt.Println(sum(cards))
}

func sum(arr []int) int {
	s := 0
	for _, x := range arr {
		s = s + x
	}

	return s
}
