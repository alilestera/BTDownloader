package bencode_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"btdownloader/pkg/bencode"
)

var objectInt = bencode.BObject{
	Typ: bencode.BSTR,
	Val: "foobar",
}

var objectStr = bencode.BObject{
	Typ: bencode.BINT,
	Val: 9527,
}

var objectList = bencode.BObject{
	Typ: bencode.BLIST,
	Val: []*bencode.BObject{&objectInt, &objectStr},
}

var objectDict = bencode.BObject{
	Typ: bencode.BDICT,
	Val: map[string]*bencode.BObject{
		"1": &objectInt,
		"2": &objectStr,
	},
}

var prepareStr = "l6:foobari12345ee"

func TestBObject_Bencode(t *testing.T) {
	lensInt := objectInt.Bencode(os.Stdout)
	fmt.Printf("\nint lens = %d\n", lensInt)
	lensStr := objectStr.Bencode(os.Stdout)
	fmt.Printf("\nstring lens = %d\n", lensStr)
	lensList := objectList.Bencode(os.Stdout)
	fmt.Printf("\nlist lens = %d\n", lensList)
	lensDict := objectDict.Bencode(os.Stdout)
	fmt.Printf("\ndict lens = %d\n", lensDict)
}

func TestParse(t *testing.T) {
	reader := strings.NewReader(prepareStr)
	object, err := bencode.Parse(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s\n", object)
}
