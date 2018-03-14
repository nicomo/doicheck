package doicheck

import (
	"strings"
)

// IsValid checks that a string is structured as a valid DOI number
// which doesn't mean much, really...
// see https://www.doi.org/doi_handbook/2_Numbering.html#2.2
func IsValid(doi string) bool {

	if !strings.HasPrefix(doi, "10.") {
		return false
	}
	s1 := strings.SplitAfterN(doi, ".", 2)
	if len(s1[1]) == 0 {
		return false
	}
	s2 := strings.SplitAfterN(s1[1], "/", 2)
	if len(s2) != 2 {
		return false
	}
	if len(s2[0]) == 0 || len(s2[1]) == 0 {
		return false
	}
	return true
}
