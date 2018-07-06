// Package variablelengthquantity implements a simple VLQ encoding/decoding
package variablelengthquantity

import "fmt"

// DecodeVarint returns the decoded in
func DecodeVarint(in []byte) ([]uint32, error) {
	if in[len(in)-1]&0x80 > 0 {
		return nil, fmt.Errorf("Incomplete sequence")
	}

	var u []uint32

	for i := 0; i < len(in); i++ {
		var b []byte
		for in[i]&0x80 > 0 {
			b = append(b, in[i])
			i++
		}
		b = append(b, in[i])
		u = append(u, decodeBytes(b))
	}

	return u, nil
}

// EncodeVarint returns the encoded in
func EncodeVarint(in []uint32) []byte {
	var b []byte

	for _, u := range in {
		b = append(b, encodeUint(u)...)
	}

	return b
}

func encodeUint(u uint32) []byte {
	if u == 0 {
		return []byte{0}
	}

	var b []byte

	for u > 0 {
		cur := byte(u & 0x7f)
		u >>= 7
		b = append([]byte{cur | 0x80}, b...)
	}

	b[len(b)-1] = b[len(b)-1] & 0x7f

	return b
}

func decodeBytes(in []byte) uint32 {
	var u uint32

	bytesCount := len(in)

	for i, b := range in {
		last := uint32(b & 0x7F)
		position := uint32((bytesCount - i - 1) * 7)
		u |= last << position
	}

	return u
}
