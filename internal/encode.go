package b64

var rfc4648 string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

func enc(b byte) byte {
	return rfc4648[b]
}

// encode bytes with b64
func EncodeBytes(in []byte) []byte {
	var (
		buf, tail, padding, off, i int
		out                        []byte
	)

	// output is fixed length
	// add padding if not even octets

	buf = (len(in) / 3) * 4
	tail = len(in) % 3
	if tail != 0 {
		out = make([]byte, buf+4)
		padding = 3 - tail
	} else {
		out = make([]byte, buf)
	}

	// output offset
	off = 0
	i = 0

	// encode octets of three to octets of four
	for i < len(in)-tail {
		out[off] = in[i] >> 2
		out[off+1] = (in[i]&0x3)<<4 + in[i+1]>>4
		out[off+2] = in[i+1]&0xf<<2 + in[i+2]>>6
		out[off+3] = in[i+2] & 0x3f
		off += 4
		i += 3
	}

	// handle padding
	if padding != 0 {
		out[off] = in[i] >> 2
		out[off+1] = (in[i] & 0x3) << 4
		out[off+2] = 0x40
		out[off+3] = 0x40

		if padding == 1 {
			out[off+1] += in[i+1] >> 4
			out[off+2] = (in[i+1] & 0xf << 2)
		}
	}

	for i := range out {
		out[i] = enc(out[i])
	}

	return out
}

// encode string with b64 to string
func EncodeStr(s string) string {
	return string(EncodeBytes([]byte(s)))
}
