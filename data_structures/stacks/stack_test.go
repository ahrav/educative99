package stacks

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"abbaca", "ca"},
		{"azxxzy", "ay"},
		{"abbbacdddbac", "abacdbac"},
		{"", ""},
		{"abcdefg", "abcdefg"},
		{"aabbbccccdddd", "b"},
	}

	for _, tc := range testCases {
		result := removeDuplicates(tc.input)
		if result != tc.expected {
			t.Errorf("removeDuplicates(%q) = %q; expected %q", tc.input, result, tc.expected)
		}
	}
}

func TestMinRemoveParens(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))(", ""},
		{"(a(b(c)d)", "a(b(c)d)"},
		{"(ab(c)", "ab(c)"},
		{"", ""},
		{"abcdefg", "abcdefg"},
		{"((()", "()"},
		{"()", "()"},
	}

	for _, tc := range testCases {
		result := minRemoveParens(tc.input)
		if result != tc.expected {
			t.Errorf("minRemoveParens(%q) = %q; expected %q", tc.input, result, tc.expected)
		}
	}
}
