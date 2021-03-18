package main

import (
	"testing"

	"github.com/mmkader85/hello/morestrings"
)

func TestReverseRunes(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "Hello!",
			want: "!olleH",
		},
		{
			in:   "",
			want: "",
		},
	}

	for _, testCase := range testCases {
		got := morestrings.ReverseRunes(testCase.in)
		if got != testCase.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", testCase.in, got, testCase.want)
		}
	}
}
