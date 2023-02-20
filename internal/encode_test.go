package b64_test

import (
	"testing"

	b64 "github.com/samharju/b64/internal"
)

func TestEncode(t *testing.T) {

	type tc struct {
		input    string
		expected string
	}

	cases := []tc{
		{
			"light work.",
			"bGlnaHQgd29yay4=",
		},
		{
			"light work",
			"bGlnaHQgd29yaw==",
		},
		{
			"light wor",
			"bGlnaHQgd29y",
		},
		{
			"light wo",
			"bGlnaHQgd28=",
		},
		{
			"light w",
			"bGlnaHQgdw==",
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {

			if got := b64.EncodeStr(c.input); got != c.expected {
				t.Fatalf("want: '%s', got: '%s'", c.expected, got)
			}

		})
	}
}
