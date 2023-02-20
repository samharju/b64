package b64

import (
	"fmt"
)

// error returned when decoding chars not in rfc4648 table
type InputCharError byte

func (ice InputCharError) Error() string {
	return fmt.Sprintf("invalid char: %c", ice)
}

// error returned trying to decode strings with wrong number of octets
type InputLenError int

func (lene InputLenError) Error() string {
	return fmt.Sprintf("input has invalid length: %d", lene)
}

func decode(b byte) (byte, error) {
	if b == '=' {
		return 0, nil
	}
	for i := range rfc4648 {
		if rfc4648[i] == b {
			return byte(i), nil
		}
	}
	return 0, InputCharError(b)
}

// decode b64 encoded bytes
func DecodeBytes(in []byte) ([]byte, error) {

	if len(in)%4 != 0 {
		return []byte{}, InputLenError(len(in))
	}
	var (
		out                  []byte
		buf, off, i, padding int
		j, k                 byte
		err                  error
	)

	// output is fixed length

	buf = (len(in) / 4) * 3
	out = make([]byte, buf)

	for i < len(in) {
		j, err = decode(in[i])
		if err != nil {
			return out, fmt.Errorf("position %d: %w", i, err)
		}

		k, err = decode(in[i+1])
		if err != nil {
			return out, fmt.Errorf("position %d: %w", i+1, err)
		}
		out[off] = j<<2 + k>>4

		j, err = decode(in[i+2])
		if err != nil {
			return out, fmt.Errorf("position %d: %w", i+2, err)
		}
		out[off+1] = k&0xf<<4 + j>>2

		k, err = decode(in[i+3])
		if err != nil {
			return out, fmt.Errorf("position %d: %w", i+3, err)
		}
		out[off+2] = j&0x3<<6 + k

		off += 3
		i += 4
	}

	// check padding
	for i = 1; i <= 2; i++ {
		if in[len(in)-i] == '=' {
			padding++
		}
	}

	return out[:len(out)-padding], nil
}

// decode b64 encoded string
func DecodeStr(s string) (string, error) {
	o, err := DecodeBytes([]byte(s))
	if err != nil {
		return "", err
	}
	return string(o), nil

}
