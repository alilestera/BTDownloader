package bencode_test

import (
	"bytes"
	"strings"
	"testing"

	"btdownloader/pkg/bencode"
)

var listObject = &bencode.BObject{
	Typ: bencode.BLIST,
	Val: []*bencode.BObject{
		{Typ: bencode.BSTR, Val: "1"},
		{Typ: bencode.BSTR, Val: "12"},
		{Typ: bencode.BSTR, Val: "123"},
		{Typ: bencode.BSTR, Val: "1234"},
		{Typ: bencode.BSTR, Val: "12345"},
	},
}

func listBObjectString() string {
	buf := bytes.Buffer{}
	_ = listObject.Bencode(&buf)
	s := buf.String()
	return s
}

func TestUnmarshal(t *testing.T) {
	i := &[]string{}
	s := listBObjectString()
	reader := strings.NewReader(s)
	err := bencode.Unmarshal(reader, i)
	if err != nil {
		t.Fatal(err)
	}
}
