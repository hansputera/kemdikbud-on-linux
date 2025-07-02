package utils

import (
	"regexp"
	"strings"
)

func ExtractVersion(s string) string {
	re := regexp.MustCompile(`\d+.\w+`)
	for x := range strings.SplitSeq(s, " ") {
		if re.MatchString(x) {
			return x
		}
	}

	return ""
}
