package bencode_test

import (
	"fmt"
	"os"
	"testing"

	"btdownloader/pkg/bencode"
)

var (
	number = 123
	str    = "foobar"
)

func TestEncodeInt(t *testing.T) {
	w := bencode.EncodeInt(os.Stdout, number)
	fmt.Printf("\nw = %d\n", w)
}

func TestEncodeString(t *testing.T) {
	w := bencode.EncodeString(os.Stdout, str)
	fmt.Printf("\nw = %d\n", w)
}
