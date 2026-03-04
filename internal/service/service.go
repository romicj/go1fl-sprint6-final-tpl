package service

import "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"

func isMorse(s string) bool {
	for _, ch := range s {
		if ch != '.' && ch != '-' && ch != ' ' {
			return false
		}
	}
	return true
}

func Convert(s string) string {
	isMorse := isMorse(s)
	if isMorse {
		return morse.ToText(s)
	}
	return morse.ToMorse(s)
}
