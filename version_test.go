package version

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		goVer     string
		wantValid bool
	}{
		{"meow", false},
		{"go1.", false},
		{"go0", false},
		{"go01", false},
		{"go1", true},
		{"go10", true},
		{"go2", true},
		{"go01", false},
		{"go1.1", true},
		{"go1.01", false},
		{"go1.10", true},
		{"go1.10beta2", true},
		{"go1.11beta123", true},
		{"go1.10beta0", false},
		{"go1.10rc2", true},
		{"go1.10rc0", false},
		{"go1.rc1", false},
		{"go2.123rc123", true},
		{"go1.beta", false},
		{"gobeta", false},
		{"go1.1betarc1", false},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test=%d", idx), func(t *testing.T) {
			gotValid := IsValid(tc.goVer)
			if gotValid != tc.wantValid {
				t.Errorf("IsValid(%q) want=%t got=%t", tc.goVer, tc.wantValid, gotValid)
			}
		})
	}
}
