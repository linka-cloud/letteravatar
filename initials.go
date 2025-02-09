package letteravatar

import (
	"strings"
	"unicode/utf8"
)

func Initials(name string) string {
	var letters string
	parts := strings.Fields(name)
	for _, v := range parts {
		l, _ := utf8.DecodeRuneInString(v)
		letters += string(l)
	}
	return letters
}
