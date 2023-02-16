package bencode_test

import (
	"fmt"
	"strings"
	"testing"

	"btdownloader/pkg/bencode"
)

var (
	strInput = "6:foobar"
	numInput = "i147258e"
)

func TestDecodeString(t *testing.T) {
	reader := strings.NewReader(strInput)
	s, err := bencode.DecodeString(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("val = %s \n", s)
}

func TestDecodeInt(t *testing.T) {
	reader := strings.NewReader(numInput)
	n, err := bencode.DecodeInt(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("val = %d \n", n)
}
