package utils

import "regexp"

func IsSymbolValid(symbol string) bool {
	return regexp.MustCompile(`^[A-Z0-9\-]+$`).MatchString(symbol)
}
