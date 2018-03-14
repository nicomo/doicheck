package doicheck

import (
	"reflect"
	"testing"
)

type doiTest struct {
	input    string
	expected bool
}

func TestIsValid(t *testing.T) {

	var doiTests = []doiTest{
		{
			input:    "",
			expected: false,
		},
		{
			input:    "az.azeaze/azeaze",
			expected: false,
		},
		{
			input:    "10.",
			expected: false,
		},
		{
			input:    "10.1234/",
			expected: false,
		},
		{
			input:    "10.1234/",
			expected: false,
		},
		{
			input:    "10.1080/0950236X.2017.1412948",
			expected: true,
		},
		{
			input:    "10.1080/pàç!uioç><///:;",
			expected: true,
		},
	}

	for _, test := range doiTests {
		actual := IsValid(test.input)
		if reflect.DeepEqual(test.expected, actual) {
			t.Logf("PASS: got %v", test.expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %v", test.input, test.expected, actual)
		}
	}
}
