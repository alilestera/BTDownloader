package bencode

import (
	"bufio"
	"errors"
	"io"
)

var (
	ErrNum = errors.New("wrong decoding number")
	ErrCol = errors.New("wrong decoding colon")
	ErrEpI = errors.New("wrong decoding sign 'i'")
	ErrEpE = errors.New("wrong decoding sign 'e'")
)

// DecodeString is to decode a string from a io.Reader
func DecodeString(r io.Reader) (val string, err error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	num, lens := readDecimal(br)
	if lens == 0 {
		return val, ErrNum
	}
	b, err := br.ReadByte()
	if err != nil {
		return val, err
	}
	if b != ':' {
		return val, ErrCol
	}
	buf := make([]byte, num)
	_, err = io.ReadAtLeast(br, buf, num)
	val = string(buf)
	return
}

func DecodeInt(r io.Reader) (val int, err error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	b, err := br.ReadByte()
	if err != nil {
		return val, err
	}
	if b != 'i' {
		return val, ErrEpI
	}
	val, _ = readDecimal(br)
	b, err = br.ReadByte()
	if err != nil {
		return val, err
	}
	if b != 'e' {
		return val, ErrEpE
	}
	return
}

// readDecimal is to read a decimal integer from a bufio.Reader
// It is used by Decode... functions
func readDecimal(w *bufio.Reader) (num, lens int) {
	for {
		bs, err := w.Peek(1)
		if err != nil {
			return
		}
		b := bs[0]
		if b < '0' || b > '9' {
			break
		}
		num = num*10 + int(b-'0')
		lens++
		w.ReadByte()
	}
	return
}
