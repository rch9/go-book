package comma

import (
	"go-book/chapter3/exercises/10/comma"
	"strings"
)

func Comma(s string) string {
	var mantissa string
	var sign string

	if s[0] == '-' {
		sign = "-"
		s = s[1:]
	}

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		mantissa = s[dot:]
		s = s[:dot]
	}

	return sign + comma.Comma(s) + mantissa
}
