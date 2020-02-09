package version

import (
	"regexp"
)

var r *regexp.Regexp

func init() {
	r = regexp.MustCompile(`^go[1-9][0-9]*(\.[1-9][0-9]*(|(beta|rc)[1-9][0-9]*)|)$`)
}

// IsValid returns true if the version of go is a valid version.
func IsValid(v string) bool {
	return r.MatchString(v)
}
