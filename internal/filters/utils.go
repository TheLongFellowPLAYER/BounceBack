package filters

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrOddOrZero = errors.New("data length is odd or equal zero")
)

func FormatStringerSlice[T fmt.Stringer](s []T) string {
	slice := make([]string, 0, len(s))
	for _, d := range s {
		slice = append(slice, d.String())
	}
	return FormatStringSlice(slice)
}

func FormatStringSlice(s []string) string {
	return "[" + strings.Join(s, ", ") + "]"
}

func xorDecrypt(key []byte, data []byte) []byte {
	out := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		out[i] = data[i] ^ key[i%len(key)]
	}
	return out
}

func netbiosDecode(data []byte) ([]byte, error) {
	if len(data)%2 != 0 || len(data) == 0 {
		return nil, ErrOddOrZero
	}
	d := bytes.ToUpper(data)
	for i := 0; i < len(d); i += 2 {
		d[i/2] = ((d[i] - byte('A')) << 4) + //nolint:gomnd
			((d[i+1] - byte('A')) & 0xF) //nolint:gomnd
	}
	return d[:len(d)/2], nil
}

func matchByMask(s string, m string) bool {
	start := m[0] == '*'
	end := m[len(m)-1] == '*'
	switch {
	case start && end:
		return strings.Contains(s, m[1:len(m)-1])
	case start:
		return strings.HasSuffix(s, m[1:])
	case end:
		return strings.HasPrefix(s, m[:len(m)-1])
	default:
		return s == m
	}
}

func checksum8(data []byte) uint8 {
	var out uint8
	for _, b := range data {
		out += b
	}
	return out
}
