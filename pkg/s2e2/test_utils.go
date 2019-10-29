package s2e2

import "testing"

func compareTokens(actual, expected []token, test *testing.T) {
	if len(actual) != len(expected) {
		test.Errorf("Wrong output size %v instead of %v", len(actual), len(expected))
		return
	}
	for i, value := range actual {
		if value != expected[i] {
			test.Errorf("Wrong output token %v instead of %v", value, expected[i])
			return
		}
	}
}
