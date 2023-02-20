package b64_test

import (
	"testing"

	b64 "github.com/samharju/b64/internal"
)

func TestDecode(t *testing.T) {

	type tc struct {
		input       string
		expected    string
		expectedErr bool
		errStr      string
	}
	cases := []tc{
		{
			"bGlnaHQgdw==",
			"light w",
			false,
			"",
		},
		{
			"bGlnaHQgd28=",
			"light wo",
			false,
			"",
		},
		{
			"bGlnaHQgd29y",
			"light wor",
			false,
			"",
		},
		{
			"pasda",
			"",
			true,
			"input has invalid length: 5",
		},
		{
			"asda[]{}",
			"",
			true,
			"position 4: invalid char: [",
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {

			got, err := b64.DecodeStr(c.input)

			if c.expectedErr {
				if err == nil {
					t.Fatal("expected error but got nil")
				} else if c.errStr != err.Error() {
					t.Fatalf("want: '%s', got: '%s'", c.errStr, err)
				}
			}

			if err != nil && !c.expectedErr {
				t.Fatalf("did not expect error: %s", err)
			}

			if got != c.expected {
				t.Fatalf("%s | want: '%s', got: '%s'", c.input, c.expected, got)
			}
		})

	}

}
