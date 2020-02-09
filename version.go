package version

import (
	"regexp"
)

var validator = regexp.MustCompile(`^go[1-9][0-9]*(\.[1-9][0-9]*(|(beta|rc)[1-9][0-9]*)|)$`)

// IsValid returns true if the version of go is a valid version.
func IsValid(v string) bool {
	return validator.MatchString(v)
}
