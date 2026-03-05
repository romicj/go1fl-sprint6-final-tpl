package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(s string) bool {
	return strings.ContainsFunc(s, func(r rune) bool {
		return r == '.' || r == '-' || r == ' '
	})
}

func Convert(s string) string {
	isMorse := isMorse(s)
	if isMorse {
		return morse.ToText(s)
	}
	return morse.ToMorse(s)
}
