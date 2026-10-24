package util

import "testing"

func TestFormat(t *testing.T) {
	cases := map[Cents]string{0: "$0.00", 5: "$0.05", 1999: "$19.99", 100000: "$1000.00"}
	for in, want := range cases {
		if got := Format(in); got != want {
			t.Errorf("Format(%d) = %q, want %q", in, got, want)
		}
	}
}
