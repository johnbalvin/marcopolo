package utils

import (
	"regexp"
	"strings"
)

var regexSpace = regexp.MustCompile(`[\s ]+`)

func RemoveSpace(value string) string {
	return regexSpace.ReplaceAllString(strings.TrimSpace(value), " ")
}
