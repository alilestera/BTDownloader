package bencode

import (
	"bufio"
	"io"
	"strconv"
)

// EncodeString is to encode a string to a io.Writer
func EncodeString(w io.Writer, val string) int {
	lens := len(val)
	bw := bufio.NewWriter(w)
	wLen := writeDecimal(bw, lens)
	bw.WriteByte(':')
	bw.WriteString(val)
	wLen += 1 + lens
	err := bw.Flush()
	if err != nil {
		return 0
	}
	return wLen
}

// EncodeInt is to encode a integer to a io.Writer
func EncodeInt(w io.Writer, val int) int {
	bw := bufio.NewWriter(w)
	bw.WriteByte('i')
	wLen := writeDecimal(bw, val)
	bw.WriteByte('e')
	wLen += 2
	err := bw.Flush()
	if err != nil {
		return 0
	}
	return wLen
}

// writeDecimal is to write a decimal integer to a bufio.Writer
func writeDecimal(w *bufio.Writer, val int) int {
	strNum := strconv.Itoa(val)
	w.WriteString(strNum)
	return len(strNum)
}
