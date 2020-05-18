package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// ContainsAll !
func ContainsAll(sliceToCheck []string, slice []string) bool {
	if len(sliceToCheck) == 0 {
		return false
	}
outsideFor:
	for _, v := range sliceToCheck {
		for _, v2 := range slice {
			if v == v2 {
				continue outsideFor
			}
		}
		return false
	}
	return true
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

// NormalizeString !
func NormalizeString(toNormalize string) string {
	var t = transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, toNormalize)
	return strings.TrimSpace(strings.ToLower(result))
}
