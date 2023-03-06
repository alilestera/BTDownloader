package bencode

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Name  string   `bencode:"name"`
	Age   int      `bencode:"age"`
	Cards []string `bencode:"cards"`
}

type Identifier struct {
	ID   string `bencode:"id"`
	User User   `bencode:"user"`
}

// SimpleStruct for testing marshal struct
type SimpleStruct struct {
	num int    `bencode:"num"`
	str string `bencode:"str"'`
}

var sliceInts = []*BObject{
	{Typ: BINT, Val: 11},
	{Typ: BINT, Val: 12},
	{Typ: BINT, Val: 13},
}

var sliceLists = []*BObject{
	{Typ: BLIST, Val: []*BObject{
		{Typ: BINT, Val: 1},
	}},
	{Typ: BLIST, Val: []*BObject{
		{Typ: BINT, Val: 1},
		{Typ: BINT, Val: 2},
	}},
	{Typ: BLIST, Val: []*BObject{
		{Typ: BINT, Val: 1},
		{Typ: BINT, Val: 2},
		{Typ: BINT, Val: 3},
	}},
}

// userObjects include a slice of string
var userObjects = map[string]*BObject{
	"name": {Typ: BSTR, Val: "yellow"},
	"age":  {Typ: BINT, Val: 20},
	"cards": {Typ: BLIST, Val: []*BObject{
		{Typ: BSTR, Val: "card1"},
		{Typ: BSTR, Val: "card2"},
		{Typ: BSTR, Val: "card3"},
	}},
}

// identifierObject include a struct 'User'
var identifierObject = map[string]*BObject{
	"id":   {Typ: BSTR, Val: "666"},
	"user": {Typ: BDICT, Val: userObjects},
}

// TestUnmarshalSlice mainly for testing a slice of ints
func TestUnmarshalSlice(t *testing.T) {
	s := &[]int{}
	p := reflect.ValueOf(s)
	l := reflect.MakeSlice(p.Type().Elem(), len(sliceInts), len(sliceInts))
	p.Elem().Set(l)
	err := unmarshalList(p, sliceInts)
	if err != nil {
		t.Fatal(err)
	}
	target := &[]int{11, 12, 13}
	assert.Equal(t, s, target, "The two int slice should be equal")
}

func TestUnmarshalSliceOfSlices(t *testing.T) {
	s := &[][]int{}
	p := reflect.ValueOf(s)
	l := reflect.MakeSlice(p.Type().Elem(), len(sliceLists), len(sliceLists))
	p.Elem().Set(l)
	err := unmarshalList(p, sliceLists)
	if err != nil {
		t.Fatal(err)
	}
	target := &[][]int{
		{1},
		{1, 2},
		{1, 2, 3},
	}
	assert.Equal(t, s, target, "The two 2d int slice should be equal")
}

// TestUnmarshalStructUser mainly for testing a struct with slice
func TestUnmarshalStructUser(t *testing.T) {
	s := &User{}
	p := reflect.ValueOf(s)
	err := unmarshalDict(p, userObjects)
	if err != nil {
		t.Fatal(err)
	}
	target := &User{
		Name: "yellow",
		Age:  20,
		Cards: []string{
			"card1",
			"card2",
			"card3",
		},
	}
	assert.Equal(t, s, target, "The two user object should be equal")
}

// TestUnmarshalStructIdentifier mainly for testing a struct with struct
func TestUnmarshalStructIdentifier(t *testing.T) {
	s := &Identifier{}
	p := reflect.ValueOf(s)
	err := unmarshalDict(p, identifierObject)
	if err != nil {
		t.Fatal(err)
	}
	target := &Identifier{
		ID: "666",
		User: User{
			Name: "yellow",
			Age:  20,
			Cards: []string{
				"card1",
				"card2",
				"card3",
			},
		},
	}
	assert.Equal(t, s, target, "The two identifier object should be equal")
}

// TestMarshalString for testing marshal string
// the function logic of marshal integer is same as this function logic
func TestMarshalString(t *testing.T) {
	s := "abcde"
	p := reflect.ValueOf(s)
	buffer := bytes.Buffer{}
	lens := marshalValue(&buffer, p)
	if lens == 0 {
		t.Fatalf("marshal string '%s' failed", s)
	}
	target := "5:abcde"
	assert.Equal(t, buffer.String(), target, "The two string should be equal")
}

func TestMarshalSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	p := reflect.ValueOf(s)
	buffer := bytes.Buffer{}
	lens := marshalValue(&buffer, p)
	if lens == 0 {
		t.Fatalf("marshal slice '%v' failed", s)
	}
	target := "li1ei2ei3ei4ei5ee"
	assert.Equal(t, buffer.String(), target, "The two string should be equal")
}

func TestMarshalStruct(t *testing.T) {
	s := SimpleStruct{num: 33, str: "abc"}
	p := reflect.ValueOf(s)
	buffer := bytes.Buffer{}
	lens := marshalValue(&buffer, p)
	if lens == 0 {
		t.Fatalf("marshal struct '%v' failed", s)
	}
	target := "d3:numi33e3:str3:abce"
	assert.Equal(t, buffer.String(), target, "The two string should be equal")
}
