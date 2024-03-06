package asn

import "strings"

func (asn Asn) IsForbidden(currentName string) bool {
	for _, asnName := range asn.ForbiddenNames {
		if strings.Contains(currentName, asnName) {
			return true
		}
	}
	return false
}

func (asn Asn) IsPriority(currentName string) (int, bool, string) {
	for i, priorityName := range asn.PrioritiesNames {
		if strings.Contains(currentName, priorityName) {
			return i, true, priorityName
		}
	}
	return 0, false, ""
}
