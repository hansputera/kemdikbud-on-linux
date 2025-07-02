package utils

import (
	"regexp"
	"strconv"
)

func ExtractNumbersOnly(s string, m int) []int {
	t := []int{}
	re := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(s, m)
	for _, x := range matches {
		num, err := strconv.Atoi(x)
		if err != nil {
			continue
		}

		t = append(t, num)
	}

	return t
}
