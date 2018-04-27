package comma

import (
	"bytes"
	"fmt"
)

func Comma(s string) string {
	var buf bytes.Buffer

	smod3 := len(s) % 3

	for i := 0; i < len(s); i++ {
		if ((i-smod3)%3 == 0 || i == smod3) && i != 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%c", s[i])
	}

	return buf.String()
}
