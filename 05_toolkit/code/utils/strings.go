package utils

import (
	"fmt"
	"strings"
)

// MakeExcited transforms a sentence to all caps with an exclamation mark
func MakeExcited(origString string) string {
	excited := fmt.Sprintf(strings.ToUpper(origString) + "!")
	return excited
}
