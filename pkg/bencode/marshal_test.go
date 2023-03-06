package bencode_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

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

type Small struct {
	Phone   int    `bencode:"phone"`
	Address string `bencode:"address"`
}

type ComplexStruct struct {
	ID     int    `bencode:"id"`
	Name   string `bencode:"name"`
	List   []int  `bencode:"list"`
	Record Small  `bencode:"record"`
}

func listBObjectString() string {
	buf := bytes.Buffer{}
	_ = listObject.Bencode(&buf)
	s := buf.String()
	return s
}

// TestUnmarshal testing unmarshal a pointer to slice of strings
func TestUnmarshal(t *testing.T) {
	i := &[]string{}
	s := listBObjectString()
	reader := strings.NewReader(s)
	err := bencode.Unmarshal(reader, i)
	if err != nil {
		t.Fatal(err)
	}
	target := &[]string{
		"1",
		"12",
		"123",
		"1234",
		"12345",
	}
	assert.Equal(t, i, target, "The two string slice should be equal")
}

// TestMarshalComplexStruct testing marshal complex struct, which have four kind
// types: int, string, list, dict
func TestMarshalComplexStruct(t *testing.T) {
	s := &ComplexStruct{
		ID:   147,
		Name: "258",
		List: []int{3, 6, 9},
		Record: Small{
			Phone:   134679,
			Address: "nothing",
		},
	}
	buffer := bytes.Buffer{}
	lens := bencode.Marshal(&buffer, s)
	if lens == 0 {
		t.Fatalf("marshal struct '%v' failed", s)
	}
	target := "d2:idi147e4:name3:2584:listli3ei6ei9ee6:recordd5:phonei134679e7:address7:nothingee"
	assert.Equal(t, buffer.String(), target, "The two string slice should be equal")
}
